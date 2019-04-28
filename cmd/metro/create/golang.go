package create

import (
	"context"
	"go/build"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/docker/docker/api/types"
)

const golangDockerfile = `
COPY . /go/.
RUN /bin/sh /go/create.sh `

func createGolang(opt *Option) error {
	var (
		src        []string // build context
		usrPkg     string
		dockerfile string
	)

	//
	// resolve usrPkg
	//
	if opt.isDir {
		usrPkg = strings.TrimPrefix(opt.SrcPath, build.Default.GOPATH+"/src/")
	} else {
		usrPkg = "app"
	}

	//
	// add Dockerfile
	//
	if tmp, err := ioutil.TempFile("", ""); err == nil {
		p := tmp.Name()
		defer os.Remove(p)
		tmp.WriteString("FROM " + opt.stRepo + golangDockerfile + usrPkg)
		dockerfile = filepath.Base(p)

		src = append(src, p)
	} else {
		return err
	}

	//
	// collecting source(s)
	//
	if tmp, err := ioutil.TempDir("", ""); err == nil {
		defer os.RemoveAll(tmp)

		var dst string

		if opt.isDir {
			dst = path.Join(tmp, "src", usrPkg)
			os.MkdirAll(dst, os.ModePerm)
		} else {
			dst = path.Join(tmp, "/src/app/app.go")
		}

		cp(opt.SrcPath, dst)
		src = append(src, path.Join(tmp, "src"))
	} else {
		return err
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
