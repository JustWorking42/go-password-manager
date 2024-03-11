package commands

import (
	"context"
	"os"
	"sync"
	"testing"

	"github.com/JustWorking42/go-password-manager/internal/client/passwordreader"
	"github.com/JustWorking42/go-password-manager/internal/client/repository"
	"github.com/JustWorking42/go-password-manager/internal/client/session"
	"github.com/JustWorking42/go-password-manager/internal/common/defens"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRegisterCommand(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := repository.NewMockRepository(mockCtrl)
	mockSession := session.NewMockSessionManager(mockCtrl)
	mockPasswordReader := passwordreader.NewMockPasswordReader(mockCtrl)

	passwordreader.SetPasswordReader(mockPasswordReader)
	repository.SetRepository(mockRepo)
	session.SetSessionManager(mockSession)

	testCases := []struct {
		name          string
		login         string
		password      string
		expectedError bool
		setup         func()
	}{
		{
			name:          "Successful registration",
			login:         "testuser",
			password:      "testpass",
			expectedError: false,
			setup: func() {
				mockRepo.EXPECT().Register(gomock.Any(), gomock.Any(), gomock.Any()).Return("jwt", nil)
				mockSession.EXPECT().SetSession(gomock.Any()).Return(nil)
				mockPasswordReader.EXPECT().ReadPassword().Return([]byte("testpass"), nil)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.setup()
			cmd := RegisterCmd
			oldStdin := os.Stdin
			r, w, _ := os.Pipe()
			os.Stdin = r
			wg := sync.WaitGroup{}
			wg.Add(1)
			go func() {
				defer wg.Done()
				defer w.Close()
				w.Write([]byte(tc.login))
				w.Write([]byte("\n"))
			}()
			defer func() { os.Stdin = oldStdin }()

			err := cmd.Execute()

			wg.Wait()
			if (err != nil) != tc.expectedError {
				t.Errorf("Expected error: %v, got: %v", tc.expectedError, err)
			}
		})
	}
}

func TestLoginCommand(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := repository.NewMockRepository(mockCtrl)
	mockSession := session.NewMockSessionManager(mockCtrl)
	mockPasswordReader := passwordreader.NewMockPasswordReader(mockCtrl)

	passwordreader.SetPasswordReader(mockPasswordReader)
	repository.SetRepository(mockRepo)
	session.SetSessionManager(mockSession)

	testCases := []struct {
		name          string
		login         string
		password      string
		expectedError bool
		setup         func()
	}{
		{
			name:          "Successful login",
			login:         "testuser",
			password:      "testpass",
			expectedError: false,
			setup: func() {
				mockRepo.EXPECT().Login(gomock.Any(), gomock.Any(), gomock.Any()).Return("jwt", nil)
				mockSession.EXPECT().SetSession(gomock.Any()).Return(nil)
				mockPasswordReader.EXPECT().ReadPassword().Return([]byte("testpass"), nil)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.setup()
			cmd := LoginCmd
			oldStdin := os.Stdin
			r, w, _ := os.Pipe()
			os.Stdin = r
			wg := sync.WaitGroup{}
			wg.Add(1)
			go func() {
				defer wg.Done()
				defer w.Close()
				w.Write([]byte(tc.login))
				w.Write([]byte("\n"))
			}()
			defer func() { os.Stdin = oldStdin }()

			err := cmd.Execute()

			wg.Wait()
			if (err != nil) != tc.expectedError {
				t.Errorf("Expected error: %v, got: %v", tc.expectedError, err)
			}
		})
	}
}

func TestAddPasswordCommand(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := repository.NewMockRepository(mockCtrl)
	mockSession := session.NewMockSessionManager(mockCtrl)

	repository.SetRepository(mockRepo)
	session.SetSessionManager(mockSession)

	testCases := []struct {
		name          string
		passwordName  string
		passwordLogin string
		passwordValue string
		key           string
		expectedError bool
		setup         func()
	}{
		{
			name:          "Successful add password",
			passwordName:  "testName",
			passwordLogin: "testLogin",
			passwordValue: "testPassword",
			key:           "testKey",
			expectedError: false,
			setup: func() {
				mockRepo.EXPECT().AddPassword(gomock.Any(), gomock.Any()).Return(nil)
				mockSession.EXPECT().GetSession().Return(session.Session{JWT: "jwt"}, nil)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.setup()
			cmd := AddPasswordCmd
			cmd.SetArgs([]string{tc.passwordName, tc.passwordLogin, tc.passwordValue, tc.key})

			err := cmd.ExecuteContext(context.Background())

			if (err != nil) != tc.expectedError {
				t.Errorf("Expected error: %v, got: %v", tc.expectedError, err)
			}
		})
	}
}

func TestGetPasswordCommand(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := repository.NewMockRepository(mockCtrl)
	mockSession := session.NewMockSessionManager(mockCtrl)

	repository.SetRepository(mockRepo)
	session.SetSessionManager(mockSession)

	testCases := []struct {
		name          string
		passwordName  string
		key           string
		expectedError bool
		setup         func(repo *repository.MockRepository)
	}{
		{
			name:          "Successful get password",
			passwordName:  "testName",
			key:           "testKey",
			expectedError: false,
			setup: func(repo *repository.MockRepository) {
				mockSession.EXPECT().GetSession().Return(session.Session{JWT: "jwt"}, nil)
				encryptedLogin, err := defens.Encrypt([]byte("testKey"), "testLogin")
				assert.NoError(t, err)
				encryptedPassword, err := defens.Encrypt([]byte("testKey"), "testPassword")
				assert.NoError(t, err)
				mockRepo.EXPECT().GetPassword(gomock.Any(), gomock.Any()).Return(repository.Password{
					Name:     "testName",
					Login:    encryptedLogin,
					Password: encryptedPassword,
				}, nil)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.setup(mockRepo)
			cmd := GetPasswordCmd
			cmd.SetArgs([]string{tc.passwordName, tc.key})

			err := cmd.ExecuteContext(context.Background())

			if (err != nil) != tc.expectedError {
				t.Errorf("Expected error: %v, got: %v", tc.expectedError, err)
			}
		})
	}
}

func TestAddCardCommand(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := repository.NewMockRepository(mockCtrl)
	mockSession := session.NewMockSessionManager(mockCtrl)

	repository.SetRepository(mockRepo)
	session.SetSessionManager(mockSession)

	testCases := []struct {
		name          string
		cardName      string
		cardNumber    string
		cardCVC       string
		cardDate      string
		cardFI        string
		key           string
		expectedError bool
		setup         func()
	}{
		{
			name:          "Successful add card",
			cardName:      "testCard",
			cardNumber:    "1234567890123456",
			cardCVC:       "123",
			cardDate:      "12/24",
			cardFI:        "Test FI",
			key:           "testKey",
			expectedError: false,
			setup: func() {
				mockRepo.EXPECT().AddCard(gomock.Any(), gomock.Any()).Return(nil)
				mockSession.EXPECT().GetSession().Return(session.Session{JWT: "jwt"}, nil)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.setup()
			cmd := AddCardCmd
			cmd.SetArgs([]string{tc.cardName, tc.cardNumber, tc.cardCVC, tc.cardDate, tc.cardFI, tc.key})

			err := cmd.ExecuteContext(context.Background())

			if (err != nil) != tc.expectedError {
				t.Errorf("Expected error: %v, got: %v", tc.expectedError, err)
			}
		})
	}
}

func TestGetCardCommand(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := repository.NewMockRepository(mockCtrl)
	mockSession := session.NewMockSessionManager(mockCtrl)

	repository.SetRepository(mockRepo)
	session.SetSessionManager(mockSession)

	testCases := []struct {
		name          string
		cardName      string
		key           string
		expectedError bool
		setup         func(repo *repository.MockRepository)
	}{
		{
			name:          "Successful get card",
			cardName:      "testCard",
			key:           "testKey",
			expectedError: false,
			setup: func(repo *repository.MockRepository) {
				mockSession.EXPECT().GetSession().Return(session.Session{JWT: "jwt"}, nil)
				encryptedCardNumber, err := defens.Encrypt([]byte("testKey"), "1234567890123456")
				assert.NoError(t, err)
				encryptedCardCVC, err := defens.Encrypt([]byte("testKey"), "123")
				assert.NoError(t, err)
				encryptedCardDate, err := defens.Encrypt([]byte("testKey"), "12/24")
				assert.NoError(t, err)
				encryptedCardFI, err := defens.Encrypt([]byte("testKey"), "Test FI")
				assert.NoError(t, err)
				mockRepo.EXPECT().GetCard(gomock.Any(), gomock.Any()).Return(repository.Card{
					CardName:   "testCard",
					CardNumber: encryptedCardNumber,
					CardCVC:    encryptedCardCVC,
					CardDate:   encryptedCardDate,
					CardFI:     encryptedCardFI,
				}, nil)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.setup(mockRepo)
			cmd := GetCardCmd
			cmd.SetArgs([]string{tc.cardName, tc.key})

			err := cmd.ExecuteContext(context.Background())

			if (err != nil) != tc.expectedError {
				t.Errorf("Expected error: %v, got: %v", tc.expectedError, err)
			}
		})
	}
}

func TestAddNoteCommand(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := repository.NewMockRepository(mockCtrl)
	mockSession := session.NewMockSessionManager(mockCtrl)

	repository.SetRepository(mockRepo)
	session.SetSessionManager(mockSession)

	testCases := []struct {
		name          string
		noteName      string
		noteContent   string
		key           string
		expectedError bool
		setup         func()
	}{
		{
			name:          "Successful add note",
			noteName:      "testNote",
			noteContent:   "This is a test note.",
			key:           "testKey",
			expectedError: false,
			setup: func() {
				mockRepo.EXPECT().AddNote(gomock.Any(), gomock.Any()).Return(nil)
				mockSession.EXPECT().GetSession().Return(session.Session{JWT: "jwt"}, nil)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.setup()
			cmd := AddNoteCmd
			cmd.SetArgs([]string{tc.noteName, tc.noteContent, tc.key})

			err := cmd.ExecuteContext(context.Background())

			if (err != nil) != tc.expectedError {
				t.Errorf("Expected error: %v, got: %v", tc.expectedError, err)
			}
		})
	}
}

func TestGetNoteCommand(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := repository.NewMockRepository(mockCtrl)
	mockSession := session.NewMockSessionManager(mockCtrl)

	repository.SetRepository(mockRepo)
	session.SetSessionManager(mockSession)

	testCases := []struct {
		name          string
		noteName      string
		key           string
		expectedError bool
		setup         func(repo *repository.MockRepository)
	}{
		{
			name:          "Successful get note",
			noteName:      "testNote",
			key:           "testKey",
			expectedError: false,
			setup: func(repo *repository.MockRepository) {
				mockSession.EXPECT().GetSession().Return(session.Session{JWT: "jwt"}, nil)
				encryptedNote, err := defens.Encrypt([]byte("testKey"), "This is a test note.")
				assert.NoError(t, err)
				mockRepo.EXPECT().GetNote(gomock.Any(), gomock.Any()).Return(repository.Note{
					NoteName: "testNote",
					Note:     encryptedNote,
				}, nil)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.setup(mockRepo)
			cmd := GetNoteCmd
			cmd.SetArgs([]string{tc.noteName, tc.key})

			err := cmd.ExecuteContext(context.Background())

			if (err != nil) != tc.expectedError {
				t.Errorf("Expected error: %v, got: %v", tc.expectedError, err)
			}
		})
	}
}

func TestAddBinaryDataCommand(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := repository.NewMockRepository(mockCtrl)
	mockSession := session.NewMockSessionManager(mockCtrl)

	repository.SetRepository(mockRepo)
	session.SetSessionManager(mockSession)
	file, err := os.CreateTemp("", "testDataFile")
	assert.NoError(t, err)

	testCases := []struct {
		name          string
		dataName      string
		dataFile      string
		key           string
		expectedError bool
		setup         func()
	}{
		{
			name:          "Successful add binary data",
			dataName:      "testData",
			dataFile:      file.Name(),
			key:           "testKey",
			expectedError: false,
			setup: func() {
				mockRepo.EXPECT().AddBytes(gomock.Any(), gomock.Any()).Return(nil)
				mockSession.EXPECT().GetSession().Return(session.Session{JWT: "jwt"}, nil)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.setup()
			cmd := AddBinaryDataCmd
			cmd.SetArgs([]string{tc.dataName, tc.dataFile, tc.key})

			err := cmd.ExecuteContext(context.Background())

			if (err != nil) != tc.expectedError {
				t.Errorf("Expected error: %v, got: %v", tc.expectedError, err)
			}
		})
	}
}

func TestGetBinaryDataCommand(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := repository.NewMockRepository(mockCtrl)
	mockSession := session.NewMockSessionManager(mockCtrl)

	repository.SetRepository(mockRepo)
	session.SetSessionManager(mockSession)

	testCases := []struct {
		name          string
		dataName      string
		key           string
		expectedError bool
		setup         func(repo *repository.MockRepository)
	}{
		{
			name:          "Successful get binary data",
			dataName:      "testData",
			key:           "testKey",
			expectedError: false,
			setup: func(repo *repository.MockRepository) {
				mockSession.EXPECT().GetSession().Return(session.Session{JWT: "jwt"}, nil)
				encryptedData, err := defens.Encrypt([]byte("testKey"), "This is test binary data.")
				assert.NoError(t, err)
				mockRepo.EXPECT().GetBytes(gomock.Any(), gomock.Any()).Return(repository.BinaryData{
					BytesName: "testData",
					Value:     []byte(encryptedData),
				}, nil)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.setup(mockRepo)
			cmd := GetBinaryDataCmd
			cmd.SetArgs([]string{tc.dataName, tc.key})

			err := cmd.ExecuteContext(context.Background())

			if (err != nil) != tc.expectedError {
				t.Errorf("Expected error: %v, got: %v", tc.expectedError, err)
			}
		})
	}
}
