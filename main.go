package main

import (
	"fmt"
	"container/list"
	"net"
	"bufio"
)

var clients *list.List

func main() {
	fmt.Println("Chat server started!")
	clients = list.New()

	server, err := net.Listen("tcp",":8080")
	if err != nil {
		fmt.Printf("Error: %s",err.Error())
		return
	}

	for {
		client, err := server.Accept()
		if err != nil {
			fmt.Printf("Error: %s",err.Error())
			return
		}
		fmt.Printf("Новый пользователь: %s",client.RemoteAddr())
		clients.PushBack(client)

		go handleClient(client)
	}
}

func handleClient(socket net.Conn)  {
	for {
		buffer, err := bufio.NewReader(socket).ReadString('\n')
		if err != nil {
			fmt.Println("Пользователь отключился!")
			socket.Close()
			return
		}

		for i:= clients.Front(); i!=nil;i = i.Next() {
			fmt.Fprint(i.Value.(net.Conn),buffer)
		}
	}
}