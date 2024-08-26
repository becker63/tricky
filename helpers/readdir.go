package helpers

import (
	"fmt"
	"os"
	"sync"

	p "path"

	cwalk "github.com/becker63/tricky/helpers/cwalk"
)

var tempMap sync.Map
var Gdirpath = ""

func RecusiveReadDir(dirpath string) {
	err := cwalk.Walk(dirpath, walkFunc)
	Gdirpath = dirpath
	fmt.Println(Gdirpath)

	if err != nil {
		fmt.Printf("Error : %s\n", err.Error())
	}

	// TODO: not this
	var RealMap = make(map[string][]byte)
	tempMap.Range(func(k, v any) bool {
		RealMap[k.(string)] = v.([]byte)
		return true
	})
	fmt.Println(RealMap["world_nether/DIM-1/region/r.-1.-1.mca"])
}

func walkFunc(path string, info os.FileInfo, err error) error {

	var abspath = p.Join(Gdirpath, path)
	filecontent, err := os.ReadFile(abspath)
	if err != nil {
		// fmt.Println(err)
	}
	tempMap.Store(path, filecontent)
	return nil
}
