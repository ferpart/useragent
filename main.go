package main

import (
	"fmt"
	"os"

	"github.com/ferpart/useragent/csv"
	"github.com/ferpart/useragent/useragentmap"
)

const (
	inPath  = "in"
	csvPath = inPath + "/%s"
)

func main() {
	fileInfos, err := os.ReadDir(inPath)
	if err != nil {
		panic(err)
	}

	for _, fileInfo := range fileInfos {
		if err := run(fileInfo.Name()); err != nil {
			panic(err)
		}
	}
}

func run(fileName string) error {
	filePath := fmt.Sprintf(csvPath, fileName)
	userAgentMap, err := useragentmap.Load(filePath)
	if err != nil {
		return err
	}

	generalizedUserAgentMap := userAgentMap.Generalize()
	userAgentRecords := generalizedUserAgentMap.To2dSortedArray()
	if err = csv.Save(fileName, userAgentRecords); err != nil {
		return err
	}

	return nil
}
