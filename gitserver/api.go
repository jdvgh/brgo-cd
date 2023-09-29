package gitserver

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	// serviceName needs to be egual to the k8s service name or the istio virtualservice name.
	serviceName = "clonerepo"
	// externalServicePort needs to be egual to the k8s service port or the istio virtualservice port.
	externalServicePort = ":50051"
	address             = serviceName + externalServicePort
)

type Client struct {
	client CloneRepoServiceClient
	conn   *grpc.ClientConn
}

// NewClient creates a client for the git service using the default address.
func NewClient() (*Client, error) {
	return NewClientWithAddress(address)
}

// NewClientWithAddress creates a client for the git service using the specified address.
func NewClientWithAddress(address string) (*Client, error) {
	dialOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	log.Printf("Setting up grpc.Dial to %s", address)
	conn, err := grpc.Dial(address, dialOpts...)
	if err != nil {
		log.Printf("Setting up grpc.Dial to %s failed: %v", address, err)
		return nil, err
	}

	client := NewCloneRepoServiceClient(conn)
	log.Printf("Setting up NewCloneRepoServiceClient  %v conn: %v", client, conn)
	return &Client{
		client: client,
		conn:   conn,
	}, nil
}

// Close closes the connection to the service.
func (c Client) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

func (c Client) CloneRepo(ctx context.Context, req *CloneRepoRequest) (*CloneRepoResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*60)
	defer cancel()
	return c.client.CloneRepo(ctx, req)
}
