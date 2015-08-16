package main

import (
	"log"
	"net/http"
	"os"
	"path"

	"github.com/docopt/docopt-go"
)

const Version = "0.1.0"
const Usage = `
  Usage:
    alfred [--directory=<dir>] [--bind=<addr>]
    alfred -h | --help
    alfred --version
  Options:
    --directory=<dir>   directory to serve [default: /]
    --bind=<addr>       bind address  [default: 0.0.0.0:3000]
    -h, --help          output help information
    -v, --version       output version
`

// Serves a single directory to handle all incoming requests.
type Server struct {
	dir string
}

// ServeHTTP implementation.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	filePath := s.dir + path.Clean(r.URL.Path)
	if _, err := os.Stat(filePath); err == nil {
		http.ServeFile(w, r, filePath)
		return
	}

	http.ServeFile(w, r, s.dir)
}

func main() {
	args, err := docopt.Parse(Usage, nil, true, Version, false)
	check(err)

	addr := args["--bind"].(string)
	dir := args["--directory"].(string)

	log.Println("binding to", addr)
	log.Println("serving", dir)

	check(http.ListenAndServe(addr, &Server{dir}))
}

func check(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
