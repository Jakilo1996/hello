// Package hello Package 是 hello 项目的 hello 模块.
// 作者: Jakilo
// 日期: 2024/10/22 17:14
// 邮箱: 1192788767@qq.com
// GitHub: Jakilo1996
package main

import (
	"flag"
	"github.com/Jakilo1996/hello"
	"os"
)

var name string

func init() {
	flag.StringVar(&name, "name", "world", "name to say hello")
}

func main() {
	flag.Parse()
	msg := hello.Hello(name)
	_, err := os.Stdout.WriteString(msg)
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
}
