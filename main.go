package main

import (
	// Standard library packages

	"log"
	"net"
	"net/http"
	"text/template"

	// Third party packages

	"github.com/julienschmidt/httprouter"
	"github.com/skratchdot/open-golang/open"
)

// go get github.com/toqueteos/webbrowser

func main() {
	l, err := net.Listen("tcp", "localhost:3000")
	if err != nil {
		log.Fatal(err)
	}
	r := httprouter.New()
	// Add a handler on /test
	r.GET("/test", testHandler)

	open.RunWith("http://localhost:3000/test", "chrome")
	// Start the blocking server loop
	http.Serve(l, r)
}

func testHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t, _ := template.ParseFiles("view/index.html", "view/head.html", "view/foot.html")
	//执行主模版,主要调用的方法
	t.ExecuteTemplate(w, "layout", nil)
}
