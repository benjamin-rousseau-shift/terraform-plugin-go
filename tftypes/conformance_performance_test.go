// Copyright IBM Corp. 2020, 2026
// SPDX-License-Identifier: MPL-2.0

package tftypes

import (
	"testing"
)

func conformanceTree(depth int) Object {
	o := Object{AttributeTypes: map[string]Type{"title": String}}
	if depth > 0 {
		o.AttributeTypes["rules"] = List{ElementType: conformanceTree(depth - 1)}
	}
	return o
}
func TestUsableAsDoesNotAllocateForBuiltInTypes(t *testing.T) {
	typ := conformanceTree(10)
	var other Type = typ
	a := testing.AllocsPerRun(20, func() {
		if !typ.UsableAs(other) {
			t.Fatal("identical types must remain usable")
		}
	})
	if a != 0 {
		t.Fatalf("built-in type conformance allocates %.0f unnecessary dynamic-type interface boxes", a)
	}
}
func BenchmarkUsableAsDeepObject(b *testing.B) {
	typ := conformanceTree(10)
	var other Type = typ
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if !typ.UsableAs(other) {
			b.Fatal("types unexpectedly differ")
		}
	}
}
