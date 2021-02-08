package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	os.Args = []string{"multi_values",
		"--stringSclice", "parsed1", "--stringSclice", "parsed3",
		"--float64Sclice", "13.3", "--float64Sclice", "15.5",
		"--int64Sclice", "13", "--int64Sclice", "15",
		"--intSclice", "13", "--intSclice", "15",
	}
	os.Setenv("ENV_CLI_STRING_SLICE", "foo,bar")
	os.Setenv("ENV_CLI_FLOAT64_SLICE", "23.3,24.4")
	os.Setenv("ENV_CLI_INT64_SLICE", "23,24")
	os.Setenv("ENV_CLI_INT_SLICE", "23,24")
	app := cli.NewApp()
	app.Name = "multi_values"
	app.Flags = []cli.Flag{
		&cli.StringSliceFlag{Name: "stringSclice", EnvVars: []string{"ENV_CLI_STRING_SLICE"}},
		&cli.Float64SliceFlag{Name: "float64Sclice", EnvVars: []string{"ENV_CLI_FLOAT64_SLICE"}},
		&cli.Int64SliceFlag{Name: "int64Sclice", EnvVars: []string{"ENV_CLI_INT64_SLICE"}},
		&cli.IntSliceFlag{Name: "intSclice", EnvVars: []string{"ENV_CLI_INT_SLICE"}},
	}
	app.Action = func(ctx *cli.Context) error {
		for i, v := range ctx.FlagNames() {
			fmt.Printf("%d-%s %#v\n", i, v, ctx.Value(v))
		}
		return ctx.Err()
	}

	_ = app.Run(os.Args)

	// bug: slice flag value don't append to default values from ENV or file
	/*
		0-float64Sclice cli.Float64Slice{slice:[]float64{23.3, 24.4, 13.3, 15.5}, hasBeenSet:true}	// bug: expect []float64{13.3, 15.5}
		1-int64Sclice cli.Int64Slice{slice:[]int64{23, 24, 13, 15}, hasBeenSet:true}				// bug: expect []int64{13, 15}
		2-intSclice cli.IntSlice{slice:[]int{23, 24, 13, 15}, hasBeenSet:true}						// bug: expect []int{13, 15}
		3-stringSclice cli.StringSlice{slice:[]string{"parsed1", "parsed3"}, hasBeenSet:true}
	*/
}
