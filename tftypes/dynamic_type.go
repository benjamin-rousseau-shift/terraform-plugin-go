// Copyright IBM Corp. 2020, 2026
// SPDX-License-Identifier: MPL-2.0

package tftypes

// isDynamicPseudoType avoids allocating an interface box for the wildcard on
// every built-in conformance check. Non-built-in implementations keep their Is
// dispatch, including embedded implementations and pointer types.
func isDynamicPseudoType(typ Type) bool {
	switch v := typ.(type) {
	case primitive:
		return v.name == DynamicPseudoType.name
	case Object, List, Map, Set, Tuple:
		return false
	default:
		return typ.Is(DynamicPseudoType)
	}
}
