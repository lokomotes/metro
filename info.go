package metro

import (
	"bufio"
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"io"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/docker/docker/api/types/container"
	log "github.com/sirupsen/logrus"
	context "golang.org/x/net/context"
)

var (
	metroContID      string
	metroContName    string
	metroContNetMode container.NetworkMode
	isMetroInCont    bool
)

func updateInfo() {
	makeRandID := func() string {
		id := make([]byte, 31)
		if _, err := io.ReadFull(rand.Reader, id); err != nil {
			log.Fatal(err)
		}

		return "zz" + hex.EncodeToString(id)
	}

	// resolve Metro container ID
	if runtime.GOOS == "windows" {
		func() {
			cmd := exec.Command("cmd", "/C", "sc", "query", "cexecsvc")
			if err := cmd.Run(); err != nil {
				metroContID = makeRandID()
				isMetroInCont = false
				return
			}

			cmd = exec.Command("cmd", "/C", "hostname")
			var out bytes.Buffer
			cmd.Stdout = &out
			if err := cmd.Run(); err != nil {
				log.Fatalf("Failed to resolve container ID: %v", err)
			}
			metroContID = strings.TrimSpace(out.String())
			isMetroInCont = true
		}()
	} else {
		func() {
			cgroup, err := os.Open("/proc/self/cgroup")
			defer cgroup.Close()
			if err != nil {
				metroContID = makeRandID()
				isMetroInCont = false
				return
			}

			scanner := bufio.NewScanner(cgroup)
			for scanner.Scan() {
				line := scanner.Text()
				if !strings.Contains(line, "docker") {
					continue
				}
				metroContID = strings.Split(line, "docker/")[1]
				isMetroInCont = true
				break
			}
		}()
	}

	// inspect Metro container
	{
		info, err := DckrCli.ContainerInspect(context.Background(), metroContID)
		if err != nil {
			log.Fatalf("Failed to inspect Metro container: %v", err)
		}
		metroContName = info.Name[1:]
		metroContNetMode = info.HostConfig.NetworkMode
	}
}
