package config

import (
	"strings"

	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

type Config struct {
	Server struct {
		ServiceApi string
		AdminApi   string
	}
	Mongodb struct {
		Url string
	}
}

var C Config

func MustRead(c *cli.Context) error {

	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")     // adding home directory as first search path
	viper.SetEnvPrefix("gogame") // so viper.AutomaticEnv will get matching envvars starting with O2M_
	viper.AutomaticEnv()         // read in environment variables that match
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if c.GlobalString("config") != "" {
		viper.SetConfigFile(c.GlobalString("config"))
	}

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	if err := viper.Unmarshal(&C); err != nil {
		return err
	}
	var err error
	if C.Id, err = core.IDFromString(C.IdRaw); err != nil {
		return err
	}
	if err := C.KeyStoreBaby.PubKey.UnmarshalText([]byte(C.KeyStoreBaby.PubKeyRaw)); err != nil {
		return err
	}
	C.KeyStoreBaby.PubKeyComp = C.KeyStoreBaby.PubKey.Compress()
	return nil
}
