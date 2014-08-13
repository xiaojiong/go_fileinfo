package fileinfo

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Files struct {
	FileList []*FileInfo
	Count    int
	Size     int64
	Line     int64
}

func (files *Files) ToJson() string {

	r, err := json.Marshal(files)

	if err != nil {
		fmt.Println(err)
		return "{}"
	}
	return string(r)
}

func (files *Files) add(file FileInfo) {
	files.Count++
	files.Line += int64(file.LineNum)
	files.Size += file.Size

	files.FileList = append(files.FileList, &file)
}

type FileInfo struct {
	FileName string
	Path     string
	LineNum  int
	Size     int64
}

func (file *FileInfo) SetLineNum(num int) {
	file.LineNum = num
}
func (file *FileInfo) SetSize(size int64) {
	file.Size = size
}

func InitFileInfo(path string) *FileInfo {
	p := new(FileInfo)
	p.Path = path
	fileInfo, _ := os.Stat(path)
	p.FileName = fileInfo.Name()
	return p
}

func PathFiles(path string) []FileInfo {
	fileInfo := make([]FileInfo, 0)

	//遍历文件夹并把文件或文件夹名称加入相应的slice
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {

		if info.IsDir() == false {
			f := InitFileInfo(path)
			f.SetSize(info.Size())
			fileInfo = append(fileInfo, *f)
		}
		return err
	})
	if err != nil {
		panic(err)
	}
	return fileInfo
}

func Parse(path string) Files {
	files := Files{Count: 0, Size: 0, Line: 0}

	fileNames := PathFiles(path)

	for i := 0; i < len(fileNames); i++ {
		fileNames[i].SetLineNum(FileLine(&fileNames[i].Path))

		files.add(fileNames[i])
	}
	return files
}
