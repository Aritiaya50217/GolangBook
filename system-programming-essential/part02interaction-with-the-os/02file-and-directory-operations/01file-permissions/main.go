package main

import (
	"fmt"
	"os"
)

func main() {
	fileInfo, err := os.Stat("example.txt")
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}

	permission := fileInfo.Mode().Perm()
	permissionString := fmt.Sprintf("%o", permission)
	fmt.Printf("Permissions: %s\n", permissionString)
}
