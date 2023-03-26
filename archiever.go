package main

import (
	"archive/tar"
	//"compress/gzip"
	"log"
	"os"
)

func TarFile(path string, FilesToTar *[]FileInfoPath) {
	//fmt.Println(os.Getwd())
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

func GzFile() {}
