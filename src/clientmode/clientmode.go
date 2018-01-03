package clientmode

import (
	"../portconnect"
	"log"
	"time"
	"net"
)

const timeout = 5
const waitTime = 1

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

/*
	funcName: InitClientThroughPoxy
	funcDesc: connect client port with server port
	funcParams: lAddr(string): server <ip:port>
				rAddr(srring): client <ip:port>
				pAddr(string): proxy-server <ip:port>
				pUser(string): proxy-auth-user
				pPass(string): proxy-auth-password
*/

func InitClientThroughPoxy(lAddr, rAddr, pAddr string) {
	for {
		pConn, rConn := createTunnal(lAddr, rAddr, pAddr)
		time.Sleep(waitTime * time.Second)
		portconnect.Forward(pConn, rConn)
	}
}

func createTunnal(lAddr, rAddr, pAddr string) (net.Conn,net.Conn) {
	var rConn,pConn net.Conn
	var err error
	httpHeader := make([]byte, 256)
	httpHeaderLen := 0
	for {
		pConn, err = net.Dial("tcp", pAddr)
		if err == nil {
			pConn.Write([]byte("CONNECT " + lAddr + " HTTP/1.1\r\n\r\n"))
			time.Sleep(waitTime * time.Second)
			for {
				httpHeaderLen, _ = pConn.Read(httpHeader)
				if httpHeaderLen > 0{
					log.Println ("proxy header length: ", httpHeaderLen)
					log.Println (pAddr + ": ", string(httpHeader))
					break
				}
			}
			log.Println("[→]", "connect [" + lAddr + "] success.")
			break
		} else {
			log.Println("[x]", "connect target address [" + pAddr + "] faild. retry... ")
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
	return pConn, rConn
}