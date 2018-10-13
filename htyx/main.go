package main

import (
	//"encoding/json"
	"errors"
	//"fmt"
	"net/http"
	//_ "net/http/pprof"
	//"os"
	"time"

	"htyx/config"
	"htyx/model"

	"htyx/router"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	cfg = pflag.StringP("config", "c", "", "apiserver config file path.")
)

func main() {
	pflag.Parse()
	// init config
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}
	// init db
	model.Init()
	// test db
	go func() {
		if err := model.DB.Self.Ping(); err != nil {
			log.Fatal("The DB no ping", err)
		}
		log.Info("The DB successfully.")
	}()
	defer model.Close()
	// Set gin mode.
	gin.SetMode(viper.GetString("runmode"))
	// Create the Gin engine.
	g := gin.New()
	// Routes
	g = router.Load(g, gin.Recovery()) //还可以加入全局通用中间件

	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Info("The router has been deployed successfully.")
	}()
	// Start to listening the incoming requests.
	/*
		cert := viper.GetString("tls.cert")
		key := viper.GetString("tls.key")
		if cert != "" && key != "" {
			go func() {
				log.Infof("Start to listening the incoming requests on https address: %s", viper.GetString("tls.addr"))
				log.Info(http.ListenAndServeTLS(viper.GetString("tls.addr"), cert, key, g).Error())
			}()
		}
	*/
	log.Infof("Start to listening the incoming requests on http address: %s", viper.GetString("addr"))
	log.Info(http.ListenAndServe(viper.GetString("addr"), g).Error())
}
func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		// Ping the server by sending a GET request to `/health`.
		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep for a second to continue the next ping.
		log.Info("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}
