package repository

type Password struct {
	Name     string
	Login    string
	Password string
}

type Card struct {
	CardName   string
	CardNumber string
	CardCVC    string
	CardDate   string
	CardFI     string
}

type Note struct {
	NoteName string
	Note     string
}

type BinaryData struct {
	BytesName string
	Value     []byte
}
