package find

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

type gitbucketResolver struct {
	path string
}

func (g gitbucketResolver) Resolve() (string, error) {
	res, err := http.Get(g.path)
	if err != nil {
		log.Printf("http request failed: %+v\n", err)
		return "", err
	}
	if res.StatusCode != 200 {
		log.Printf("File not found. Please verify file path or API.")
		return "", errors.New("Invalid Path")
	}
	byteArray, _ := ioutil.ReadAll(res.Body)
	return string(byteArray), nil
}
