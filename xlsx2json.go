package main

import (
	//"fmt"
	"log"
	"encoding/json"
	"os"

	"github.com/tealeg/xlsx"
	//github.com/microcosm-cc/bluemonday
)

type ExcelFile struct {
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

// OpenFile returns an xlsx.File
//	TODO: Parameterize the file input
//	TODO: Consider reading from another endpoint
func OpenFile() (*xlsx.File, error) {
	filename := "/go/src/github.com/richard8thday/xlsx2json/input.xlsx"

	// Open the file
	f, err := xlsx.OpenFile(filename)

	return f, err
}

// DecodeFile returns the contents of an xlsx.File as JSON
//	TODO: Consider additional attributes for ExcelFile
func DecodeFile(f *xlsx.File) ([]byte, error) {
	// Create an ExcelFile to return
	ef := ExcelFile {
		Filename: "",
		Sheets: []Sheet{},
	}

	// Iterate over sheets in the file
	for _, sheet := range f.Sheets {
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
		ef.Sheets = append(ef.Sheets, s)
	}

	// JSON-encode the struct
	j, err := json.Marshal(ef)
	
	// Return the JSON encoding
	return j, err
}

// WriteJson sends json-encoded file contents to an endpoint
//	TODO: Everything
func WriteJson(json []byte) error {
	os.Stdout.Write(json)

	return nil
}

// TODO: Define input and output endpoints (read from Confluence? A stream? Write to a location?)
// TODO: Import bluemonday for sanitization? Might not be necessary if the input is internal
// TODO: Write a function to handle errors
func main() {
	// Get a file
	f, err := OpenFile()
	if err != nil {
		log.Fatal(err)
	}

	// Convert the file
	j, err := DecodeFile(f)
	if err != nil {
		log.Fatal(err)
	}
	
	// Write the file somewhere
	err = WriteJson(j)
	if err != nil {
		log.Fatal(err)
	}
}