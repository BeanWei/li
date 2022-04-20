package sqlx

import "strings"

// escape escapes w with the default escape character ('/'),
// to be used by the pattern matching functions below.
// The second return value indicates if w was escaped or not.
func escape(w string) (string, bool) {
	var n int
	for i := range w {
		if c := w[i]; c == '%' || c == '_' || c == '\\' {
			n++
		}
	}
	// No characters to escape.
	if n == 0 {
		return w, false
	}
	var b strings.Builder
	b.Grow(len(w) + n)
	for i := range w {
		if c := w[i]; c == '%' || c == '_' || c == '\\' {
			b.WriteByte('\\')
		}
		b.WriteByte(w[i])
	}
	return b.String(), true
}
