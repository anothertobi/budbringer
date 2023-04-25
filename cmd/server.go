package cmd

import (
	"fmt"
	"log"
	"net"

	"github.com/spf13/cobra"

	"github.com/anothertobi/budbringer/internal"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts budbringer in server mode",
	Long:  "Starts budbringer in server mode, printing all incoming messages and sending stdin input back",
	Run: func(cmd *cobra.Command, args []string) {
		startServer()
	},
}

var ListenAddress string
var ListenPort int

func init() {
	rootCmd.AddCommand(serverCmd)

	serverCmd.Flags().StringVarP(&ListenAddress, "address", "a", "0.0.0.0", "The server address to listen on")
	serverCmd.Flags().IntVarP(&ListenPort, "port", "p", 9999, "The server port to listen")
}

func startServer() {
	listenAddressIP := net.ParseIP(ListenAddress)
	if listenAddressIP == nil {
		log.Fatalf("error parsing address: %s", ListenAddress)
	}

	addr := &net.TCPAddr{
		IP:   listenAddressIP,
		Port: ListenPort,
	}

	input := make(chan []byte)
	go internal.ReadStdin(input)

	listenTCP(addr, input)
}

func listenTCP(addr *net.TCPAddr, input chan []byte) {
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatalf("error listening on %s:%d", addr.IP, addr.Port)
	}
	defer listener.Close()

	fmt.Printf("waiting for ✉️  on %s:%d\n\n", addr.IP, addr.Port)

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Printf("error accepting TCP connection: %v", err)
		}

		fmt.Printf("📥 from %s\n\n", conn.RemoteAddr())

		internal.HandleConnection(conn, input)
	}
}
