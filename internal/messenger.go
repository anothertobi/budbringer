package internal

import (
	"context"
	"fmt"
	"net"
)

// HandleConnection handles incoming and outgoing TCP connections for the messenger
func HandleConnection(conn *net.TCPConn, input chan []byte) {
	ctx, cancel := context.WithCancelCause(context.Background())
	defer cancel(nil)

	go HandleAnswer(ctx, conn, input)

	for {
		buff := make([]byte, 1000000)
		buffLen, err := conn.Read(buff)
		if err != nil {
			fmt.Printf("error reading from TCP connection: %v\n", err)
			cancel(err)
			return
		}

		fmt.Print(string(buff[0:buffLen]))
	}
}

// HandleAnswer handles answers for the messenger
func HandleAnswer(ctx context.Context, conn *net.TCPConn, input chan []byte) {
	for {
		select {
		case answer := <-input:
			_, err := conn.Write(answer)
			if err != nil {
				fmt.Println(err)
			}
		case <-ctx.Done():
			return
		}
	}
}
