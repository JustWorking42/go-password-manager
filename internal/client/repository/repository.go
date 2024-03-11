package repository

import (
	"context"
)

var repository Repository

type Repository interface {
	Close()
	Register(ctx context.Context, login, password string) (string, error)
	Login(ctx context.Context, login, password string) (string, error)
	AddPassword(ctx context.Context, pass Password) error
	GetPassword(ctx context.Context, name string) (Password, error)
	AddCard(ctx context.Context, card Card) error
	GetCard(ctx context.Context, cardName string) (Card, error)
	AddNote(ctx context.Context, note Note) error
	GetNote(ctx context.Context, noteName string) (Note, error)
	AddBytes(ctx context.Context, binaryData BinaryData) error
	GetBytes(ctx context.Context, bytesName string) (BinaryData, error)
}

func SetRepository(rep Repository) {
	repository = rep
}

func GetRepository() Repository {
	return repository
}
