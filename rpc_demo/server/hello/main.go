package main

import (
	"fmt"
	"net"
	"net/rpc"
)

func main() {
	fmt.Println("hello")
	err1 := rpc.RegisterName("helle", new(Helle))
	if err1 != nil {
		fmt.Println(err1)
	}

	listener, err2 := net.Listen("tcp", "127.0.0.1:8800")
	if err2 != nil {
		fmt.Println(err2)
	}

	defer listener.Close()

	for {
		conn, err3 := listener.Accept()
		if err3 != nil {
			fmt.Println(err3)
		}
		rpc.ServeConn(conn)
	}

}

type Helle struct {
}

func (this Helle) SayHelle(req string, res *string) error {
	fmt.Println(req)
	*res = "hello this is helle" + req
	return nil
}
