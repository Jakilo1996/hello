// Package hello Package 是 hello 项目的 hello 模块.
// 作者: Jakilo
// 日期: 2024/10/22 16:11
// 邮箱: 1192788767@qq.com
// GitHub: Jakilo1996
package hello

import "fmt"

// ExampleFunc 是一个示例函数.
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return fmt.Sprintf("Hello %s!", name)
}
