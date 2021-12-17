package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

const (
	configAddress   = "app.address"
	configPort      = "app.port"
	configTLSEnable = "app.tls.enable"
	configTLSCert   = "app.tls.cert"
	configTLSKey    = "app.tls.key"
)

var config *viper.Viper

func init() {
	config = viper.New()
	config.AddConfigPath("configs")
	config.SetConfigName("config")
	config.SetConfigType("yaml")
	config.SetDefault(configAddress, "localhost")
	config.SetDefault(configPort, 8080)
	config.SetDefault(configTLSEnable, false)
}

func main() {
	g := gin.Default()
	g.GET("/ping", ping)
	listenAddress := fmt.Sprintf("%s:%d", config.GetString(configAddress), config.GetInt(configPort))
	if config.GetBool(configTLSEnable) {
		g.RunTLS(listenAddress, config.GetString(configTLSCert), config.GetString(configTLSKey))
		return
	}
	g.Run(listenAddress)
}

func ping(ctx *gin.Context) {
	ctx.Writer.WriteString("pong")
}
