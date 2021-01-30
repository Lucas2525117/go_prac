package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//获取命令行参数
	list := os.Args
	if len(list) != 3 {
		fmt.Println("usage: xxx srcFile dstFile")
		return
	}

	srcFileName := list[1]
	dstFileName := list[2]
	if srcFileName == dstFileName {
		fmt.Println("源文件和目的文件名字不能相同")
		return
	}

	//只读方式打开源文件
	srcFile, err1 := os.Open(srcFileName)
	if err1 != nil {
		fmt.Println("err1 = ", err1)
		return
	}

	//新建目的文件
	dstFile, err2 := os.Create(dstFileName)
	if err2 != nil {
		fmt.Println("err2 = ", err2)
		return
	}

	defer srcFile.Close()
	defer dstFile.Close()

	//从源文件读物内容，往目的文件写，读多少写多少
	buf := make([]byte, 4*1024)
	for {
		//往buf里面读
		num, err3 := srcFile.Read(buf)
		if err3 != nil {
			if err3 == io.EOF {
				fmt.Println("文件拷贝完成")
				break
			}
		}

		//往目的文件写数据
		dstFile.Write(buf[:num])
	}
}
