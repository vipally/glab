package main

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "cli_bug",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "bool",
				Value:   true,
				Aliases: []string{"b"},
			},
			&cli.StringFlag{
				Name:    "string",
				Aliases: []string{"s"},
				Value:   "default",
			},
			&cli.IntFlag{
				Name:    "int",
				Aliases: []string{"i"},
				Value:   1,
			},
			&cli.StringSliceFlag{
				Name:    "stringSlice",
				Aliases: []string{"ss"},
				Value:   cli.NewStringSlice("default1", "default2"),
			},
			&cli.Float64SliceFlag{
				Name:    "float64Slice",
				Aliases: []string{"f64s"},
				Value:   cli.NewFloat64Slice(1.1, 2.2),
			},
			&cli.Int64SliceFlag{
				Name:    "int64Slice",
				Aliases: []string{"i64s"},
				Value:   cli.NewInt64Slice(1, 2),
			},
			&cli.IntSliceFlag{
				Name:    "intSlice",
				Aliases: []string{"is"},
				Value:   cli.NewIntSlice(1, 2),
			},
		},
		Action: func(ctx *cli.Context) error {
			fmt.Println("FlagNames", ctx.FlagNames())
			for i, v := range ctx.FlagNames() {
				fmt.Printf("%d %s %+v\n", i, v, ctx.Value(v))
			}
			err := ctx.Err()
			fmt.Println("error", err)
			if err == nil {
				cli.ShowAppHelp(ctx)
			}
			return nil
		},
	}
	args := []string{
		"cli_bug.exe",
		"--string", "setstring1",
		"--bool", "false",
		"-b", "false",
		"--string", "setstring1",
		"-s", "setstring2",
		"--int", "13",
		"-i", "14",
		"--stringSlice", "stringSlice",
		"-ss", "ss",
		"--float64Slice", "13.3",
		"-f64s", "14.4",
		"--int64Slice", "13",
		"-i64s", "14",
		"--intSlice", "13",
		"-is", "14",
	}
	fmt.Printf("args: %+v\n\n", args)
	if err := app.Run(args); err != nil {
		panic(err)
	}
}
