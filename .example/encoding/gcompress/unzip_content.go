package main

import (
	"fmt"

	"github.com/gogf/gf/encoding/gcompress"
	"github.com/gogf/gf/os/gfile"
)

func main() {
	err := gcompress.UnZipContent(
		gfile.GetBinContents(`D:\Workspace\Go\GOPATH\src\github.com\gogf\gf\geg\encoding\gcompress\data.zip`),
		`D:\Workspace\Go\GOPATH\src\github.com\gogf\gf\geg`,
	)
	fmt.Println(err)
}