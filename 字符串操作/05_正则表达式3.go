package main

import (
	"fmt"
	"regexp"
)

func main() {
	buf := `
	<!DOCTYPE html>
<html lang="zh-CN">
<head>
	<title>Go语言标准库文档中文版 | Go语言中文网 | Golang中文社区 | Golang中国</title>
	<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
	<meta http-equiv="X-UA-Compatible" content="IE=edge, chrome=1">
	<meta charset="utf-8">
	<link rel="shortcut icon" href="/static/img/go.ico">
	<link rel="apple-touch-icon" type="image/png" href="/static/img/logo2.png">
	<meta name="author" content="polaris <polaris@studygolang.com>">
	<meta name="keywords" content="中文, 文档, 标准库, Go语言,Golang,Go社区,Go中文社区,Golang中文社区,Go语言社区,Go语言学习,学习Go语言,Go语言学习园地,Golang 中国,Golang中国,Golang China, Go语言论坛, Go语言中文网">
	<meta name="description" content="Go语言文档中文版，Go语言中文网，中国 Golang 社区，Go语言学习园地，致力于构建完善的 Golang 中文社区，Go语言爱好者的学习家园。分享 Go 语言知识，交流使用经验">
</head>

<div>测试1</div>
<div>测试2
你好
麦克
</div>
<div>测试3</div>
<div>测试4</div>

<frameset cols="15,85">
	<frame src="/static/pkgdoc/i.html">
	<frame name="main" src="/static/pkgdoc/main.html" tppabs="main.html" >
	<noframes>
	</noframes>
</frameset>
</html>
	`
	//1. 解析正则表达式
	//reg := regexp.MustCompile(`<div>(.*)</div>`) //+表示前一个字符1次或多次
	reg := regexp.MustCompile(`<div>(?s:(.*?))</div>`)
	if reg == nil {
		fmt.Println("reg = ", reg)
		return
	}

	//2. 提取关键信息
	s := reg.FindAllStringSubmatch(buf, -1) //-1表示匹配所有
	fmt.Println("s = ", s)

	//3. 过滤
	for _, data := range s {
		//fmt.Println("data[0] = ", data[0]) //带标签

		fmt.Println("data[1] = ", data[1]) //不带标签
	}
}
