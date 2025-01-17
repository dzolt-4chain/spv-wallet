package destinations

import (
	"github.com/bitcoin-sv/spv-wallet/actions"
	"github.com/bitcoin-sv/spv-wallet/config"
	"github.com/bitcoin-sv/spv-wallet/server/routes"
	"github.com/gin-gonic/gin"
)

// Action is an extension of actions.Action for this package
type Action struct {
	actions.Action
}

// NewHandler creates the specific package routes
func NewHandler(appConfig *config.AppConfig, services *config.AppServices) (routes.OldBasicEndpointsFunc, routes.OldAPIEndpointsFunc) {
	action := &Action{actions.Action{AppConfig: appConfig, Services: services}}

	basicEndpoints := routes.OldBasicEndpointsFunc(func(router *gin.RouterGroup) {
		basicDestinationGroup := router.Group("/destination")
		basicDestinationGroup.GET("", action.get)
		basicDestinationGroup.POST("/count", action.count)
		basicDestinationGroup.GET("/search", action.search)
		basicDestinationGroup.POST("/search", action.search)
	})

	apiEndpoints := routes.OldAPIEndpointsFunc(func(router *gin.RouterGroup) {
		apiDestinationGroup := router.Group("/destination")
		apiDestinationGroup.POST("", action.create)
		apiDestinationGroup.PATCH("", action.update)
	})

	return basicEndpoints, apiEndpoints
}
