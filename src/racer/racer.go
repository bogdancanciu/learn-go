package racer

import(
	"fmt"
	"time"
	"net/http"
)

const Timeout = 10 * time.Second
var TimeoutErr = fmt.Errorf("Request timeout")

func Race(a, b string) (winner string, error error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(Timeout):
		return "", TimeoutErr
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}