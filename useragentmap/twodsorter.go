package useragentmap

import "strconv"

type TwoDSorter [][]string

func (t TwoDSorter) Len() int { return len(t) }

func (t TwoDSorter) Less(i, j int) bool {
	iValue, err := strconv.Atoi(t[i][1])
	if err != nil {
		panic(err)
	}

	jValue, err := strconv.Atoi(t[j][1])
	if err != nil {
		panic(err)
	}

	return iValue > jValue
}

func (t TwoDSorter) Swap(i, j int) { t[i], t[j] = t[j], t[i] }
