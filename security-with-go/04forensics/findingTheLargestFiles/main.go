package main

import (
	"container/list"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type FileNode struct {
	FullPath string
	Info     os.FileInfo
}

func insertSorted(fileList *list.List, fileNode FileNode) {
	if fileList.Len() == 0 {
		// If list is empty , just insert and return
		fileList.PushFront(fileNode)
		return
	}
	/*
		- fileList.Front() → ดึง element ตัวแรกของลิสต์
		- วนลูปไปเรื่อย ๆ จนถึงท้ายลิสต์ (element.Next())
		- element.Value.(FileNode) → cast ค่าในลิสต์กลับมาเป็น FileNode
		- ถ้า ไฟล์ใหม่ (fileNode) มีขนาดเล็กกว่า ไฟล์ใน element ตอนนี้ → InsertBefore(fileNode, element)
	*/

	for element := fileList.Front(); element != nil; element = element.Next() {
		if fileNode.Info.Size() < element.Value.(FileNode).Info.Size() {
			fileList.InsertBefore(fileNode, element)
		}
	}
	fileList.PushBack(fileNode) // ถ้าไม่เจอไฟล์ที่ใหญ่กว่าเลย → PushBack ไปท้าย
}

func getFilesInDirRecursivelyBySize(fileList *list.List, path string) {
	dirFiles, err := ioutil.ReadDir(path)
	if err != nil {
		log.Println("Error reading directory: " + err.Error())
	}

	for _, dirFile := range dirFiles {
		fullpath := filepath.Join(path, dirFile.Name())
		if dirFile.IsDir() {
			getFilesInDirRecursivelyBySize(
				fileList,
				filepath.Join(path, dirFile.Name()),
			)
		} else if dirFile.Mode().IsRegular() {
			insertSorted(
				fileList,
				FileNode{FullPath: fullpath, Info: dirFile},
			)
		}
	}
}

func main() {
	fileList := list.New()
	getFilesInDirRecursivelyBySize(fileList, "/home")

	for element := fileList.Front(); element != nil; element = element.Next() {
		fmt.Println("permission : ", element.Value.(FileNode).Info.Mode())
		fmt.Printf("size : %d\n", element.Value.(FileNode).Info.Size())
		fmt.Printf("full path : %s\n", element.Value.(FileNode).FullPath)
	}
}
