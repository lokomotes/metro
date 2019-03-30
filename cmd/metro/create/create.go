package create

import "errors"

// Create creates metro activity image.
func Create(opt *Option) (string, error) {
	if err := opt.resolve(); err != nil {
		return "", err
	}

	var err error

	switch opt.Runtime {
	default:
		return "", errors.New("unexpected runtime identifier")
	case NodeJS:
		err = createNodeJS(opt)
	case Golang:
		err = createGolang(opt)
	}

	if err != nil {
		return "", err
	}

	return opt.OutRepo, nil
}
