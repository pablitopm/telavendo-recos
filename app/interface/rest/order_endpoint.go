package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"telavendo.com.ar/telavendo-recos/app/domain/model"
	"telavendo.com.ar/telavendo-recos/app/registry"
	"telavendo.com.ar/telavendo-recos/app/usecase"
)

func SaveOrderProducts(c *gin.Context) {
	req := struct {
		StoreID int    `json:"store_id"`
		Event   string `json:"event"`
		ID      int    `json:"id"`
	}{}
	err := c.BindJSON(&req)
	if err != nil {
		log.Error("error", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	ctn := c.MustGet("ctn").(*registry.Container)
	orderUseCase := ctn.Resolve("order-usecase").(usecase.OrderUsecase)
	userPurchaseUseCase := ctn.Resolve("user-purchase-usecase").(usecase.UserPurchaseUsecase)

	order, err := orderUseCase.GetOrder(req.ID)
	for _, product := range order.Products {

		userPurchase := &model.UserPurchase{}
		userPurchase.UserID = order.Customer.ID
		userPurchase.ProductID = product.ProductID
		userPurchase.PurchasedTimes++

		userPurchaseUseCase.RegisterUserPurchase(*userPurchase)
	}
	/*err = order.Validate()
	if err != nil {
		log.Error("Did not pass validations", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}*/

	//ctn := c.MustGet("ctn").(*registry.Container)
	//useCase := ctn.Resolve("user-purchase-usecase").(usecase.UserPurchaseUsecase)
	//userPurchase, _ := useCase.StartGame(game)

	c.JSON(http.StatusCreated, nil)
}
