package libvirt

import (
	"fmt"
	"path/filepath"

	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
)

// This step copies the virtual disk that will be used as the
// hard drive for the virtual machine.
type stepCopyDisk struct{}

func (s *stepCopyDisk) Run(state multistep.StateBag) multistep.StepAction {
	config := state.Get("config").(*config)
	isoPath := state.Get("iso_path").(string)
	ui := state.Get("ui").(packer.Ui)
	path := filepath.Join(config.OutputDir, "disk.img")
	name := "disk.img"

	command := []string{
		"convert",
		"-O", config.DiskType,
		isoPath,
		path,
	}

	if config.DiskImage == false {
		return multistep.ActionContinue
	}

	ui.Say("Copying hard drive...")
	_, _, err := qemuImg(command...)
	if err != nil {
		err := fmt.Errorf("Error creating hard drive: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	state.Put("disk_filename", name)

	return multistep.ActionContinue
}

func (s *stepCopyDisk) Cleanup(state multistep.StateBag) {}
