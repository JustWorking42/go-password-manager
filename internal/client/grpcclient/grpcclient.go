package grpcclient

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/JustWorking42/go-password-manager/internal/client/repository"
	"github.com/JustWorking42/go-password-manager/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

type PassGRPCClient struct {
	grpc proto.PassManagerClient
	conn *grpc.ClientConn
}

func InitAndGetPassGRPCClient(ctx context.Context, config Config) (*PassGRPCClient, error) {

	pemServerCA, err := os.ReadFile(config.CACertPath)
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}

	clientCert, err := tls.LoadX509KeyPair(config.CertPath, config.KeyPath)
	if err != nil {
		return nil, err
	}

	configs := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      certPool,
	}

	conn, err := grpc.DialContext(ctx, config.Address, grpc.WithTransportCredentials(credentials.NewTLS(configs)))
	if err != nil {
		return nil, err
	}

	return &PassGRPCClient{grpc: proto.NewPassManagerClient(conn), conn: conn}, nil
}

func (c *PassGRPCClient) Close() {
	c.conn.Close()
}

func (c *PassGRPCClient) Register(ctx context.Context, login, password string) (string, error) {
	var header metadata.MD = make(metadata.MD)
	_, err := c.grpc.Register(ctx, &proto.Creds{Login: login, Password: password}, grpc.Header(&header))
	if err != nil {
		return "", err
	}

	token, ok := header["authorization"]
	if !ok {
		return "", errors.New("authorization header not found")
	}

	jwt := strings.TrimPrefix(token[0], "Bearer ")

	return jwt, nil
}

func (c *PassGRPCClient) Login(ctx context.Context, login, password string) (string, error) {
	var header metadata.MD = make(metadata.MD)
	_, err := c.grpc.Login(ctx, &proto.Creds{Login: login, Password: password}, grpc.Header(&header))
	if err != nil {
		return "", err
	}

	token, ok := header["authorization"]
	if !ok {
		return "", errors.New("authorization header not found")
	}

	jwt := strings.TrimPrefix(token[0], "Bearer ")

	return jwt, nil
}

func (c *PassGRPCClient) AddPassword(ctx context.Context, pass repository.Password) error {

	_, err := c.grpc.AddPassword(ctx, &proto.Password{
		ServiceName:     pass.Name,
		ServiceLogin:    pass.Login,
		ServicePassword: pass.Password,
	})
	if err != nil {
		return err
	}

	return nil
}

func (c *PassGRPCClient) GetPassword(ctx context.Context, name string) (repository.Password, error) {
	pass, err := c.grpc.GetPassword(ctx, &proto.GetPasswordRequest{ServiceName: name})
	if err != nil {
		return repository.Password{}, err
	}

	return repository.Password{
		Name:     pass.ServiceName,
		Login:    pass.ServiceLogin,
		Password: pass.ServicePassword,
	}, nil

}

func (c *PassGRPCClient) AddCard(ctx context.Context, card repository.Card) error {
	_, err := c.grpc.AddCard(ctx, &proto.Card{
		CardName:   card.CardName,
		CardNumber: card.CardNumber,
		CardCVC:    card.CardCVC,
		CardDate:   card.CardDate,
		CardFI:     card.CardFI,
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *PassGRPCClient) GetCard(ctx context.Context, cardName string) (repository.Card, error) {
	card, err := c.grpc.GetCard(ctx, &proto.GetCardRequest{CardName: cardName})
	if err != nil {
		return repository.Card{}, err
	}
	return repository.Card{
		CardName:   card.CardName,
		CardNumber: card.CardNumber,
		CardCVC:    card.CardCVC,
		CardDate:   card.CardDate,
		CardFI:     card.CardFI,
	}, nil
}

func (c *PassGRPCClient) AddNote(ctx context.Context, note repository.Note) error {
	_, err := c.grpc.AddNote(ctx, &proto.Note{
		NoteName: note.NoteName,
		Note:     note.Note,
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *PassGRPCClient) GetNote(ctx context.Context, noteName string) (repository.Note, error) {
	note, err := c.grpc.GetNote(ctx, &proto.GetNoteRequest{NoteName: noteName})
	if err != nil {
		return repository.Note{}, err
	}
	return repository.Note{
		NoteName: note.NoteName,
		Note:     note.Note,
	}, nil
}

func (c *PassGRPCClient) AddBytes(ctx context.Context, binaryData repository.BinaryData) error {
	_, err := c.grpc.AddBytes(ctx, &proto.Bytes{
		BytesName: binaryData.BytesName,
		Value:     binaryData.Value,
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *PassGRPCClient) GetBytes(ctx context.Context, bytesName string) (repository.BinaryData, error) {
	bytes, err := c.grpc.GetBytes(ctx, &proto.GetBytesRequest{BytesName: bytesName})
	if err != nil {
		return repository.BinaryData{}, err
	}
	return repository.BinaryData{
		BytesName: bytes.BytesName,
		Value:     bytes.Value,
	}, nil
}
