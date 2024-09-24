package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	grytaEndpoint string
	bindAddr      string
	fakeResponse  string
)

func init() {
	flag.StringVar(&grytaEndpoint, "gryta-endpoint", envOrDefault("GRYTA_ENDPOINT", "troll-gryta"), "endpoint where troll-gryta is running")
	flag.StringVar(&fakeResponse, "fake-response", os.Getenv("FAKE_RESPONSE"), "run in dev mode")
	flag.StringVar(&bindAddr, "bind-address", ":8080", "ip:port where http requests are served")

	flag.Parse()
}

func main() {
	w := New(grytaEndpoint)

	if fakeResponse == "" {
		go w.Run()
	} else {
		for _, v := range strings.Split(fakeResponse, ",") {
			w.instances[v] = true
		}
	}

	http.HandleFunc("/instances", func(wr http.ResponseWriter, _ *http.Request) {
		wr.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(wr).Encode(w.Instances())
	})

	if err := http.ListenAndServe(bindAddr, nil); err != nil {
		log.Fatal(err)
	}
}

type Watcher struct {
	grytaEndpoint string
	instances     map[string]bool
}

func New(grytaEndpoint string) *Watcher {
	return &Watcher{
		instances:     make(map[string]bool),
		grytaEndpoint: grytaEndpoint,
	}
}

func (w *Watcher) Run() {
	for {
		fmt.Println("Checking grytaendpoint", w.grytaEndpoint)

		resp, err := http.Get(fmt.Sprintf("http://%s/", w.grytaEndpoint))
		if err != nil {
			fmt.Println("Error getting instances", err)
			continue
		}

		b, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response", err)
			continue
		}

		resp.Body.Close()

		found := string(b)
		fmt.Println("Found instances", found)

		w.instances[found] = true
		time.Sleep(1 * time.Second)
	}
}

func (w *Watcher) Instances() []string {
	ret := []string{}
	for k := range w.instances {
		ret = append(ret, k)
	}
	return ret
}

func envOrDefault(key, def string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return def
}
