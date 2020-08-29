package find

import (
	"io/ioutil"
	"log"
	"net/http"
)

type gitbucketResolver struct {
	path string
}

func (g gitbucketResolver) Resolve() string {
	res, err := http.Get(g.path)
	if err != nil {
		log.Fatalf("http request failed: %+v\n", err)
	}
	if res.StatusCode != 200 {
		log.Fatalln("File not found. Please verify file path or API.")
	}
	byteArray, _ := ioutil.ReadAll(res.Body)
	return string(byteArray)
}
