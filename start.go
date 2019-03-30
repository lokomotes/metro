package metro

import (
	code "net/http"

	api "github.com/lokomotes/metro/api"

	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

func (desc *instDesc) start(name string, msg string) (string, error) {
	station := &api.Station{
		Id:   api.GenerateID(),
		Name: name,
	}

	if err := desc.transmit(api.Signal{
		Dst:     station,
		Control: api.Signal_START,
		Message: msg,
	}); err != nil {
		return "", err
	}

	return station.GetId(), nil
}

func startHandler(ctx context.Context, in *api.StartRequest) (*api.Response, error) {
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

	logger.Info("Start is requested")

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

	flowID, err := desc.start(station.GetName(), in.GetMessage())

	switch err {
	default:
		res.Code = code.StatusInternalServerError
		logger.Warn(err)
		return res, nil
	case errNExists:
		logger.Fatal("it should not be happend")
	case nil:
		logger.WithField(
			"flow", api.TruncateID(flowID),
		).Info("new flow is started")
	}

	return res, nil
}

// Start creates and run entry point Station
func (h *RouterHandle) Start(ctx context.Context, in *api.StartRequest) (*api.Response, error) {
	return startHandler(ctx, in)
}

// Start creates and run entry point Station
func (h *CtlHandle) Start(ctx context.Context, in *api.StartRequest) (*api.Response, error) {
	return startHandler(ctx, in)
}
