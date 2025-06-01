package tablemgr

import (
	"bytes"
	"fmt"
)

type TableType int

const EncodeTable TableType = 1
const DecodeTable TableType = 2

func GenerateTable(table []int, name string, tableType TableType, euroValue string) string {
	var output string = ""
	for i := 0; i < 16; i++ {
		var line string = ""
		for j := 0; j < 16; j++ {
			line += fmt.Sprintf("0x%02x,", table[16*i+j])
		}
		start := 16 * i
		end := 16*i + 15
		output += fmt.Sprintf("\t\t\t%s // 0x%02x..0x%02x\n", line, start, end)
	}
	var outBuff bytes.Buffer
	if tableType == EncodeTable {
		outBuff.WriteString(fmt.Sprintf("\t%s: unicodeToEbcdicMap{\n", name))
	} else {
		outBuff.WriteString(fmt.Sprintf("\t%s: {\n", name))
	}
	if euroValue == "" {
		outBuff.WriteString("\t\tHasEuroPatch: false,\n")
	} else {
		outBuff.WriteString("\t\tHasEuroPatch: true,\n")
		outBuff.WriteString(fmt.Sprintf("\t\tEuroChar:     %s,\n", euroValue))
	}
	if tableType == EncodeTable {
		outBuff.WriteString(fmt.Sprintln("\t\tMap: []byte{"))
	} else {
		outBuff.WriteString(fmt.Sprintln("\t\tMap: []rune{"))
	}
	outBuff.WriteString(output)
	outBuff.WriteString("\t\t}\n\t},\n")

	return outBuff.String()
}
