package useragentmap

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/mileusna/useragent"
)

const (
	mapKeyTemplate = "%s/%s"
	sevenPlus      = "7plus_android"
	fetchTv        = "FetchTV"
)

type UserAgentMap map[string]int

func Load(filePath string) (UserAgentMap, error) {
	csvFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer func() { _ = csvFile.Close() }()

	csvReader := csv.NewReader(csvFile)

	if _, err := csvReader.Read(); err != nil {
		return nil, err
	}

	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	csvMap := make(map[string]int)
	for _, line := range records {
		if len(line) != 2 {
			return nil, fmt.Errorf("line is more than 2 elements long")
		}

		userAgentCount, err := strconv.Atoi(line[1])
		if err != nil {
			return nil, err
		}

		csvMap[line[0]] = userAgentCount
	}

	return csvMap, nil
}

func (u UserAgentMap) Generalize() UserAgentMap {
	generalizedUserAgentMap := make(map[string]int)
	for userAgent, count := range u {
		parsedUserAgent := useragent.Parse(userAgent)

		mapKey := generateMapKey(&parsedUserAgent)

		generalizedUserAgentMap[mapKey] += count
	}
	return generalizedUserAgentMap
}

func (u UserAgentMap) To2dSortedArray() [][]string {
	outputArray := make([][]string, len(u))

	var iter int
	for key, value := range u {
		outputArray[iter] = []string{key, fmt.Sprintf("%d", value)}
		iter++
	}

	sort.Sort(TwoDSorter(outputArray))
	return outputArray
}

func generateMapKey(userAgent *useragent.UserAgent) string {
	if userAgent.String == "" {
		return "empty_user_agent"
	}

	if userAgent.Bot {
		return "bot"
	}

	userAgentName, userAgentVersion := userAgent.OS, userAgent.OSVersionNoShort()
	if userAgentVersion == "" {
		userAgentVersion = userAgent.OSVersion
	}

	if userAgentName == "" || userAgentVersion == "" {
		userAgentName, userAgentVersion = userAgent.Name, userAgent.VersionNoShort()
	}

	if userAgentVersion == "" {
		userAgentVersion = userAgent.Version
	}

	if userAgentName != "" && userAgentVersion != "" {
		return fmt.Sprintf(mapKeyTemplate, userAgentName, userAgentVersion)
	}

	if strings.Contains(userAgent.Name, sevenPlus) {
		return sevenPlus
	}

	if strings.Contains(userAgent.String, fetchTv) {
		splitUserAgent := strings.Split(userAgent.String, " ")
		fetchTvVersion := splitUserAgent[len(splitUserAgent)-2]

		return fmt.Sprintf(mapKeyTemplate, fetchTv, fetchTvVersion)
	}

	return ""
}
