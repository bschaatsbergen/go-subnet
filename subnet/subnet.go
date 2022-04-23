package subnet

import (
	"math"
	"net"
)

// GetAvailableHostAddresses returns the number of available host addresses in a subnet.
func GetAvailableHostAddressesCount(base *net.IPNet) int {
	mask := base.Mask

	parentLen, addrLen := mask.Size()
	hostLen := addrLen - parentLen

	// The number of available host addresses is the number of host addresses minus the network address and broadcast address.
	return int(math.Pow(2, float64(hostLen)) - 2)
}
