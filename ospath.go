package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path"
	"strings"
)

var FileMask = "%-48v %4.2f %v\n"
var units = map[int]string{0: "bytes", 1: "KB", 2: "MB", 3: "GB"}

type FileInfoPath struct {
	Info         fs.FileInfo
	RelativePath string
}

func OpenFile(path string) (file *os.File) {
	// Open file and return the object
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func ListFilesToWrite(wwwpath string) (files []FileInfoPath) {
	wwwpath = path.Clean(wwwpath)
	files = ListDir(wwwpath)
	totalBytes := 0
	for idx, file := range files {
		totalBytes += int(file.Info.Size())
		if wwwpath == file.RelativePath {
			files[idx].RelativePath = ""
		} else {
			files[idx].RelativePath = strings.Replace(file.RelativePath, wwwpath+"/", "", 1)
		}
	}
	totalSize, totalUnit := FormatFileSize(int64(totalBytes))
	fmt.Println("Listing files to write:")
	fmt.Println(strings.Repeat("-", 59))
	for _, file := range files {
		size, unit := FormatFileSize(file.Info.Size())

		_, save := file.GetPaths(wwwpath)
		fmt.Printf(FileMask, save, size, unit)

	}
	fmt.Println(strings.Repeat("-", 59))
	Printfln(FileMask, "Total:", totalSize, totalUnit)
	fmt.Println()
	return
}

func (file FileInfoPath) GetPaths(start string) (openPath, savePath string) {
	/*
		Preparing PATH values to open file and to save file as path
		OPEN PATH - path to the file in the file system, we read it's content and add it to the archive
		SAVE PATH - relative path to the file, which is used to display it inside archive

		example 1:
			we have provided path: /tmp/test1/test2/file1.txt
			we add file1.txt to archive root
		example 2:
			we have provided path /tmp/test/test2
			we add file.txt and any other file in the folder test2 to archive (file2.txt, ...) root
		example 3:
			we have provided path /tmp/test/
			we add folder test2 with it's content to archive. if there are any other folders or files - we add them to the root of the archive
	*/
	// BUILDING SAVE PATH
	if file.RelativePath == "" {
		savePath = file.Info.Name()
	} else {
		savePath = path.Join(file.RelativePath, file.Info.Name())
	}
	// BUILDING OPEN PATH
	if start == "" || start == "." {
		openPath = savePath
	} else if strings.Contains(start, file.Info.Name()) {
		openPath = start
	} else {
		openPath = path.Join(start, savePath)
	}
	return
}

func ListDir(startPath string) (filesInDirectory []FileInfoPath) {
	/*
		Print files in the directory
		Return total size of the directory content in bytes and slice of the files
	*/
	pathObject := OpenFile(startPath)
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
					if entry.Name() != ".git" {
						filesInDirectoryBuffer := ListDir(startPath + "/" + entry.Name())
						filesInDirectory = append(filesInDirectory, filesInDirectoryBuffer...)
					}
				} else {
					filesInDirectory = append(filesInDirectory, FileInfoPath{entry, startPath})
				}
			}
		}
	} else {
		filesInDirectory = append(filesInDirectory, FileInfoPath{pathObjectStat, ""})
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
