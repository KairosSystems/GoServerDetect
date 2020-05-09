package GoServerDetect

import (
	"encoding/json"
	"log"
	"net"
	"strconv"
)

type ServerResponse struct {
	Ip string `json:"server"`
	Data []byte `json:"data"`
}

func CreateServer(port int, psk string, response []byte) error {
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
		if request == psk {
			log.Printf("%s validated\n", addr)
			data, err := json.Marshal(&ServerResponse{
				Ip:   addr.(*net.UDPAddr).IP.String(),
				Data: response,
			})
			_, err = pc.WriteTo(data, addr)
			if err != nil {
				log.Println(err)
			} else {
				log.Printf("%s response returned\n", addr)
			}
		} else {
			log.Printf("%s with invalid payload %s\n", addr, request)
		}
	}
}
