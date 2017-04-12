package cleanarch

import (
	"reflect"
	"testing"
)

func TestValidator_Validate(t *testing.T) {
	testCases := []struct{
		Path string
		IsValid bool
	}{
		{"tests/valid-simple", true},
		{"tests/invalid-infra-in-domain-import", false},
		{"tests/invalid-app-in-domain-import", false},
		{"tests/invalid-import-between-modules-apps", false},
		{"tests/valid-cross-module-deps", true},
	}

	for _, c := range testCases {
		t.Run(c.Path, func(t *testing.T) {
			validator := NewValidator()
			valid, errors := validator.Validate(c.Path)

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
	testCases := []struct{
		Path                 string
		ExpectedFileMetadata LayerMetadata
	}{
		{
			Path: "/tests/valid-simple/app/payment.go",
			ExpectedFileMetadata: LayerMetadata{
				Module: "valid-simple",
				Layer: LayerApplication,
			},
		},
		{
			Path: "/tests/valid-simple/domain/payment.go",
			ExpectedFileMetadata: LayerMetadata{
				Module: "valid-simple",
				Layer: LayerDomain,
			},
		},
		{
			Path: "/tests/valid-simple/infrastructure/payment.go",
			ExpectedFileMetadata: LayerMetadata{
				Module: "valid-simple",
				Layer: LayerInfrastructure,
			},
		},
		{
			Path: "/tests/valid-simple/interfaces/payment.go",
			ExpectedFileMetadata: LayerMetadata{
				Module: "valid-simple",
				Layer: LayerInterfaces,
			},
		},
		{
			Path: "/tests/valid-simple/app/payments/payment.go",
			ExpectedFileMetadata: LayerMetadata{
				Module: "valid-simple",
				Layer: LayerApplication,
			},
		},
		{
			Path: "/tests/app/domain/payments/payment.go",
			ExpectedFileMetadata: LayerMetadata{
				Module: "app",
				Layer: LayerDomain,
			},
		},
	}

	for _, c := range testCases {
		t.Run(c.Path, func(t *testing.T) {
			metadata := ParsePath(c.Path)

			if !reflect.DeepEqual(metadata, c.ExpectedFileMetadata) {
				t.Errorf("invalid metadata: %+v, expected %+v", metadata, c.ExpectedFileMetadata)
			}
		})
	}
}
