package rest

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
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
	//first we save each product for user
	for _, product := range order.Products {
		userPurchaseUseCase.RegisterUserPurchase(order.Customer.ID, product.ProductID)
	}
	//then we iterate the combination to add weight to combination
	c.JSON(http.StatusOK, nil)
}

func GetProductCombination(c *gin.Context) {
	productID, _ := strconv.Atoi(c.Param("productID"))
	quantity, _ := strconv.Atoi(c.Param("quantity"))
	if quantity == 0 {
		quantity = 5 //basic number
	}
	fmt.Println(productID)
	//ctn := c.MustGet("ctn").(*registry.Container)
	//productCombinationUseCase := ctn.Resolve("product-combination-usecase").(usecase.UserPurchaseUsecase)

	//productRanking, err := productCombinationUseCase.GetProductCombinationTop(quantity)

	c.JSON(http.StatusOK, nil)
}
