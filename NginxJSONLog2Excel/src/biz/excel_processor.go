package biz

import (
    "fmt"
    "github.com/360EntSecGroup-Skylar/excelize"
    "log"
    "strconv"
)

type ExcelData struct {
    columnTitleList []string
    file            *excelize.File
    FileName        string
}

const sheetName = "Nginx Log"

func NewExcelFile(fileName string) *ExcelData {
    newObj := new(ExcelData)
    newObj.FileName = fileName
    newObj.file = excelize.NewFile()
    //newObj.file.NewSheet(sheetName)
    newObj.file.SetSheetName("Sheet1", sheetName)
    newObj.columnTitleList = make([]string, 0)
    return newObj
}

func (excelData *ExcelData) Save() {
    err := excelData.file.SaveAs(excelData.FileName)
    if err != nil {
        log.Fatalln(err)
    }
}

func (excelData *ExcelData) WriteAllData(mapObjList []map[string]string) {
    for idx, obj := range mapObjList {
        excelData.WriteData(obj, idx)
    }
}

func (excelData *ExcelData) setCellValueEx(sheet, axis string, value interface{}) {

    // 内容尝试转数字
    val, err := strconv.ParseFloat(fmt.Sprintf("%v", value), 64)

    // 转成功则按数字使用
    if err == nil {
      value = val
    }

    err = excelData.file.SetCellValue(sheetName, axis, value)
    if err != nil {
        log.Fatalln("Cell SetCellValue Error: ", err)
    }
}

func (excelData *ExcelData) WriteData(mapObj map[string]string, objIndex int) {

    for sKey, sVal := range mapObj {
        columnIdx := indexOfString(sKey, excelData.columnTitleList)
        if columnIdx == -1 {
            excelData.columnTitleList = append(excelData.columnTitleList, sKey)
            columnIdx = len(excelData.columnTitleList) - 1
            err := excelData.file.SetCellValue(sheetName, columnToTitleCell(columnIdx), sKey)
            if err != nil {
                log.Fatalln("Title SetCellValue Error: ", err)
            }
        }
        cellAxis := rowColumnToCell(objIndex, columnIdx)
        excelData.setCellValueEx(sheetName, cellAxis, sVal)
    }
}

func columnToTitleCell(columnIdx int) string {
    return string('A'+rune(columnIdx)) + "1"
}

func rowColumnToCell(row, column int) string {
    row += 2
    return string('A'+rune(column)) + strconv.Itoa(row)
}

func indexOfString(str string, strArray []string) int {
    retIndex := -1
    for idx, s := range strArray {
        if str == s {
            retIndex = idx
        }
    }
    return retIndex
}
