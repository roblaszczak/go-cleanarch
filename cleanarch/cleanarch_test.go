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
	}{
		{"../examples/valid-simple", true},
		{"../examples/invalid-infra-in-domain-import", false},
		{"../examples/invalid-app-to-domain-import", false},
		{"../examples/invalid-cross-module-deps", false},
		{"../examples/valid-cross-module-deps", true},
	}

	for _, c := range testCases {
		t.Run(c.Path, func(t *testing.T) {
			validator := cleanarch.NewValidator()
			valid, errors, err := validator.Validate(c.Path)
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

func TestParsePath(t *testing.T) {
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
			metadata := cleanarch.ParsePath(c.Path)

			if !reflect.DeepEqual(metadata, c.ExpectedFileMetadata) {
				t.Errorf("invalid metadata: %+v, expected %+v", metadata, c.ExpectedFileMetadata)
			}
		})
	}
}
