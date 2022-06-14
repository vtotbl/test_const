package sender

import (
	"context"
	"net/http"
	"net/url"
	"sync"
)

type workerSender interface {
	Send(ctx context.Context, url *url.URL) (*http.Response, error)
}

type worker struct {
	ctx    context.Context
	queue  <-chan *url.URL
	errC   chan<- error
	rspC   chan<- *http.Response
	sender workerSender
	count  int
	wg     sync.WaitGroup
}

func NewWorker(ctx context.Context, queue <-chan *url.URL, count int) (*worker, <-chan *http.Response, <-chan error) {
	rspC := make(chan *http.Response)
	errC := make(chan error)
	w := &worker{
		ctx:    ctx,
		queue:  queue,
		sender: Default(),
		count:  count,
		errC:   errC,
		rspC:   rspC,
		wg:     sync.WaitGroup{},
	}

	return w, rspC, errC
}

// SetSender Можно установить свой sender, чтобы изменить логику отправки
func (w *worker) SetSender(sender workerSender) {
	w.sender = sender
}

func (w *worker) Run() {
	go func() {
		for i := 0; i < w.count; i++ {
			w.wg.Add(1)
			go w.send()
		}

		w.wg.Wait() // ждем пока все вокреры завершат свою работу и только потом закрываем каналы
		close(w.rspC)
		close(w.errC)
	}()
}

func (w *worker) send() {
	defer func() {
		w.wg.Done()
	}()

	for {
		select {
		case <-w.ctx.Done():
			return
		default:
			u, ok := <-w.queue
			if ok {
				rsp, err := w.sender.Send(w.ctx, u)
				if err != nil {
					w.errC <- err
					return
				}

				if rsp != nil {
					w.rspC <- rsp
				}
			} else {
				return
			}
		}
	}
}
