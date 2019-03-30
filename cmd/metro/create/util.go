package create

import (
	"io/ioutil"
	"os"

	"github.com/docker/docker/client"
	"github.com/mholt/archiver"
	"github.com/otiai10/copy"
)

var dockerCli *client.Client

func cp(src string, dst string) error {
	return copy.Copy(src, dst)
	// i, err := os.Open(src)
	// defer i.Close()
	// if err != nil {
	// 	return err
	// }

	// o, err := os.Open(dst)
	// defer o.Close()
	// if err != nil {
	// 	return err
	// }

	// _, err = io.Copy(o, i)
	// if err != nil {
	// 	return err
	// }

	// // o.Sync()

	// return nil
}

func tarring(src []string) (*os.File, error) {
	tmp, err := ioutil.TempFile("", "*.tar")
	if err != nil {
		return nil, err
	}

	tar := archiver.Tar{OverwriteExisting: true}
	if err = tar.Archive(src, tmp.Name()); err != nil {
		return nil, err
	}

	return tmp, nil
}

func getDockerCli() (*client.Client, error) {
	var err error
	if dockerCli == nil {
		dockerCli, err = client.NewClientWithOpts(client.WithVersion("1.39"))
		if err != nil {
			return nil, err
		}
	}
	return dockerCli, nil
}
