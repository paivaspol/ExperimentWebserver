package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// Implements the server.
type Server struct {
	filename string
}

// Hello implements the handler for HTTP request.
func (s *Server) Hello(w http.ResponseWriter, r *http.Request) {
	filename := filepath.Base(s.filename)
	log.Printf("%v\n", filename)

	data, err := ioutil.ReadFile(s.filename)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func main() {
	filename := flag.String("file", "novalid", "the file to be served")
	port := flag.Int("port", -1, "The port for the server to bind to")
	flag.Parse()
	if _, err := os.Stat(*filename); os.IsNotExist(err) {
		log.Fatalf("given file %v not exists\n", *filename)
	}

	s := &Server{filename: *filename}
	http.HandleFunc("/", s.Hello)
	hostport := fmt.Sprintf(":%v", *port)
	http.ListenAndServe(hostport, nil)
}
