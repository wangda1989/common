package common

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"reflect"
)

func LoadExcel2Struct(filePath, sheetName string, sRow, sCol, eCol int, holder interface{}) error {
	ef, err := excelize.OpenFile(filePath)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer ef.Close()
	rows, err := ef.GetRows(sheetName) // 所有行
	if err != nil {
		return err
	}

	for _, row := range rows[sRow:] {
		tp := reflect.TypeOf(holder).Elem().Elem().Elem() // 结构体的类型
		val := reflect.New(tp)                            // 创建一个新的结构体对象

		for i := sCol; i <= eCol; i++ {
			field := val.Elem().Field(i) // 第idx个字段的反射Value
			cellValue := row[i]          // 第idx个字段对应的Excel数据
			field.SetString(cellValue)   // 将Excel数据保存到结构体对象的对应字段中
		}

		listV := reflect.ValueOf(holder)
		listV.Elem().Set(reflect.Append(listV.Elem(), val)) // 将结构体对象添加到holder中
	}

	return nil
}
