package travis_demo

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Server(port int)  {
	http.HandleFunc("/", indexHandler)
	log.Printf("server start at: :%d", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	log.Fatal(err)
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	log.Printf("[%s] %s", req.URL, req.RemoteAddr)
	io.WriteString(w, "Hello, world!\n")
}
