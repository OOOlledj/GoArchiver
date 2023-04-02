package main

import (
	"flag"
	"fmt"

	//"os"

	//"fmt"
	"strings"
)

func main() {

	archMode := flag.String("m", "", "Compression algorithm (targz, zip)")
	archOut := flag.String("o", "", "Output archive name, '.tar.gz' and '.zip' in the name are not required")
	//helpFlag := flag.Bool("h", true, "Display help and exit")
	flag.Usage = func() {
		fmt.Println("Usage:")
		fmt.Println("  gfarch -m=targz -o=test <file1> <file2> <folder1>...")
		fmt.Println("  gfarch -m=zip -o=test <file1> <file2> <folder1>...")
		fmt.Println("Options:")
		flag.PrintDefaults()
	}
	flag.Parse()

	if *archMode == "" || (*archMode != "targz" && *archMode != "zip") {
		panic("Please, provide -m option with proper value!")
	} else {
		*archOut = strings.TrimRight(*archOut, ".tar.gz")
		*archOut = strings.TrimRight(*archOut, ".zip")
	}

	if *archOut == "" {
		panic("Please, provide -o option with output file name!")
	}

	var files []FileInfoPath
	for _, file := range flag.Args() {
		files = append(files, ListFilesToWrite(file)...)
	}
	PrintFilesToWrite(files)

	if *archMode == "targz" {
		TarGzFile(*archOut, &files)
	} else if *archMode == "zip" {
		ZipFile(*archOut, &files)
	}
}

// func runAllTests() {
// 	test("./", "dot-slash")
// 	test(".", "dot")
// 	test("./test", "dot-slash-dir")
// 	test("test", "dir")
// 	test("/home/ooolledj/Documents/go-projects/GoFolderArch/test/", "abs-path")
// 	test("go.mod", "single-file")
// 	test("./go.mod", "single-dot-slash-file")
// 	test("test/tester/lic.docx", "file-in-dir")
// 	test("/home/ooolledj/Documents/go-projects/GoFolderArch/test/tester/lic.docx", "file-in-dir-abs")
// }

// func test(testpath, name string) {
// 	ListFilesToWrite(testpath)
// 	files := ListFilesToWrite(testpath)
// 	TarGzFile(name, &files)
// 	ZipFile(name, &files)
// }
