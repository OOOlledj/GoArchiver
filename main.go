package main

import (
	"fmt"
)

var path string = "test"

func main() {
	totalBytes, files := ListDir(path)
	totalSize, totalUnit := FormatFileSize(totalBytes)

	fmt.Println("Listing files to write:")
	for _, file := range files {
		size, unit := FormatFileSize(file.Info.Size())
		fmt.Printf(FileMask, file.Path+"/"+file.Info.Name(), size, unit)
	}
	Printfln("Total: %4.2f %v", totalSize, totalUnit)

	TarGzFile(path, &files)
	ZipFile(path, &files)
}
