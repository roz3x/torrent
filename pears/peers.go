// Package pears ...
// the package will contain the pears implimentation
// and will do the fetch out the pears from
// torrent server response
package pears

import (
	"encoding/binary"
	"net"
)

//Pear the ips and ports of the file
//sharers
type Pear struct {
	IP   net.IP
	Port uint16
}

//GetPears ... gets the pears
func GetPears(payload []byte) ([]Pear, error) {
	pears := []Pear{}

	//Parseing for the pears
	sizePears := len(payload) / 6
	pears = make([]Pear, sizePears)
	for i := 0; i < sizePears; i++ {
		offset := i * 6
		pears[i].IP = net.IP(payload[offset : offset+4])
		pears[i].Port = binary.BigEndian.Uint16(payload[offset+4 : offset+6])
	}
	return pears, nil
}
