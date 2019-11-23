package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type application struct {
	infoLog *log.Logger
	secret  *string
}

func (app *application) greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Secret is: %s", *app.secret)
	app.infoLog.Printf("secret is => %s", *app.secret)
}

func main() {

	// Load config
	viper.SetDefault("secret", "nosecret")
	viper.SetDefault("logLevel", "NOINFO")

	viper.SetConfigName("config.yaml")
	viper.AddConfigPath("./")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	viper.WatchConfig()

	// Set secret
	secret := viper.GetString("secret")

	// Set logging parameters
	logger := log.New(os.Stdout, viper.GetString("logLevel")+"\t", log.Ldate|log.Ltime)

	app := &application{
		infoLog: logger,
		secret:  &secret,
	}

	// Reload config
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		secret = viper.GetString("secret")
		app.infoLog = log.New(os.Stdout, viper.GetString("logLevel")+"\t", log.Ldate|log.Ltime)
	})

	// Run server
	http.HandleFunc("/", app.greet)
	http.ListenAndServe(":8080", nil)
}
