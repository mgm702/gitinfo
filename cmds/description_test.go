package cmds

import (
	"testing"

	"github.com/urfave/cli"
)

func TestDescriptionCmd(t *testing.T) {
	var cmdTest = []struct {
		Name    string
		Aliases []string
		Usage   string
	}{
		{"description", []string{"d"}, "Description of a project"},
	}

	app := cli.NewApp()
	DescriptionCmd(app)

	for i := 0; i < len(app.Commands); i += 1 {
		if app.Commands[i].Name != cmdTest[i].Name {
			t.Errorf("fixInput: expected %d, actual %d", app.Commands[i].Name, cmdTest[i].Name)
		}

		for y := 0; y < len(app.Commands[i].Aliases); y += 1 {
			if app.Commands[i].Aliases[y] != cmdTest[i].Aliases[y] {
				t.Errorf("fixInput: expected %d, actual %d", app.Commands[i].Aliases[y], cmdTest[i].Aliases[y])
			}
		}

		if app.Commands[i].Usage != cmdTest[i].Usage {
			t.Errorf("fixInput: expected %d, actual %d", app.Commands[i].Usage, cmdTest[i].Usage)
		}
	}
}
