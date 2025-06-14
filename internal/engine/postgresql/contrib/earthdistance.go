// Code generated by sqlc-pg-gen. DO NOT EDIT.

package contrib

import (
	"github.com/istforks/sqlc/internal/sql/ast"
	"github.com/istforks/sqlc/internal/sql/catalog"
)

var funcsEarthdistance = []*catalog.Function{
	{
		Name:       "earth",
		Args:       []*catalog.Argument{},
		ReturnType: &ast.TypeName{Name: "double precision"},
	},
	{
		Name: "earth_box",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "earth"},
			},
			{
				Type: &ast.TypeName{Name: "double precision"},
			},
		},
		ReturnType: &ast.TypeName{Name: "cube"},
	},
	{
		Name: "earth_distance",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "earth"},
			},
			{
				Type: &ast.TypeName{Name: "earth"},
			},
		},
		ReturnType: &ast.TypeName{Name: "double precision"},
	},
	{
		Name: "gc_to_sec",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "double precision"},
			},
		},
		ReturnType: &ast.TypeName{Name: "double precision"},
	},
	{
		Name: "geo_distance",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "point"},
			},
			{
				Type: &ast.TypeName{Name: "point"},
			},
		},
		ReturnType: &ast.TypeName{Name: "double precision"},
	},
	{
		Name: "latitude",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "earth"},
			},
		},
		ReturnType: &ast.TypeName{Name: "double precision"},
	},
	{
		Name: "ll_to_earth",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "double precision"},
			},
			{
				Type: &ast.TypeName{Name: "double precision"},
			},
		},
		ReturnType: &ast.TypeName{Name: "earth"},
	},
	{
		Name: "longitude",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "earth"},
			},
		},
		ReturnType: &ast.TypeName{Name: "double precision"},
	},
	{
		Name: "sec_to_gc",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "double precision"},
			},
		},
		ReturnType: &ast.TypeName{Name: "double precision"},
	},
}

func Earthdistance() *catalog.Schema {
	s := &catalog.Schema{Name: "pg_catalog"}
	s.Funcs = funcsEarthdistance
	return s
}
