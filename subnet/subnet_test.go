package subnet

import (
	"net"
	"testing"
)

func TestGetAvailableHostAddressesCount(t *testing.T) {
	type Case struct {
		Base     string
		Expected int
	}

	cases := []Case{
		{
			Base:     "10.0.0.0/16",
			Expected: 65534,
		},
		{
			Base:     "192.0.2.0/24",
			Expected: 254,
		},
		{
			Base:     "66.100.50.0/11",
			Expected: 2097150,
		},
	}

	for _, testCase := range cases {
		_, base, _ := net.ParseCIDR(testCase.Base)
		result := GetAvailableHostAddressesCount(base)

		if result != testCase.Expected {
			t.Errorf("Expected %d, got %d", testCase.Expected, result)
		}
	}
}
