package cmd

import (
	"fmt"
	pb "grpcapp/proto"
	service "grpcapp/service"
	"net"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var (
	port string
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "server cmd for grpcapp",
	Long:  "server cmd for grpcapp",
	Run: func(cmd *cobra.Command, args []string) {
		RunServer()
	},
}

func init() {
	RootCmd.AddCommand(serverCmd)
	serverCmd.Flags().StringVarP(&port, "port", "p", "8080", "port to listen")
}

// RunServer run the server cmd
func RunServer() {
	var (
		l   net.Listener
		err error
	)
	if l, err = net.Listen("tcp", ":"+port); err != nil {
		fmt.Printf("listen error: %v\n", err)
		return
	}
	s := grpc.NewServer()
	pb.RegisterItemServiceServer(s, &service.ItemService{})
	pb.RegisterOrderServiceServer(s, &service.OrderService{})
	fmt.Printf("listening on :%s\n", port)
	if err = s.Serve(l); err != nil {
		fmt.Printf("serve error: %v\n", err)
	}
}
