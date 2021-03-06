package libvirt

import (
	"fmt"
	"time"

	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
)

// stepBootWait waits the configured time period.
type stepBootWait struct{}

func (s *stepBootWait) Run(state multistep.StateBag) multistep.StepAction {
	config := state.Get("config").(*config)
	ui := state.Get("ui").(packer.Ui)

	if int64(config.bootWait) > 0 {
		ui.Say(fmt.Sprintf("Waiting %s for boot...", config.bootWait))
		time.Sleep(config.bootWait)
	}

	return multistep.ActionContinue
}

func (s *stepBootWait) Cleanup(state multistep.StateBag) {}
