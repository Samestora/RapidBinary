package main

import (
	"flag"
	"fmt"
	"os"
	"rapidbinary/internal/server"
)

func main() {
	mode := flag.String("mode", "", "Mode: serve or send")
	port := flag.Int("port", 8080, "Port for server")
	upload := flag.Bool("upload", false, "Enable uploads")
	useLog := flag.Bool("log", false, "Enable logging")

	flag.Usage = func() {
		fmt.Println("RapidBinary Tool")
		flag.PrintDefaults()
	}

	flag.Parse()

	if *mode == "serve" {
		server.RunServer(*port, ".", "./uploads", *upload, *useLog)
	} else {
		flag.Usage()
		os.Exit(1)
	}
}
