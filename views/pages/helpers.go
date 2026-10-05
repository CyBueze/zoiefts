package pages

import "strconv"

func formatNaira(kobo int64) string {
	naira := kobo / 100
	s := strconv.FormatInt(naira, 10)
	if len(s) <= 3 {
		return s
	}
	var out []byte
	for i, c := range []byte(s) {
		if i > 0 && (len(s)-i)%3 == 0 {
			out = append(out, ',')
		}
		out = append(out, c)
	}
	return string(out)
}