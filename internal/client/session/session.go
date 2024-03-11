package session

import (
	"encoding/json"

	"github.com/dgraph-io/badger/v4"
)

const sessionKey = "session"

var sessiomManager SessionManager

type SessionManager interface {
	GetSession() (Session, error)
	SetSession(session Session) error
	Close() error
}

type SessionManagerBadger struct {
	db *badger.DB
}

type Session struct {
	JWT string `json:"jwt"`
}

func InitAndGetSessionManager(config *Config) (*SessionManagerBadger, error) {

	opts := badger.DefaultOptions(config.SessionPath).
		WithEncryptionKey([]byte(config.DBSecret)).
		WithIndexCacheSize(5 << 20).WithLoggingLevel(badger.ERROR)

	db, err := badger.Open(opts)
	if err != nil {
		return nil, err
	}

	return &SessionManagerBadger{db: db}, nil
}

func SetSessionManager(sm SessionManager) {
	sessiomManager = sm
}

func GetSessionManager() SessionManager {
	return sessiomManager
}

func (m *SessionManagerBadger) GetSession() (Session, error) {

	var sessionBytes []byte
	err := m.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(sessionKey))
		if err != nil {
			return err
		}
		sessionBytes, err = item.ValueCopy(nil)
		return err
	})
	if err != nil {
		return Session{}, err
	}
	var session Session
	err = json.Unmarshal(sessionBytes, &session)
	if err != nil {
		return Session{}, err
	}
	return session, nil
}

func (m *SessionManagerBadger) SetSession(session Session) error {
	sessionBytes, err := json.Marshal(session)
	if err != nil {
		return err
	}
	err = m.db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(sessionKey), sessionBytes)
	})
	if err != nil {
		return err
	}
	return nil
}

func (m *SessionManagerBadger) Close() error {
	err := m.db.Close()
	return err
}
