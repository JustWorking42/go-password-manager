// Package repository provides an interface for managing password storage and retrieval.
package repository

import (
	"context"
)

// Repository is an singleton.
var repository Repository

// Repository is an interface for managing password storage and retrieval.
type Repository interface {
	// Close closes the repository.
	Close()
	Register(ctx context.Context, login, password string) (string, error)
	Login(ctx context.Context, login, password string) (string, error)
	// AddPassword adds a new password to the repository.
	AddPassword(ctx context.Context, pass Password) error
	GetPassword(ctx context.Context, name string) (Password, error)
	// AddCard adds a new card to the repository.
	AddCard(ctx context.Context, card Card) error
	GetCard(ctx context.Context, cardName string) (Card, error)
	// AddNote adds a new note to the repository.
	AddNote(ctx context.Context, note Note) error
	GetNote(ctx context.Context, noteName string) (Note, error)
	// AddBytes adds new binary data to the repository.
	AddBytes(ctx context.Context, binaryData BinaryData) error
	GetBytes(ctx context.Context, bytesName string) (BinaryData, error)
}

// SetRepository sets the repository instance.
func SetRepository(rep Repository) {
	repository = rep
}

// GetRepository returns the current repository instance.
func GetRepository() Repository {
	return repository
}
