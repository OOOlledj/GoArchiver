package main

import (
	"archive/tar"
	//"fmt"
	"io"
	//"compress/gzip"
	"archive/zip"
	"log"
	"os"
)

func TarFile(path string, FilesToTar *[]FileInfoPath) {
	if _, err := os.Stat(path + ".tar"); err == nil {
		log.Fatalln("Archieve file already exists!")
	}
	f, err := os.Create("test.tar")
	if err != nil {
		log.Fatalln(err)
	}
	tw := tar.NewWriter(f)
	for _, file := range *FilesToTar {
		hdr := &tar.Header{
			Name: file.GetRelativePath(),
			Size: file.Info.Size(),
			Mode: int64(file.Info.Mode()),
		}
		if err := tw.WriteHeader(hdr); err != nil {
			log.Fatalln(err)
		}
		fileBytes, _ := os.ReadFile(file.GetRelativePath())
		if _, err := tw.Write(fileBytes); err != nil {
			log.Fatalln(err)
		}
	}
	if err := tw.Close(); err != nil {
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
		fileFullName := file.Path + "/" + file.Info.Name()
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

func GzFile() {}
