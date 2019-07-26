package general

import (
	"log"
	"os/exec"
)

func RunCMD(name string, args ...string) *exec.Cmd {
	var err error

	cmd := exec.Command(name, args...)
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()
	cmd.Start()

	go func() {
		for {
			buf := make([]byte, 1024)
			n, err := stderr.Read(buf)

			if n > 0 {
				log.Printf("READ OUT %s", string(buf[:n]))
			}

			if n == 0 {
				break
			}

			if err != nil {
				log.Printf("READ OUT %v", err)
				return
			}
		}
	}()
	go func() {
		for {
			buf := make([]byte, 1024)
			n, err := stdout.Read(buf)

			if n == 0 {
				break
			}

			if n > 0 {
				log.Printf("read out %s", string(buf[:n]))
			}

			if n == 0 {
				break
			}

			if err != nil {
				log.Printf("read out %v", err)
				return
			}

		}
	}()

	err = cmd.Wait()
	if err != nil {
		log.Printf("cmd err is %v", err)
	}
	return cmd
}
