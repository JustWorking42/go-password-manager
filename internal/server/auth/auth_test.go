package auth

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
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
			if (err != nil) != tt.wantErr {
				t.Errorf("NewToken() error = %v, wantErr %v", err, tt.wantErr)
				return
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
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
