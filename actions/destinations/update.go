package destinations

import (
	"net/http"

	"github.com/bitcoin-sv/spv-wallet/engine"
	"github.com/bitcoin-sv/spv-wallet/engine/spverrors"
	"github.com/bitcoin-sv/spv-wallet/mappings"
	"github.com/bitcoin-sv/spv-wallet/server/auth"
	"github.com/gin-gonic/gin"
)

// update will update an existing model
// Update Destination godoc
// @Summary		Update destination
// @Description	Update destination
// @Tags		Destinations
// @Produce		json
// @Param		UpdateDestination body UpdateDestination false " "
// @Success		200 {object} models.Destination "Updated Destination"
// @Failure		400	"Bad request - Error while parsing UpdateDestination from request body"
// @Failure 	500	"Internal Server Error - Error while updating destination"
// @Router		/v1/destination [patch]
// @Security	x-auth-xpub
func (a *Action) update(c *gin.Context) {
	reqXPubID := c.GetString(auth.ParamXPubHashKey)

	var requestBody UpdateDestination
	if err := c.Bind(&requestBody); err != nil {
		spverrors.ErrorResponse(c, spverrors.ErrCannotBindRequest, a.Services.Logger)
		return
	}
	if requestBody.ID == "" && requestBody.Address == "" && requestBody.LockingScript == "" {
		spverrors.ErrorResponse(c, spverrors.ErrOneOfTheFieldsIsRequired, a.Services.Logger)
		return
	}

	// Get the destination
	var destination *engine.Destination
	var err error
	if requestBody.ID != "" {
		destination, err = a.Services.SpvWalletEngine.UpdateDestinationMetadataByID(
			c.Request.Context(), reqXPubID, requestBody.ID, requestBody.Metadata,
		)
	} else if requestBody.Address != "" {
		destination, err = a.Services.SpvWalletEngine.UpdateDestinationMetadataByAddress(
			c.Request.Context(), reqXPubID, requestBody.Address, requestBody.Metadata,
		)
	} else {
		destination, err = a.Services.SpvWalletEngine.UpdateDestinationMetadataByLockingScript(
			c.Request.Context(), reqXPubID, requestBody.LockingScript, requestBody.Metadata,
		)
	}
	if err != nil {
		spverrors.ErrorResponse(c, err, a.Services.Logger)
		return
	}

	contract := mappings.MapOldToDestinationContract(destination)
	c.JSON(http.StatusOK, contract)
}
