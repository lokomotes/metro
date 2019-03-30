package metro

import (
	"context"
	code "net/http"

	api "github.com/lokomotes/metro/api"

	log "github.com/sirupsen/logrus"
)

// Load creates Station
func (h *CtlHandle) Load(ctx context.Context, in *api.LoadRequest) (*api.Response, error) {
	var (
		res     = &api.Response{Code: code.StatusOK}
		station = in.GetStation()
		desc    = newInstDesc(in.GetUserID(), station.GetImage())
	)

	logger := log.WithFields(log.Fields{
		"userID": desc.userID,
		"image":  desc.image,
		"name":   station.GetName(),
	})

	logger.Info("Load is requested")

	err := newInstance(desc, nil)

	switch err {
	default:
		res.Code = code.StatusInternalServerError
		logger.Warn(err)
		return res, nil
	case errExists:
	case errNExists:
		res.Code = code.StatusNotFound
		logger.Warn(err)
		return res, nil
	case nil:
		logger.Info("new instance is created")
	}

	return res, nil
}
