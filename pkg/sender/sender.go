package sender

import (
	"context"
	"net/http"
	"net/url"
)

type sender struct {
}

func Default() *sender {
	return &sender{}
}

func (s *sender) Send(ctx context.Context, url *url.URL) (*http.Response, error) {
	rspC := make(chan *http.Response)
	errC := make(chan error)

	go s.send(ctx, url, rspC, errC)

	select {
	case <-ctx.Done(): // если контекст уже отменили, то выйдем из функции, горутина send корректно завершится
		return nil, nil
	case err := <-errC:
		return nil, err
	case rsp := <-rspC:
		return rsp, nil
	}
}

func (s *sender) send(ctx context.Context, url *url.URL, rspC chan<- *http.Response, errC chan<- error) {
	defer func() {
		close(rspC)
		close(errC)
	}()

	rsp, err := http.Get(url.String())
	if err != nil {
		errC <- err
		return
	}

	select {
	case <-ctx.Done():
		return
	default:
		rspC <- rsp
		return
	}
}
