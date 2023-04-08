package main

import (
	// "archive/tar"
	"archive/zip"
	"fmt"
	"io"

	// "compress/gzip"
	// "fmt"
	// "io"
	"log"
	"os"
	"path"
)

func UnZip(archivePath, outPath string) {
	if _, err := os.Stat(archivePath); err != nil {
		log.Fatalln("There is no file with provided name!")
	}
	if _, err := os.Stat(outPath); os.IsNotExist(err) {
		fmt.Println("Creating output directory...")
		os.MkdirAll(outPath, 0764)
	}

	zr, err := zip.OpenReader(archivePath) // zip reader
	if err != nil {
		fmt.Println(os.Getwd())
		log.Fatalln("Cannot open archive:", err)
	}
	fmt.Println("Writing files from ZIP:")
	for _, file := range zr.File {
		filePath := path.Join(outPath, file.Name)

		if file.FileInfo().IsDir() {
			os.MkdirAll(filePath, 0764)
			continue
		}
		dstFile, _ := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		fileInArchive, _ := file.Open()
		fmt.Println(filePath)
		io.Copy(dstFile, fileInArchive)
		if _, err := io.Copy(dstFile, fileInArchive); err != nil {
			log.Fatalln("Error while writing file to destination!")
		}
		dstFile.Close()
		fileInArchive.Close()
	}
}
