package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Starting kokken")

	w := New()
	w.Run()
}

type Watcher struct {
	instances map[string]bool
}

func New() *Watcher {
	return &Watcher{
		instances: make(map[string]bool),
	}
}

func (w *Watcher) Run() {
	for {
		fmt.Println("Checking...")
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
