package main

import (
	"fmt"
	"io/ioutil"
)

// 将content存储到 filePath
func saveFileToPath(content, filePath string) {
	ioutil.WriteFile(filePath, []byte(content), 0644)
	fmt.Println("File Saved:", filePath)
}
