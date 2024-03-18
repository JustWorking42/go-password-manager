// Package credintails provide function to create TransportCredentials for use with gRPC.
package credintails

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"

	"github.com/JustWorking42/go-password-manager/internal/server/grpcserver"
	"google.golang.org/grpc/credentials"
)

// Credintails creates a TransportCredentials for use with gRPC.
func Credentials(serverConfig grpcserver.Config) (credentials.TransportCredentials, error) {
	pemClientCA, err := os.ReadFile(serverConfig.CACertPath)
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemClientCA) {
		return nil, fmt.Errorf("failed to add client CA's certificate")
	}

	serverCert, err := tls.LoadX509KeyPair(serverConfig.CertPath, serverConfig.KeyPath)
	if err != nil {
		return nil, err
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	}

	return credentials.NewTLS(config), nil
}
