package mysqltest

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type CredentialManager interface {
	GetPassword(username string) (string, error)
}

type credentialManager struct {
	credentials map[string]string
}

func NewCredentialManager(credFile string) (CredentialManager, error) {
	m, err := loadJsonMap(credFile)
	if err != nil {
		return nil, err
	}
	return &credentialManager{credentials: m}, nil
}

func loadJsonMap(filename string) (map[string]string, error) {
	b, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}

	return m, nil
}

func (c *credentialManager) GetPassword(username string) (string, error) {
	password, ok := c.credentials[username]
	if !ok {
		return "", errors.New(fmt.Sprintf("no password found for user %s", username))
	}
	return password, nil
}
