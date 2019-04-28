package create

import (
	"context"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"github.com/docker/docker/api/types"
)

const dotnetDockerfile = `
COPY . /usr/station/App/.
RUN /bin/sh /usr/station/create.sh`

func createDotNet(opt *Option) error {
	var (
		src        []string
		dockerfile string
	)

	//
	// add Dockerfile
	//
	if tmp, err := ioutil.TempFile("", ""); err == nil {
		p := tmp.Name()
		defer os.Remove(p)
		tmp.WriteString("FROM " + opt.stRepo + dotnetDockerfile)
		dockerfile = filepath.Base(p)

		src = append(src, p)
	} else {
		return err
	}

	//
	// collecting source(s)
	//
	if opt.isDir {
		fs, err := ioutil.ReadDir(opt.SrcPath)
		if err != nil {
			return err
		}

		for _, f := range fs {
			src = append(src, path.Join(opt.SrcPath, f.Name()))
		}
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

	bRes, err := cli.ImageBuild(ctx, content, types.ImageBuildOptions{
		Tags:       []string{opt.OutRepo},
		Dockerfile: dockerfile,
		Remove:     true,
	})
	if err != nil {
		return err
	}
	defer bRes.Body.Close()

	io.Copy(os.Stdout, bRes.Body)

	return nil
}
