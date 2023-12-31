package myrpc

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func StartService(service interface{}, host string) error {
	err := rpc.Register(service)
	if err != nil {
		return err
	}

	listener, err := net.Listen("tcp", host)
	if err != nil {
		return err
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}

func NewClient(host string) (*rpc.Client, error) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}

	return jsonrpc.NewClient(conn), nil
}
