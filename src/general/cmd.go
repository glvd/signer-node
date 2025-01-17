package general

import (
	"log"
	"os/exec"
)

// RunCMD ...
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
				log.Printf("[CMD OUT] %s", string(buf[:n]))
			}

			if n == 0 {
				break
			}

			if err != nil {
				log.Printf("[CMD ERROR] %v", err)
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
				log.Printf("[CMD OUT] %s", string(buf[:n]))
			}

			if n == 0 {
				break
			}

			if err != nil {
				log.Printf("[CMD ERROR]%v", err)
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
