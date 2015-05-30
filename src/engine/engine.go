package engine

import (
	"config"
	"fmt"
	"github.com/nporsche/np-golang-logging"
	"net/http"
	"runtime/debug"
)

var logger = logging.MustGetLogger("engine")

func Run(handlers map[string]http.HandlerFunc) {
	for url, handler := range handlers {
		http.Handle(url, safeHandler(url, handler))
	}

	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.This.Listen), nil); err != nil {
		logger.Fatal("ListenAndServe:" + err.Error())
	}
	logger.Debug("Run finished")
}

func safeHandler(url string, rHandler http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		defer faultRecovery()
		rHandler(rw, r)
	}
}

func faultRecovery() {
	if x := recover(); x != nil {
		logger.Error("Uncaught exception=[%v] stack=[%s]", x, string(debug.Stack()))
	}
}
