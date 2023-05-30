package web

import (
	"errors"
	"net/http"
	"path"
	"strings"
)

type item struct {
	fs     http.FileSystem
	path   string
	prefix string
	index  string
}

type FileSystem struct {
	items []*item
	//items map[string]*item
}

func (f *FileSystem) Put(path string, fs http.FileSystem, prefix string, index string) {
	f.items = append(f.items, &item{fs: fs, path: path, prefix: prefix, index: index})
}

func (f *FileSystem) Open(name string) (file http.File, err error) {
	for _, ff := range f.items {
		//fn := path.Join(ff.prefix, name)
		if ff.path == "" || strings.Index(name, ff.path) == 0 {
			file, err = ff.fs.Open(path.Join(ff.prefix, name))
			if err != nil {
				//尝试默认页
				file, err = ff.fs.Open(path.Join(ff.prefix, ff.index))
			}
			return
		}
	}
	return nil, errors.New("not found")
}
