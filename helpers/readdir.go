package helpers

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	cwalk "github.com/becker63/tricky/helpers/cwalk"
	"github.com/puzpuzpuz/xsync/v3"
)

type FileContent struct {
	isDir bool
	// initialized to zero
	content []byte
}

var tempMap = xsync.NewMapOf[string, FileContent]()

func RecusiveReadDir(dirpath string) *xsync.MapOf[string, FileContent] {
	err := cwalk.WalkWithSymlinks(dirpath, walkFunc)

	if err != nil {
		fmt.Printf("Error : %s\n", err.Error())
	}

	return tempMap
}

func walkFunc(path string, info os.FileInfo, err error, root string) error {
	if err != nil {
		return err
	}
	abs := filepath.Join(root, path)
	if !info.IsDir() {
		// open the file
		file, err := os.Open(abs)
		if err != nil {
			return err
		}
		defer file.Close()

		// Read the entire file into a byte slice
		data, err := io.ReadAll(file)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return err
		}
		tempMap.Store(path, FileContent{isDir: false, content: data})
		return nil
	}
	// if is a dir, we can check by going key==value
	tempMap.Store(path, FileContent{isDir: true})
	return nil
}
