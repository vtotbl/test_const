package main

import (
	"github.com/vtotbl/test_const.git/internal/handler"
	"github.com/vtotbl/test_const.git/internal/usecase"
	"github.com/vtotbl/test_const.git/pkg/rate_limiter"
	"log"
	"net/http"
)

func main() {
	senderUCase := usecase.NewSendReqs()
	rateLimiter := rate_limiter.NewLimiter(100) //условие задачи в 100 одновременных запросов

	h := handler.NewHandler(senderUCase, rateLimiter)
	http.HandleFunc("/send", h.SendRequests)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
