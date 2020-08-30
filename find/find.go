package find

import (
	"log"
	"os"
)

//Option is option for Finder
type Option struct {
	Src string
}

// Finder find by path from
// 1) filesystem
// 2) gitbucket
type Finder struct {
	opt      Option
	resolver Resolver
}

// Resolver resolve path to string
type Resolver interface {
	Resolve() (string, error)
}

func resolve(r Resolver) (string, error) {
	return r.Resolve()
}

// Find return resolved body
func Find(opts Option, path string) string {
	finder := Finder{opt: opts}
	v, e := finder.find(path)
	if e != nil {
		log.Printf("err: %v\n", e)
		os.Exit(1)
	}
	return v
}

func (f *Finder) find(path string) (string, error) {
	var r Resolver
	switch f.opt.Src {
	case "gitbucket":
		r = gitbucketResolver{path: path}
	case "filesystem":
		r = fileSystemResolver{path: path}
	}
	return resolve(r)
}

func (f *Finder) readOpt() {
	f.opt.Src = "gitbucket"
}
