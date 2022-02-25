package HttpReq

import (
	"fmt"
	http "net/http"
	sc "sync"
)

func GetReq(s string) (string, error) {
	resp, err := http.Get(s)
	if err != nil {
		return "", nil
	}
	return resp.Status, nil
}

func GoGetReq(wg *sc.WaitGroup) {
	defer wg.Done()
	http.Get("http://google.com")
	fmt.Println("Requesting...")
}
