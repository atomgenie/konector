package ssh

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/atomgenie/konector/github"
)

// SaveKeys Save SSH keys in authorized_keys
func SaveKeys(keys github.APISSHKeys) error {
	homeDir := os.Getenv("HOME")

	if homeDir == "" {
		return fmt.Errorf("HOME environment variable is not defined")
	}

	var outData strings.Builder
	for _, sshKey := range keys {
		rawSSH := sshKey.Key
		outData.WriteString(rawSSH)
		outData.WriteRune('\n')
	}

	err := ioutil.WriteFile(path.Join(homeDir, ".ssh/authorized_keys"), []byte(outData.String()), 0644)

	return err
}
