package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
)

func main() {
	args := os.Args[1:]

	if len(args) < 2 {
		log.Fatal("missing arguments")
	}
	/* #nosec */
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	server := &http.Server{Addr: ":8090", Handler: nil, ReadHeaderTimeout: time.Second * 5}
	http.HandleFunc("/kill", func(w http.ResponseWriter, r *http.Request) {
		err := cmd.Process.Kill()
		if err != nil {
			fmt.Println("error killing process", err)
		}
		_, _ = w.Write([]byte(string("OK")))
		w.WriteHeader(http.StatusAccepted)
		err = server.Shutdown(r.Context())
		if err != nil {
			fmt.Println("error shuting down", err)
		}
	})
	err := server.ListenAndServe()
	_ = cmd.Process.Kill()
	if err != nil {
		log.Fatal(err)
	}
}
