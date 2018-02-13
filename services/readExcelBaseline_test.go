package services

import (
	"testing"
	"github.com/tealeg/xlsx"

	"fmt"
	"reflect"
)

func testExcelHeaderParsing(t *testing.T){
	file:="test-data/test-baseline.xlsx"
	xlFile, err := xlsx.OpenFile(file)
	if err != nil {
		fmt.Println("error reading")
	}
	sheet := xlFile.Sheets[0]
	fmt.Println(reflect.TypeOf(sheet))
}