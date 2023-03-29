package main

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"io"
	"log"
	"os"
)

func TarGzFile(path string, FilesToTar *[]FileInfoPath) {
	/*
		add files from slice to .tar archive
	*/
	if _, err := os.Stat(path + ".tar.gz"); err == nil {
		log.Fatalln("Archieve file already exists!")
	}
	f, err := os.Create("test.tar.gz")
	if err != nil {
		log.Fatalln(err)
	}
	gw := gzip.NewWriter(f) //gzip writer
	tw := tar.NewWriter(gw) //tar writer
	for _, file := range *FilesToTar {
		// header (hdr) is required for writing to .tar
		hdr := &tar.Header{
			Name: file.GetRelativePath(),
			Size: file.Info.Size(),
			Mode: int64(file.Info.Mode()),
		}
		// write the header
		if err := tw.WriteHeader(hdr); err != nil {
			log.Fatalln(err)
		}
		// write the file
		fileBytes, _ := os.ReadFile(file.GetRelativePath())
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

func ZipFile(path string, FilesToZip *[]FileInfoPath) {
	if _, err := os.Stat(path + ".zip"); err == nil {
		log.Fatalln("ZIP file already exists!")
	}
	f, err := os.Create("test.zip")
	if err != nil {
		log.Fatalln(err)
	}
	zw := zip.NewWriter(f)
	for _, file := range *FilesToZip {
		fileFullName := file.GetRelativePath()
		f, err := os.Open(fileFullName)
		if err != nil {
			log.Fatalln(err)
		}
		defer f.Close()
		w, err := zw.Create(fileFullName)
		if err != nil {
			log.Fatalln(err)
		}
		if _, err := io.Copy(w, f); err != nil {
			log.Fatalln(err)
		}
	}
	if err := zw.Close(); err != nil {
		log.Fatalln(err)
	}
}
