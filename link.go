package metro

import (
	"context"
	code "net/http"

	api "github.com/lokomotes/metro/api"

	log "github.com/sirupsen/logrus"
)

// Link connects two stations between caller and requested station
func (h *RouterHandle) Link(ctx context.Context, in *api.LinkRequest) (*api.Response, error) {
	var (
		res   = &api.Response{Code: code.StatusOK}
		token = (*token)(in.GetToken())
		srcSt = in.GetSrc()
		dstSt = in.GetDst()
		msg   = in.GetMessage()
	)

	srcDesc, ok := token.getDesc()
	if !ok {
		log.Warn(errInvTkn)
		res.Code = code.StatusUnauthorized
		return res, nil
	}
	srcSt.Image = srcDesc.image
	dstSt.Id = srcSt.GetId()

	dstDesc := &instDesc{
		userID: srcDesc.userID,
		image:  dstSt.GetImage(),
	}

	logger := log.WithFields(log.Fields{
		"token": ((*api.Token)(token)).ToShort(),
		"flow":  srcSt.ToShort(),
		"src":   srcSt.ToString(),
		"dst":   dstSt.ToString(),
	})

	logger.Info("Link is requested")

	err := newInstance(dstDesc, &api.Signal{
		Src:     srcSt,
		Dst:     dstSt,
		Control: api.Signal_LINKED,
		Message: msg,
	})

	switch err {
	default:
		res.Code = code.StatusInternalServerError
		logger.Warn(err)
		return res, nil
	case errExists:
	case nil:
		logger.Info("new instance is created")
	}

	logger.Info("linked")

	return res, nil
}
