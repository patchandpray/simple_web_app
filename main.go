// from https://grisha.org/blog/2017/04/27/simplistic-go-web-app-part-2/
package main

// import daemon package
import (
    "github.com/patchandpray/simple_web_app/daemon"
)

// declare var assetsPath; type is string
var assetsPath string

// function to process flags, takes no input, returns daemon.Config
// package flag implements command-line flag parsing

func processFlags() *daemon.Config {
    // cfg is pointer to daemon.Config
    cfg := &daemon.Config{}

    // declaration of program input flags
    flag.StringVar(&cfg.ListenSpec,
                   "listen",
                   "localhost:3000",
                   "HTTP listen spec")
    flag.StringVar(&cfg.Db.ConnectString,
                   "db-connect",
                   "host=/var/run/postgresql dbname=gowebapp sslmode=disable",
                   "DB Connect String")
    flag.StringVar(&assetsPath,
                   "assets-path",
                   "assets",
                   "Path to assets dir")

    // parse the flags
    flag.Parse()
    return cfg

// function to setup http assets, takes cfg as input
func setupHttpAssets(cfg *daemon.Config) {
    log.Printf("Assets served from %q.", assetsPath)
    cfg.UI.Assets = http.Dir(assetsPath)
}

func main() {
    cfg := processFlags()

    setupHttpAssets(cfg)

    if err := daemon.Run(cfg); err !=nil {
        log.Printf("error in main(): %v", err)
    }
}
