package components

import "strconv"

// itoaLiteral renders an int as a plain string for embedding into
// Datastar expressions (e.g. "$amount = 5000").
func itoaLiteral(n int) string {
	return strconv.Itoa(n)
}

// formatThousands renders 5000 as "5,000" for display on preset buttons.
func formatThousands(n int) string {
	s := strconv.Itoa(n)
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