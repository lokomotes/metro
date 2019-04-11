package create

import (
	"context"
	"errors"
	"io/ioutil"
	"os"
	"path"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
)

func createDotNet(opt *Option) error {
	var (
		src []string
	)

	//
	// collecting srource(s)
	//
	if opt.isDir {
		return errors.New("Not implemented")
	} else {
		tmp, err := ioutil.TempDir("", "")
		if err != nil {
			return err
		}
		defer os.RemoveAll(tmp)

		dst := path.Join(tmp, "Program.cs")
		cp(opt.SrcPath, dst)

		src = append(src, dst)
	}

	//
	// tarring content
	//
	content, err := tarring(src)
	if err != nil {
		return err
	}
	defer os.Remove(content.Name())

	//
	// create
	//

	cli, err := getDockerCli()
	if err != nil {
		return err
	}

	ctx := context.Background()

	cRes, err := cli.ContainerCreate(ctx, &container.Config{
		Image: opt.stRepo,
	}, nil, nil, "")
	if err != nil {
		return err
	}
	defer cli.ContainerRemove(ctx, cRes.ID, types.ContainerRemoveOptions{})

	err = cli.CopyToContainer(
		ctx, cRes.ID, "/usr/station/App", content, types.CopyToContainerOptions{})
	if err != nil {
		return err
	}

	_, err = cli.ContainerCommit(ctx, cRes.ID, types.ContainerCommitOptions{
		Reference: opt.OutRepo,
	})
	if err != nil {
		return err
	}

	return nil
}
