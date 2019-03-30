package metro

import (
	"net"
	"strconv"

	api "github.com/lokomotes/metro/api"

	"github.com/docker/docker/client"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// DckrCli is docker client
var DckrCli *client.Client
var serveOpts ServeOptions

// ServeOptions holds parameters to serving a Metro server
type ServeOptions struct {
	Host string
	Port uint16
}

// Serve starts Metro server
func Serve(opt *ServeOptions) {
	serveOpts = *opt
	addr := serveOpts.getServerAddress()
	version := "1.39"

	cli, err := client.NewClientWithOpts(client.WithVersion(version))
	if err != nil {
		panic(err)
	}
	DckrCli = cli
	log.WithFields(log.Fields{
		"version": cli.ClientVersion(),
	}).Info("the Docker client created")

	updateInfo()

	log.WithFields(log.Fields{
		"ID":   api.TruncateID(metroContID),
		"Name": metroContName,
	}).Info("the Metro server inspected")

	serveOptFields := log.Fields{
		"Host": serveOpts.Host,
		"Port": serveOpts.Port,
	}

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.WithFields(serveOptFields).Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	api.RegisterRouterServer(s, &RouterHandle{})
	api.RegisterCtlServer(s, &CtlHandle{})

	reflection.Register(s)
	log.WithFields(serveOptFields).Info("starting the Metro server")
	if err := s.Serve(lis); err != nil {
		log.WithFields(serveOptFields).Fatalf("failed to serve: %v", err)
	}
}

func (opts *ServeOptions) getServerAddress() string {
	return opts.Host + ":" + strconv.Itoa(int(opts.Port))
}
