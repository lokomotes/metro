package metro

import (
	"context"
	code "net/http"

	api "github.com/lokomotes/metro/api"

	log "github.com/sirupsen/logrus"
)

func (desc *instDesc) transmit(signal api.Signal) error {
	body, ok := desc.getBody()
	if !ok {
		return errNExists
	}

	body.transmit <- signal
	return nil
}

// Transmit deliver messages to other server or Metro through Metro server steram
func (h *RouterHandle) Transmit(ctx context.Context, in *api.TransmitRequest) (*api.Response, error) {
	var (
		res   = &api.Response{Code: code.StatusOK}
		token = (*token)(in.GetToken())
		srcSt = in.GetSrc()
		dstSt = in.GetDst()
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

	logger.Info("Transmit is requested")

	dstDesc.transmit(api.Signal{
		Src:     srcSt,
		Dst:     dstSt,
		Message: in.GetMessage(),
		Control: api.Signal_MESSAGE,
	})

	logger.WithFields(log.Fields{
		"msg": in.GetMessage(),
	}).Info("message is transmitted")

	return res, nil
}
