package param

import (
	"fmt"
	"reflect"
)

var paramUnionType = reflect.TypeOf(APIUnion{})

// VariantFromUnion can be used to extract the present variant from a param union type.
// A param union type is a struct with an embedded field of [APIUnion].
func VariantFromUnion(u reflect.Value) (any, error) {
	if u.Kind() == reflect.Ptr {
		u = u.Elem()
	}

	if u.Kind() != reflect.Struct {
		return nil, fmt.Errorf("param: cannot extract variant from non-struct union")
	}

	isUnion := false
	nVariants := 0
	variantIdx := -1
	for i := 0; i < u.NumField(); i++ {
		if !u.Field(i).IsZero() {
			nVariants++
			variantIdx = i
		}
		if u.Field(i).Type() == paramUnionType {
			isUnion = u.Type().Field(i).Anonymous
		}
	}

	if !isUnion {
		return nil, fmt.Errorf("param: cannot extract variant from non-union")
	}

	if nVariants > 1 {
		return nil, fmt.Errorf("param: cannot extract variant from union with multiple variants")
	}

	if nVariants == 0 {
		return nil, fmt.Errorf("param: cannot extract variant from union with no variants")
	}

	return u.Field(variantIdx).Interface(), nil
}

// namePrinter is an interface for types that have a Name method
type namePrinter interface {
	Name() string
}

// Ensure ToolUnionUnionParam implements namePrinter
var _ namePrinter = (*ToolUnionUnionParam)(nil)

// ToolUnionUnionParam defines the interface that tool parameters must implement
type ToolUnionUnionParam interface {
	implementsToolUnionUnionParam()
	Name() string
}
