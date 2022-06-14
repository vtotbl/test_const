package usecase

import (
	"context"
	"net/http"
	"net/url"
	"testConstant/pkg/sender"
	"time"
)

const workersCount int = 4

type sendReqsUCase struct {
}

func NewSendReqs() *sendReqsUCase {
	return &sendReqsUCase{}
}

func (u *sendReqsUCase) Send(ctx context.Context, strUrls []string) ([]*http.Response, error) {
	queue := make(chan *url.URL, len(strUrls))
	errC := make(chan error)
	senderCtx, cancel := context.WithTimeout(ctx, time.Second)

	defer func() {
		cancel()
		close(errC)
	}()

	go u.fillQueue(strUrls, queue, errC)

	workers, rspC, wErrC := sender.NewWorker(senderCtx, queue, workersCount)
	workers.Run()

	rsps := make([]*http.Response, 0, len(strUrls))

	for {
		select {
		case err := <-wErrC: // ошибка пришла с воркера
			return nil, err
		case err := <-errC: // ошибка пришла заполнения очереди
			return nil, err
		case rsp := <-rspC:
			rsps = append(rsps, rsp)
			if len(rsps) == len(strUrls) {
				return rsps, nil
			}
		}
	}
}

// fillQueue Парсинг url адресов и добавление в очередь
func (u *sendReqsUCase) fillQueue(strUrls []string, queue chan<- *url.URL, errC chan<- error) {
	defer close(queue)
	for _, strUrl := range strUrls {
		getUrl, err := url.Parse(strUrl)
		if err != nil {
			errC <- err
			return
		}
		queue <- getUrl
	}

	return
}
