package cmd

import (
	"fmt"
	"log"
	"reflect"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var RootCmd = &cobra.Command{
	Use:   "gmusic",
	Short: "gmusic",
	Run: func(_ *cobra.Command, _ []string) {

	},
}

func initConfig() {

}

func init() {
	cobra.OnInitialize(initConfig)

	flags := []Flag{
		{
			Name:  "email",
			Desc:  "email",
			Value: "",
			Short: "e",
		},
		{

			Name:  "password",
			Desc:  "password",
			Value: "",
			Short: "p",
		},
	}

	BindFlags(RootCmd, flags, true)
}

func BindFlags(cmd *cobra.Command, flags []Flag, persistent bool) {
	for _, f := range flags {

		fs := cmd.Flags()

		if persistent {
			fs = cmd.PersistentFlags()
		}

		typeOf := reflect.TypeOf(f.Value)
		typeName := fmt.Sprintf("%s", typeOf)

		switch typeName {
		case "string":
			val := f.Value.(string)
			fs.StringP(f.Name, f.Short, val, f.Desc)
		case "uint":
			val := f.Value.(uint)
			fs.UintP(f.Name, f.Short, val, f.Desc)
		case "int":
			val := f.Value.(int)
			fs.IntP(f.Name, f.Short, val, f.Desc)
		case "bool":
			val := f.Value.(bool)
			fs.BoolP(f.Name, f.Short, val, f.Desc)
		case "[]string":
			val := f.Value.([]string)
			fs.StringSliceP(f.Name, f.Short, val, f.Desc)
		case "[]int":
			val := f.Value.([]int)
			fs.IntSliceP(f.Name, f.Short, val, f.Desc)
		default:
			log.Fatalf("Flag type %s is invalid", typeName)
		}

		viper.BindPFlag(f.Name, fs.Lookup(f.Name))
	}
}
