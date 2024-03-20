// Package grpcserver implements a gRPC server.
package grpcserver

import (
	"context"

	"github.com/JustWorking42/go-password-manager/internal/server/auth"
	"github.com/JustWorking42/go-password-manager/internal/server/storage"
	"github.com/JustWorking42/go-password-manager/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// PassGRPCServer implements a gRPC server.
type PassGRPCServer struct {
	proto.UnimplementedPassManagerServer
	db   storage.Storage
	auth *auth.Auth
}

// NewPassGRPCServer creates a new PassGRPCServer.
func NewPassGRPCServer(db storage.Storage, auth *auth.Auth) *PassGRPCServer {
	return &PassGRPCServer{
		db:   db,
		auth: auth,
	}
}

var _ proto.PassManagerServer = (*PassGRPCServer)(nil)

// Register handles the registration of a new user.
func (s *PassGRPCServer) Register(ctx context.Context, req *proto.Creds) (*emptypb.Empty, error) {

	if req.Login == "" || req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "login and password must not be empty")
	}

	isEnabled, err := s.db.IsLoginEnabled(ctx, req.Login)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if !isEnabled {
		return nil, status.Error(codes.InvalidArgument, "login is not enabled")
	}
	passHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	id, err := s.db.AddUser(ctx, storage.NewUser(req.Login, passHash))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	jwt, err := s.auth.NewToken(id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())

	}

	md := metadata.Pairs("authorization", "Bearer "+jwt)
	grpc.SetHeader(ctx, md)
	return &emptypb.Empty{}, nil

}

// Login handles the login of a user.
func (s *PassGRPCServer) Login(ctx context.Context, req *proto.Creds) (*emptypb.Empty, error) {

	if req.Login == "" || req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "login and password must not be empty")
	}

	user, err := s.db.GetUser(ctx, req.Login)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	err = bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(req.Password))
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid login or password")
	}

	jwt, err := s.auth.NewToken(user.ID)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	md := metadata.Pairs("authorization", "Bearer "+jwt)
	grpc.SetHeader(ctx, md)
	return &emptypb.Empty{}, nil
}

// AddPassword handles the adding of a new password.
func (s *PassGRPCServer) AddPassword(ctx context.Context, req *proto.Password) (*emptypb.Empty, error) {
	primitiveID, err := getPrimitiveID(ctx)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err = s.db.AddPassword(ctx, primitiveID, storage.PasswordData{
		ServiceName:     req.ServiceName,
		ServiceLogin:    req.ServiceLogin,
		ServicePassword: req.ServicePassword,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil
}

// GetPassword handles the getting of a password.
func (s *PassGRPCServer) GetPassword(ctx context.Context, req *proto.GetPasswordRequest) (*proto.Password, error) {

	primitiveID, err := getPrimitiveID(ctx)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	passData, err := s.db.GetPassword(ctx, primitiveID, req.ServiceName)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &proto.Password{
		ServiceName:     passData.ServiceName,
		ServiceLogin:    passData.ServiceLogin,
		ServicePassword: passData.ServicePassword,
	}, nil
}

// AddCard handles the adding of a new card.
func (s *PassGRPCServer) AddCard(ctx context.Context, req *proto.Card) (*emptypb.Empty, error) {
	primitiveID, err := getPrimitiveID(ctx)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err = s.db.AddCard(ctx, primitiveID, storage.CardData{
		CardName: req.CardName,
		Number:   req.CardNumber,
		CVC:      req.CardCVC,
		Date:     req.CardDate,
		FI:       req.CardFI,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil
}

// GetCard handles the getting of a card.
func (s *PassGRPCServer) GetCard(ctx context.Context, req *proto.GetCardRequest) (*proto.Card, error) {
	primitiveID, err := getPrimitiveID(ctx)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	cardData, err := s.db.GetCard(ctx, primitiveID, req.CardName)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &proto.Card{
		CardName:   cardData.CardName,
		CardNumber: cardData.Number,
		CardCVC:    cardData.CVC,
		CardDate:   cardData.Date,
		CardFI:     cardData.FI,
	}, nil
}

// AddNote handles the adding of a new note.
func (s *PassGRPCServer) AddNote(ctx context.Context, req *proto.Note) (*emptypb.Empty, error) {
	primitiveID, err := getPrimitiveID(ctx)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err = s.db.AddNote(ctx, primitiveID, storage.Note{
		Name: req.NoteName,
		Text: req.Note,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil
}

// GetNote handles the getting of a note.
func (s *PassGRPCServer) GetNote(ctx context.Context, req *proto.GetNoteRequest) (*proto.Note, error) {
	primitiveID, err := getPrimitiveID(ctx)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	noteData, err := s.db.GetNote(ctx, primitiveID, req.NoteName)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &proto.Note{
		NoteName: noteData.Name,
		Note:     noteData.Text,
	}, nil
}

// AddByte handles the adding of a new byte.
func (s *PassGRPCServer) AddBytes(ctx context.Context, req *proto.Bytes) (*emptypb.Empty, error) {
	primitiveID, err := getPrimitiveID(ctx)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err = s.db.AddBytes(ctx, primitiveID, storage.BinaryData{
		Name: req.BytesName,
		Data: req.Value,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil
}

// GetBytes handles the getting of a byte.
func (s *PassGRPCServer) GetBytes(ctx context.Context, req *proto.GetBytesRequest) (*proto.Bytes, error) {

	primitiveID, err := getPrimitiveID(ctx)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	binaryData, err := s.db.GetBytes(ctx, primitiveID, req.BytesName)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &proto.Bytes{
		BytesName: binaryData.Name,
		Value:     binaryData.Data,
	}, nil
}

func getPrimitiveID(ctx context.Context) (primitive.ObjectID, error) {
	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		return primitive.ObjectID{}, status.Error(codes.InvalidArgument, "no metadata")
	}

	ids := md.Get("id")
	if len(ids) == 0 {
		return primitive.ObjectID{}, status.Error(codes.InvalidArgument, "no id in metadata")
	}

	primitiveID, err := primitive.ObjectIDFromHex(ids[0])
	if err != nil {
		return primitive.ObjectID{}, status.Error(codes.InvalidArgument, err.Error())
	}

	return primitiveID, nil
}
