package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/devenbhooshan/kitchendb/pkg/server"
	"github.com/devenbhooshan/kitchendb/pkg/sql"
)

var options struct {
	listenAddress   string
	responseCommand string
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage:  %s [options]\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.StringVar(&options.listenAddress, "listen", "127.0.0.1:15432", "Listen address")
	flag.Parse()

	ln, err := net.Listen("tcp", options.listenAddress)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Listening on", ln.Addr())

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Accepted connection from", conn.RemoteAddr())

		b := server.NewPgFortuneBackend(conn, sql.NewKitchenSQLEngine())
		go func() {
			err := b.Run()
			if err != nil {
				log.Println(err)
			}
			log.Println("Closed connection from", conn.RemoteAddr())
		}()
	}
}
