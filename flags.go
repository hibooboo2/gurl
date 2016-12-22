package main

import (
	"flag"
	"os"
	"strings"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigType("json")
	viper.SetConfigName(".gurl")
	viper.AddConfigPath("./")
	viper.AddConfigPath(os.Getenv("HOME"))
	viper.SetEnvPrefix("GURL_CLI")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	authCheck := true
	flag.BoolVar(&authCheck, "verifyauth", false, "Pass to verify that auth is correct at base level")
	flag.BoolVar(&client.Verbose, "v", false, "Verbose logging")
	flag.StringVar(&method, "m", `get`, "method for type for endpoint (GET/PUT/POST/DELETE)")
	flag.StringVar(&endpoint, "e", "", "Endpoint to use")
	flag.StringVar(&payloadString, "d", "{}", "Json payload for the request to be used")
	flag.StringVar(&client.Host, "h", "", "Host overide flag: ex: example.com")
	flag.Parse()

	if client.Host == "" {
		client.Host = viper.GetString("host")
	}
	client.AuthToken = viper.GetString("token")
}
