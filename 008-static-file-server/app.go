package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
)

/**
port: 运行的端口
dirName:静态资源的文件路径
*/
func runFileServer(port, dirName string) {
	log.Printf("will start server  on %s and Directory %s", port, dirName)
	http.ListenAndServe(":"+port, http.FileServer(http.Dir(dirName)))
}

func main() {
	//args:=os.Args

	if len(os.Args) == 3 {
		port := os.Args[1]
		dirName := os.Args[2]
		runFileServer(port, dirName)
	} else {
		currentDir, _ := os.Getwd()
		port := 80
		dirName := currentDir
		runFileServer(strconv.Itoa(port), dirName)
	}
	//fmt.Println(currentDir)
	//for idx,arg := range args {
	//	fmt.Println(idx,arg)
	//}
}
