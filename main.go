package main

// import (
// 	"fmt"
// )

func main() {
	//runAllTests()
}

func runAllTests() {
	test("./", "dot-slash")
	test(".", "dot")
	test("./test", "dot-slash-dir")
	test("test", "dir")
	test("/home/ooolledj/Documents/go-projects/GoFolderArch/test/", "abs-path")
	test("go.mod", "single-file")
	test("./go.mod", "single-dot-slash-file")
	test("test/tester/lic.docx", "file-in-dir")
	test("/home/ooolledj/Documents/go-projects/GoFolderArch/test/tester/lic.docx", "file-in-dir-abs")
}

func test(testpath, name string) {
	ListFilesToWrite(testpath)
	files := ListFilesToWrite(testpath)
	TarGzFile(testpath, name, &files)
	ZipFile(testpath, name, &files)
}
