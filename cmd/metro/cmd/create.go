package cmd

import (
	"errors"
	"unicode"

	. "github.com/lokomotes/metro/cmd/metro/create"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	createOptOutRepo        string
	createOptRuntime        string
	createOptStationVersion string
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create FILE|DIR",
	Short: "Create metro function image",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires at least one arg")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		r, rv := resolveRuntime(createOptRuntime)

		repo, err := Create(&Option{
			OutRepo:        createOptOutRepo,
			Runtime:        r,
			RuntimeVersion: rv,
			SrcPath:        args[0],
			StationVersion: createOptStationVersion,
		})
		if err != nil {
			log.Fatal(err)
		}

		log.WithFields(log.Fields{
			"repo": repo,
		}).Info("New Metro function is created")
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringVarP(&createOptOutRepo, "output", "o", "", "created image repository name (FILE or DIR name by default)")
	createCmd.Flags().StringVarP(&createOptRuntime, "runtime", "r", "", "runtime for create (decided from FILE extension by default)")
	createCmd.Flags().StringVar(&createOptStationVersion, "station-version", "", "version of station (default \"latest\")")
}

func resolveRuntime(runtime string) (Runtime, string) {
	var (
		r string
		v string
	)
	for i, c := range runtime {
		if !unicode.IsDigit(c) {
			continue
		}

		r = runtime[:i]
		v = runtime[i:]
		break
	}

	if len(r) == 0 {
		r = runtime
	}

	switch r {
	default:
		return Auto, ""
	case "node":
		fallthrough
	case "nodejs":
		return NodeJS, v

	case "go":
		fallthrough
	case "golang":
		return Golang, v
	}
}
