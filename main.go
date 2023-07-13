package main

import (
	"fmt"

	"github.com/ferpart/useragent/csv"
	"github.com/ferpart/useragent/useragentmap"
)

const (
	csvPath = "data/%s.csv"
)

func main() {
	playbackTypes := []string{"dash", "hls"}

	for _, playbackType := range playbackTypes {
		if err := run(playbackType); err != nil {
			panic(err)
		}
	}
}

func run(playbackType string) error {
	filePath := fmt.Sprintf(csvPath, playbackType)
	userAgentMap, err := useragentmap.Load(filePath)
	if err != nil {
		return err
	}

	generalizedUserAgentMap := userAgentMap.Generalize()
	userAgentRecords := generalizedUserAgentMap.To2dSortedArray()
	if err = csv.Save(playbackType, userAgentRecords); err != nil {
		return err
	}

	return nil
}
