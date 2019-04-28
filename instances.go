package metro

import (
	"context"
	"strconv"
	"strings"

	api "github.com/lokomotes/metro/api"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	mounttypes "github.com/docker/docker/api/types/mount"
	volumetypes "github.com/docker/docker/api/types/volume"
)

type token api.Token

var (
	// userID:imageName:instBody
	instances = make(map[string]map[string]instBody)

	containers = make(map[string]instDesc)
)

type instDesc struct {
	userID string
	image  string
}

type instBody struct {
	contID   string
	transmit chan api.Signal
}

func newInstDesc(userID, image string) *instDesc {
	if !strings.Contains(image, ":") {
		image = image + ":latest"
	}

	return &instDesc{
		userID: userID,
		image:  image,
	}
}

func (token *token) getDesc() (instDesc, bool) {
	desc, ok := containers[((*api.Token)(token)).GetId()]
	return desc, ok
}

func (desc *instDesc) getBody() (instBody, bool) {
	body, ok := instances[desc.userID][desc.image]
	return body, ok
}

// user ID 넘겨주고 없으면 volume 만들어서 마운트?
func (desc *instDesc) createInstance() (string, error) {
	if _, err := DckrCli.VolumeInspect(context.Background(), desc.userID); err != nil {
		DckrCli.VolumeCreate(context.Background(), volumetypes.VolumeCreateBody{
			Name: desc.userID,
		})
	}

	res, err := DckrCli.ContainerCreate(context.Background(), &container.Config{
		Image: desc.image,
		Env: []string{
			"LOKO_METRO_HOST=" + metroContName,
			"LOKO_METRO_PORT=" + strconv.Itoa(int(serveOpts.Port)),
		},
	}, &container.HostConfig{
		NetworkMode: metroContNetMode,
		Mounts: []mounttypes.Mount{
			mounttypes.Mount{
				Type:   mounttypes.TypeVolume,
				Source: desc.userID,
				Target: "/usr/cargo",
			},
		},
	}, nil, "")

	if err != nil {
		return "", err
	}

	return res.ID, err
}

func startInstance(id string) error {
	return DckrCli.ContainerStart(
		context.Background(), id,
		types.ContainerStartOptions{},
	)
}

func newInstance(desc *instDesc, sig *api.Signal) error {
	pool, ok := instances[desc.userID]
	if !ok {
		inst := make(map[string]instBody)
		instances[desc.userID] = inst
		pool = inst
	}

	if body, ok := pool[desc.image]; ok {
		if sig != nil {
			body.transmit <- *sig
		}
		return errExists
	}

	tc := make(chan api.Signal, 3)
	if sig != nil {
		tc <- *sig
	}

	pool[desc.image] = instBody{transmit: tc}

	contID, err := desc.createInstance()

	if err != nil {
		delete(pool, desc.image)

		if i, ok := err.(interface{ NotFound() bool }); ok && i.NotFound() {
			return errNExists
		}

		return err
	}

	body, _ := pool[desc.image]
	body.contID = contID
	containers[contID] = *desc

	err = startInstance(contID)
	if err != nil {
		return err
	}

	return nil
}
