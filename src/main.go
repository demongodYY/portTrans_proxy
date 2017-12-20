package main 
import (
	"./servermode"
	"./clientmode"
	"flag"
	"os"
)
func main() {
	flagServerMode := flag.Bool("s", false, "use server mode")
	flagClientMode := flag.Bool("c", false, "use client mode")
	flagTransMode := flag.Bool("t", false, "use trans mode")
	flagProxyClientMode := flag.Bool("pc", false, "use proxy client mode")

	flag.Parse()
	switch {
	case *flagServerMode == true:
		if len(os.Args) != 4 {
			flag.Usage()
			os.Exit(1)
		}
		lPort := os.Args[2]
		rPort := os.Args[3]
		servermode.InitServer(lPort,rPort)
	case *flagClientMode == true: 
		if len(os.Args) != 4{
			flag.Usage()
			os.Exit(1)
		}
		lAddr := os.Args[2]
		rAddr := os.Args[3]
		clientmode.InitClient(lAddr, rAddr)
	case *flagProxyClientMode == true:
		if len(os.Args) != 5{
			flag.Usage()
			os.Exit(1)
		}
		lAddr := os.Args[2]
		rAddr := os.Args[3]
		pAddr := os.Args[4]
		clientmode.InitClientThroughPoxy(lAddr, rAddr, pAddr)
	case *flagTransMode == true:
		if len(os.Args) != 4 {
			flag.Usage()
			os.Exit(1)
		}
		lPort := os.Args[2]
		rAddr := os.Args[3]
		servermode.Port2host(lPort, rAddr)
	default:
		flag.Usage()
	}	
}


