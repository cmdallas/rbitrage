package config

import (
	"os"
	"testing"
)

const (
	badConfigPath = "/fake/path/lkjlkasjalkdj"
	dirConfigPath = "/"
)

var (
	goPath         string = os.Getenv("GOPATH")
	goodConfigPath string = goPath + "/src/github.com/cmdallas/rbitrage/examples/config/.rbitrage.yaml"
)

func TestValidateConfigPath(t *testing.T) {
	if err := ValidateConfigPath(goodConfigPath); err != nil {
		t.Fail()
	}
	if err := ValidateConfigPath(badConfigPath); err == nil {
		t.Fail()
	}
	if err := ValidateConfigPath(dirConfigPath); err == nil {
		t.Fail()
	}
}

func TestNewConfig(t *testing.T) {
	c, err := NewConfig(goodConfigPath)
	switch {
	case err != nil:
		t.Log(err)
		t.Fail()
	case c.Applications != nil:
		return
	}

	_, err = NewConfig(dirConfigPath)
	if err == nil {
		t.Fail()
	}

	_, err = NewConfig(badConfigPath)
	if err == nil {
		t.Fail()
	}
}
