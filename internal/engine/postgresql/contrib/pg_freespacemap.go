// Code generated by sqlc-pg-gen. DO NOT EDIT.

package contrib

import (
	"github.com/istforks/sqlc/internal/sql/ast"
	"github.com/istforks/sqlc/internal/sql/catalog"
)

var funcsPgFreespacemap = []*catalog.Function{
	{
		Name: "pg_freespace",
		Args: []*catalog.Argument{
			{
				Name: "rel",
				Type: &ast.TypeName{Name: "regclass"},
			},
		},
		ReturnType: &ast.TypeName{Name: "record"},
	},
	{
		Name: "pg_freespace",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "regclass"},
			},
			{
				Type: &ast.TypeName{Name: "bigint"},
			},
		},
		ReturnType: &ast.TypeName{Name: "smallint"},
	},
}

func PgFreespacemap() *catalog.Schema {
	s := &catalog.Schema{Name: "pg_catalog"}
	s.Funcs = funcsPgFreespacemap
	return s
}
