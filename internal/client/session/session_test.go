package session

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetSession(t *testing.T) {
	config := &Config{
		SessionPath: "/tmp/testdb",
		DBSecret:    "testsecrettestse",
	}
	manager, err := InitAndGetSessionManager(config)
	assert.NoError(t, err)
	SetSessionManager(manager)

	sm := GetSessionManager()

	testCases := []struct {
		name          string
		session       Session
		expectedError bool
	}{
		{
			name:          "Set Session",
			session:       Session{JWT: "test_jwt"},
			expectedError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := sm.SetSession(tc.session)
			if tc.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}

	err = sm.Close()
	assert.NoError(t, err)
}

func TestGetSession(t *testing.T) {
	config := &Config{
		SessionPath: "/tmp/testdb",
		DBSecret:    "testsecrettestse",
	}
	manager, err := InitAndGetSessionManager(config)
	assert.NoError(t, err)
	SetSessionManager(manager)

	sm := GetSessionManager()

	err = sm.SetSession(Session{JWT: "test_jwt"})
	assert.NoError(t, err)

	testCases := []struct {
		name          string
		expectedError bool
	}{
		{
			name:          "Get Session",
			expectedError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := sm.GetSession()
			if tc.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}

	err = sm.Close()
	assert.NoError(t, err)
}
