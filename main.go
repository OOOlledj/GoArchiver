package main

// import(
//     "fmt"
// )

func main() {
	totalBytes := ListDir("./test/")
	//print directory total
	totalSize, totalUnit := HandleSize(totalBytes)
	Printfln("Total: %4.2f %v", totalSize, totalUnit)
}
