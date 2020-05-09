package GoServerDetect

import (
	"time"
	"net"
	"strconv"
)

func DiscoverServer(port int, psk string) ([]byte, error) {
	ServerAddr,err := net.ResolveUDPAddr("udp","0.0.0.0:" + strconv.Itoa(port))
	if err != nil {
		return []byte{}, err
	}
	LocalAddr, err := net.ResolveUDPAddr("udp", ":0")
	if err != nil {
		return []byte{}, err
	}
	Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
	if err != nil {
		return []byte{}, err
	}
	defer Conn.Close()

	_, err = Conn.Write([]byte(psk))
	if err != nil {
		return []byte{}, err
	}
	time.Sleep(time.Second * 1)

	buf := make([]byte, 1024)
	n,_,err := Conn.ReadFrom(buf)
	if err != nil {
		return []byte{}, err
	}
	Conn.Close()
	return buf[:n], nil
}