package connection

import (
	"bufio"
	"context"
	"fmt"
	"os/exec"

	"github.com/rivo/tview"
)

func Share(ctx context.Context, screen *tview.TextView, stop chan bool) {
	cmd := exec.CommandContext(ctx, "bash", "-c", "playit -s")

	out, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}

	if err = cmd.Start(); err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(out)

	for scanner.Scan() {
		select {
		case <-ctx.Done():
			_ = cmd.Process.Kill()

			return

		default:
			fmt.Fprintln(screen, scanner.Text())
		}
	}

	<-stop

	if err = cmd.Wait(); err != nil {
		panic(err)
	}
}
