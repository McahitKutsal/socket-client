package main

import (
	"bufio"
	"fmt"
	"main/cryption"
	"net"
	"os"
)

const key string = "b3bdcf847fa2e61810470a91328a84d2456335b7cb2ec3b949098ec02179be7a"

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8098")
	messages := make(chan string)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	connbuf := bufio.NewReader(conn)

	go func() {
		for {
			str, err := connbuf.ReadString('@')
			if err != nil {
				break
			}

			if len(str) > 0 {
				messages <- str
			}
		}
	}()
	go func() {
		for {
			select {
			case msg := <-messages:
				client := msg[:10]
				decrypted := cryption.Decrypt(msg[10:], key)
				fmt.Println(client + decrypted)
			}
		}
	}()
	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		encrypted := cryption.Encrypt(text, key)
		encrypted = encrypted + "@"
		fmt.Fprintf(conn, encrypted)
	}
}
