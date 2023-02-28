package models

const (
	CLOSED   = 0
	OPEN     = 1
	FILTERED = 2
)

const (
	TCP = 0
	UDP = 1
)

type Port struct {
	Number  int
	Type    int
	Status  int
	Service string
}

var OpenPorts []Port
