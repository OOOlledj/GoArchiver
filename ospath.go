package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	//"strings"
)

var FileMask = "%-48v %4.2f %v\n"
var units = map[int]string{0: "bytes", 1: "KB", 2: "MB", 3: "GB"}

type FileInfoPath struct {
	Info fs.FileInfo
	Path string
}

func (file FileInfoPath) GetRelativePath() (path string) {
	/*
		get relative path of file or directory
	*/
	// handle cases "file.txt"
	if file.Path != file.Info.Name() {
		path = file.Path + "/" + file.Info.Name()
		// handle cases "/home/ooolledj/file.txt"
	} else {
		path = file.Path
	}
	return path
}

func OpenFile(path string) (file *os.File) {
	// Open file and return the object
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func ListDir(path string) (totalBytes int64, filesInDirectory []FileInfoPath) {
	/*
		Print files in the directory
		Return total size of the directory content in bytes and slice of the files
	*/
	pathObject := OpenFile(path)
	pathObjectStat, err := pathObject.Stat()
	if err != nil {
		log.Fatalln(err)
	}
	if pathObjectStat.IsDir() {
		files, err := pathObject.Readdir(0)
		if err != nil {
			log.Fatalln(err)
		} else {
			// inspect all files in the directory
			for _, entry := range files {
				// if entry is a directory, recursively start to inspect it
				if entry.IsDir() {
					// Buffer is used to store file references in the entry, if it is a directory
					totalBytesBuffer, filesInDirectoryBuffer := ListDir(path + "/" + entry.Name())
					filesInDirectory = append(filesInDirectory, filesInDirectoryBuffer...)
					totalBytes += totalBytesBuffer
					// if enry is file, simply add it to list and add file size to total
				} else {
					filesInDirectory = append(filesInDirectory, FileInfoPath{entry, path})
					totalBytes += entry.Size()
				}
			}
		}
	} else {
		filesInDirectory = append(filesInDirectory, FileInfoPath{pathObjectStat, path})
		totalBytes = pathObjectStat.Size()
	}
	pathObject.Close()
	return
}

func FormatFileSize(bytes int64) (fBytes float64, unit string) {
	/*
		Return file representation representation as calculated size and units
		ex. 2048 bytes -> 2.0 KB, etc
	*/
	fBytes = float64(bytes)
	count := 0
	for range []int{0, 1, 2, 3} { // see map "units"
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
