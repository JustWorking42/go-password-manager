package grpcclient

import (
	"context"
	"errors"
	"testing"

	"github.com/JustWorking42/go-password-manager/internal/client/repository"
	"github.com/JustWorking42/go-password-manager/proto"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
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
		wantErr       bool
		mockSetup     func(mockClient *proto.MockPassManagerClient)
	}{
		{
			name:          "Success",
			login:         "testLogin",
			password:      "testPassword",
			expectedToken: "expectedToken",
			wantErr:       false,
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
		{
			name:     "ErrorNoAuthorizationHeader",
			login:    "testLogin",
			password: "testPassword",
			wantErr:  true,
			mockSetup: func(mockClient *proto.MockPassManagerClient) {
				mockClient.EXPECT().Register(
					gomock.Any(),
					gomock.Any(),
					gomock.Any(),
				).DoAndReturn(func(ctx context.Context, in *proto.Creds, opts ...grpc.CallOption) (*emptypb.Empty, error) {
					return nil, nil
				})
			},
		},
		{
			name:     "ErrorGRPC",
			login:    "testLogin",
			password: "testPassword",
			wantErr:  true,
			mockSetup: func(mockClient *proto.MockPassManagerClient) {
				mockClient.EXPECT().Register(
					gomock.Any(),
					gomock.Any(),
					gomock.Any(),
				).DoAndReturn(func(ctx context.Context, in *proto.Creds, opts ...grpc.CallOption) (*emptypb.Empty, error) {
					return nil, errors.New("test error")
				})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup(mockClient)
			token, err := client.Register(context.Background(), tt.login, tt.password)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedToken, token)
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
		wantErr       bool
		mockSetup     func(mockClient *proto.MockPassManagerClient)
	}{
		{
			name:          "Success",
			login:         "testLogin",
			password:      "testPassword",
			expectedToken: "expectedToken",
			wantErr:       false,
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
		{
			name:     "ErrorNoAuthorizationHeader",
			login:    "testLogin",
			password: "testPassword",
			wantErr:  true,
			mockSetup: func(mockClient *proto.MockPassManagerClient) {
				mockClient.EXPECT().Login(
					gomock.Any(),
					gomock.Any(),
					gomock.Any(),
				).DoAndReturn(func(ctx context.Context, in *proto.Creds, opts ...grpc.CallOption) (*emptypb.Empty, error) {
					return nil, nil
				})
			},
		},
		{
			name:     "ErrorGRPC",
			login:    "testLogin",
			password: "testPassword",
			wantErr:  true,
			mockSetup: func(mockClient *proto.MockPassManagerClient) {
				mockClient.EXPECT().Login(
					gomock.Any(),
					gomock.Any(),
					gomock.Any(),
				).DoAndReturn(func(ctx context.Context, in *proto.Creds, opts ...grpc.CallOption) (*emptypb.Empty, error) {
					return nil, errors.New("test error")
				})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup(mockClient)
			token, err := client.Login(context.Background(), tt.login, tt.password)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedToken, token)
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
		wantErr   bool
		pass      repository.Password
		mockSetup func(mockClient *proto.MockPassManagerClient)
	}{
		{
			name:    "Success",
			wantErr: false,
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
		{
			name: "Error",
			pass: repository.Password{
				Name:     "testService",
				Login:    "testLogin",
				Password: "testPassword",
			},
			wantErr: true,
			mockSetup: func(mockClient *proto.MockPassManagerClient) {
				mockClient.EXPECT().AddPassword(
					gomock.Any(),
					gomock.Any(),
				).DoAndReturn(func(ctx context.Context, in *proto.Password, opts ...grpc.CallOption) (*emptypb.Empty, error) {
					return nil, errors.New("test error")
				})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup(mockClient)
			err := client.AddPassword(context.Background(), tt.pass)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
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
		wantErr          bool
		expectedPassword repository.Password
		mockSetup        func(mockClient *proto.MockPassManagerClient)
	}{
		{
			name:    "Success",
			wantErr: false,
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
		{
			name:    "Error",
			wantErr: true,
			mockSetup: func(mockClient *proto.MockPassManagerClient) {
				mockClient.EXPECT().GetPassword(
					gomock.Any(),
					gomock.Any(),
				).DoAndReturn(func(ctx context.Context, in *proto.GetPasswordRequest, opts ...grpc.CallOption) (*proto.Password, error) {
					return nil, errors.New("test error")
				})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup(mockClient)
			password, err := client.GetPassword(context.Background(), tt.name)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, repository.Password{}, password)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedPassword.Name, password.Name)
				assert.Equal(t, tt.expectedPassword.Login, password.Login)
				assert.Equal(t, tt.expectedPassword.Password, password.Password)
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
		wantErr   bool
		card      repository.Card
		mockSetup func(mockClient *proto.MockPassManagerClient)
	}{
		{
			name:    "Success",
			wantErr: false,
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
		{
			name: "Error",
			card: repository.Card{
				CardName:   "testCard",
				CardNumber: "1234567890123456",
				CardCVC:    "123",
				CardDate:   "12/24",
				CardFI:     "Test Bank",
			},
			wantErr: true,
			mockSetup: func(mockClient *proto.MockPassManagerClient) {
				mockClient.EXPECT().AddCard(
					gomock.Any(),
					gomock.Any(),
				).DoAndReturn(func(ctx context.Context, in *proto.Card, opts ...grpc.CallOption) (*emptypb.Empty, error) {
					return nil, errors.New("test error")
				})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup(mockClient)
			err := client.AddCard(context.Background(), tt.card)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
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
		wantErr      bool
		expectedCard repository.Card
		mockSetup    func(mockClient *proto.MockPassManagerClient)
	}{
		{
			name:     "Success",
			cardName: "testCard",
			wantErr:  false,
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
		{
			name:     "Error",
			cardName: "testCard",
			wantErr:  true,
			mockSetup: func(mockClient *proto.MockPassManagerClient) {
				mockClient.EXPECT().GetCard(
					gomock.Any(),
					gomock.Any(),
				).DoAndReturn(func(ctx context.Context, in *proto.GetCardRequest, opts ...grpc.CallOption) (*proto.Card, error) {
					return nil, errors.New("test error")
				})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup(mockClient)
			card, err := client.GetCard(context.Background(), tt.cardName)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, repository.Card{}, card)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedCard.CardName, card.CardName)
				assert.Equal(t, tt.expectedCard.CardNumber, card.CardNumber)
				assert.Equal(t, tt.expectedCard.CardCVC, card.CardCVC)
				assert.Equal(t, tt.expectedCard.CardDate, card.CardDate)
				assert.Equal(t, tt.expectedCard.CardFI, card.CardFI)
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
		wantErr   bool
		note      repository.Note
		mockSetup func(mockClient *proto.MockPassManagerClient)
	}{
		{
			name:    "Success",
			wantErr: false,
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
		{
			name: "Error",
			note: repository.Note{
				NoteName: "testNote",
				Note:     "This is a test note.",
			},
			wantErr: true,
			mockSetup: func(mockClient *proto.MockPassManagerClient) {
				mockClient.EXPECT().AddNote(
					gomock.Any(),
					gomock.Any(),
				).DoAndReturn(func(ctx context.Context, in *proto.Note, opts ...grpc.CallOption) (*emptypb.Empty, error) {
					return nil, errors.New("test error")
				})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup(mockClient)
			err := client.AddNote(context.Background(), tt.note)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
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
		wantErr      bool
		expectedNote repository.Note
		mockSetup    func(mockClient *proto.MockPassManagerClient)
	}{
		{
			name:     "Success",
			noteName: "testNote",
			wantErr:  false,
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
		{
			name:     "Error",
			noteName: "testNote",
			wantErr:  true,
			mockSetup: func(mockClient *proto.MockPassManagerClient) {
				mockClient.EXPECT().GetNote(
					gomock.Any(),
					gomock.Any(),
				).DoAndReturn(func(ctx context.Context, in *proto.GetNoteRequest, opts ...grpc.CallOption) (*proto.Note, error) {
					return nil, errors.New("test error")
				})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup(mockClient)
			note, err := client.GetNote(context.Background(), tt.noteName)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, repository.Note{}, note)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedNote.NoteName, note.NoteName)
				assert.Equal(t, tt.expectedNote.Note, note.Note)
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
		wantErr    bool
		binaryData repository.BinaryData
		mockSetup  func(mockClient *proto.MockPassManagerClient)
	}{
		{
			name:    "Success",
			wantErr: false,
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
		{
			name: "Error",
			binaryData: repository.BinaryData{
				BytesName: "testBytes",
				Value:     []byte("test binary data"),
			},
			wantErr: true,
			mockSetup: func(mockClient *proto.MockPassManagerClient) {
				mockClient.EXPECT().AddBytes(
					gomock.Any(),
					gomock.Any(),
				).DoAndReturn(func(ctx context.Context, in *proto.Bytes, opts ...grpc.CallOption) (*emptypb.Empty, error) {
					return nil, errors.New("test error")
				})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup(mockClient)
			err := client.AddBytes(context.Background(), tt.binaryData)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
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
		wantErr       bool
		expectedBytes repository.BinaryData
		mockSetup     func(mockClient *proto.MockPassManagerClient)
	}{
		{
			name:      "Success",
			bytesName: "testBytes",
			wantErr:   false,
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
		{
			name:      "Error",
			bytesName: "testBytes",
			wantErr:   true,
			mockSetup: func(mockClient *proto.MockPassManagerClient) {
				mockClient.EXPECT().GetBytes(
					gomock.Any(),
					gomock.Any(),
				).DoAndReturn(func(ctx context.Context, in *proto.GetBytesRequest, opts ...grpc.CallOption) (*proto.Bytes, error) {
					return nil, errors.New("test error")
				})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup(mockClient)
			bytes, err := client.GetBytes(context.Background(), tt.bytesName)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, repository.BinaryData{}, bytes)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedBytes.BytesName, bytes.BytesName)
				assert.Equal(t, tt.expectedBytes.Value, bytes.Value)
			}
		})
	}
}
