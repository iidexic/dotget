package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type Options struct {
	maindir  string
	datafile string
}

func main() {
	dig()
}

func pathContents(path string) {
	dirpathA := filepath.Dir(path)
	readA, e := os.ReadDir(dirpathA)
	autoerr(e)
	if len(readA) > 0 {
		for _, v := range readA {

			info, e := v.Info()
			autoerr(e)
			typ := v.Type()
			fmt.Printf("name: %s,type: %s, typetype: %s, typestr: %s, mode: %s,\n", info.Name(), typ, typ.Type(), typ.String(), info.Mode())
		}
	}
}

func GetConfig() Options {

}
func autoerr(e error) {
	if e != nil {
		panic(e)
	}
}
