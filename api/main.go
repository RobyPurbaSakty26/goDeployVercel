package main

import (
	"myapp/api/handler"
	"net/http"
)

type res struct {
	w http.ResponseWriter
}

func main() {
	wr := res{}
	r := &http.Request{}
	handler.Handler(wr.w, r)
}
