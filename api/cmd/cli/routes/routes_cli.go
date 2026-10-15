package routes

import (
	"finance/internal/logger"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type RouteList struct {
	Method string
	Path   string
	File   string
}

func validatePath(informedPath string) error {
	if _, err := os.Stat(informedPath); err != nil {
		if os.IsNotExist(err) {
			logger.Error(fmt.Sprintf("Caminho: %s não existe", informedPath), err)
			return os.ErrNotExist
		}

		logger.Error(fmt.Sprintf("Erro no caminho: %s", informedPath), err)
		return err
	}

	return nil
}

func readFileRoutes(filePath string) ([]RouteList, error) {
	logger.Info("Lendo arquivo de rotas:", "file", filePath)

	fileSet := token.NewFileSet()

	file, err := parser.ParseFile(
		fileSet,
		filePath,
		nil,
		parser.SkipObjectResolution,
	)
	if err != nil {
		logger.Error("Erro ao analisar o arquivo: ", err)
		return nil, err
	}

	routes := make([]RouteList, 0)

	ast.Inspect(file, func(node ast.Node) bool {
		call, ok := node.(*ast.CallExpr)
		if !ok {
			return true
		}

		selector, ok := call.Fun.(*ast.SelectorExpr)
		if !ok || selector.Sel.Name != "HandleFunc" {
			return true
		}

		if len(call.Args) == 0 {
			return true
		}

		routeLiteral, ok := call.Args[0].(*ast.BasicLit)
		if !ok || routeLiteral.Kind != token.STRING {
			return true
		}

		pattern, err := strconv.Unquote(routeLiteral.Value)
		if err != nil {
			return true
		}

		method, path, found := strings.Cut(pattern, " ")
		if !found {
			path = method
			method = "ANY"
		}

		routes = append(routes, RouteList{
			Method: method,
			Path:   path,
			File:   filePath,
		})

		return true

	})

	return routes, nil
}

func readPathRoutes(rootPath string) ([]RouteList, error) {
	modulesPath := filepath.Join(rootPath, "internal", "modules")

	if err := validatePath(modulesPath); err != nil {
		logger.Error("Erro ao conferir se o caminho de módulos existe:", err)
		return nil, err
	}

	modules, err := os.ReadDir(modulesPath)
	if err != nil {
		logger.Error("Erro ao acessar o caminho dos módulos:", err)
		return nil, err
	}

	routes := make([]RouteList, 0)

	for _, module := range modules {
		if !module.IsDir() {
			continue
		}

		moduleName := module.Name()
		routesPath := filepath.Join(modulesPath, moduleName, "routes")

		if _, err := os.Stat(routesPath); err != nil {
			if os.IsNotExist(err) {
				continue
			}

			logger.Error(fmt.Sprintf("Erro ao acessar as rotas do módulo: %s", moduleName), err)
			return nil, err
		}

		err := filepath.WalkDir(
			routesPath,
			func(path string, entry fs.DirEntry, walkErr error) error {
				if walkErr != nil {
					logger.Error("walkErr", walkErr)
					return walkErr
				}

				if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".go") {
					return nil
				}

				fileRoutes, err := readFileRoutes(path)
				if err != nil {
					return err
				}

				routes = append(routes, fileRoutes...)

				return nil
			},
		)

		if err != nil {
			return nil, err
		}
	}

	return routes, nil
}

func ListRoutes(rootPath string) error {
	routes, err := readPathRoutes(rootPath)
	if err != nil {
		return err
	}

	fmt.Printf("%-8s %-40s\n", "MÉTODO", "ROTA")
	fmt.Printf("%-8s %-40s\n", "--------", "----")

	for _, route := range routes {
		fmt.Printf(
			"%-8s %-40s\n",
			route.Method,
			route.Path,
		)
	}

	return nil
}
