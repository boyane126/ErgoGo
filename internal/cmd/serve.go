package cmd

import (
	"ErgoGo/bootstrap"
	"ErgoGo/pkg/config"
	"ErgoGo/pkg/console"
	"ErgoGo/pkg/logger"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var CmdServe = &cobra.Command{
	Use:   "serve",
	Short: "Start web server",
	Run:   runWeb,
	Args:  cobra.NoArgs,
}

func runWeb(cmd *cobra.Command, args []string) {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	bootstrap.SetupRoute(router)

	console.Success(fmt.Sprintf("INFO Server running on [http://127.0.0.1:%s].\n", config.Get("app.port")))
	console.Warning("Press Ctrl+C to stop the server")

	err := router.Run(":" + config.Get("app.port"))
	if err != nil {
		logger.ErrorString("CMD", "serve", err.Error())
		console.Exit("Unable to start server, error:" + err.Error())
	}
}
