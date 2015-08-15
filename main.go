package main

import (
	"log"
	"net/http"

	"github.com/docopt/docopt-go"
)

const Version = "1.0.0"
const Usage = `
  Usage:
    serve <file> [--bind=<addr>]
    serve -h | --help
    serve --version
  Options:
    --bind=<addr>       bind address [default: 0.0.0.0:3000]
    -h, --help          output help information
    -v, --version       output version
`

// Serves a single file to handle all incoming requests.
type Server struct {
	file string
}

// ServeHTTP implementation.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, s.file)
}

func main() {
	args, err := docopt.Parse(Usage, nil, true, Version, false)
	check(err)

	addr := args["--bind"].(string)
	file := args["<file>"].(string)

	log.Println("binding to", addr)
	log.Println("serving", file)

	check(http.ListenAndServe(addr, &Server{file}))
}

func check(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
