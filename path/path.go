package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func main() {
	//最终的目的是获取当前路径

	getPath()
	getPathT()
	getPathTr()
}

func getPath() {
	//获取到执行目录
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("获取当前目录失败:%v", err)
		return
	}
	fmt.Printf("当前路径:%v", currentDir)
}
func getPathT() {
	//获取到缓存程序的路径
	executable, err := os.Executable()
	if err != nil {
		log.Fatalf("获取当前目录失败:%v", err)
		return
	}
	fmt.Printf("当前路径:%v\n", executable)
}
func getPathTr() {
	//获取到执行go程序的路径
	pc, executable, line, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("获取当前目录失败")
		return
	}
	currentDir := filepath.Dir(executable)
	fmt.Printf("当前路径:%v\n", executable)
	fmt.Printf("当前路径:%v\n", currentDir)
	fmt.Printf("当前计数器:%v\n", pc)
	fmt.Printf("当前行号:%v\n", line)
}
