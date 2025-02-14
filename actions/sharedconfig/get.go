package sharedconfig

import (
	"net/http"
	"sync"

	"github.com/bitcoin-sv/spv-wallet/models"
	"github.com/bitcoin-sv/spv-wallet/models/response"
	"github.com/gin-gonic/gin"
)

// sharedConfig will return the shared configuration
// Get shared config godoc
// @Summary		Get shared config - Use (GET) /api/v1/configs/shared instead.
// @Description	This endpoint has been deprecated. Use (GET) /api/v1/configs/shared instead.
// @Tags		Configurations
// @Produce		json
// @Success		200 {object} models.SharedConfig "Shared configuration"
// @DeprecatedRouter  /v1/shared-config [get]
// @Security	x-auth-xpub
func (a *Action) oldGet(c *gin.Context) {
	makeConfig := sync.OnceValue(func() models.SharedConfig {
		return models.SharedConfig{
			PaymailDomains: a.AppConfig.Paymail.Domains,
			ExperimentalFeatures: map[string]bool{
				"pike_contacts_enabled": a.AppConfig.ExperimentalFeatures.PikeContactsEnabled,
				"pike_payment_enabled":  a.AppConfig.ExperimentalFeatures.PikePaymentEnabled,
			},
		}
	})

	c.JSON(http.StatusOK, makeConfig())
}

// sharedConfig will return the shared configuration
// Get shared config godoc
// @Summary		Get shared config
// @Description	Get shared config
// @Tags		Configurations
// @Produce		json
// @Success		200 {object} response.SharedConfig "Shared configuration"
// @Router		/api/v1/configs/shared [get]
// @Security	x-auth-xpub
func (a *Action) get(c *gin.Context) {
	makeConfig := sync.OnceValue(func() response.SharedConfig {
		return response.SharedConfig{
			PaymailDomains: a.AppConfig.Paymail.Domains,
			ExperimentalFeatures: map[string]bool{
				"pikeContactsEnabled": a.AppConfig.ExperimentalFeatures.PikeContactsEnabled,
				"pikePaymentEnabled":  a.AppConfig.ExperimentalFeatures.PikePaymentEnabled,
			},
		}
	})

	c.JSON(http.StatusOK, makeConfig())
}
