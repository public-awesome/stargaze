package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	args := os.Args[1:]

	if len(args) < 2 {
		log.Fatal("missing arguments")
	}

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	server := &http.Server{Addr: ":8090", Handler: nil}
	http.HandleFunc("/kill", func(w http.ResponseWriter, r *http.Request) {
		err := cmd.Process.Kill()
		if err != nil {
			fmt.Println("error killing process", err)
		}
		server.Shutdown(r.Context())
	})
	err := server.ListenAndServe()
	cmd.Process.Kill()
	if err != nil {
		log.Fatal(err)
	}
}
