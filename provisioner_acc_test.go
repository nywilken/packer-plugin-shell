package shell_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/packer/command"
	"github.com/hashicorp/packer/helper/tests"
)

func TestBuildShellProvisionerWithBuildVariablesSharing(t *testing.T) {
	testshelper.AccTestPreValidate(t)

	UUID, _ := uuid.GenerateUUID()
	os.Setenv("PACKER_RUN_UUID", UUID)
	c := &command.BuildCommand{
		Meta: testshelper.TestMetaFile(t),
	}

	file := "provisioner.shell." + UUID + ".txt"
	defer testshelper.CleanupFiles(file)

	args := []string{
		filepath.Join("./test-fixtures", "shell-provisioner.json"),
	}
	if code := c.Run(args); code != 0 {
		testshelper.FatalCommand(t, c.Meta)
	}

	if !testshelper.FileExists(file) {
		t.Errorf("Expected to find %s", file)
	} else {
		helper := testshelper.AWSHelper{
			Region:  "us-east-1",
			AMIName: "packer-test-shell-interpolate",
		}
		helper.CleanUpAmi(t)
	}
}
