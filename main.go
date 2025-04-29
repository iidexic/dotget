package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type path struct {
	isdir    bool
	path     string
	contents map[string]path
}

func main() {
	dig()
}

func pathContents() {
	dirpathA := path{path: filepath.Dir(".\\pathA")}
	readA, e := os.ReadDir(dirpathA.path)
	if e != nil {
		panic(e)
	}
	if len(readA) > 0 {
		for _, v := range readA {

			info, e := v.Info()
			autoerr(e)
			typ := v.Type()
			fmt.Printf("name: %s,type: %s, typetype: %s, typestr: %s, mode: %s,\n", info.Name(), typ, typ.Type(), typ.String(), info.Mode())
		}
	}
}

func autoerr(e error) {
	if e != nil {
		panic(e)
	}
}
