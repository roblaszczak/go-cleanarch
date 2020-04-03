package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/roblaszczak/go-cleanarch/cleanarch"
)

var domainAliases = []string{"domain", "entities"}
var applicationAliases = []string{"app", "application", "usecases", "usecase", "use_cases"}
var interfacesAliases = []string{"interfaces", "interface", "adapters", "adapter"}
var infrastructureAliases = []string{"infrastructure", "infra"}

func main() {
	ignoredPackages := sliceFlag{}

	ignoreTests := flag.Bool("ignore-tests", false, "if flag is passed *_test.go files will be not checked")
	debug := flag.Bool("debug", false, "debug mode")
	flag.Var(
		&ignoredPackages,
		"ignore-package",
		"provided packages can be imported to any layer, "+
			"for example you can use`-ignore-package github.com/roblaszczak/go-cleanarch/infrastructure` to import "+
			"this package to the domain",
	)
	domain := flag.String("domain", "", "name of the domain layer")
	application := flag.String("application", "", "name of the application layer")
	interfaces := flag.String("interfaces", "", "name of the interfaces layer")
	infrastructure := flag.String("infrastructure", "", "name of the infrastructure layer")

	flag.Parse()
	var path string

	aliases := make(map[string]cleanarch.Layer)

	addAliases(aliases, *domain, domainAliases, cleanarch.LayerDomain)
	addAliases(aliases, *application, applicationAliases, cleanarch.LayerApplication)
	addAliases(aliases, *interfaces, interfacesAliases, cleanarch.LayerInterfaces)
	addAliases(aliases, *infrastructure, infrastructureAliases, cleanarch.LayerInfrastructure)

	if *debug {
		cleanarch.Log.SetOutput(os.Stderr)
	}

	if len(flag.Args()) > 1 {
		path = flag.Args()[1]
	} else {
		var err error
		path, err = filepath.Abs(flag.Arg(0))
		if err != nil {
			panic(err)
		}
	}

	fmt.Printf("[cleanarch] checking %s\n", path)

	validator := cleanarch.NewValidator(aliases)
	isValid, errors, err := validator.Validate(path, *ignoreTests, ignoredPackages)
	if err != nil {
		panic(err)
	}

	if !isValid {
		for _, err := range errors {
			fmt.Println(err.Error())
		}

		fmt.Println("Uncle Bob is not happy.")
		os.Exit(1)
	}

	os.Exit(0)
}

func addAliases(aliases map[string]cleanarch.Layer, name string, names []string, layer cleanarch.Layer) {
	if len(name) > 0 {
		aliases[name] = layer
	} else {
		for _, n := range names {
			aliases[n] = layer
		}
	}
}
