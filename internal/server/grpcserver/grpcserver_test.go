package grpcserver_test

import (
	"context"
	"errors"
	"testing"

	"github.com/JustWorking42/go-password-manager/internal/server/auth"
	"github.com/JustWorking42/go-password-manager/internal/server/grpcserver"
	"github.com/JustWorking42/go-password-manager/internal/server/storage"
	"github.com/JustWorking42/go-password-manager/proto"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/metadata"
)

const testId = "65f6dd78ca00201eb0afbd14"

func TestRegister(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := storage.NewMockStorage(ctrl)
	auth := auth.NewAuth("1")
	s := grpcserver.NewPassGRPCServer(mockStorage, auth)

	tests := []struct {
		name    string
		req     *proto.Creds
		setup   func()
		wantErr bool
	}{
		{
			name: "Success",
			req: &proto.Creds{
				Login:    "test",
				Password: "password",
			},
			setup: func() {
				mockStorage.EXPECT().IsLoginEnabled(gomock.Any(), "test").Return(true, nil)
				mockStorage.EXPECT().AddUser(gomock.Any(), gomock.Any()).Return(primitive.NewObjectID(), nil)
			},
			wantErr: false,
		},
		{
			name: "LoginNotEnabled",
			req: &proto.Creds{
				Login:    "test",
				Password: "password",
			},
			setup: func() {
				mockStorage.EXPECT().IsLoginEnabled(gomock.Any(), "test").Return(false, nil)
			},
			wantErr: true,
		},
		{
			name: "RegisterWithEmptyLogin",
			req: &proto.Creds{
				Login:    "",
				Password: "password",
			},
			setup:   func() {},
			wantErr: true,
		},
		{
			name: "RegisterWithEmptyPassword",
			req: &proto.Creds{
				Login:    "test",
				Password: "",
			},
			setup:   func() {},
			wantErr: true,
		},
		{
			name: "RegisterWithExistingLogin",
			req: &proto.Creds{
				Login:    "existing_login",
				Password: "password",
			},
			setup: func() {
				mockStorage.EXPECT().IsLoginEnabled(gomock.Any(), "existing_login").Return(true, nil)
				mockStorage.EXPECT().AddUser(gomock.Any(), gomock.Any()).Return(primitive.NewObjectID(), nil)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			_, err := s.Register(context.Background(), tt.req)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := storage.NewMockStorage(ctrl)
	auth := auth.NewAuth("1")
	s := grpcserver.NewPassGRPCServer(mockStorage, auth)

	tests := []struct {
		name    string
		req     *proto.Creds
		setup   func()
		wantErr bool
	}{
		{
			name: "Success",
			req: &proto.Creds{
				Login:    "test",
				Password: "password",
			},
			setup: func() {
				mockStorage.EXPECT().GetUser(gomock.Any(), "test").Return(storage.User{
					Login:        "test",
					PasswordHash: []byte("$2a$10$xsgRoU4q4ke15ovQI1eC9Od5NojMbTXhvNhwMq24qtz1jgZZhOLuK"),
				}, nil)
			},
			wantErr: false,
		},
		{
			name: "IncorrectPassword",
			req: &proto.Creds{
				Login:    "test",
				Password: "wrong_password",
			},
			setup: func() {
				mockStorage.EXPECT().GetUser(gomock.Any(), "test").Return(storage.User{
					Login:        "test",
					PasswordHash: []byte("$2a$10$xsgRoU4q4ke15ovQI1eC9Od5NojMbTXhvNhwMq24qtz1jgZZhOLuK"),
				}, nil)
			},
			wantErr: true,
		},
		{
			name: "NonExistentUser",
			req: &proto.Creds{
				Login:    "nonexistent",
				Password: "password",
			},
			setup: func() {
				mockStorage.EXPECT().GetUser(gomock.Any(), "nonexistent").Return(storage.User{}, errors.New("user not found"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			_, err := s.Login(context.Background(), tt.req)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestAddPassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := storage.NewMockStorage(ctrl)
	auth := auth.NewAuth("1")
	s := grpcserver.NewPassGRPCServer(mockStorage, auth)

	tests := []struct {
		name    string
		req     *proto.Password
		setup   func()
		wantErr bool
	}{
		{
			name: "Success",
			req: &proto.Password{
				ServiceName:     "service",
				ServiceLogin:    "login",
				ServicePassword: "password",
			},
			setup: func() {
				mockStorage.EXPECT().AddPassword(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "StorageError",
			req: &proto.Password{
				ServiceName:     "service",
				ServiceLogin:    "login",
				ServicePassword: "password",
			},
			setup: func() {
				mockStorage.EXPECT().AddPassword(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("storage error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			ctx := metadata.NewOutgoingContext(context.Background(), metadata.MD{
				"id": []string{testId},
			})
			_, err := s.AddPassword(ctx, tt.req)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestGetPassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := storage.NewMockStorage(ctrl)
	auth := auth.NewAuth("1")
	s := grpcserver.NewPassGRPCServer(mockStorage, auth)

	tests := []struct {
		name    string
		req     *proto.GetPasswordRequest
		setup   func()
		wantErr bool
	}{
		{
			name: "Success",
			req: &proto.GetPasswordRequest{
				ServiceName: "service",
			},
			setup: func() {
				mockStorage.EXPECT().GetPassword(gomock.Any(), gomock.Any(), "service").Return(storage.PasswordData{
					ServiceName:     "service",
					ServiceLogin:    "login",
					ServicePassword: "password",
				}, nil)
			},
			wantErr: false,
		},
		{
			name: "StorageError",
			req: &proto.GetPasswordRequest{
				ServiceName: "service",
			},
			setup: func() {
				mockStorage.EXPECT().GetPassword(gomock.Any(), gomock.Any(), "service").Return(storage.PasswordData{}, errors.New("storage error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			ctx := metadata.NewOutgoingContext(context.Background(), metadata.MD{
				"id": []string{testId},
			})
			_, err := s.GetPassword(ctx, tt.req)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestAddCard(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := storage.NewMockStorage(ctrl)
	auth := auth.NewAuth("1")
	s := grpcserver.NewPassGRPCServer(mockStorage, auth)

	tests := []struct {
		name    string
		req     *proto.Card
		setup   func()
		wantErr bool
	}{
		{
			name: "Success",
			req: &proto.Card{
				CardName:   "card",
				CardNumber: "123456789012345",
				CardCVC:    "123",
				CardDate:   "12/24",
				CardFI:     "123456",
			},
			setup: func() {
				mockStorage.EXPECT().AddCard(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "StorageError",
			req: &proto.Card{
				CardName:   "card",
				CardNumber: "123456789012345",
				CardCVC:    "123",
				CardDate:   "12/24",
				CardFI:     "123456",
			},
			setup: func() {
				mockStorage.EXPECT().AddCard(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("storage error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			ctx := metadata.NewOutgoingContext(context.Background(), metadata.MD{
				"id": []string{testId},
			})
			_, err := s.AddCard(ctx, tt.req)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestGetCard(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := storage.NewMockStorage(ctrl)
	auth := auth.NewAuth("1")
	s := grpcserver.NewPassGRPCServer(mockStorage, auth)

	tests := []struct {
		name    string
		req     *proto.GetCardRequest
		setup   func()
		wantErr bool
	}{
		{
			name: "Success",
			req: &proto.GetCardRequest{
				CardName: "test_card",
			},
			setup: func() {
				mockStorage.EXPECT().GetCard(gomock.Any(), gomock.Any(), "test_card").Return(storage.CardData{
					CardName: "test_card",
					Number:   "123456789012345",
					CVC:      "123",
					Date:     "12/24",
					FI:       "123456",
				}, nil)
			},
			wantErr: false,
		},
		{
			name: "CardNotFound",
			req: &proto.GetCardRequest{
				CardName: "nonexistent_card",
			},
			setup: func() {
				mockStorage.EXPECT().GetCard(gomock.Any(), gomock.Any(), "nonexistent_card").Return(storage.CardData{}, errors.New("card not found"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			ctx := metadata.NewOutgoingContext(context.Background(), metadata.MD{
				"id": []string{testId},
			})
			_, err := s.GetCard(ctx, tt.req)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestAddNote(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := storage.NewMockStorage(ctrl)
	auth := auth.NewAuth("1")
	s := grpcserver.NewPassGRPCServer(mockStorage, auth)

	tests := []struct {
		name    string
		req     *proto.Note
		setup   func()
		wantErr bool
	}{
		{
			name: "Success",
			req: &proto.Note{
				NoteName: "test_note",
				Note:     "This is a test note",
			},
			setup: func() {
				mockStorage.EXPECT().AddNote(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "StorageError",
			req: &proto.Note{
				NoteName: "test_note",
				Note:     "This is a test note",
			},
			setup: func() {
				mockStorage.EXPECT().AddNote(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("storage error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			ctx := metadata.NewOutgoingContext(context.Background(), metadata.MD{
				"id": []string{testId},
			})
			_, err := s.AddNote(ctx, tt.req)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestGetNote(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := storage.NewMockStorage(ctrl)
	auth := auth.NewAuth("1")
	s := grpcserver.NewPassGRPCServer(mockStorage, auth)

	tests := []struct {
		name    string
		req     *proto.GetNoteRequest
		setup   func()
		wantErr bool
	}{
		{
			name: "Success",
			req: &proto.GetNoteRequest{
				NoteName: "test_note",
			},
			setup: func() {
				mockStorage.EXPECT().GetNote(gomock.Any(), gomock.Any(), "test_note").Return(storage.Note{
					Name: "test_note",
					Text: "This is a test note",
				}, nil)
			},
			wantErr: false,
		},
		{
			name: "NoteNotFound",
			req: &proto.GetNoteRequest{
				NoteName: "nonexistent_note",
			},
			setup: func() {
				mockStorage.EXPECT().GetNote(gomock.Any(), gomock.Any(), "nonexistent_note").Return(storage.Note{}, errors.New("note not found"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			ctx := metadata.NewOutgoingContext(context.Background(), metadata.MD{
				"id": []string{testId},
			})
			_, err := s.GetNote(ctx, tt.req)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestAddBytes(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := storage.NewMockStorage(ctrl)
	auth := auth.NewAuth("1")
	s := grpcserver.NewPassGRPCServer(mockStorage, auth)

	tests := []struct {
		name    string
		req     *proto.Bytes
		setup   func()
		wantErr bool
	}{
		{
			name: "Success",
			req: &proto.Bytes{
				BytesName: "test_binary",
				Value:     []byte("This is a test binary data"),
			},
			setup: func() {
				mockStorage.EXPECT().AddBytes(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "StorageError",
			req: &proto.Bytes{
				BytesName: "test_binary",
				Value:     []byte("This is a test binary data"),
			},
			setup: func() {
				mockStorage.EXPECT().AddBytes(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("storage error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			ctx := metadata.NewOutgoingContext(context.Background(), metadata.MD{
				"id": []string{testId},
			})
			_, err := s.AddBytes(ctx, tt.req)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestGetBytes(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := storage.NewMockStorage(ctrl)
	auth := auth.NewAuth("1")
	s := grpcserver.NewPassGRPCServer(mockStorage, auth)

	tests := []struct {
		name    string
		req     *proto.GetBytesRequest
		setup   func()
		wantErr bool
	}{
		{
			name: "Success",
			req: &proto.GetBytesRequest{
				BytesName: "test_binary",
			},
			setup: func() {
				mockStorage.EXPECT().GetBytes(gomock.Any(), gomock.Any(), "test_binary").Return(storage.BinaryData{
					Name: "test_binary",
					Data: []byte("This is a test binary data"),
				}, nil)
			},
			wantErr: false,
		},
		{
			name: "BytesNotFound",
			req: &proto.GetBytesRequest{
				BytesName: "nonexistent_binary",
			},
			setup: func() {
				mockStorage.EXPECT().GetBytes(gomock.Any(), gomock.Any(), "nonexistent_binary").Return(storage.BinaryData{}, errors.New("binary data not found"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			ctx := metadata.NewOutgoingContext(context.Background(), metadata.MD{
				"id": []string{testId},
			})
			_, err := s.GetBytes(ctx, tt.req)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}
