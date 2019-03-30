package cmd

import (
	"context"
	"errors"
	"time"

	metro "github.com/lokomotes/metro/api"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// loadCmd represents the load command
var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "Load creates Station",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires at least one arg")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		image := args[0]

		conn, err := grpc.Dial(getServerAddress(), grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Failed to connect Metro server: %v", err)
		}
		defer conn.Close()
		cli := metro.NewCtlClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		res, err := cli.Load(ctx, &metro.LoadRequest{
			Station: &metro.Station{
				Image: image,
			},
			UserID: username,
		})
		if err != nil {
			log.Fatalf("Failed to load: %v", err)
		}

		switch code := res.GetCode(); code {
		case 200:
			log.WithFields(log.Fields{
				"image": image,
			}).Info("Loaded new Station")
		case 404:
			log.WithField("image", image).Warn("No such image")
		default:
			log.WithField("code", code).Warn("Responded unknown error")
		}
	},
}

func init() {
	rootCmd.AddCommand(loadCmd)

}
