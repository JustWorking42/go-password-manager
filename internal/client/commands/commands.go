package commands

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/JustWorking42/go-password-manager/internal/client/passwordreader"
	"github.com/JustWorking42/go-password-manager/internal/client/repository"
	"github.com/JustWorking42/go-password-manager/internal/client/session"
	"github.com/JustWorking42/go-password-manager/internal/common/defens"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/metadata"
)

func readLogin() string {
	fmt.Println("Enter login:")
	login, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSuffix(login, "\n")

}

func readPassword() string {
	fmt.Println("Enter password:")
	password, err := passwordreader.GetPasswordReader().ReadPassword()
	if err != nil {
		log.Fatal(err)
	}
	return string(password)

}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func encryptData(key, data string) string {
	encryptedData, err := defens.Encrypt([]byte(key), data)
	handleError(err)
	return encryptedData
}

func decryptData(key, data string) string {
	decryptedData, err := defens.Decrypt([]byte(key), data)
	handleError(err)
	return decryptedData
}

func getSessionCtx(ctx context.Context) context.Context {
	session, err := session.GetSessionManager().GetSession()
	newCtx := metadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+session.JWT)

	handleError(err)
	return newCtx
}

var RegisterCmd = &cobra.Command{
	Use:   "register",
	Short: "Register a new user",
	Run: func(cmd *cobra.Command, args []string) {

		login := readLogin()
		password := readPassword()

		client := repository.GetRepository()

		jwt, err := client.Register(context.Background(), login, password)
		handleError(err)

		err = session.GetSessionManager().SetSession(session.Session{JWT: jwt})
		handleError(err)
		log.Println("User registered successfully and session saved.")
	},
}

var LoginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to the system",
	Run: func(cmd *cobra.Command, args []string) {

		login := readLogin()
		password := readPassword()

		client := repository.GetRepository()

		jwt, err := client.Login(context.Background(), login, password)
		handleError(err)

		err = session.GetSessionManager().SetSession(session.Session{JWT: jwt})
		handleError(err)
		log.Println("User logged in successfully and session saved.")
	},
}

var AddPasswordCmd = &cobra.Command{
	Use:   "add-pass [name] [login] [password] [key]",
	Short: "Add a new password",
	Args:  cobra.ExactArgs(4),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		login := args[1]
		password := args[2]
		key := args[3]

		ctx := getSessionCtx(context.Background())

		client := repository.GetRepository()
		encryptedPassword := encryptData(key, password)
		encryptedLogin := encryptData(key, login)

		err := client.AddPassword(ctx, repository.Password{
			Name:     name,
			Login:    encryptedLogin,
			Password: encryptedPassword,
		})
		handleError(err)
	},
}

var GetPasswordCmd = &cobra.Command{
	Use:   "get-pass [name] [key]",
	Short: "Get a password",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		key := args[1]

		ctx := getSessionCtx(context.Background())

		client := repository.GetRepository()
		password, err := client.GetPassword(ctx, name)
		handleError(err)

		decryptedPassword := decryptData(key, password.Password)
		decryptedLogin := decryptData(key, password.Login)

		log.Println(password.Name)
		log.Println(decryptedLogin)
		log.Println(decryptedPassword)
	},
}

var AddCardCmd = &cobra.Command{
	Use:   "add-card [cardName] [cardNumber] [cardCVC] [cardDate] [cardFI] [key]",
	Short: "Add a new card",
	Args:  cobra.ExactArgs(6),
	Run: func(cmd *cobra.Command, args []string) {
		cardName := args[0]
		cardNumber := args[1]
		cardCVC := args[2]
		cardDate := args[3]
		cardFI := args[4]
		key := args[5]

		ctx := getSessionCtx(context.Background())

		client := repository.GetRepository()
		encryptedCardNumber := encryptData(key, cardNumber)
		encryptedCardCVC := encryptData(key, cardCVC)
		encryptedCardDate := encryptData(key, cardDate)
		encryptedCardFI := encryptData(key, cardFI)

		err := client.AddCard(ctx, repository.Card{
			CardName:   cardName,
			CardNumber: encryptedCardNumber,
			CardCVC:    encryptedCardCVC,
			CardDate:   encryptedCardDate,
			CardFI:     encryptedCardFI,
		})
		handleError(err)
	},
}

var GetCardCmd = &cobra.Command{
	Use:   "get-card [cardName] [key]",
	Short: "Get card information",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		cardName := args[0]
		key := args[1]

		ctx := getSessionCtx(context.Background())

		client := repository.GetRepository()
		card, err := client.GetCard(ctx, cardName)
		handleError(err)

		decryptedCardNumber := decryptData(key, card.CardNumber)
		decryptedCardCVC := decryptData(key, card.CardCVC)
		decryptedCardDate := decryptData(key, card.CardDate)
		decryptedCardFI := decryptData(key, card.CardFI)

		log.Println("Card Name:", card.CardName)
		log.Println("Card Number:", decryptedCardNumber)
		log.Println("Card CVC:", decryptedCardCVC)
		log.Println("Card Date:", decryptedCardDate)
		log.Println("Card FI:", decryptedCardFI)
	},
}

var AddNoteCmd = &cobra.Command{
	Use:   "add-note [noteName] [noteContent] [key]",
	Short: "Add a new note",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		noteName := args[0]
		noteContent := args[1]
		key := args[2]

		ctx := getSessionCtx(context.Background())

		client := repository.GetRepository()
		encryptedNote := encryptData(key, noteContent)

		err := client.AddNote(ctx, repository.Note{
			NoteName: noteName,
			Note:     encryptedNote,
		})
		handleError(err)
	},
}

var GetNoteCmd = &cobra.Command{
	Use:   "get-note [noteName] [key]",
	Short: "Get note content",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		noteName := args[0]
		key := args[1]

		ctx := getSessionCtx(context.Background())

		client := repository.GetRepository()
		note, err := client.GetNote(ctx, noteName)
		handleError(err)

		decryptedNote := decryptData(key, note.Note)

		log.Println("Note Name:", note.NoteName)
		log.Println("Note Content:", decryptedNote)
	},
}

var AddBinaryDataCmd = &cobra.Command{
	Use:   "add-binary [dataName] [dataFile] [key]",
	Short: "Add binary data",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		dataName := args[0]
		dataFile := args[1]
		key := args[2]

		ctx := getSessionCtx(context.Background())

		client := repository.GetRepository()
		data, err := os.ReadFile(dataFile)
		handleError(err)

		encryptedData := encryptData(key, string(data))

		err = client.AddBytes(ctx, repository.BinaryData{
			BytesName: dataName,
			Value:     []byte(encryptedData),
		})
		handleError(err)
	},
}

var GetBinaryDataCmd = &cobra.Command{
	Use:   "get-binary [dataName] [key]",
	Short: "Get binary data",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		dataName := args[0]
		key := args[1]

		ctx := getSessionCtx(context.Background())

		client := repository.GetRepository()
		binaryData, err := client.GetBytes(ctx, dataName)
		handleError(err)

		decryptedData := decryptData(key, string(binaryData.Value))

		log.Println("Data Name:", binaryData.BytesName)
		log.Println("Data Content:", string(decryptedData))
	},
}
