package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
)

var temp = ""

func main() {
	filePath := ""
	if len(os.Args) < 2 {
		fmt.Printf("请将文件拖到该软件图标上，即可自动进行运算\n勿双击运行\n\n按任意键退出...")
		wait()
		os.Exit(1)
	}
	filePath = os.Args[1]

	fmt.Printf("正在处理文件：%s\n\n请耐心等候...\n\n", filePath)

	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		fmt.Println("Error: ", err)
		wait()
		os.Exit(1)
	}

	md5h := md5.New()
	io.Copy(md5h, file)
	fmt.Printf("MD5:\t %x\n", md5h.Sum([]byte("")))

	sha1h := sha1.New()
	io.Copy(sha1h, file)
	fmt.Printf("SHA1:\t %x\n", sha1h.Sum([]byte("")))

	fmt.Printf("\n\n按任意键退出...")

	wait()
	os.Exit(0)
}

func wait() {
	fmt.Scanln(&temp)
}
