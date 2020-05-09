package GoServerDetect

import (
	"log"
	"net"
	"strconv"
)

func CreateServer(port int, psk string, res []byte) error {
	pc,err := net.ListenPacket("udp4", ":" + strconv.Itoa(port))
	if err != nil {
		return err
	}
	defer pc.Close()

	for {
		buf := make([]byte, 1024)
		n,addr,err := pc.ReadFrom(buf)
		if err != nil {
			log.Println(err)
		}
		request := string(buf[:n])
		log.Printf("%s sent this: %s\n\n", addr, request)
		if request == psk {
			_, err = pc.WriteTo(res, addr)
			if err != nil {
				log.Println(err)
			}
		}
	}
}