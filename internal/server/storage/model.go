// Package storage contains all the models used in the storage.
package storage

import "go.mongodb.org/mongo-driver/bson/primitive"

// User represents a user.
type User struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Login          string             `bson:"login"`
	PasswordHash   []byte             `bson:"pass_hash"`
	SavedPasswords []PasswordData     `bson:"passwords"`
	Notes          []Note             `bson:"notes"`
	BinaryDatas    []BinaryData       `bson:"binaryDatas"`
	Cards          []CardData         `bson:"cards"`
}

// PasswordData represents a password.
type PasswordData struct {
	ServiceName     string `bson:"serviceName"`
	ServiceLogin    string `bson:"serviceLogin"`    // Encrypted
	ServicePassword string `bson:"servicePassword"` // Encrypted
}

// / Note represents a note.
type Note struct {
	Name string `bson:"name"`
	Text string `bson:"text"` // Encrypted
}

// BinaryData represents a binary data.
type BinaryData struct {
	Name string `bson:"name"`
	Data []byte `bson:"data"` // Encrypted
}

// CardData represents a card.
type CardData struct {
	CardName string `bson:"cardName"`
	Number   string `bson:"number"` // Encrypted
	CVC      string `bson:"cvc"`    // Encrypted
	Date     string `bson:"date"`   // Encrypted
	FI       string `bson:"fi"`     // Encrypted
}

// NewUser creates a new user.
func NewUser(login string, passwordHash []byte) User {
	return User{
		Login:          login,
		PasswordHash:   passwordHash,
		SavedPasswords: []PasswordData{},
		Notes:          []Note{},
		BinaryDatas:    []BinaryData{},
		Cards:          []CardData{},
	}
}
