// Code generated by "stringer -type=SAFI"; DO NOT EDIT.

package zebra

import "strconv"

const _SAFI_name = "FRR_ZAPI5_SAFI_UNICASTFRR_ZAPI5_SAFI_MULTICASTFRR_ZAPI5_SAFI_MPLS_VPNFRR_ZAPI5_SAFI_ENCAPFRR_ZAPI5_SAFI_EVPNFRR_ZAPI5_SAFI_LABELED_UNICASTFRR_ZAPI5_SAFI_FLOWSPECFRR_ZAPI5_SAFI_MAX"

var _SAFI_index = [...]uint8{0, 22, 46, 69, 89, 108, 138, 161, 179}

func (i SAFI) String() string {
	i -= 1
	if i >= SAFI(len(_SAFI_index)-1) {
		return "SAFI(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _SAFI_name[_SAFI_index[i]:_SAFI_index[i+1]]
}