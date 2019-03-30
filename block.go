package metro

import (
	"context"
	code "net/http"

	api "github.com/lokomotes/metro/api"

	log "github.com/sirupsen/logrus"
)

// Block prevent transmiting signal to src from dst
func (h *RouterHandle) Block(ctx context.Context, in *api.BlockRequest) (*api.Response, error) {
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

	logger.Info("Block is requested")

	dstDesc.transmit(api.Signal{
		Src:     srcSt,
		Dst:     dstSt,
		Control: api.Signal_BLOCKED,
		Message: msg,
	})

	logger.Info("blocked")

	return res, nil
}
