package csv

import (
	"encoding/csv"
	"fmt"
	"os"
)

const (
	outPath = "out"
	csvPath = outPath + "/%s"
)

func Save(fileName string, records [][]string) error {
	if err := os.MkdirAll(outPath, os.ModePerm); err != nil {
		return err
	}

	filePath := fmt.Sprintf(csvPath, fileName)
	csvFile, err := os.Create(filePath)
	if err != nil {
		return err
	}

	defer func() { _ = csvFile.Close() }()

	csvWriter := csv.NewWriter(csvFile)

	if err = csvWriter.Write([]string{"normalized_user_agent", "count"}); err != nil {
		return err
	}

	if err = csvWriter.WriteAll(records); err != nil {
		return err
	}

	return nil
}
