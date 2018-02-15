package services

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"path/filepath"
	"platformOps-EC/models"
)

func LoadFromExcel(file string) (b models.Baseline, controls []models.Control) {
	return loadBaseline(file), loadControl(file)
}

func loadBaseline(file string) (b models.Baseline) {
	name := filepath.Base(file)
	return models.Baseline{Name: name}
}

func parseHeaders(sheet *xlsx.Sheet)(int, int, int, int, int, int,int){

	var startrow, reqIdInd,  CategoryInd, RequirementsInd, DiscussionInd, CheckTextInd, FixTextInd int

	rows := sheet.Rows

	for i, r := range rows{
		if r.Cells[0].Value =="Req #"{
			startrow = i
			cells := r.Cells
			for n, c:= range cells{
				switch c.Value {
				case "Req #":{
					reqIdInd = n
				}
				case "Category":{
					CategoryInd=n
				}
				case "Requirements":{
					RequirementsInd =n
				}
				case "Discussion":{
					DiscussionInd =n
				}
				case "Check Text":{
					CheckTextInd =n
				}
				case "Fix Text":{
					FixTextInd =n
				}

				}
			}

		}
	}
	return startrow, reqIdInd, CategoryInd, RequirementsInd, DiscussionInd, CheckTextInd, FixTextInd


}
func loadControl(file string) (controls []models.Control) {
	xlFile, err := xlsx.OpenFile(file)
	if err != nil {
		fmt.Println("error reading")
	}
	sheet := xlFile.Sheets[0]
	length := len(sheet.Rows)

	// Removing header in excel sheet
	startrow, reqIdInd, CategoryInd, RequirementsInd, DiscussionInd, CheckTextInd, FixTextInd := parseHeaders(sheet)
	rows := sheet.Rows[startrow : length-1]

	for _, row := range rows {

		cells := row.Cells
		fmt.Printf("about to read %v\n", cells[0])

		reqId, err := cells[reqIdInd].Int()
		fmt.Printf("my ReqId %v\n", cells[0])
		if err != nil {
			fmt.Println("error reading reqId")
			//fmt.Println(err)
		}

		//Need to dynamically determine the ones we need
		control := models.Control{
			ReqId: reqId,
			CisId: cells[reqIdInd].String(),
			Category: cells[CategoryInd].String(),
			Requirement: cells[RequirementsInd].String(),
			Discussion: cells[DiscussionInd].String(),
			CheckText: cells[CheckTextInd].String(),
			FixText: cells[FixTextInd].String(),
			RowDesc: cells[reqIdInd].String(),
		}


		controls = append(controls, control)
		fmt.Println("appended", control.ReqId)
	}

	return controls
}