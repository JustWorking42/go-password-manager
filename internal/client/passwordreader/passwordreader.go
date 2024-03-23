package passwordreader

import (
	"os"

	"golang.org/x/term"
)

var passwordReader PasswordReader

type PasswordReader interface {
	ReadPassword() ([]byte, error)
}

func SetPasswordReader(pr PasswordReader) {
	passwordReader = pr
}

func GetPasswordReader() PasswordReader {
	return passwordReader
}

type TermPasswordReader struct{}

func (t *TermPasswordReader) ReadPassword() ([]byte, error) {
	return term.ReadPassword(int(os.Stdin.Fd()))
}
