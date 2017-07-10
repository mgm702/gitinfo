package cmds

import (
	"testing"

	"github.com/urfave/cli"
)

func TestForksCmd(t *testing.T) {
	var cmdTest = []struct {
		Name    string
		Aliases []string
		Usage   string
	}{
		{"forks", []string{"f"}, "Forks for a project"},
	}

	app := cli.NewApp()
	ForksCmd(app)

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
