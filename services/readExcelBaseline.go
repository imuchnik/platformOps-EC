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


func loadControl(file string) (controls []models.Control) {
	xlFile, err := xlsx.OpenFile(file)
	if err != nil {
		fmt.Println("error reading")
	}
	sheet := xlFile.Sheets[0]
	length := len(sheet.Rows)

	// Removing header in excel sheet
	//TODO: this is brittle, because the spreadsheets vary
	//TODO: determine the row dynamically

	rows := sheet.Rows[2 : length-1]

	for _, row := range rows {

		cells := row.Cells
		fmt.Printf("about to read %v\n", cells[0])

		reqId, err := cells[0].Int()
		fmt.Printf("my ReqId %v\n", cells[0])
		if err != nil {
			fmt.Println("error reading reqId")
			//fmt.Println(err)
		}
		fmt.Println("why is this a problem?",  cells[0],)
		//TODO: same as above the columns vary
		//Need to dynamically determine the ones we need
		control := models.Control{
			ReqId: reqId,
			CisId: cells[0].String(),
			Category: cells[1].String(),
			Requirement: cells[2].String(),
			Discussion: cells[3].String(),
			CheckText: cells[4].String(),
			FixText: cells[5].String(),
			RowDesc: cells[0].String(),
		}

		fmt.Println("where do I die?", control.FixText)

		controls = append(controls, control)
		fmt.Println("appended", control.ReqId)
	}

	return controls
}