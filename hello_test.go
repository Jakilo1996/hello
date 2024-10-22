// Package hello Package 是 hello 项目的 hello 模块.
// 作者: Jakilo
// 日期: 2024/10/22 16:17
// 邮箱: 1192788767@qq.com
// GitHub: Jakilo1996
package hello_test

import (
	"fmt"
	"github.com/Jakilo1996/hello"
	"testing"
)

func TestHello(t *testing.T) {
	data := "jack"
	expected := fmt.Sprintf("hello %s!", data)
	result := hello.Hello(data)

	if result != expected {
		t.Fatalf("expected result %s, but got %s", expected, result)
	}
}
