package main

import (
	"flag"
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
	flag.Parse()
	if _, err := os.Stat(*filename); os.IsNotExist(err) {
		log.Fatalf("given file not exists\n")
	}

	s := &Server{filename: *filename}
	http.HandleFunc("/", s.Hello)
	http.ListenAndServe(":8000", nil)
}
