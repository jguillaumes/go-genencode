package tablemgr

import (
	"bytes"
	"fmt"
)

func GenerateEBCDICtoASCII(table []int, name string, withEuro bool) string {
	var output string = ""
	for i := 0; i < 16; i++ {
		var line string = ""
		for j := 0; j < 16; j++ {
			line += fmt.Sprintf("0x%02x,", table[16*i+j])
		}
		start := 16 * i
		end := 16*i + 15
		if i < 15 {
			output += fmt.Sprintf("\t\t\t%s // 0x%02x..0x%02x\n", line, start, end)
		} else {
			line = line[:len(line)-1]
			output += fmt.Sprintf("\t\t\t%s  // 0x%02x..0x%02x\n", line, start, end)
		}
	}
	var outBuff bytes.Buffer
	outBuff.WriteString(fmt.Sprintf("\t%s: unicodeToEbcdicMap{\n", name))
	outBuff.WriteString(fmt.Sprintf("\t\tHasEuroPatch: %t\n", withEuro))
	outBuff.WriteString(fmt.Sprintln("\t\tEuroChar:     0x9F,"))
	outBuff.WriteString(fmt.Sprintln("\t\tMap: []byte{"))
	outBuff.WriteString(output)
	outBuff.WriteString("\t\t)\n\t},\n")

	return outBuff.String()
}
