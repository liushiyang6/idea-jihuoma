package main

import (
	"archive/zip"
	"fmt"
	"github.com/atotto/clipboard"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	downloadFile("http://idea.medeming.com/jets/images/jihuoma.zip")
}

func downloadFile(url string) {
	fmt.Println("开始获取激活码...")
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	f, err := os.Create("jihuoma.zip")
	if err != nil {
		panic(err)
	}
	io.Copy(f, res.Body)

	r, err := zip.OpenReader(f.Name())

	if err != nil {
		fmt.Println(err)
		return
	}
	for _, k := range r.Reader.File {
		if !strings.HasPrefix(k.Name, "2018.2֮") {
			continue
		}
		fmt.Println("开始解压....")
		if k.FileInfo().IsDir() {
			err := os.MkdirAll(k.Name, 0644)
			if err != nil {
				fmt.Println(err)
			}
			continue
		}
		r, err := k.Open()
		if err != nil {
			fmt.Println(err)
			continue
		}

		defer r.Close()
		NewFile, err := os.Create("激活码.txt")
		if err != nil {
			fmt.Println(err)
			continue
		}
		io.Copy(NewFile, r)
		NewFile.Close()

		bytes, err := ioutil.ReadFile("激活码.txt")
		if err != nil {
			fmt.Println("error : %s", err)
			return
		}
		fmt.Println("开始读取....")
		s := string(bytes)
		fmt.Println("激活码:" + s)
		fmt.Println("开始复制到剪切板....")
		clipboard.WriteAll(s)
		fmt.Println("成功!")
	}
	f.Close()
	r.Close()
	os.Remove("./jihuoma.zip")
}
