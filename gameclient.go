package main

import (
    "net"
    "os"
	"bufio"
	"fmt"
	"strings"
)

var maze *Maze

func main() {
    conn, _ := net.Dial("tcp", ":6666")
	client := NewClient(conn)
	reader := bufio.NewReader(os.Stdin)
	input := ""
	
	go func(client *Client) {
		for {
			select{
			case in := <- client.incoming:
				fmt.Println("INCOMING:")
				fmt.Println(in)
				maze = MazeFromJSON(in)
				maze.Print()
			}
		}
	} (client)
	
	for {
	    fmt.Println("Input:")
	    input, _ = reader.ReadString('\n')
	    input = strings.TrimSpace(input)
		fmt.Println("bef sending")
		client.outgoing <- input
		fmt.Println("aft sending")	
    }
	
    
}