// Package storage provides an interface for storage operations in the password manager application.
package storage

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Storage defines the interface for storage operations.
type Storage interface {
	// AddPassword adds a new password to the storage.
	AddPassword(ctx context.Context, id primitive.ObjectID, data PasswordData) error
	// GetPassword retrieves a password by user ID and service title.
	GetPassword(ctx context.Context, id primitive.ObjectID, serviceTitle string) (PasswordData, error)
	// AddUser adds a new user to the storage.
	AddUser(ctx context.Context, user User) (primitive.ObjectID, error)
	// IsLoginEnabled checks if a login is enabled in the storage.
	IsLoginEnabled(ctx context.Context, login string) (bool, error)
	// GetUser retrieves a user by their login.
	GetUser(ctx context.Context, login string) (User, error)
	// AddCard adds a new card to the storage.
	AddCard(ctx context.Context, id primitive.ObjectID, card CardData) error
	// GetCard retrieves a card by user ID and card name.
	GetCard(ctx context.Context, id primitive.ObjectID, cardName string) (CardData, error)
	// AddNote adds a new note to the storage.
	AddNote(ctx context.Context, id primitive.ObjectID, note Note) error
	// GetNote retrieves a note by user ID and note name.
	GetNote(ctx context.Context, id primitive.ObjectID, noteName string) (Note, error)
	// AddBytes adds binary data to the storage.
	AddBytes(ctx context.Context, id primitive.ObjectID, binaryData BinaryData) error
	// GetBytes retrieves binary data by user ID and binary name.
	GetBytes(ctx context.Context, id primitive.ObjectID, binaryName string) (BinaryData, error)
	// Close closes the storage connection.
	Close(ctx context.Context)
}
