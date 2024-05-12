package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nahdukesaba/be-assignment/helpers"
	"github.com/nahdukesaba/be-assignment/repo"
	"github.com/nahdukesaba/be-assignment/services"
	"github.com/nahdukesaba/be-assignment/services/request"
)

type PaymentHandler struct {
	paymentService services.Payments
}

func NewPaymentHandler(db *repo.DB) *PaymentHandler {
	return &PaymentHandler{
		paymentService: services.NewPaymentService(db),
	}
}

// Send Send balance to other account
//
//	@Summary		Send
//	@Description	Send balance to other account
//	@Tags			payment
//	@Accept			json
//	@Produce		json
//	@Param			payment	body		request.TransactionRequest	true	"Send balance"
//	@Success		200		{object}	helpers.Message
//	@Failure		400		{object}	helpers.Message
//	@Failure		404		{object}	helpers.Message
//	@Failure		500		{object}	helpers.Message
//	@Router			/api/payment/send [post]
func (ph *PaymentHandler) Send(c *gin.Context) {
	if !helpers.ProtectedHandler(c) {
		log.Println("Send Unauthorized!")
		c.JSON(http.StatusUnauthorized, helpers.UnauthorizedMessage)
	}

	form := new(request.TransactionRequest)
	if err := json.NewDecoder(c.Request.Body).Decode(&form); err != nil {
		log.Printf("Error reading body request: %v\n", err)
		c.JSON(http.StatusBadRequest, helpers.BadRequesetMessage)
		return
	}

	form.Username = helpers.GetusernameFromToken(c)
	if err := ph.paymentService.SendBalance(c, form); err != nil {
		log.Printf("SendBalance error : %v\n", err)
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, helpers.SuccessMessage)
}

// Withdraw Withdraw account balance
//
//	@Summary		Withdraw
//	@Description	withdraw account balance
//	@Tags			payment
//	@Accept			json
//	@Produce		json
//	@Param			payment	body		request.TransactionRequest	true	"withdraw account balance"
//	@Success		200		{object}	helpers.Message
//	@Failure		400		{object}	helpers.Message
//	@Failure		404		{object}	helpers.Message
//	@Failure		500		{object}	helpers.Message
//	@Router			/api/payment/withdraw [post]
func (ph *PaymentHandler) Withdraw(c *gin.Context) {
	if !helpers.ProtectedHandler(c) {
		log.Println("Withdraw Unauthorized!")
		c.JSON(http.StatusUnauthorized, helpers.UnauthorizedMessage)
	}

	form := new(request.TransactionRequest)
	if err := json.NewDecoder(c.Request.Body).Decode(&form); err != nil {
		log.Printf("Error reading body request: %v\n", err)
		c.JSON(http.StatusBadRequest, helpers.BadRequesetMessage)
		return
	}

	form.Username = helpers.GetusernameFromToken(c)
	if err := ph.paymentService.WithdrawBalance(c, form); err != nil {
		log.Printf("WithdrawBalance error : %v\n", err)
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, helpers.SuccessMessage)
}

// GetAccountsUser Get all account from specified user
//
//	@Summary		GetAccountsUser
//	@Description	get all account from specified user
//	@Tags			account
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	helpers.Message
//	@Failure		400	{object}	helpers.Message
//	@Failure		404	{object}	helpers.Message
//	@Failure		500	{object}	helpers.Message
//	@Router			/api/account [get]
func (ph *PaymentHandler) GetAccountsUser(c *gin.Context) {
	if !helpers.ProtectedHandler(c) {
		log.Println("Unauthorized!")
		c.JSON(http.StatusUnauthorized, helpers.UnauthorizedMessage)
	}

	username := helpers.GetusernameFromToken(c)
	result, err := ph.paymentService.GetAllAccountsUser(c, username)
	if err != nil {
		log.Printf("GetAllAccounts error : %v\n", err)
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"data": result})
}

// GetTransactionsAccount Get all transactions from specified account
//
//	@Summary		GetTransactionsAccount
//	@Description	get all transactions from specified account
//	@Tags			account
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Account ID"
//	@Success		200	{object}	helpers.Message
//	@Failure		400	{object}	helpers.Message
//	@Failure		404	{object}	helpers.Message
//	@Failure		500	{object}	helpers.Message
//	@Router			/api/account/:account_id [get]
func (ph *PaymentHandler) GetTransactionsAccount(c *gin.Context) {
	if !helpers.ProtectedHandler(c) {
		log.Println("Unauthorized!")
		c.JSON(http.StatusUnauthorized, helpers.UnauthorizedMessage)
	}

	form := &request.TransactionsAccountRequest{
		Username:  helpers.GetusernameFromToken(c),
		AccountID: c.Param("account_id"),
	}
	result, err := ph.paymentService.GetAllTransactionsAccount(c, form)
	if err != nil {
		log.Printf("GetAllAccounts error : %v\n", err)
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"data": result})
}
