package main

import (
	"fmt"
	"os"

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
				EnvVars: []string{"ENV_BOOL1", "ENV_BOOL2"},
			},
			&cli.StringFlag{
				Name:    "string",
				Aliases: []string{"s"},
				Value:   "default",
				EnvVars: []string{"ENV_STRING1", "ENV_STRING2"},
			},
			&cli.IntFlag{
				Name:    "int",
				Aliases: []string{"i"},
				Value:   1,
				EnvVars: []string{"ENV_INT1", "ENV_INT2"},
			},
			&cli.StringSliceFlag{
				Name:    "stringSlice",
				Aliases: []string{"ss"},
				Value:   cli.NewStringSlice("default1", "default2"),
				EnvVars: []string{"ENV_STRINGSLICE1", "ENV_STRINGSLICE2"},
			},
			&cli.Float64SliceFlag{
				Name:    "float64Slice",
				Aliases: []string{"f64s"},
				Value:   cli.NewFloat64Slice(1.1, 2.2),
				EnvVars: []string{"ENV_FLOAT64SLICE1", "ENV_FLOAT64SLICE2"},
			},
			&cli.Int64SliceFlag{
				Name:    "int64Slice",
				Aliases: []string{"i64s"},
				Value:   cli.NewInt64Slice(1, 2),
				EnvVars: []string{"ENV_INT64SLICE1", "ENV_INT64SLICE2"},
			},
			&cli.IntSliceFlag{
				Name:    "intSlice",
				Aliases: []string{"is"},
				Value:   cli.NewIntSlice(1, 2),
				EnvVars: []string{"ENV_INTSLICE1", "ENV_INTSLICE2"},
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
	envs := []string{
		"ENV_BOOL1", "ENV_BOOL2",
		"ENV_STRING1", "ENV_STRING2",
		"ENV_INT1", "ENV_INT2",
		"ENV_STRINGSLICE1", "ENV_STRINGSLICE2",
		"ENV_FLOAT64SLICE1", "ENV_FLOAT64SLICE2",
		"ENV_INT64SLICE1", "ENV_INT64SLICE2",
		"ENV_INTSLICE1", "ENV_INTSLICE2",
	}
	for i, v := range envs {
		os.Setenv(v, fmt.Sprintf("%d", i%2))
	}
	fmt.Printf("args: %+v\n\n", args)
	if err := app.Run(args); err != nil {
		panic(err)
	}

	// Flags defalt value change with the env values
	/*
		FlagNames [b bool f64s float64Slice i i64s int int64Slice intSlice is s ss string stringSlice]
		0 b true
		1 bool true
		2 f64s {slice:[0 13.3] hasBeenSet:true}			// bug: Evn value and command-line never append and never allow both
		3 float64Slice {slice:[0 13.3] hasBeenSet:true}
		4 i 13
		5 i64s {slice:[0 13] hasBeenSet:true}
		6 int 13
		7 int64Slice {slice:[0 13] hasBeenSet:true}
		8 intSlice {slice:[0 13] hasBeenSet:true}
		9 is {slice:[0 13] hasBeenSet:true}
		10 s setstring1
		11 ss {slice:[stringSlice] hasBeenSet:true}
		12 string setstring1
		13 stringSlice {slice:[stringSlice] hasBeenSet:true}
		error <nil>
		NAME:
		   cli_bug - A new cli application

		USAGE:
		   lab24.exe [global options] command [command options] [arguments...]

		COMMANDS:
		   help, h  Shows a list of commands or help for one command

		GLOBAL OPTIONS:
		   --bool, -b                          (default: false) [%ENV_BOOL1%, %ENV_BOOL2%]
		   --string value, -s value            (default: "0") [%ENV_STRING1%, %ENV_STRING2%] // bug: default value never change with env values
		   --int value, -i value               (default: 0) [%ENV_INT1%, %ENV_INT2%]
		   --stringSlice value, --ss value     (default: "stringSlice") [%ENV_STRINGSLICE1%, %ENV_STRINGSLICE2%] // bug: default value never change with env values and command-line
		   --float64Slice value, --f64s value  (default: 0, 13.3) [%ENV_FLOAT64SLICE1%, %ENV_FLOAT64SLICE2%]
		   --int64Slice value, --i64s value    (default: 0, 13) [%ENV_INT64SLICE1%, %ENV_INT64SLICE2%]
		   --intSlice value, --is value        (default: 0, 13) [%ENV_INTSLICE1%, %ENV_INTSLICE2%]
		   --help, -h                          show help (default: false)
	*/
}
