package main

import "flag"

var (
	port int
)

func init() {
	flag.IntVar(&port, "port", 8080, "Defines the application HTTP port")
}

func readAllFlags() {
	flag.Parse()
}
