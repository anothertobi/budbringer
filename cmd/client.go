package cmd

import (
	"fmt"
	"log"
	"net"

	"github.com/anothertobi/budbringer/internal"
	"github.com/spf13/cobra"
)

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Starts budbringer in client mode",
	Long:  "Starts budbringer in client mode, sending stdin input and printing all incoming messages",
	Run: func(cmd *cobra.Command, args []string) {
		startClient()
	},
}

var ServerAddress string
var ServerPort int

func init() {
	rootCmd.AddCommand(clientCmd)

	clientCmd.Flags().StringVarP(&ServerAddress, "address", "a", "::1", "The server address to connect to")
	clientCmd.Flags().IntVarP(&ServerPort, "port", "p", 9999, "The server port to connect to")
}

func startClient() {
	serverAddressIP := net.ParseIP(ServerAddress)
	if serverAddressIP == nil {
		log.Fatalf("error parsing address: %s", ServerAddress)
	}

	addr := &net.TCPAddr{
		IP:   serverAddressIP,
		Port: ListenPort,
	}

	input := make(chan []byte)
	go internal.ReadStdin(input)

	dialTCP(addr, input)
}

func dialTCP(addr *net.TCPAddr, input chan []byte) {
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Fatalf("error connecting to %s:%d", addr.IP, addr.Port)
	}
	defer conn.Close()

	fmt.Printf("📤 to %s\n\n", addr)

	internal.HandleConnection(conn, input)
}
