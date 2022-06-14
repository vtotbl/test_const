package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testConstant/internal/handler/requests"
	"testConstant/internal/handler/responses"
	"testConstant/internal/handler/validation"
)

type senderUCase interface {
	Send(ctx context.Context, strUrls []string) ([]*http.Response, error)
}

type rateLimiter interface {
	CanDoWork() bool
	Done()
}

type handler struct {
	senderUCase senderUCase
	rateLimiter rateLimiter
}

func NewHandler(senderUCase senderUCase, rateLimiter rateLimiter) *handler {
	return &handler{
		senderUCase: senderUCase,
		rateLimiter: rateLimiter,
	}
}

func (h *handler) SendRequests(w http.ResponseWriter, r *http.Request) {
	defer h.rateLimiter.Done()

	if !h.rateLimiter.CanDoWork() {
		http.Error(w, "the server is overloaded", http.StatusServiceUnavailable)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "request method must be POST", http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("request body error: %s", err.Error()), http.StatusBadRequest)
		return
	}

	var sendReq requests.SendReq
	err = json.Unmarshal(body, &sendReq)
	if err != nil {
		http.Error(w, fmt.Sprintf("request body error: %s", err.Error()), http.StatusBadRequest)
		return
	}

	err = validation.SenReq(sendReq)
	if err != nil {
		http.Error(w, fmt.Sprintf("validation error: %s", err.Error()), http.StatusBadRequest)
		return
	}

	rsps, err := h.senderUCase.Send(r.Context(), sendReq.Urls)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var sendRsps responses.SendRsp
	for _, rsp := range rsps {
		var rspBody []byte
		rspBody, err = io.ReadAll(rsp.Body)
		if err != nil {
			fmt.Println(err.Error())
			http.Error(w, "internal sever error", http.StatusInternalServerError)
			return
		}
		sendRsps.Responses = append(sendRsps.Responses, string(rspBody))
	}
	sendRspBytes, err := json.Marshal(sendRsps)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "internal sever error", http.StatusInternalServerError)
		return
	}

	_, err = fmt.Fprintf(w, string(sendRspBytes))
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "internal sever error", http.StatusInternalServerError)
		return
	}
}
