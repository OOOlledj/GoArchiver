package main

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
)

func TarGzFile(archName string, FilesToTar *[]FileInfoPath) {
	/*
		add files from slice to .tar.gz archive
	*/
	archFullName := archName + ".tar.gz"
	if _, err := os.Stat(archFullName); err == nil {
		log.Fatalln("Archieve file already exists!")
	}
	f, err := os.Create(archFullName)
	if err != nil {
		log.Fatalln(err)
	}
	gw := gzip.NewWriter(f) //gzip writer
	tw := tar.NewWriter(gw) //tar writer
	// fmt.Println(*FilesToTar)
	fmt.Printf("Writing files to %v...\n", archFullName)
	for _, file := range *FilesToTar {
		// header (hdr) is required for writing to .tar
		openPath, savePath := file.GetPaths(file.StartPath)
		hdr := &tar.Header{
			Name: savePath,
			Size: file.Info.Size(),
			Mode: int64(file.Info.Mode()),
		}
		// write the header
		if err := tw.WriteHeader(hdr); err != nil {
			log.Fatalln(err)
		}
		fileBytes, _ := os.ReadFile(openPath)
		// fmt.Println("-- ", savePath)
		if _, err := tw.Write(fileBytes); err != nil {
			log.Fatalln(err)
		}
	}
	if err := tw.Close(); err != nil { // close tar writer
		log.Fatalln(err)
	}
	if err := gw.Close(); err != nil { // close gzip writer
		log.Fatalln(err)
	}
	fmt.Printf("Files successfully written to %v\n", archFullName)
}

func ZipFile(archName string, FilesToZip *[]FileInfoPath) {
	/*
		add files from slice to .zip archive
	*/
	archFullName := archName + ".zip"
	if _, err := os.Stat(archFullName); err == nil {
		log.Fatalln("ZIP file already exists!")
	}
	f, err := os.Create(archFullName)
	if err != nil {
		log.Fatalln(err)
	}
	zw := zip.NewWriter(f)
	fmt.Println("Writing files to ZIP...")
	for _, file := range *FilesToZip {
		openPath, savePath := file.GetPaths(file.StartPath)
		f, err := os.Open(openPath)
		if err != nil {
			log.Fatalln(err)
		}
		defer f.Close()
		w, err := zw.Create(savePath)
		if err != nil {
			log.Fatalln(err)
		}
		// fmt.Println("-- ", savePath)
		if _, err := io.Copy(w, f); err != nil {
			log.Fatalln(err)
		}
	}
	if err := zw.Close(); err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Files successfully written to %v\n", archFullName)
}
