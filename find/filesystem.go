package find

import (
	"fmt"
	"io/ioutil"
)

type fileSystemResolver struct {
	path string
}

func (f fileSystemResolver) Resolve() (string, error) {
	res, err := ioutil.ReadFile(f.path)
	if err != nil {
		fmt.Printf("err: %+v\n", err)
		return "", err
	}
	return string(res), nil
}
