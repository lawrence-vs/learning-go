package parser

import (
	"github.com/xuri/excelize/v2"
)

type ParsedData struct {
	DeviceName     string
	PlcDescription string
	Type           string
	DPAddress      string
	Symbol         string
	Description    string
}

func ParseExcel(filename string) ([]ParsedData, error) {
	f, err := excelize.OpenFile(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var parsedData []ParsedData

	sheets := f.GetSheetList()
	for _, sheetName := range sheets {
		rows, err := f.GetRows(sheetName)
		if err != nil {
			return nil, err
		}
		for _, row := range rows[1:] {
			// Adjust parsing based on the columns available
			if len(row) >= 5 {
				parsedData = append(parsedData, ParsedData{
					DeviceName:     row[0],
					PlcDescription: row[1],
					Type:           row[2],
					DPAddress:      row[3],
					Symbol:         row[4],
					Description:    row[5],
				})
			}
		}
	}

	return parsedData, nil
}
