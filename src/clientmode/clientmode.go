package clientmode

import (
	"../portconnect"
	"log"
	"time"
	"net"
)

const timeout = 5

/*
	funcName: InitClient
	funcDesc: connect client port with server port
	funcParams: lAddr(string): server <ip:port>
				rAddr(srring): client <ip:port>
*/

func InitClient(lAddr, rAddr string) {
	for {
		var lConn, rConn net.Conn
		var err error
		for {
			lConn, err = net.Dial("tcp", lAddr)
			if err == nil {
				log.Println("[→]", "connect [" + lAddr + "] success.")
				break
			} else {
				log.Println("[x]", "connect target address [" + lAddr + "] faild. retry... ")
				time.Sleep(timeout * time.Second)
			}
		
		}
		for {
			rConn, err = net.Dial("tcp", rAddr)
			if err == nil {
				log.Println("[→]", "connect [" + rAddr + "] success.")
				break
			} else {
				log.Println("[x]", "connect target address [" + rAddr + "] faild. retry...")
				time.Sleep(timeout * time.Second)
			}
		}
		portconnect.Forward(lConn, rConn)
	}
}
