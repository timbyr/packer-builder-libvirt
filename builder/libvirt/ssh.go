package libvirt

import (
	"fmt"

	gossh "code.google.com/p/go.crypto/ssh"
	"github.com/mitchellh/multistep"
	commonssh "github.com/mitchellh/packer/common/ssh"
	"github.com/mitchellh/packer/communicator/ssh"
)

func sshAddress(state multistep.StateBag) (string, error) {
	guest_ip := state.Get("guest_ip").(string)
	sshHostPort := state.Get("config").(*config).SSHPort
	return fmt.Sprintf("%s:%d", guest_ip, sshHostPort), nil
}

func sshConfig(state multistep.StateBag) (*gossh.ClientConfig, error) {
	config := state.Get("config").(*config)

	auth := []gossh.AuthMethod{
		gossh.Password(config.SSHPassword),
		gossh.KeyboardInteractive(
			ssh.PasswordKeyboardInteractive(config.SSHPassword)),
	}

	if config.SSHKeyPath != "" {
		signer, err := commonssh.FileSigner(config.SSHKeyPath)
		if err != nil {
			return nil, err
		}

		auth = append(auth, gossh.PublicKeys(signer))
	}

	return &gossh.ClientConfig{
		User: config.SSHUser,
		Auth: auth,
	}, nil
}
