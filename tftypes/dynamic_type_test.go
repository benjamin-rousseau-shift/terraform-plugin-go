// Copyright IBM Corp. 2020, 2026
// SPDX-License-Identifier: MPL-2.0

package tftypes

import "testing"

type embeddedDynamicType struct{ Type }

func (embeddedDynamicType) Is(Type) bool { return true }

func TestDynamicTypeCheckPreservesIsDispatch(t *testing.T) {
	s := String
	o := Object{AttributeTypes: map[string]Type{"name": String}}
	cases := []Type{String, Number, Bool, DynamicPseudoType, o, List{ElementType: String}, Map{ElementType: String}, Set{ElementType: String}, Tuple{ElementTypes: []Type{String}}, &s, &o, embeddedDynamicType{String}}
	for _, typ := range cases {
		if got, want := isDynamicPseudoType(typ), typ.Is(DynamicPseudoType); got != want {
			t.Fatalf("%T: got %v, want %v", typ, got, want)
		}
	}
}
