package handlers

import (
	"log"
	"net/http"
)

type bye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *bye {
	return &bye{l}
}

func (g *bye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Byeee"))
}
