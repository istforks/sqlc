// Code generated by sqlc-pg-gen. DO NOT EDIT.

package contrib

import (
	"github.com/istforks/sqlc/internal/sql/ast"
	"github.com/istforks/sqlc/internal/sql/catalog"
)

var funcsLtree = []*catalog.Function{
	{
		Name: "_lt_q_regex",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ltree[]"},
			},
			{
				Type: &ast.TypeName{Name: "lquery[]"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "_lt_q_rregex",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "lquery[]"},
			},
			{
				Type: &ast.TypeName{Name: "ltree[]"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "_ltq_extract_regex",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ltree[]"},
			},
			{
				Type: &ast.TypeName{Name: "lquery"},
			},
		},
		ReturnType: &ast.TypeName{Name: "ltree"},
	},
	{
		Name: "_ltq_regex",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ltree[]"},
			},
			{
				Type: &ast.TypeName{Name: "lquery"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "_ltq_rregex",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "lquery"},
			},
			{
				Type: &ast.TypeName{Name: "ltree[]"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "_ltree_extract_isparent",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ltree[]"},
			},
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
		},
		ReturnType: &ast.TypeName{Name: "ltree"},
	},
	{
		Name: "_ltree_extract_risparent",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ltree[]"},
			},
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
		},
		ReturnType: &ast.TypeName{Name: "ltree"},
	},
	{
		Name: "_ltree_isparent",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ltree[]"},
			},
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "_ltree_r_isparent",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
			{
				Type: &ast.TypeName{Name: "ltree[]"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "_ltree_r_risparent",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
			{
				Type: &ast.TypeName{Name: "ltree[]"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "_ltree_risparent",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ltree[]"},
			},
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "_ltxtq_exec",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ltree[]"},
			},
			{
				Type: &ast.TypeName{Name: "ltxtquery"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "_ltxtq_extract_exec",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ltree[]"},
			},
			{
				Type: &ast.TypeName{Name: "ltxtquery"},
			},
		},
		ReturnType: &ast.TypeName{Name: "ltree"},
	},
	{
		Name: "_ltxtq_rexec",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ltxtquery"},
			},
			{
				Type: &ast.TypeName{Name: "ltree[]"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "index",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer"},
	},
	{
		Name: "index",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
			{
				Type: &ast.TypeName{Name: "integer"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer"},
	},
	{
		Name: "lca",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
		},
		ReturnType: &ast.TypeName{Name: "ltree"},
	},
	{
		Name: "lca",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
		},
		ReturnType: &ast.TypeName{Name: "ltree"},
	},
	{
		Name: "lca",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
		},
		ReturnType: &ast.TypeName{Name: "ltree"},
	},
	{
		Name: "lca",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
		},
		ReturnType: &ast.TypeName{Name: "ltree"},
	},
	{
		Name: "lca",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
		},
		ReturnType: &ast.TypeName{Name: "ltree"},
	},
	{
		Name: "lca",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
		},
		ReturnType: &ast.TypeName{Name: "ltree"},
	},
	{
		Name: "lca",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
		},
		ReturnType: &ast.TypeName{Name: "ltree"},
	},
	{
		Name: "lca",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ltree[]"},
			},
		},
		ReturnType: &ast.TypeName{Name: "ltree"},
	},
	{
		Name: "lquery_in",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "cstring"},
			},
		},
		ReturnType: &ast.TypeName{Name: "lquery"},
	},
	{
		Name: "lquery_out",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "lquery"},
			},
		},
		ReturnType: &ast.TypeName{Name: "cstring"},
	},
	{
		Name: "lquery_send",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "lquery"},
			},
		},
		ReturnType: &ast.TypeName{Name: "bytea"},
	},
	{
		Name: "lt_q_regex",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
			{
				Type: &ast.TypeName{Name: "lquery[]"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "lt_q_rregex",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "lquery[]"},
			},
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "ltq_regex",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
			{
				Type: &ast.TypeName{Name: "lquery"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "ltq_rregex",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "lquery"},
			},
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "ltree2text",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
		},
		ReturnType: &ast.TypeName{Name: "text"},
	},
	{
		Name: "ltree_addltree",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
		},
		ReturnType: &ast.TypeName{Name: "ltree"},
	},
	{
		Name: "ltree_addtext",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
			{
				Type: &ast.TypeName{Name: "text"},
			},
		},
		ReturnType: &ast.TypeName{Name: "ltree"},
	},
	{
		Name: "ltree_cmp",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer"},
	},
	{
		Name: "ltree_eq",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "ltree_ge",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "ltree_gist_in",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "cstring"},
			},
		},
		ReturnType: &ast.TypeName{Name: "ltree_gist"},
	},
	{
		Name: "ltree_gist_out",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ltree_gist"},
			},
		},
		ReturnType: &ast.TypeName{Name: "cstring"},
	},
	{
		Name: "ltree_gt",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "ltree_in",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "cstring"},
			},
		},
		ReturnType: &ast.TypeName{Name: "ltree"},
	},
	{
		Name: "ltree_isparent",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "ltree_le",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "ltree_lt",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "ltree_ne",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "ltree_out",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
		},
		ReturnType: &ast.TypeName{Name: "cstring"},
	},
	{
		Name: "ltree_risparent",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "ltree_send",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
		},
		ReturnType: &ast.TypeName{Name: "bytea"},
	},
	{
		Name: "ltree_textadd",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "text"},
			},
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
		},
		ReturnType: &ast.TypeName{Name: "ltree"},
	},
	{
		Name: "ltxtq_exec",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
			{
				Type: &ast.TypeName{Name: "ltxtquery"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "ltxtq_in",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "cstring"},
			},
		},
		ReturnType: &ast.TypeName{Name: "ltxtquery"},
	},
	{
		Name: "ltxtq_out",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ltxtquery"},
			},
		},
		ReturnType: &ast.TypeName{Name: "cstring"},
	},
	{
		Name: "ltxtq_rexec",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ltxtquery"},
			},
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "ltxtq_send",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ltxtquery"},
			},
		},
		ReturnType: &ast.TypeName{Name: "bytea"},
	},
	{
		Name: "nlevel",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer"},
	},
	{
		Name: "subltree",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
			{
				Type: &ast.TypeName{Name: "integer"},
			},
			{
				Type: &ast.TypeName{Name: "integer"},
			},
		},
		ReturnType: &ast.TypeName{Name: "ltree"},
	},
	{
		Name: "subpath",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
			{
				Type: &ast.TypeName{Name: "integer"},
			},
		},
		ReturnType: &ast.TypeName{Name: "ltree"},
	},
	{
		Name: "subpath",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "ltree"},
			},
			{
				Type: &ast.TypeName{Name: "integer"},
			},
			{
				Type: &ast.TypeName{Name: "integer"},
			},
		},
		ReturnType: &ast.TypeName{Name: "ltree"},
	},
	{
		Name: "text2ltree",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "text"},
			},
		},
		ReturnType: &ast.TypeName{Name: "ltree"},
	},
}

func Ltree() *catalog.Schema {
	s := &catalog.Schema{Name: "pg_catalog"}
	s.Funcs = funcsLtree
	return s
}
