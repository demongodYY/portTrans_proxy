package portconnect

import (
	// "io/ioutil"
	"io"
	"log"
	"net"
	"sync"
)

func Forward(conn1 net.Conn, conn2 net.Conn) {
	log.Printf("[+] start transmit. [%s],[%s] <-> [%s],[%s] \n", conn1.LocalAddr().String(), conn1.RemoteAddr().String(), conn2.LocalAddr().String(), conn2.RemoteAddr().String())
	var wg sync.WaitGroup
	// wait tow goroutines
	wg.Add(2)
	go connCopy(conn1, conn2, &wg)
	go connCopy(conn2, conn1, &wg)
	//blocking when the wg is locked
	wg.Wait()
}

func connCopy(conn1 net.Conn, conn2 net.Conn, wg *sync.WaitGroup) {
	log.Println("[←!!!]", "begin the connect at local:["+conn1.LocalAddr().String()+"] and remote:["+conn1.RemoteAddr().String()+"]")
	io.Copy(conn1, conn2)
	conn1.Close()
	log.Println("[←]", "close the connect at local:["+conn1.LocalAddr().String()+"] and remote:["+conn1.RemoteAddr().String()+"]")
	wg.Done()
}