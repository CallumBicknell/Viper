package scanner

import (
	"fmt"
	"net"

	"github.com/CallumBicknell/Viper/pkg/models"
)

type PortScanner struct {
	Host      string
	EveryPort bool
	Top100    bool
}

var Top100Ports = []int{7, 9, 13, 21, 22, 23, 25, 26, 37, 53, 79, 80, 81, 88, 106, 110, 111, 113, 119, 135, 139, 143, 144, 179, 199, 389, 427, 443, 444, 445, 465, 513, 514, 515, 543, 544, 548, 554, 587, 631, 646, 873, 990, 993, 995, 1025, 1026, 1027, 1028, 1029, 1110, 1433, 1720, 1723, 1755, 1900, 2000, 2001, 2049, 2121, 2717, 3000, 3128, 3306, 3389, 3986, 4899, 5000, 5009, 5051, 5060, 5101, 5190, 5357, 5432, 5631, 5666, 5800, 5900, 6000, 6001, 6646, 7070, 8000, 8008, 8009, 8080, 8081, 8443, 8888, 9100, 9999, 10000, 32768, 49152, 49153, 49154, 49155, 49156, 49157}

// Rubbish port scanner :)
func (p *PortScanner) StartPortScanner() ([]models.Port, error) {
	var openPorts []models.Port

	if p.EveryPort {
		for i := 1; i <= 65535; i++ {
			_, err := p.connectToTcpPort(i)
			if err == nil {
				openPorts = append(openPorts, models.Port{
					Number:  i,
					Type:    models.TCP,
					Status:  models.OPEN,
					Service: "Unknown",
				})
			}
		}
	} else if p.Top100 {
		for i := range Top100Ports {
			_, err := p.connectToTcpPort(i)
			if err == nil {
				openPorts = append(openPorts, models.Port{
					Number:  i,
					Type:    models.TCP,
					Status:  models.OPEN,
					Service: "Unknown",
				})
			}
		}
	}

	return openPorts, nil
}

func (p *PortScanner) connectToTcpPort(port int) (*net.TCPConn, error) {
	var conn *net.TCPConn = nil
	servAddr := fmt.Sprintf("%s:%d", p.Host, port)
	tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
	if err != nil {
		goto end
	}

	conn, err = net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		goto end
	}

end:
	conn.Close()
	return conn, nil
}
