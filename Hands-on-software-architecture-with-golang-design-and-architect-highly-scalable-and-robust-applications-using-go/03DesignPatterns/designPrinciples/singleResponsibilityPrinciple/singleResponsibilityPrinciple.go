package main

import (
	"fmt"
	"io/ioutil"
)

// ทำหน้าที่รับผิดชอบในการอ่านไฟล์
func readFile(fileName string) (string, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// ทำหน้าที่รับผิดชอบในการแสดงผลข้อมูล
func printData(data string) {
	fmt.Println(data)
}

func main() {
	data, err := readFile("example.txt")
	if err != nil {
		fmt.Println("Error reading file : ", err)
		return
	}
	printData(data)
}
