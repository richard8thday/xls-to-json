package main

import (
	"fmt"
	"strconv"
	"encoding/json"
	"os"

	"github.com/tealeg/xlsx"
	//github.com/microcosm-cc/bluemonday
)

// TODO: Define input and output endpoints (read from Confluence? A stream? Write to a location?)
// TODO: Make more intuitive functions out of these test functions
// TODO: Organize the project (better structure)
// TODO: Import bluemonday for sanitization
func main() {
	//testManualParse()

	// Convert xlsx input to a struct
	f := testStruct()
	fmt.Print(f)

	// Encode the struct as json
	j := json.NewEncoder(os.Stdout).Encode(
		&f,
	)

	// Print out our result
	fmt.Print(j)
}

type File struct {
	Filename	string	`json:"filename"`
	Sheets		[]Sheet	`json:"sheets"`
}

type Sheet struct {
	Rows		[]Row	`json:"rows"`
}

type Row struct {
	Cells		[]Cell	`json:"cells"`
}

type Cell struct {
	RowIndex	int		`json:"row_index"`
	CellIndex	int		`json:"cell_index"`
	Value		string	`json:"value"`
}

func testStruct() File {
	// Test parse and print results
	// TODO: Make it a function
	// TODO: Parameterize the file input, probably
	excelFileName := "/go/src/github.com/richard8thday/xls-to-json/input.xlsx"

	// Open the file
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Print(err)
	}

	f := File {
		Filename: excelFileName,
		Sheets: []Sheet{},
	}

	// Iterate over sheets in the file
	for _, sheet := range xlFile.Sheets {
		rowIndex := 0

		// Create a Sheet
		s := Sheet {
			Rows: []Row{},
		}

		// Iterate over rows in the sheet
		for _, row := range sheet.Rows {
			cellIndex := 0

			// Create a Row
			r := Row {
				Cells: []Cell{},
			}

			// Iterate over cells in the row
			for _, cell := range row.Cells {
				text, _ := cell.String()

				// Create a Cell
				c := Cell {
					CellIndex: cellIndex,
					RowIndex: rowIndex,
					Value: text,
				}

				// Add this Cell to the Row
				r.Cells = append(r.Cells, c)

				// Increment cell index
				cellIndex++
			}

			// Add this Row to the Sheet
			s.Rows = append(s.Rows, r)

			// Increment row index
			rowIndex++
		}

		// Add this Sheet to the File
		f.Sheets = append(f.Sheets, s)
	}

	return f
}

// Doing our own parsing sucks, but I wanted to write this out to evaluate it
func testManualParse() {
	// Test parse and print results
	// TODO: Make it a function
	excelFileName := "/go/src/github.com/richard8thday/xls-to-json/input.xlsx"

	// Open the file
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Print(err)
	}

	
	// TODO: Finalize json structure...
	/*
	{
		"filename": "inpux.xlsx",
		"sheets": [{
			"rows": [{
				"cells": [{
					"cell_index": 0
					"row_index": 0
					"value": "1"
				}, {...}]
			}, {...}]
		}, {...}]
	}
	*/

	//Define a variable to hold 
	output := `{"filename": "` + excelFileName + `", "sheets": [`
	sheetIndex := 0

	// Iterate over sheets
	for _, sheet := range xlFile.Sheets {
		rowIndex := 0

		// Begin writing the sheet
		if sheetIndex > 0 {
			output += ","
		}
		output += `{"rows": [`

		// Iterate over rows in the sheet
		for _, row := range sheet.Rows {
			cellIndex := 0
			
			// Begin writing the row
			if rowIndex > 0 {
				output += ","
			}
			output += `{"cells": [`

			// Iterate over cells in the row
			for _, cell := range row.Cells {
				// Get the contents of the cell
				text, _ := cell.String()
				text += ""
				// TODO: check for error
				
				// Write the cell
				if cellIndex > 0 {
					output += ","
				}
				//output += `{"cell_index":` + strconv.Itoa(cellIndex) + `,"row_index":` + strconv.Itoa(rowIndex) + `,"value":"` + text + `"}`
				output += `{"cell_index":` + strconv.Itoa(cellIndex) + `,"row_index":` + strconv.Itoa(rowIndex) + `,"value":"` + "" + `"}`

				// Increment cell index
				cellIndex++
			}

			// Finish writing the row
			output += "]}"

			// Increment row index
			rowIndex++
		}

		// Finish writing the sheet
		output += "]}"

		// Increment sheet index
		sheetIndex++
	}

	output += "]}"

	fmt.Print(output)
}