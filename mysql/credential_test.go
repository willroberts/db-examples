package mysqltest

import "testing"

func TestLoadJSON(t *testing.T) {
	c := &credentialManager{}
	err := c.LoadJSON(".credentials")
	if err != nil {
		t.FailNow()
	}
}

func TestGetPassword(t *testing.T) {
	c := &credentialManager{}
	_ = c.LoadJSON(".credentials")
	if _, err := c.GetPassword("testuser"); err != nil {
		t.FailNow()
	}
}
