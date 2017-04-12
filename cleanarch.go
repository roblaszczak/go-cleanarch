package cleanarch

import (
	"fmt"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Layer string

var Log = log.New(ioutil.Discard, "[cleanarch] ", log.LstdFlags|log.Lshortfile)

const (
	LayerDomain Layer = "domain"
	LayerApplication Layer = "application"
	LayerInfrastructure Layer = "infrastructure"
	LayerInterfaces Layer = "interfaces"
)

var LayersHierarchy = map[Layer]int {
	LayerDomain: 1,
	LayerApplication: 2,
	LayerInfrastructure: 3,
	LayerInterfaces: 4,
}

var LayersAliases = map[string]Layer{
	// Domain
	"domain": LayerDomain,
	"entities": LayerDomain,

	// Application
	"app": LayerApplication,
	"application": LayerApplication,
	"usecases": LayerApplication,
	"usecase": LayerApplication,
	"use_cases": LayerApplication,

	// Infrastructure
	"infrastructure": LayerInfrastructure,
	"infra": LayerInfrastructure,

	// Interfaces
	"interfaces": LayerInterfaces,
	"interface": LayerInterfaces,
}

// todo - layers aliases

func NewValidator() *Validator {
	filesMetadata := make(map[string]LayerMetadata, 0)
	return &Validator{filesMetadata:filesMetadata}
}

type ValidationError error

type Validator struct {
	filesMetadata map[string]LayerMetadata
}

func (v *Validator) Validate(root string) (bool, []ValidationError) {
	// todo - return erros?
	valid := true

	errors := []ValidationError{}

	filepath.Walk(root, func(path string, fi os.FileInfo, err error) error {
		if fi.IsDir() {
			return nil
		}

		if !strings.HasSuffix(path, ".go") {
			return nil
		}

		if strings.Contains(path, "/vendor/") {
			// todo - better check and flag
			return nil
		}

		if strings.Contains(path, ".glide") {
			// todo - better check
			return nil
		}

		fset := token.NewFileSet()


		f, err := parser.ParseFile(fset, path, nil, parser.ImportsOnly)
		if err != nil {
			panic(err)
		}

		Log.Println("processing: ", path)
		importerMeta := v.fileMetadata(path)
		Log.Printf("metadata: %+v\n", importerMeta)

		if importerMeta.Layer == "" || importerMeta.Module == "" {
			// todo - error from meta parser?
			Log.Printf("WARN: cannot parse metadata for file %s, meta: %+v\n", path, importerMeta)
			return nil
		}

		Log.Println(f.Name, f.Package, f.Scope)
		for _, imp := range f.Imports {
			importPath := imp.Path.Value
			importPath = strings.TrimSuffix(importPath, `"`)
			importPath = strings.TrimPrefix(importPath, `"`)

			importMeta := v.fileMetadata(importPath)
			Log.Printf("import: %s, importMeta: %+v\n", importPath, importMeta)

			if importMeta.Module == importerMeta.Module {
				if LayersHierarchy[importMeta.Layer] > LayersHierarchy[importerMeta.Layer] {
					err := fmt.Errorf(
						"you cannot import %s layer (%s) to %s layer (%s)",
						importMeta.Layer, importPath,
						importerMeta.Layer, path,
					)
					errors = append(errors, err)
					valid = false
				}
			} else if importMeta.Layer != "" {
				if importMeta.Layer != LayerInterfaces && importerMeta.Layer != LayerInfrastructure {
					// todo - better handling
					err := fmt.Errorf(
						"ERROR: trying to import %s layer (%s) to %s layer (%s) between %s and %s modules, you can only import interfaces layer to infrastructure layer",
						importMeta.Layer, importPath,
						importerMeta.Layer, path,
						importMeta.Module, importerMeta.Module,
					)
					errors = append(errors, err)
					valid = false
				}
			}
		}

		return nil
	})

	return valid, errors
}

func (v *Validator) fileMetadata(path string) LayerMetadata {
	if metadata, ok := v.filesMetadata[path]; ok {
		return metadata
	}

	v.filesMetadata[path] = ParsePath(path)
	return v.filesMetadata[path]
}

type LayerMetadata struct {
	Module string
	Layer Layer
}

func ParsePath(path string) LayerMetadata {
	pathParts := strings.Split(path, "/")
	Log.Println(pathParts)

	metadata := LayerMetadata{}

	for i := len(pathParts)-1; i >= 0; i-- {
		pathPart := pathParts[i]

		// we assume that the path upper the layer is module name
		if metadata.Layer != "" {
			metadata.Module = pathPart
			break
		}

		for alias, layer := range LayersAliases {
			if pathPart == alias {
				metadata.Layer = layer
				continue
			}
		}
	}

	return metadata
}