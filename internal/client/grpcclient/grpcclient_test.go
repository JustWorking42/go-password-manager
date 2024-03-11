package grpcclient

import (
	"context"
	"testing"

	"github.com/JustWorking42/go-password-manager/internal/client/repository"
	"github.com/JustWorking42/go-password-manager/proto"
	"github.com/golang/mock/gomock"
	"golang.org/x/exp/slices"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestRegister(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := proto.NewMockPassManagerClient(ctrl)

	client := &PassGRPCClient{grpc: mockClient}

	tests := []struct {
		name          string
		login         string
		password      string
		expectedToken string
		mockSetup     func(mockClient *proto.MockPassManagerClient)
	}{
		{
			name:          "Success",
			login:         "testLogin",
			password:      "testPassword",
			expectedToken: "expectedToken",
			mockSetup: func(mockClient *proto.MockPassManagerClient) {
				mockClient.EXPECT().Register(
					gomock.Any(),
					gomock.Any(),
					gomock.Any(),
				).DoAndReturn(func(ctx context.Context, in *proto.Creds, opts ...grpc.CallOption) (*emptypb.Empty, error) {
					opts[0].(grpc.HeaderCallOption).HeaderAddr.Append("authorization", "Bearer expectedToken")
					return nil, nil
				})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.mockSetup(mockClient)
			token, err := client.Register(context.Background(), tt.login, tt.password)

			if err != nil {
				t.Errorf("Register returned an error: %v", err)
			}

			if token != tt.expectedToken {
				t.Errorf("Expected token %q, got %q", tt.expectedToken, token)
			}
		})
	}
}

func TestLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := proto.NewMockPassManagerClient(ctrl)

	client := &PassGRPCClient{grpc: mockClient}

	tests := []struct {
		name          string
		login         string
		password      string
		expectedToken string
		mockSetup     func(mockClient *proto.MockPassManagerClient)
	}{
		{
			name:          "Success",
			login:         "testLogin",
			password:      "testPassword",
			expectedToken: "expectedToken",
			mockSetup: func(mockClient *proto.MockPassManagerClient) {
				mockClient.EXPECT().Login(
					gomock.Any(),
					gomock.Any(),
					gomock.Any(),
				).DoAndReturn(func(ctx context.Context, in *proto.Creds, opts ...grpc.CallOption) (*emptypb.Empty, error) {
					opts[0].(grpc.HeaderCallOption).HeaderAddr.Append("authorization", "Bearer expectedToken")
					return nil, nil
				})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.mockSetup(mockClient)
			token, err := client.Login(context.Background(), tt.login, tt.password)

			if err != nil {
				t.Errorf("Login returned an error: %v", err)
			}

			if token != tt.expectedToken {
				t.Errorf("Expected token %q, got %q", tt.expectedToken, token)
			}
		})
	}
}

func TestAddPassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := proto.NewMockPassManagerClient(ctrl)

	client := &PassGRPCClient{grpc: mockClient}

	tests := []struct {
		name      string
		pass      repository.Password
		mockSetup func(mockClient *proto.MockPassManagerClient)
	}{
		{
			name: "Success",
			pass: repository.Password{
				Name:     "testService",
				Login:    "testLogin",
				Password: "testPassword",
			},
			mockSetup: func(mockClient *proto.MockPassManagerClient) {
				mockClient.EXPECT().AddPassword(
					gomock.Any(),
					gomock.Any(),
				).DoAndReturn(func(ctx context.Context, in *proto.Password, opts ...grpc.CallOption) (*emptypb.Empty, error) {
					return &emptypb.Empty{}, nil
				})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.mockSetup(mockClient)
			err := client.AddPassword(context.Background(), tt.pass)

			if err != nil {
				t.Errorf("AddPassword returned an error: %v", err)
			}
		})
	}
}

func TestGetPassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := proto.NewMockPassManagerClient(ctrl)

	client := &PassGRPCClient{grpc: mockClient}

	tests := []struct {
		name             string
		expectedPassword repository.Password
		mockSetup        func(mockClient *proto.MockPassManagerClient)
	}{
		{
			name: "Success",
			expectedPassword: repository.Password{
				Name:     "testService",
				Login:    "testLogin",
				Password: "testPassword",
			},
			mockSetup: func(mockClient *proto.MockPassManagerClient) {
				mockClient.EXPECT().GetPassword(
					gomock.Any(),
					gomock.Any(),
				).DoAndReturn(func(ctx context.Context, in *proto.GetPasswordRequest, opts ...grpc.CallOption) (*proto.Password, error) {
					return &proto.Password{
						ServiceName:     "testService",
						ServiceLogin:    "testLogin",
						ServicePassword: "testPassword",
					}, nil
				})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.mockSetup(mockClient)
			password, err := client.GetPassword(context.Background(), tt.name)

			if err != nil {
				t.Errorf("GetPassword returned an error: %v", err)
			}

			if password.Name != tt.expectedPassword.Name || password.Login != tt.expectedPassword.Login || password.Password != tt.expectedPassword.Password {
				t.Errorf("Expected password %v, got %v", tt.expectedPassword, password)
			}
		})
	}
}

func TestAddCard(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := proto.NewMockPassManagerClient(ctrl)

	client := &PassGRPCClient{grpc: mockClient}

	tests := []struct {
		name      string
		card      repository.Card
		mockSetup func(mockClient *proto.MockPassManagerClient)
	}{
		{
			name: "Success",
			card: repository.Card{
				CardName:   "testCard",
				CardNumber: "1234567890123456",
				CardCVC:    "123",
				CardDate:   "12/24",
				CardFI:     "Test Bank",
			},
			mockSetup: func(mockClient *proto.MockPassManagerClient) {
				mockClient.EXPECT().AddCard(
					gomock.Any(),
					gomock.Any(),
				).DoAndReturn(func(ctx context.Context, in *proto.Card, opts ...grpc.CallOption) (*emptypb.Empty, error) {
					return &emptypb.Empty{}, nil
				})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.mockSetup(mockClient)
			err := client.AddCard(context.Background(), tt.card)

			if err != nil {
				t.Errorf("AddCard returned an error: %v", err)
			}
		})
	}
}

func TestGetCard(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := proto.NewMockPassManagerClient(ctrl)

	client := &PassGRPCClient{grpc: mockClient}

	tests := []struct {
		name         string
		cardName     string
		expectedCard repository.Card
		mockSetup    func(mockClient *proto.MockPassManagerClient)
	}{
		{
			name:     "Success",
			cardName: "testCard",
			expectedCard: repository.Card{
				CardName:   "testCard",
				CardNumber: "1234567890123456",
				CardCVC:    "123",
				CardDate:   "12/24",
				CardFI:     "Test Bank",
			},
			mockSetup: func(mockClient *proto.MockPassManagerClient) {
				mockClient.EXPECT().GetCard(
					gomock.Any(),
					gomock.Any(),
				).DoAndReturn(func(ctx context.Context, in *proto.GetCardRequest, opts ...grpc.CallOption) (*proto.Card, error) {
					return &proto.Card{
						CardName:   "testCard",
						CardNumber: "1234567890123456",
						CardCVC:    "123",
						CardDate:   "12/24",
						CardFI:     "Test Bank",
					}, nil
				})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.mockSetup(mockClient)
			card, err := client.GetCard(context.Background(), tt.cardName)

			if err != nil {
				t.Errorf("GetCard returned an error: %v", err)
			}

			if card.CardName != tt.expectedCard.CardName || card.CardNumber != tt.expectedCard.CardNumber || card.CardCVC != tt.expectedCard.CardCVC || card.CardDate != tt.expectedCard.CardDate || card.CardFI != tt.expectedCard.CardFI {
				t.Errorf("Expected card %v, got %v", tt.expectedCard, card)
			}
		})
	}
}

func TestAddNote(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := proto.NewMockPassManagerClient(ctrl)

	client := &PassGRPCClient{grpc: mockClient}

	tests := []struct {
		name      string
		note      repository.Note
		mockSetup func(mockClient *proto.MockPassManagerClient)
	}{
		{
			name: "Success",
			note: repository.Note{
				NoteName: "testNote",
				Note:     "This is a test note.",
			},
			mockSetup: func(mockClient *proto.MockPassManagerClient) {
				mockClient.EXPECT().AddNote(
					gomock.Any(),
					gomock.Any(),
				).DoAndReturn(func(ctx context.Context, in *proto.Note, opts ...grpc.CallOption) (*emptypb.Empty, error) {
					return &emptypb.Empty{}, nil
				})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.mockSetup(mockClient)
			err := client.AddNote(context.Background(), tt.note)

			if err != nil {
				t.Errorf("AddNote returned an error: %v", err)
			}
		})
	}
}

