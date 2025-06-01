package tablemgr

import (
	"bufio"
	"io"
	"log"
	"regexp"
	"strconv"
)

const table_regex = "0x([0-9a-f]{2})\\s+0x([0-9a-f]{4})\\s+.*"

func GenerateDecoder(in io.Reader) ([]int, int, int, int, error) {
	re, err := regexp.Compile(table_regex)
	if err != nil {
		return nil, 0, 0, 0, err
	}
	table := make([]int, 256)

	var oklines = 0
	var kolines = 0
	var comments = 0

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		if line[0] != '#' {
			parts := re.FindStringSubmatch(line)
			if len(parts) != 3 {
				log.Printf("Malformed/unparsable line: [%s]", line)
				kolines++
			} else {
				i, _ := strconv.ParseUint(parts[1], 16, 8)
				v, _ := strconv.ParseUint(parts[2], 16, 16)
				table[int(i)] = int(v)
				oklines++
			}
		} else {
			comments++
		}
	}

	return table, oklines, kolines, comments, nil
}

func GenerateEncoder(decoderTable []int) ([]int, error) {
	table := make([]int, 256)

	// Iterate over the decoder slice content
	for i, v := range decoderTable {
		table[v] = i
	}
	return table, nil
}
