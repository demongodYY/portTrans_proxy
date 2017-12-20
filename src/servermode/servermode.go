package servermode

import (
	"../portconnect"
	"log"
	"time"
	"net"
)
const timeout = 5

/*
	funcName: InitServer
	funcDesc: create server to listen local && remote connect
	funcParams: lPort(string): local <port>
				rPort(srring): remote <port>
*/

func InitServer(lPort, rPort string) {
	lSock := createListenPort("0.0.0.0:"+lPort)
	rSock := createListenPort("0.0.0.0:"+rPort)
	log.Println("[√]", "listen port:", lPort, "and", rPort, "success. waiting for client...")
	for {
		lConn := accept(lSock)
		rConn := accept(rSock)
		if lConn == nil || rConn == nil {
			log.Println("[x]", "accept client faild. retry in ", timeout, " seconds. ")
			time.Sleep(timeout * time.Second)
			continue
		}

		portconnect.Forward(lConn, rConn)
	}
}

func createListenPort (strIp string) net.Listener{
	listenPort, err := net.Listen("tcp", strIp)
	if err!= nil {
		log.Fatalln("[x]", "listen address [" + strIp + "] faild.")
	}
	log.Println("[√]", "start listen at address:[" + strIp + "]")
	return listenPort
}

func accept(listener net.Listener) net.Conn {
	conn, err := listener.Accept()
	if err != nil {
		log.Println("[x]", "accept connect ["+conn.RemoteAddr().String()+"] faild.", err.Error())
		return nil
	}
	log.Println("[√]", "accept a new client. remote address:["+conn.RemoteAddr().String()+"], local address:["+conn.LocalAddr().String()+"]")
	return conn
}
