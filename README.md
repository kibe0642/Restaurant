# Restaurant



This is a backend restaurant management system written in golang
package main

import (
	"fmt"
	"log"

	"github.com/xuri/excelize/v2"
)

// MapColumns reads specific columns from the source file and writes them to the target file with custom headers.
func MapColumns(sourceFile, targetFile string, columnMapping map[string]string) error {
	// Open the source file
	source, err := excelize.OpenFile(sourceFile)
	if err != nil {
		return fmt.Errorf("failed to open source file: %w", err)
	}
	defer source.Close()

	// Create a new target file
	target := excelize.NewFile()

	// Get the first sheet from the source file
	sourceSheet := source.GetSheetName(0)
	if sourceSheet == "" {
		return fmt.Errorf("no sheets found in source file")
	}

	// Add a new sheet to the target file
	targetSheet := "Mapped Data"
	target.NewSheet(targetSheet)

	// Write custom headers to the target file
	row := 1
	colIndex := 1
	for _, customHeader := range columnMapping {
		cell, _ := excelize.CoordinatesToCellName(colIndex, row)
		target.SetCellValue(targetSheet, cell, customHeader)
		colIndex++
	}

	// Map and copy data
	sourceRows, err := source.GetRows(sourceSheet)
	if err != nil {
		return fmt.Errorf("failed to read rows from source file: %w", err)
	}

	for rowIndex, sourceRow := range sourceRows {
		// Skip the header row in the source file
		if rowIndex == 0 {
			continue
		}

		colIndex = 1
		for sourceCol, customHeader := range columnMapping {
			sourceCellIndex, _ := excelize.ColumnNameToNumber(sourceCol)

			if sourceCellIndex <= len(sourceRow) {
				// Read the value from the source cell
				value := sourceRow[sourceCellIndex-1]

				// Write the value to the target file
				cell, _ := excelize.CoordinatesToCellName(colIndex, rowIndex+1)
				target.SetCellValue(targetSheet, cell, value)
			}
			colIndex++
		}
	}

	// Save the target file
	if err := target.SaveAs(targetFile); err != nil {
		return fmt.Errorf("failed to save target file: %w", err)
	}

	return nil
}

func main() {
	// Source and target file paths
	sourceFile := "source.xlsx"
	targetFile := "target.xlsx"

	// Define the column mapping: Source column -> Custom header
	columnMapping := map[string]string{
		"A": "Customer Name",
		"B": "Order ID",
		"C": "Product",
		"D": "Quantity",
	}

	// Map and transfer columns
	if err := MapColumns(sourceFile, targetFile, columnMapping); err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println("Data mapping completed successfully!")
}
