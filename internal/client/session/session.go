// Package session provides session management.
package session

import (
	"encoding/json"

	"github.com/dgraph-io/badger/v4"
)

// SesionKey is the key of session.
const sessionKey = "session"

// SessiomManager is a singleton of session manager.
var sessiomManager SessionManager

// SessionManager is a session manager inteface.
type SessionManager interface {
	GetSession() (Session, error)
	SetSession(session Session) error
	Close() error
}

// SessionManagerBadger is a decorator of badger.
type SessionManagerBadger struct {
	db *badger.DB
}

// Session is a session struct with jwt.
type Session struct {
	JWT string `json:"jwt"`
}

// InitAndGetSessionManager create and returns a session manager with encryption.
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

// SetSessionManager saves the session manager to singleton.
func SetSessionManager(sm SessionManager) {
	sessiomManager = sm
}

// GetSessionManager returns the session manager.
func GetSessionManager() SessionManager {
	return sessiomManager
}

// GetSession returns the current session.
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

// SetSession saves the current session.
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

// Close closes the session manager conntection with badger.
func (m *SessionManagerBadger) Close() error {
	err := m.db.Close()
	return err
}
