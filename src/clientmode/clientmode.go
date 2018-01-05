package clientmode

import (
	"../portconnect"
	"log"
	"time"
	"net"
	"encoding/base64"
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
		lConn, rConn := createPipe(lAddr, rAddr)
		time.Sleep(waitTime * time.Second)
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

func InitClientThroughPoxy(lAddr, rAddr, pAddr, proxyStr string) {
	for {
		pConn, rConn := createTunnal(lAddr, rAddr, pAddr, proxyStr)
		time.Sleep(waitTime * time.Second)
		portconnect.Forward(pConn, rConn)
	}
}

func createPipe(lAddr, rAddr string ) (net.Conn,net.Conn) {
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
	return lConn, rConn
}


func createTunnal(lAddr, rAddr, pAddr, proxyStr string) (net.Conn,net.Conn) {
	var rConn,pConn net.Conn
	var err error
	httpResHeader := make([]byte, 256)
	httpResHeaderLen := 0
	for {
		pConn, err = net.Dial("tcp", pAddr)
		if err == nil {
			pConn.Write([]byte(createHttpReqHead(proxyStr, lAddr)))
			time.Sleep(waitTime * time.Second)
			for {
				httpResHeaderLen, _ = pConn.Read(httpResHeader)
				if httpResHeaderLen > 0{
					log.Println ("proxy header length: ", httpResHeaderLen)
					log.Println (pAddr + ": ", string(httpResHeader))
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

func createHttpReqHead(proxyStr, lAddr string) string {
	httpReqHeader := "CONNECT " + lAddr + " HTTP/1.1\r\n"
	if len(proxyStr) > 0 {
		proxyData := []byte(proxyStr)
		base64ProxyStr := base64.StdEncoding.EncodeToString(proxyData)
		httpReqHeader += "Proxy-Authorization: Basic " + base64ProxyStr + "\r\n"
	}
	httpReqHeader += "\r\n"
	return httpReqHeader
}