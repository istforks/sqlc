package postgresql

import (
	"github.com/istforks/sqlc/internal/sql/catalog"
)

func pgTemp() *catalog.Schema {
	return &catalog.Schema{Name: "pg_temp"}
}
