package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/chukwuka-emi/hackerNews_comments_sentiments/sentiment_analyzer/messagebroker"
)

// Server ...
type Server struct{}

type httpHandler struct{

}

func (h *httpHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
	){
	if r.Method=="GET" && r.URL.Path =="/health" {
		w.Header().Add("Content-Type","application/json")
		w.Write([]byte("Server Okay and Running!"))
	}
}

// Start HTTP Server
func (r * Server) Start(){
	port := os.Getenv("PORT")
	if port == "" {
		port = "5003"
	}
	go messagebroker.InitMessageBroker()
	httpRouter := &httpHandler{}
	log.Println("Server listening on Port",port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port),httpRouter)

	if err !=nil{
		log.Fatal("Error Starting HTTP Server",err.Error())
	}
	
}


