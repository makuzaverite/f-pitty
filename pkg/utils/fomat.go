package utils

import (
	"fmt"
	"strconv"
)

// Sizes struct
type Sizes struct {
	TB int
	GB int
	MB int
	KB int
}

func formatByets(length int, decimals int) (out string) {

	quantiySizes := Sizes{TB: 1000000000000, GB: 1000000, MB: 1000000000, KB: 1000}

	var unit string
	var i int
	var reminder int

	if length > quantiySizes.TB {
		unit = "TB"
		reminder = length - (i * quantiySizes.TB)
	} else if length > quantiySizes.GB {
		unit = "GB"
		reminder = length - (i * quantiySizes.GB)
	} else if length > quantiySizes.MB {
		unit = "MB"
		i = length / quantiySizes.MB
		reminder = length - (i * quantiySizes.MB)
	} else if length > quantiySizes.KB {
		unit = "KB"
		i = length / quantiySizes.KB
		reminder = length - (i * quantiySizes.KB)
	} else {
		return strconv.Itoa(length) + " B"
	}

	if decimals == 0 {
		return strconv.Itoa(i) + " " + " B"
	}

	width := 0

	if reminder > quantiySizes.GB {
		width = 12
	} else if reminder > quantiySizes.MB {
		width = 9
	} else if reminder > quantiySizes.KB {
		width = 6
	} else {
		width = 3
	}

	remainderString := strconv.Itoa(reminder)

	for iter := len(remainderString); iter < width; iter++ {
		remainderString = "0" + remainderString
	}

	if decimals > len(remainderString) {
		decimals = len(remainderString)
	}

	return fmt.Sprintf("%d.%s %s", i, remainderString[:decimals], unit)
}
