package magic

import (
	"os"
	"path/filepath"

	magic_struct "./authentication/layer2/layer3/typedef"
)

func Walk(root string) (route []magic_struct.FileInfo) {

	var files []magic_struct.FileInfo
	var temp magic_struct.FileInfo
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		temp.Path = path
		temp.Name = info.Name()
		temp.Size = info.Size()
		temp.ModTime = info.ModTime()
		temp.IsDir = info.IsDir()
		files = append(files, temp)
		return nil
	})
	if err != nil {
		panic(err)
	}

	return files

}
