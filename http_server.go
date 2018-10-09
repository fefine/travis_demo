package travis_demo

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Server(port int)  {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/image", imageHandler)
	log.Printf("server start at: :%d", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	log.Fatal(err)
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	log.Printf("[%s] %s", req.URL, req.RemoteAddr)
	io.WriteString(w, "Hello, world!\n")
}

func imageHandler(w http.ResponseWriter, req *http.Request)  {
	log.Printf("[%s] %s", req.URL, req.RemoteAddr)
	io.WriteString(w, "<img src='https://www.bing.com/az/hprichbg/rb/NorseBuilding_EN-CN6747879545_1920x1080.jpg' />")
}