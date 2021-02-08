package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	os.Args = []string{"multi_values",
		"--stringSclice", "parsed1,parsed2", "--stringSclice", "parsed3,parsed4",
		"--float64Sclice", "13.3,14.4", "--float64Sclice", "15.5,16.6",
		"--int64Sclice", "13,14", "--int64Sclice", "15,16",
		"--intSclice", "13,14", "--intSclice", "15,16",
	}
	os.Setenv("ENV_CLI_FLOAT64_SLICE", "23.3,24.4")
	app := cli.NewApp()
	app.Name = "multi_values"
	app.Flags = []cli.Flag{
		&cli.StringSliceFlag{Name: "stringSclice"},
		&cli.Float64SliceFlag{Name: "float64Sclice", EnvVars: []string{"ENV_CLI_FLOAT64_SLICE"}},
		&cli.Int64SliceFlag{Name: "int64Sclice"},
		&cli.IntSliceFlag{Name: "intSclice"},
	}
	app.Action = func(ctx *cli.Context) error {
		for i, v := range ctx.FlagNames() {
			fmt.Printf("%d-%s %#v\n", i, v, ctx.Value(v))
		}
		return ctx.Err()
	}

	_ = app.Run(os.Args)
}
