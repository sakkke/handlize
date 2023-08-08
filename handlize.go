package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

var queue = make(chan bool, 100)

func main() {
	cmd := os.Args[1]

	switch cmd {
	case "run":
		run()

	case "serve":
		go worker()
		serve()
	}
}

func run() {
	http.Get("http://127.0.0.1:9000/api/run")
}

func serveHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/api/run":
		queue <- true
	}
}

func worker() {
	cmdStr := os.Args[2]

	for range queue {
		cmd := exec.Command("sh", "-c", cmdStr)
		out, _ := cmd.CombinedOutput()
		fmt.Print(string(out))
	}
}

func serve() {
	http.ListenAndServe(":9000", http.HandlerFunc(serveHandler))
}