func TestGetNote(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := proto.NewMockPassManagerClient(ctrl)

	client := &PassGRPCClient{grpc: mockClient}

	tests := []struct {
		name         string
		noteName     string
		expectedNote repository.Note
		mockSetup    func(mockClient *proto.MockPassManagerClient)
	}{
		{
			name:     "Success",
			noteName: "testNote",
			expectedNote: repository.Note{
				NoteName: "testNote",
				Note:     "This is a test note.",
			},
			mockSetup: func(mockClient *proto.MockPassManagerClient) {
				mockClient.EXPECT().GetNote(
					gomock.Any(),
					gomock.Any(),
				).DoAndReturn(func(ctx context.Context, in *proto.GetNoteRequest, opts ...grpc.CallOption) (*proto.Note, error) {
					return &proto.Note{
						NoteName: "testNote",
						Note:     "This is a test note.",
					}, nil
				})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.mockSetup(mockClient)
			note, err := client.GetNote(context.Background(), tt.noteName)

			if err != nil {
				t.Errorf("GetNote returned an error: %v", err)
			}

			if note.NoteName != tt.expectedNote.NoteName || note.Note != tt.expectedNote.Note {
				t.Errorf("Expected note %v, got %v", tt.expectedNote, note)
			}
		})
	}
}

func TestAddBytes(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := proto.NewMockPassManagerClient(ctrl)

	client := &PassGRPCClient{grpc: mockClient}

	tests := []struct {
		name       string
		binaryData repository.BinaryData
		mockSetup  func(mockClient *proto.MockPassManagerClient)
	}{
		{
			name: "Success",
			binaryData: repository.BinaryData{
				BytesName: "testBytes",
				Value:     []byte("test binary data"),
			},
			mockSetup: func(mockClient *proto.MockPassManagerClient) {
				mockClient.EXPECT().AddBytes(
					gomock.Any(),
					gomock.Any(),
				).DoAndReturn(func(ctx context.Context, in *proto.Bytes, opts ...grpc.CallOption) (*emptypb.Empty, error) {
					return &emptypb.Empty{}, nil
				})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.mockSetup(mockClient)
			err := client.AddBytes(context.Background(), tt.binaryData)

			if err != nil {
				t.Errorf("AddBytes returned an error: %v", err)
			}
		})
	}
}

func TestGetBytes(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := proto.NewMockPassManagerClient(ctrl)

	client := &PassGRPCClient{grpc: mockClient}

	tests := []struct {
		name          string
		bytesName     string
		expectedBytes repository.BinaryData
		mockSetup     func(mockClient *proto.MockPassManagerClient)
	}{
		{
			name:      "Success",
			bytesName: "testBytes",
			expectedBytes: repository.BinaryData{
				BytesName: "testBytes",
				Value:     []byte("test binary data"),
			},
			mockSetup: func(mockClient *proto.MockPassManagerClient) {
				mockClient.EXPECT().GetBytes(
					gomock.Any(),
					gomock.Any(),
				).DoAndReturn(func(ctx context.Context, in *proto.GetBytesRequest, opts ...grpc.CallOption) (*proto.Bytes, error) {
					return &proto.Bytes{
						BytesName: "testBytes",
						Value:     []byte("test binary data"),
					}, nil
				})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.mockSetup(mockClient)
			bytes, err := client.GetBytes(context.Background(), tt.bytesName)

			if err != nil {
				t.Errorf("GetBytes returned an error: %v", err)
			}

			if bytes.BytesName != tt.expectedBytes.BytesName || !slices.Equal(bytes.Value, tt.expectedBytes.Value) {
				t.Errorf("Expected bytes %v, got %v", tt.expectedBytes, bytes)
			}
		})
	}
}
