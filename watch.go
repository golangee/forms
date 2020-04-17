package wtk

import (
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

type Watch struct {
	url       string
	interval  time.Duration
	ticker    *time.Ticker
	done      chan bool
	mutex     sync.Mutex
	expected  string
	listener  []func(found string, expected string)
	lastCheck time.Time
}

func NewWatch(expected string) *Watch {
	return &Watch{url: "/version", interval: 5 * time.Second, expected: expected}
}

func (a *Watch) SetInterval(d time.Duration) {
	a.interval = d
}

func (a *Watch) Start() {
	if a.ticker != nil {
		a.Stop()
	}
	a.ticker = time.NewTicker(a.interval)
	a.done = make(chan bool)

	myTicker := a.ticker // create a hard ref for the spawned goroutine
	myDone := a.done     // create a hard ref for the spawned goroutine

	go func() {
		for {
			select {
			case <-myDone:
				return
			case t := <-myTicker.C:
				go a.check(t) // spawn another goroutine to avoid deadlock if calling stop from listener
			}
		}
	}()
}

func (a *Watch) Check() {
	go a.check(time.Now())
}

func (a *Watch) check(t time.Time) {
	if t.Sub(a.lastCheck) < 5*time.Millisecond {
		return // do not check faster than that
	}
	a.lastCheck = t
	client := http.Client{Timeout: 10 * time.Second}

	res, err := client.Get(a.url)
	if err != nil {
		log.Printf("cannot get version from %s:%v", a.url, err)
		return
	}

	defer res.Body.Close()

	msg, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("cannot read version from %s:%v", a.url, err)
		return
	}

	if string(msg) != a.expected {
		a.notifyChanged(string(msg), a.expected)
	}

}

func (a *Watch) notifyChanged(found string, expected string) {
	for _, listener := range a.listener {
		if listener != nil {
			listener(found, expected)
		}
	}
}

func (a *Watch) AddListener(f func(found string, expected string)) Resource {
	a.listener = append(a.listener, f)
	idx := len(a.listener) - 1
	return NewListener(f, func() {
		a.listener[idx] = nil
	})
}

func (a *Watch) Stop() {
	if a.ticker != nil {
		a.ticker.Stop()
		a.ticker = nil
	}
	a.done <- true
	close(a.done)
	a.done = nil
}
