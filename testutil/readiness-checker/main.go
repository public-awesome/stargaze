package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func waitFor(timeout time.Duration, blocks int, url string, ch chan<- error) {
	cli := http.Client{
		Timeout: time.Second * 1,
	}
	end := time.After(timeout)
	for {
		select {
		case <-time.After(time.Second * 5):
			resp, err := cli.Get(fmt.Sprintf("%s/status", url))
			if err != nil {
				log.Printf("%s\n", err.Error())
				continue
			}
			if resp.StatusCode == http.StatusOK {
				status := &ChainStatus{}
				err = json.NewDecoder(resp.Body).Decode(&status)
				_ = resp.Body.Close()
				if err != nil {
					log.Printf("%s: error decoding response %s\n", url, err.Error())
					continue
				}

				if block, err := strconv.Atoi(status.Result.SyncInfo.LatestBlockHeight); err == nil && block > blocks {
					log.Printf("%s: chain is ready \n", url)
					ch <- nil
					return
				}
				log.Printf("%s latest block: %s \n", url, status.Result.SyncInfo.LatestBlockHeight)
			}
		case <-end:
			ch <- fmt.Errorf("timed out waiting for %s", url)
			return
		}
	}
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}

func main() {
	checkFile := os.Getenv("PLUGIN_CHECK_FILE")
	if checkFile != "" && !exists(checkFile) {
		log.Fatal("check file doesn't exists")
	}
	chains := strings.Split(os.Getenv("PLUGIN_CHAIN_LIST"), ",")

	if len(chains) == 0 || len(os.Getenv("PLUGIN_CHAIN_LIST")) == 0 {
		log.Fatal("must provide at least one chain")
	}

	timeout, err := strconv.Atoi(os.Getenv("PLUGIN_TIMEOUT"))
	if err != nil {
		log.Fatal("must provide a valid timeout")
	}

	numBlocks, err := strconv.Atoi(os.Getenv("PLUGIN_BLOCKS"))
	if err != nil || numBlocks < 5 {
		numBlocks = 5
	}

	ch := make(chan error, len(chains))

	for _, c := range chains {
		log.Printf("wait for %s %d", c, timeout)
		go waitFor(time.Duration(timeout)*time.Second, numBlocks, c, ch)
	}

	jobs := len(chains)
	for err := range ch {
		if err != nil {
			log.Fatal(err)
		}
		jobs--
		if jobs == 0 {
			return
		}
	}
}
