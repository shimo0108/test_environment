package cmd

import (
	"fmt"
	"log"
	"net"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type serverApp struct {
	*cobra.Command
	environment string
}

func newServerApp() *serverApp {
	cmd := &cobra.Command{
		Use:   "server",
		Short: "Test service server",
	}
	app := &serverApp{Command: cmd}
	cmd.Flags().StringVar(&app.environment, "environment", "local", "environment to run")
	app.RunE = func(c *cobra.Command, args []string) error {
		return app.run()
	}
	return app
}

func (s *serverApp) run() error {
	port := 50051
	listenPort, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	reflection.Register(server)
	server.Serve(listenPort)
	return nil
}
