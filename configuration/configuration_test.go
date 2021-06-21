package configuration

import (
	"os"
	"testing"
)

func TestUserConfiguration(t *testing.T) {
	user := os.Getenv("USER")
	home := os.Getenv("HOME")

	err := os.Setenv("USER", "")
	if err != nil {
		t.Errorf("error while wiping USER prior to configuration setup: %s", err.Error())
	}

	err = os.Setenv("HOME", "")
	if err != nil {
		t.Errorf("error while wiping HOME prior to configuration setup: %s", err.Error())
	}

	err = os.Setenv("USERNAME", "")
	if err != nil {
		t.Errorf("error while wiping USERNAME prior to configuration setup: %s", err.Error())
	}

	c := Load("../testdata/configuration/test.json")
	if c.User != user && c.User != os.Getenv("USER") {
		t.Errorf("User was not reset properly in the environment or configuration: %s, %s", c.User, os.Getenv("USER"))
	}

	if c.Home != home && c.User != os.Getenv("HOME") {
		t.Errorf("Home was not reset properly in the environment or configuration: %s, %s", c.Home, os.Getenv("HOME"))
	}

	if c.PS1 != "[{user}@{hostname}]: >" {
		t.Errorf("PS1 is nto set properly. Got %s", c.PS1)
	}

	c2 := Load("../testdata/configuration/test.json")
	if c2.PS1 != "[{user}@{hostname}]: >" {
		t.Errorf("PS1 is nto set properly. Got %s", c2.PS1)
	}

	c3 := Load("../testdata/configuration/turtlerc.json")
	c3.Reload()
}
