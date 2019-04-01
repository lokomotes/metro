package create

import (
	"context"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/docker/distribution/reference"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
)

// Runtime is identifier for Metro activity runtime.
type Runtime int

const (
	// Auto is an identifier that selects runtime automatically.
	Auto Runtime = iota
	// NodeJS is an identifier for Node.js.
	NodeJS
	// Golang is an identifier for Golang.
	Golang
)

// Option holds options to create Metro activity image.
type Option struct {
	Runtime        Runtime
	RuntimeVersion string
	StationVersion string
	SrcPath        string
	OutRepo        string

	isDir  bool
	stRepo string
}

func (opt *Option) resolve() error {
	//
	// resolve SrcPath
	// "." by default
	// the SrcPath become absolute
	//
	if len(opt.SrcPath) == 0 {
		opt.SrcPath = "."
	}
	if path, err := filepath.Abs(opt.SrcPath); err == nil {
		opt.SrcPath = path
	} else {
		return err
	}

	//
	// resolve Runtime
	// the Runtime is decided by extension of the SrcPath
	// Auto identifier valid only for a file
	// if the Runtime is Auto and SrcPath is a directory, return an error
	//
	if stat, err := os.Stat(opt.SrcPath); err == nil {
		opt.isDir = stat.IsDir()

		if opt.isDir {
			if opt.Runtime == Auto {
				return errors.New("the runtime must be specified for the directory")
			}
		} else if rt, err := resolveRuntime(filepath.Base(opt.SrcPath)); err == nil {
			opt.Runtime = rt
		} else {
			return err
		}
	}

	//
	// resolve OutRepo
	// check if valid repository name
	// the OutRepo decided by extension of the SrcPath if it is not specified
	// tagging "latest" by default if it is not tagged
	//
	if len(opt.OutRepo) == 0 {
		opt.OutRepo = filepath.Base(opt.SrcPath)

		if !opt.isDir {
			opt.OutRepo = strings.TrimSuffix(opt.OutRepo, filepath.Ext(opt.OutRepo))
		}
	}
	if repo, err := normalizeRepo(opt.OutRepo); err == nil {
		opt.OutRepo = repo
	} else {
		return err
	}

	//
	// resolve RuntimeVersion
	// set RuntimeVersion if it is not specified
	// default value can be seen `getDefaultVersionOfRuntime`
	//
	if len(opt.RuntimeVersion) == 0 {
		v, err := getDefaultVersionOfRuntime(opt.Runtime)
		if err != nil {
			return err
		}
		opt.RuntimeVersion = v
	}

	//
	// resolve stRepo
	// stRepo is set according to RuntimeVersion
	// tagged "latest" by default if StationVersion not specified
	// the station image is pulled if it is not exists locally
	//
	if id, err := getIdentifierOfRuntime(opt.Runtime); err == nil {
		opt.stRepo = "lokomotes/station-" + id // + opt.RuntimeVersion
	} else {
		return err
	}
	if len(opt.StationVersion) > 0 {
		opt.stRepo = opt.stRepo + ":" + opt.StationVersion
	}
	if repo, err := normalizeRepo(opt.stRepo); err == nil {
		opt.stRepo = repo
	} else {
		return err
	}
	if err := ensureStationRepo(opt.stRepo); err != nil {
		return err
	}

	return nil
}

func normalizeRepo(rwaRepo string) (string, error) {
	named, err := reference.ParseNormalizedNamed(rwaRepo)
	if err != nil {
		return "", err
	}

	if reference.IsNameOnly(named) {
		return rwaRepo + ":latest", nil
	}

	return rwaRepo, nil
}

func resolveRuntime(filename string) (Runtime, error) {
	var (
		rst Runtime
		ext = filepath.Ext(filename)[1:]
	)

	if ext == "" {
		return Auto, errors.New("the extension must be specified in the filename")
	}

	switch ext {
	default:
		rst = Auto

	case "js":
		rst = NodeJS

	case "a":
		fallthrough
	case "go":
		rst = Golang
	}

	if rst == Auto {
		return Auto, errors.New("unsupported extension: " + ext)
	}

	return rst, nil
}

func getIdentifierOfRuntime(runtime Runtime) (string, error) {
	switch runtime {
	default:
		return "", errors.New("unexpected runtime identifier")

	case NodeJS:
		return "node", nil

	case Golang:
		return "go", nil
	}
}

func getDefaultVersionOfRuntime(runtime Runtime) (string, error) {
	switch runtime {
	default:
		return "", errors.New("unexpected runtime identifier")

	case NodeJS:
		return "10", nil

	case Golang:
		return "1.11", nil
	}
}

func ensureStationRepo(repo string) error {
	cli, err := getDockerCli()
	if err != nil {
		return err
	}

	ctx := context.Background()

	imgs, err := cli.ImageList(ctx, types.ImageListOptions{
		Filters: filters.NewArgs(filters.Arg("reference", repo)),
	})
	if err != nil {
		return err
	}

	if len(imgs) > 0 {
		return nil
	}

	out, err := cli.ImagePull(ctx, repo, types.ImagePullOptions{})
	if err != nil {
		return err
	}
	defer out.Close()

	io.Copy(os.Stdout, out)

	return nil
}
