package metro

import (
	log "github.com/sirupsen/logrus"

	api "github.com/lokomotes/metro/api"
)

// Listen messages come from other stations or Metro through Metro server stream
func (h *RouterHandle) Listen(in *api.ListenRequest, stream api.Router_ListenServer) error {
	var (
		token = (*token)(in.GetToken())
	)

	desc, ok := token.getDesc()
	if !ok {
		log.Warn(errInvTkn)
		return nil
	}

	logger := log.WithFields(log.Fields{
		"token": ((*api.Token)(token)).ToShort(),
		"user":  desc.userID,
		"image": desc.image,
	})

	logger.Info("Listen is requested")

	body, ok := desc.getBody()
	if !ok {
		logger.Fatal(errNExists)
	}

	go func() {
		for {
			select {
			case <-stream.Context().Done():
				return
			case sig := <-body.transmit:
				stream.Send(&sig)
			}
		}
	}()

	<-stream.Context().Done()
	logger.Info("stops listening")
	return nil
}
