package main

import (
	"fmt"
	"grl"
	"net/http"
	"time"
)

type student struct {
	Name string
	Age  int8
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	r := grl.Default()
	r.GET("/", func(context *grl.Context) {
		context.String(http.StatusOK,"hello GRL framework!")
	})
	r.GET("/panic", func(context *grl.Context) {
		names := []string{"nice move!"}
		context.String(http.StatusOK,names[100])
	})

	r.Run(":9999")
}
