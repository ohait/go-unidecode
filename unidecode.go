package unidecode

import (
	"strings"
)

var m = make(map[int]string, 65536)

func Decode(unicode string) string {
	out := strings.Builder{}
	for _, r := range unicode {
		x, exists := m[int(r)]
		if exists {
			//log.Printf("`%c`=>%s", r, x)
			out.WriteString(x)
		} else {
			out.WriteRune(r) // leave untouched
		}
	}
	return out.String()
}
