package libchange_name

import (
	"fmt"
	"os"
	"path/filepath"
)

type PathFinder struct {
}

func NewPathFinder() *PathFinder {
	return &PathFinder{}
}

func (p *PathFinder) PathFile(path string) []string {
	files := make([]string, 0)
	err := filepath.Walk(path, func(path string, finfo os.FileInfo, err error) error {
		if finfo == nil {
			return err
		}
		if finfo.IsDir() {
			if filepath.HasPrefix(finfo.Name(), ".") {
				return filepath.SkipDir
			}

			return nil
		}

		if filepath.HasPrefix(filepath.Base(path), ".") {
			return nil
		}

		files = append(files, path)

		return nil
	})

	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}

	//for _, item := range files {
	//
	//	if strings.Contains(item, "ohl") {
	//		fmt.Println(item)
	//	}
	//
	//}

	return files
}
