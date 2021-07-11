package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	fileName := flag.String("file", "", "имя файла для парсинга")
	constType := flag.String("type", "", "тип константы")
	flag.Parse()
	if *fileName == "" || *constType == "" {
		flag.PrintDefaults()
		os.Exit(2)
	}

	fmt.Printf("start generator at file %s\n", *fileName)
	fset := token.NewFileSet()
	fpath, err := filepath.Abs(*fileName)
	if err != nil {
		log.Fatalf("wrong file %s %v", os.Args[1], err)
	}
	f, err := parser.ParseFile(fset, fpath, nil, 0)
	if err != nil {
		log.Fatalf("parse %s error %v", fpath, err)
	}

	result := map[string]int64{}

	for _, decl := range f.Decls {
		switch decl := decl.(type) {
		case *ast.GenDecl:
			switch decl.Tok {
			case token.CONST:
				for _, spec := range decl.Specs {
					vspec := spec.(*ast.ValueSpec)
					if len(vspec.Values) == 0 {
						continue
					}
					if vspec.Type == nil {
						continue
					}
					var typ string
					if ident, ok := vspec.Type.(*ast.Ident); ok {
						typ = ident.Name
					} else {
						continue
					}
					if typ != *constType {
						continue
					}
					// находим список всех констант
					for i, name := range vspec.Names {
						lit := vspec.Values[i].(*ast.BasicLit)
						val, err := strconv.Atoi(lit.Value)
						if err != nil {
							log.Fatalf("cannot get const %s value from %s", name.Name, lit.Value)
						}
						result[name.Name] = int64(val)
						fmt.Println(typ, val)
					}
				}
			}
		}
	}
	genFile, err := os.Create("main_gen.go")
	if err != nil {
		log.Fatalf("cannot create file %v", err)
	}

	src := `
package main
var statusesGenerated = map[string]OrderStatus{}
func getStatusesV2() map[string]OrderStatus {
	return statusesGenerated
}
`
	genFile.WriteString(src)
	genFile.WriteString(fmt.Sprintf("func init() {\n"))
	for k, v := range result {
		genFile.WriteString(fmt.Sprintf("\tstatusesGenerated[\"%s\"] = OrderStatus(%d)\n", k, v))
	}
	genFile.WriteString("}\n")
	genFile.Close()

}
