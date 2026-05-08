package main

import (
	"flag"
	"fmt"
	"os"
	"rapidbinary/internal/server"
	"rapidbinary/internal/shared"
)

func main() {
	mode := flag.String("mode", "", "Mode: serve or send")
	port := flag.Int("port", 8080, "Port for server")
	upload := flag.Bool("upload", false, "Enable uploads")
	useLog := flag.Bool("log", false, "Enable logging")

	flag.Usage = func() {
		shared.PrintHeader()
		fmt.Println("Usage of RapidBinary:")
		flag.PrintDefaults()
	}

	flag.Parse()

	if *mode == "serve" {
		shared.PrintHeader()
		fmt.Printf("%sStarting server on port %d...%s\n", shared.Cyan, *port, shared.Reset)
		server.RunServer(*port, ".", "./uploads", *upload, *useLog)
	} else {
		flag.Usage()
		os.Exit(1)
	}
}
