package main

import (
	"fmt"
	"github.com/xuri/excelize"
	"strconv"
)

func main(){
	xlsx := excelize.NewFile()
	
	for i:= 0 ; i < 2000 ; i++{
		xlsx.SetCellValue("Sheet1", "A"+strconv.Itoa(i+1), "lalalalalala")
	}
	
	err := xlsx.SaveAs("file/lalala.xlsx")
	if err != nil {
		fmt.Println(err.Error())
	}
	
	fmt.Println("lalalala")
}
