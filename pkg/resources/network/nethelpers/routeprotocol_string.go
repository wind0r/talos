// Code generated by "stringer -type=RouteProtocol -linecomment"; DO NOT EDIT.

package nethelpers

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ProtocolUnspec-0]
	_ = x[ProtocolRedirect-1]
	_ = x[ProtocolKernel-2]
	_ = x[ProtocolBoot-3]
	_ = x[ProtocolStatic-4]
}

const _RouteProtocol_name = "unspecredirectkernelbootstatic"

var _RouteProtocol_index = [...]uint8{0, 6, 14, 20, 24, 30}

func (i RouteProtocol) String() string {
	if i >= RouteProtocol(len(_RouteProtocol_index)-1) {
		return "RouteProtocol(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _RouteProtocol_name[_RouteProtocol_index[i]:_RouteProtocol_index[i+1]]
}