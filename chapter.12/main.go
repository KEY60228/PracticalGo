package main

import (
	"fmt"
	"log"
	"net/http"

	"go.uber.org/zap"
)

type logForwarder struct {
	l *zap.SugaredLogger
}

func (fw *logForwarder) Write(p []byte) (int, error) {
	fw.l.Errorw(string(p))
	return len(p), nil
}

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World")
	}
	http.HandleFunc("/test", handler)

	l, err := zap.NewDevelopment()
	if err != nil {
		l = zap.NewNop()
	}
	logger := l.Sugar()

	server := &http.Server{
		Addr:     "localhost:8888",
		ErrorLog: log.New(&logForwarder{l: logger}, "", 0),
	}

	logger.Fatal("server: %v", server.ListenAndServe())
}
