package main

import (
	"log"
	"net/http"
	"os"

	"resthopai/app/bundles/kittiesbundle"

	"resthopai/app/core"

	"github.com/joho/godotenv"
	"github.com/urfave/cli"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	name    = "kitties"
	v       = "undefined"
	address string // address of the server
)

func main() {
	app := cli.NewApp()
	app.Name = name
	app.Version = v
	app.Usage = "REST THO PAI web server"

	app.Commands = []cli.Command{
		{
			Name:    "server",
			Aliases: []string{"s"},
			Usage:   "start web server",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "address",
					Usage:       "address to bind",
					Value:       ":8080",
					Destination: &address,
				},
			},
			Action: func(c *cli.Context) error {
				return startServer(address)
			},
		},
		{
			Name:  "db:init",
			Usage: "init the database",
			Action: func(c *cli.Context) error {
				return initDB()
			},
		},
	}

	// Set the default action as server
	if len(os.Args) == 1 {
		os.Args = append(os.Args, "s")
	}

	app.Run(os.Args)

}

func loadConfig() *core.Config {
	err := godotenv.Load()
	if err != nil {
		log.Print(".env file not found - ")
	}

	c := &core.Config{}
	c.Fetch()

	return c
}

func initBundles(db *gorm.DB) []core.Bundle {
	return []core.Bundle{kittiesbundle.NewKittiesBundle(db)}
}

func initDB() error {
	cfg := loadConfig()
	db, err := gorm.Open(cfg.DBType, cfg.DBConnection)
	defer db.Close()

	if err != nil {
		return err
	}

	db.AutoMigrate(&kittiesbundle.Kitty{})

	db.Create(kittiesbundle.NewKitty("Gaspart", "British", "2016-07-05"))
	db.Create(kittiesbundle.NewKitty("Marcel", "European", "2014-05-02"))

	return nil
}

func startServer(addr string) error {
	c := loadConfig()

	db, err := gorm.Open(c.DBType, c.DBConnection)
	defer db.Close()

	if err != nil {
		return err
	}

	r := mux.NewRouter()
	s := r.PathPrefix("/api/v1/").Subrouter()

	for _, b := range initBundles(db) {
		for _, route := range b.GetRoutes() {
			s.HandleFunc(route.Path, route.Handler).Methods(route.Method)
		}
	}

	// Routes handling
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(addr, nil))

	return nil
}
