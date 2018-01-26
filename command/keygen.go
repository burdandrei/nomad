package command

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"strings"
)

// KeygenCommand is a Command implementation that generates an encryption
// key for use in `nomad agent`.
type KeygenCommand struct {
	Meta
}

func (c *KeygenCommand) Synopsis() string {
	return "Generates a new encryption key"
}

func (c *KeygenCommand) Help() string {
	helpText := `
Usage: nomad keygen

  Generates a new encryption key that can be used to configure the
  agent to encrypt traffic. The output of this command is already
  in the proper format that the agent expects.
`
	return strings.TrimSpace(helpText)
}

func (c *KeygenCommand) Run(_ []string) int {
	key := make([]byte, 16)
	n, err := rand.Reader.Read(key)
	if err != nil {
		c.Ui.Error(fmt.Sprintf("Error reading random data: %s", err))
		return 1
	}
	if n != 16 {
		c.Ui.Error(fmt.Sprintf("Couldn't read enough entropy. Generate more entropy!"))
		return 1
	}

	c.Ui.Output(base64.StdEncoding.EncodeToString(key))
	return 0
}
