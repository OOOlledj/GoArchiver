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

func TarGzFile(path, archName string, FilesToTar *[]FileInfoPath) {
	/*
		add files from slice to .tar archive
	*/
	if _, err := os.Stat(archName + ".tar.gz"); err == nil {
		log.Fatalln("Archieve file already exists!")
	}
	f, err := os.Create(archName + ".tar.gz")
	if err != nil {
		log.Fatalln(err)
	}
	gw := gzip.NewWriter(f) //gzip writer
	tw := tar.NewWriter(gw) //tar writer
	fmt.Println(*FilesToTar)
	for _, file := range *FilesToTar {
		// header (hdr) is required for writing to .tar
		openPath, savePath := file.GetPaths(path)
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
		fmt.Printf("Writing file to TAR\n: %v", savePath)
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
}

func ZipFile(path, archName string, FilesToZip *[]FileInfoPath) {
	if _, err := os.Stat(archName + ".zip"); err == nil {
		log.Fatalln("ZIP file already exists!")
	}
	f, err := os.Create(archName + ".zip")
	if err != nil {
		log.Fatalln(err)
	}
	zw := zip.NewWriter(f)
	for _, file := range *FilesToZip {
		openPath, savePath := file.GetPaths(path)
		f, err := os.Open(openPath)
		if err != nil {
			log.Fatalln(err)
		}
		defer f.Close()
		w, err := zw.Create(savePath)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("Writing file to ZIP\n: %v", savePath)
		if _, err := io.Copy(w, f); err != nil {
			log.Fatalln(err)
		}
	}
	if err := zw.Close(); err != nil {
		log.Fatalln(err)
	}
}
