package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	fmt.Println("hello this is client")

	conn, err1 := rpc.Dial("tcp", "127.0.0.1:8800")
	if err1 != nil {
		fmt.Println(err1)
	}
	defer conn.Close()

	var reply string
	err2 := conn.Call("helle.SayHelle", "i am client", &reply)
	// call 拿到的error 即rpc函数的返回值
	if err2 != nil {
		fmt.Println(err2)
	}

	fmt.Println(reply)
}
