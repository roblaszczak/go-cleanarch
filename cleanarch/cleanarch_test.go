package cleanarch_test

import (
	"github.com/roblaszczak/go-cleanarch/cleanarch"
	"reflect"
	"testing"
)

func TestValidator_Validate(t *testing.T) {
	testCases := []struct {
		Path    string
		IsValid bool
		IgnoreTests bool
	}{
		{"../examples/valid-simple", true, false},
		{"../examples/invalid-infra-in-domain-import", false, false},
		{"../examples/invalid-app-to-domain-import", false, false},
		{"../examples/invalid-cross-module-deps", false, false},
		{"../examples/valid-cross-module-deps", true, false},
		{"../examples/valid-imports-inside-module", true, false},
		{"../examples/invalid-imports-between-submodules", false, false},
		{"../examples/ignored-dirs", true, false},
		{"../examples/ignored-dirs", true, false},
		{"../examples/invalid-infrastructure-to-app-import-in-tests", true, true},
		{"../examples/invalid-infrastructure-to-app-import-in-tests", false, false},
	}

	for _, c := range testCases {
		t.Run(c.Path, func(t *testing.T) {
			validator := cleanarch.NewValidator()
			valid, errors, err := validator.Validate(c.Path, c.IgnoreTests)
			if err != nil {
				t.Fatal(err)
			}

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
