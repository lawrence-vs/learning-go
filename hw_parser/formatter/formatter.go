package formatter

import (
	"fmt"
	"hw_parser/config"
	"hw_parser/parser"

	"github.com/xuri/excelize/v2"
)

func WriteFormattedExcel(parsedData []parser.ParsedData, outputFile string, cfg *config.Configuration) error {
	f := excelize.NewFile()

	sheetName := "Report"
	f.NewSheet(sheetName)
	headerStyle, _ := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{Bold: true, Color: "#000000"},
		Fill: excelize.Fill{Type: "pattern", Pattern: 1, Color: []string{"#C0C0C0"}},
	})

	for i, header := range cfg.IODocument.ColumnNames {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheetName, cell, header)
		f.SetCellStyle(sheetName, cell, cell, headerStyle)
	}

	row := 2
	for _, data := range parsedData {
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", row), data.DeviceName)
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", row), data.Description)
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", row), data.Type)
		f.SetCellValue(sheetName, fmt.Sprintf("D%d", row), data.DPAddress)
		f.SetCellValue(sheetName, fmt.Sprintf("E%d", row), data.Symbol)
		f.SetCellValue(sheetName, fmt.Sprintf("F%d", row), data.Description)
		row++
	}

	return f.SaveAs(outputFile)
}
