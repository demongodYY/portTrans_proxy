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

	flag.Parse()	
	if *flagServerMode {
		if len(os.Args) != 4{
			flag.Usage()
			os.Exit(1)
		}
		lPort := os.Args[2]
		rPort := os.Args[3]
		servermode.InitServer(lPort,rPort)
	}
	if *flagClientMode {
		if len(os.Args) != 4{
			flag.Usage()
			os.Exit(1)
		}
		lAddr := os.Args[2]
		rAddr := os.Args[3]
		clientmode.InitClient(lAddr, rAddr)
	}

}