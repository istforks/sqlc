// Code generated by sqlc-pg-gen. DO NOT EDIT.

package contrib

import (
	"github.com/istforks/sqlc/internal/sql/ast"
	"github.com/istforks/sqlc/internal/sql/catalog"
)

var funcsSslinfo = []*catalog.Function{
	{
		Name:       "ssl_cipher",
		Args:       []*catalog.Argument{},
		ReturnType: &ast.TypeName{Name: "text"},
	},
	{
		Name:       "ssl_client_cert_present",
		Args:       []*catalog.Argument{},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name:       "ssl_client_dn",
		Args:       []*catalog.Argument{},
		ReturnType: &ast.TypeName{Name: "text"},
	},
	{
		Name: "ssl_client_dn_field",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "text"},
			},
		},
		ReturnType: &ast.TypeName{Name: "text"},
	},
	{
		Name:       "ssl_client_serial",
		Args:       []*catalog.Argument{},
		ReturnType: &ast.TypeName{Name: "numeric"},
	},
	{
		Name:       "ssl_extension_info",
		Args:       []*catalog.Argument{},
		ReturnType: &ast.TypeName{Name: "record"},
	},
	{
		Name:       "ssl_is_used",
		Args:       []*catalog.Argument{},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name:       "ssl_issuer_dn",
		Args:       []*catalog.Argument{},
		ReturnType: &ast.TypeName{Name: "text"},
	},
	{
		Name: "ssl_issuer_field",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "text"},
			},
		},
		ReturnType: &ast.TypeName{Name: "text"},
	},
	{
		Name:       "ssl_version",
		Args:       []*catalog.Argument{},
		ReturnType: &ast.TypeName{Name: "text"},
	},
}

func Sslinfo() *catalog.Schema {
	s := &catalog.Schema{Name: "pg_catalog"}
	s.Funcs = funcsSslinfo
	return s
}
