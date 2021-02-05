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
				Value:   false,
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
		"lab23",
		"-b", // "false",
		"--string", "setstring1",
		"--int", "13",
		"--stringSlice", "stringSlice",
		"--float64Slice", "13.3",
		"--int64Slice", "13",
		"--intSlice", "13",
	}
	fmt.Printf("args: %+v\n\n", args)
	if err := app.Run(args); err != nil {
		panic(err)
	}

	/*
		GLOBAL OPTIONS:
		   --bool, -b                          (default: false)
		   --string value, -s value            (default: "default")
		   --int value, -i value               (default: 1)
		   --stringSlice value, --ss value     (default: "stringSlice") // bug expect: (default: "default1", "default2")
		   --float64Slice value, --f64s value  (default: 13.3)			// bug expect: (default: 1.1, 2.2)
		   --int64Slice value, --i64s value    (default: 13)			// bug expect: (default: 1, 2)
		   --intSlice value, --is value        (default: 13)			// bug expect: (default: 1, 2)
		   --help, -h                          show help (default: false)
	*/
}
