package find

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
	Resolve() string
}

func resolve(r Resolver) string {
	return r.Resolve()
}

// Find return resolved body
func Find(opts Option, path string) string {
	finder := Finder{opt: opts}
	return finder.find(path)
}

func (f *Finder) find(path string) string {
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
