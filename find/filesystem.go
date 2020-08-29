package find

import (
	"fmt"
	"io/ioutil"
)

type fileSystemResolver struct {
	path string
}

func (f fileSystemResolver) Resolve() string {
	res, err := ioutil.ReadFile(f.path)
	if err != nil {
		fmt.Printf("err: %+v\n", err)
	}
	return string(res)
}
