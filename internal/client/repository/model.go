// Package repository defines the data structures used for password management.
package repository

// Password represents a stored password with its name, login, and password.
type Password struct {
	Name     string // Name of the password entry.
	Login    string // Login associated with the password.
	Password string // The stored password.
}

type Card struct {
	CardName   string // Name of the card entry.
	CardNumber string // The card number.
	CardCVC    string // The card's CVC code.
	CardDate   string // The card's expiration date.
	CardFI     string // The card issuer.
}

type Note struct {
	NoteName string // Name of the note entry.
	Note     string // The content of the note.
}

type BinaryData struct {
	BytesName string // Name of the binary data entry.
	Value     []byte // The stored binary data.
}
