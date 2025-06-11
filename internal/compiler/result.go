package compiler

import (
	"github.com/istforks/sqlc/internal/sql/catalog"
)

type Result struct {
	Catalog *catalog.Catalog
	Queries []*Query
}
