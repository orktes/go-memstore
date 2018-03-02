package parser

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os/exec"
	"strings"
)

const structComment = "memstore:generate"

type FieldInfo struct {
	Name  string
	Index bool
	Type  string
}

type StructInfo struct {
	StructName string
	Fields     []FieldInfo
}

type Parser struct {
	PkgPath    string
	PkgName    string
	Structs    []StructInfo
	AllStructs bool
}

type visitor struct {
	*Parser

	name     string
	explicit bool
}

func (p *Parser) needType(comments string) bool {
	for _, v := range strings.Split(comments, "\n") {
		if strings.HasPrefix(v, structComment) {
			return true
		}
	}
	return false
}

func (v *visitor) Visit(n ast.Node) (w ast.Visitor) {
	switch n := n.(type) {
	case *ast.Package:
		return v
	case *ast.File:
		v.PkgName = n.Name.String()
		return v

	case *ast.GenDecl:
		v.explicit = v.needType(n.Doc.Text())

		if !v.explicit && !v.AllStructs {
			return nil
		}
		return v
	case *ast.TypeSpec:
		v.name = n.Name.String()

		// Allow to specify non-structs explicitly independent of '-all' flag.
		if v.explicit {
			info := StructInfo{StructName: v.name}
			if st, ok := n.Type.(*ast.StructType); ok {
				for _, field := range st.Fields.List {
					index := field.Tag != nil && strings.Index(field.Tag.Value, "`memstore:\"index\"`") > -1
					info.Fields = append(info.Fields, FieldInfo{
						Name:  field.Names[0].Name,
						Type:  fmt.Sprintf("%s", field.Type),
						Index: index,
					})
				}
			}
			v.Structs = append(v.Structs, info)
			return nil
		}
		return v
	case *ast.StructType:
		v.Structs = append(v.Structs, StructInfo{StructName: v.name})
		return nil
	}
	return nil
}

func (p *Parser) Parse(fname string, isDir bool) error {
	var err error
	if p.PkgPath, err = getPkgPath(fname, isDir); err != nil {
		return err
	}

	fset := token.NewFileSet()
	if isDir {
		packages, err := parser.ParseDir(fset, fname, nil, parser.ParseComments)
		if err != nil {
			return err
		}

		for _, pckg := range packages {
			ast.Walk(&visitor{Parser: p}, pckg)
		}
	} else {
		f, err := parser.ParseFile(fset, fname, nil, parser.ParseComments)
		if err != nil {
			return err
		}

		ast.Walk(&visitor{Parser: p}, f)
	}
	return nil
}

func getDefaultGoPath() (string, error) {
	output, err := exec.Command("go", "env", "GOPATH").Output()
	return string(output), err
}
