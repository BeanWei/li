package sqlx

import (
	"strings"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

func ContainsFoldAny(col, substr string) *sql.Predicate {
	p := sql.P()
	return p.Append(func(b *sql.Builder) {
		w, escaped := escape(substr)
		switch b.Dialect() {
		case dialect.MySQL:
			b.WriteString("CONVERT(" + col + " using 'utf8mb4_general_ci') LIKE ")
			b.Arg("%" + strings.ToLower(w) + "%")
		case dialect.Postgres:
			b.WriteString(col + "::text ILIKE ")
			b.Arg("%" + strings.ToLower(w) + "%")
		default: // SQLite.
			var f sql.Func
			f.SetDialect(b.Dialect())
			f.Lower(col)
			b.WriteString(f.String()).WriteString(" LIKE ")
			b.Arg("%" + strings.ToLower(w) + "%")
			if escaped {
				p.WriteString(" ESCAPE ").Arg("\\")
			}
		}
	})
}
