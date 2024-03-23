package auth

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func TestNewToken(t *testing.T) {
	auth := NewAuth("testSecret")
	userId := primitive.NewObjectID()

	tests := []struct {
		name    string
		userId  primitive.ObjectID
		wantErr bool
	}{
		{"Valid User ID", userId, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := auth.NewToken(tt.userId)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestParseToken(t *testing.T) {
	auth := NewAuth("testSecret")
	userId := primitive.NewObjectID()
	tokenString, _ := auth.NewToken(userId)

	tests := []struct {
		name    string
		token   string
		want    primitive.ObjectID
		wantErr bool
	}{
		{"Valid Token", tokenString, userId, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := auth.ParseToken(tt.token)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestOnlyAuthorizedInterceptor(t *testing.T) {
	a := NewAuth("your_grpc_secret")

	testCases := []struct {
		name          string
		metadata      metadata.MD
		expectedError error
	}{
		{
			name:          "NoMetadata",
			metadata:      nil,
			expectedError: status.Errorf(codes.Unauthenticated, "metadata is not provided"),
		},
		{
			name:          "NoAuthorization",
			metadata:      metadata.Pairs("other", "value"),
			expectedError: status.Errorf(codes.Unauthenticated, "authorization is not provided"),
		},
		{
			name:          "InvalidToken",
			metadata:      metadata.Pairs("authorization", "invalid_token"),
			expectedError: status.Errorf(codes.Unauthenticated, "authorization is not valid"),
		},
		{
			name:          "ValidToken",
			metadata:      metadata.Pairs("authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjY1ZjgyZDc4Y2EwMDIwMWViMGFmYmQxNiJ9.ucsGpTnbFKRhHFTAMISr2H268pymzboldB7Esr2Z-t8"),
			expectedError: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			if tc.metadata != nil {
				ctx = metadata.NewIncomingContext(context.Background(), tc.metadata)
			}
			testHandler := func(ctx context.Context, req any) (any, error) {
				md, ok := metadata.FromOutgoingContext(ctx)
				assert.True(t, ok)
				id := md.Get("id")[0]
				assert.NotEmpty(t, id)
				return nil, nil
			}
			_, err := a.OnlyAuthorizedInterceptor(ctx, nil, nil, testHandler)
			if !errors.Is(err, tc.expectedError) {
				assert.EqualError(t, err, tc.expectedError.Error())
			}
		})
	}
}
