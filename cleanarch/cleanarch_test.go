package cleanarch_test

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/roblaszczak/go-cleanarch/cleanarch"
)

func init() {
	cleanarch.Log.SetOutput(os.Stderr)
}

func TestValidator_Validate(t *testing.T) {
	testCases := []struct {
		Path            string
		IsValid         bool
		IgnoreTests     bool
		IgnoredPackages []string
	}{
		{Path: "../examples/valid-simple", IsValid: true},
		{Path: "../examples/invalid-infra-in-domain-import", IsValid: false},
		{Path: "../examples/invalid-app-to-domain-import", IsValid: false},
		{Path: "../examples/invalid-cross-module-deps", IsValid: false},
		{Path: "../examples/valid-cross-module-deps", IsValid: true},
		{Path: "../examples/valid-imports-inside-module", IsValid: true},
		{Path: "../examples/invalid-imports-between-submodules", IsValid: false},
		{Path: "../examples/invalid-imports-between-submodules-2", IsValid: false},
		{Path: "../examples/ignored-dirs", IsValid: true},
		{Path: "../examples/ignored-dirs", IsValid: true},
		{Path: "../examples/invalid-infrastructure-to-app-import-in-tests", IsValid: true, IgnoreTests: true},
		{Path: "../examples/invalid-infrastructure-to-app-import-in-tests", IsValid: false},
		{
			Path:            "../examples/ignore-package",
			IsValid:         true,
			IgnoredPackages: []string{"github.com/roblaszczak/go-cleanarch/examples/ignore-package/app"},
		},
	}

	for _, c := range testCases {
		t.Run(c.Path, func(t *testing.T) {
			validator := cleanarch.NewValidator()
			valid, errors, err := validator.Validate(c.Path, c.IgnoreTests, c.IgnoredPackages)
			if err != nil {
				t.Fatal(err)
			}

			fmt.Println("errors: ", errors)

			if valid != c.IsValid {
				t.Errorf("path %s should be %t, but is %t", c.Path, c.IsValid, valid)
			}
			if !c.IsValid && len(errors) == 0 {
				t.Error("module is invalid, but errors are empty")
			}
		})
	}
}

func TestParseLayerMetadata(t *testing.T) {
	testCases := []struct {
		Path                 string
		ExpectedFileMetadata cleanarch.LayerMetadata
	}{
		{
			Path: "/tests/valid-simple/app/payment.go",
			ExpectedFileMetadata: cleanarch.LayerMetadata{
				Module: "valid-simple",
				Layer:  cleanarch.LayerApplication,
			},
		},
		{
			Path: "/tests/valid-simple/domain/payment.go",
			ExpectedFileMetadata: cleanarch.LayerMetadata{
				Module: "valid-simple",
				Layer:  cleanarch.LayerDomain,
			},
		},
		{
			Path: "/tests/valid-simple/infrastructure/payment.go",
			ExpectedFileMetadata: cleanarch.LayerMetadata{
				Module: "valid-simple",
				Layer:  cleanarch.LayerInfrastructure,
			},
		},
		{
			Path: "/tests/valid-simple/interfaces/payment.go",
			ExpectedFileMetadata: cleanarch.LayerMetadata{
				Module: "valid-simple",
				Layer:  cleanarch.LayerInterfaces,
			},
		},
		{
			Path: "/tests/valid-simple/app/payments/payment.go",
			ExpectedFileMetadata: cleanarch.LayerMetadata{
				Module: "valid-simple",
				Layer:  cleanarch.LayerApplication,
			},
		},
		{
			Path: "/tests/app/domain/payments/payment.go",
			ExpectedFileMetadata: cleanarch.LayerMetadata{
				Module: "app",
				Layer:  cleanarch.LayerDomain,
			},
		},
	}

	for _, c := range testCases {
		t.Run(c.Path, func(t *testing.T) {
			metadata := cleanarch.ParseLayerMetadata(c.Path)

			if !reflect.DeepEqual(metadata, c.ExpectedFileMetadata) {
				t.Errorf("invalid metadata: %+v, expected %+v", metadata, c.ExpectedFileMetadata)
			}
		})
	}
}
