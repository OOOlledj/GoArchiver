package main

import (
	"fmt"
	"os"
)

var FileMask = "%-48v %4.2f %v\n"
var units = map[int]string{0: "bytes", 1: "KB", 2: "MB", 3: "GB"}

func OpenFile(path string) (file *os.File) {
	// Open file and return the object
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func ListDir(path string) (totalBytes int64) {
	/* Print files in the directory
	   Retun total size of the directory content in bytes
	*/
	directory := OpenFile(path)
	files, err := directory.Readdir(0)
	if err != nil {
		fmt.Println(err)
	} else {
		// inspect all files in the directory
		for _, entry := range files {
			size, unit := HandleSize(entry.Size())
			if entry.IsDir() { // if object is directory, recursively start to inspect it
				totalBytes += ListDir(path + "/" + entry.Name())
			} else {
				fmt.Printf(FileMask, path+"/"+entry.Name(), size, unit)
				totalBytes += entry.Size()
			}
		}
	}
	directory.Close()
	return
}

func HandleSize(bytes int64) (fBytes float64, unit string) {
	/*
	   Format file size and return it's representation as calculated size in the units
	   ex. 2048 bytes -> 2.0 KB, etc
	*/
	fBytes = float64(bytes)
	count := 0
	for range []int{0, 1, 2, 3} { // see variable units"
		if fBytes/1024 < 1 {
			break
		} else {
			fBytes /= 1024
		}
		count++
	}
	unit = units[count]
	return
}
