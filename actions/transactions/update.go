package transactions

import (
	"net/http"

	"github.com/bitcoin-sv/spv-wallet/engine/spverrors"
	"github.com/bitcoin-sv/spv-wallet/mappings"
	"github.com/bitcoin-sv/spv-wallet/server/auth"
	"github.com/gin-gonic/gin"
)

// update will update a transaction
// Update transaction godoc
// @Summary		Update transaction - Use (PATCH) /api/v1/transactions/{id} instead.
// @Description	This endpoint has been deprecated. Use (PATCH) /api/v1/transactions/{id} instead.
// @Tags		Transactions
// @Produce		json
// @Param		UpdateTransaction body UpdateTransaction true "Pass update transaction request model in the body"
// @Success		200 {object} models.Transaction "Updated transaction"
// @Failure		400	"Bad request - Error while parsing UpdateTransaction from request body, tx not found or tx is not associated with the xpub"
// @Failure 	500	"Internal Server Error - Error while updating transaction"
// @DeprecatedRouter	/v1/transaction [patch]
// @Security	x-auth-xpub
func (a *Action) update(c *gin.Context) {

	var requestBody OldUpdateTransaction
	if err := c.Bind(&requestBody); err != nil {
		spverrors.ErrorResponse(c, spverrors.ErrCannotBindRequest, a.Services.Logger)
		return
	}
	id := requestBody.ID

	reqXPubID := c.GetString(auth.ParamXPubHashKey)

	// Get a transaction by ID
	transaction, err := a.Services.SpvWalletEngine.UpdateTransactionMetadata(
		c.Request.Context(),
		reqXPubID,
		id,
		requestBody.Metadata,
	)
	if err != nil {
		spverrors.ErrorResponse(c, err, a.Services.Logger)
		return
	} else if transaction == nil {
		spverrors.ErrorResponse(c, spverrors.ErrCouldNotFindTransaction, a.Services.Logger)
	} else if !transaction.IsXpubIDAssociated(reqXPubID) {
		spverrors.ErrorResponse(c, spverrors.ErrAuthorization, a.Services.Logger)
		return
	}

	contract := mappings.MapToOldTransactionContract(transaction)
	c.JSON(http.StatusOK, contract)
}

// update will update a transaction metadata
// Update transaction godoc
// @Summary		Update transaction metadata
// @Description	Update transaction metadata
// @Tags		Transactions
// @Produce		json
// @Param		UpdateTransactionRequest body UpdateTransactionRequest true "Pass update transaction request model in the body with updated metadata"
// @Success		200 {object} response.Transaction "Updated transaction metadata"
// @Failure		400	"Bad request - Error while parsing UpdateTransaction from request body, tx not found or tx is not associated with the xpub"
// @Failure 	500	"Internal Server Error - Error while updating transaction metadata"
// @Router		/api/v1/transactions/{id} [patch]
// @Security	x-auth-xpub
func (a *Action) updateTransactionMetadata(c *gin.Context) {
	var requestBody UpdateTransactionRequest
	if err := c.Bind(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	id := c.Param("id")

	reqXPubID := c.GetString(auth.ParamXPubHashKey)

	// Get a transaction by ID
	transaction, err := a.Services.SpvWalletEngine.UpdateTransactionMetadata(
		c.Request.Context(),
		reqXPubID,
		id,
		requestBody.Metadata,
	)
	if err != nil {
		spverrors.ErrorResponse(c, err, a.Services.Logger)
		return
	} else if transaction == nil {
		spverrors.ErrorResponse(c, spverrors.ErrCouldNotFindTransaction, a.Services.Logger)
	} else if !transaction.IsXpubIDAssociated(reqXPubID) {
		spverrors.ErrorResponse(c, spverrors.ErrAuthorization, a.Services.Logger)
		return
	}

	contract := mappings.MapToTransactionContract(transaction)
	c.JSON(http.StatusOK, contract)
}
