package connection

import (
	"bufio"
	"context"
	"fmt"
	"os/exec"

	"github.com/rivo/tview"
)

func Share(ctx context.Context, screen *tview.TextView, stop chan bool) {

	fmt.Println(<-ctx.Done())


	return
	cmd := exec.CommandContext(ctx, "bash", "-c", "playit -s")

	out, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}

	if err = cmd.Start(); err != nil {
		panic(err)
	}

	s := bufio.NewScanner(out)

	for s.Scan() {
		select {
		case <-ctx.Done():
			cmd.Process.Kill()
			return

		default:
			fmt.Fprintln(screen, s.Text())
		}
	}

	<-stop

	if err = cmd.Wait(); err != nil {
		panic(err)
	}
}
