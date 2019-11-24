package main

import (
	"biz"
	"log"
    "os"
    "path"
    "strings"
)

func main() {

    var srcFile string
    var destFile string

    // 参数个数
    paramCount := len(os.Args)

    // 没有带参数
    if paramCount < 2 {
        log.Fatalln("Error: No source file specified.")
    }

    // 只指定了源文件
    if paramCount == 2 {

        srcFile = os.Args[1]

        fileSuffix := path.Ext(srcFile)
        destFile = strings.TrimSuffix(srcFile, fileSuffix)

        destFile += ".xlsx"
    }

    // 同时指定了 源文件 和 目标文件
    if paramCount > 2 {
        destFile = os.Args[2]
    }

	data := biz.ReadJSONFile(srcFile)
	excelData := biz.NewExcelFile(destFile)
	excelData.WriteAllData(data)
	excelData.Save()

	println("Successfully imported Nginx logs into Excel.")
}
