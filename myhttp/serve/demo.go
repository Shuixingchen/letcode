package serve

import (
	"letcode/myhttp/serve/gee"
	"net/http"
)

func NewServer() {
	engine := gee.New()
	engine.GET("/demo", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("demo"))
	})

	engine.Run(":8080")
}
