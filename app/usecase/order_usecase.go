package usecase

import (
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
	"telavendo.com.ar/telavendo-recos/app/config"
	"telavendo.com.ar/telavendo-recos/app/domain/model"
)

type OrderUsecase interface {
	GetOrder(orderID int) (*model.Order, error)
}

type orderUsecase struct {
	config *config.Config
	//repo    repository.OrderRepository
	//service *service.OrderService
}

func NewOrderUsecase() *orderUsecase {
	return &orderUsecase{
		config: config.CreateConfig(),
		//repo:    repo,
		//service: service,
	}
}

func (o *orderUsecase) GetOrder(orderID int) (*model.Order, error) {

	client := http.Client{}
	url := fmt.Sprintf("%sorders/%d", o.config.BaseUrl, orderID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Error(err.Error)
		return nil, err
	}
	req.Header.Set("Authentication", o.config.BearerToken)
	req.Header.Set("User-Agent", o.config.UserAgent)

	resp, err := client.Do(req)
	defer resp.Body.Close()

	if err != nil {
		log.Error(err.Error)
		return nil, err
	}

	order := &model.Order{}
	err = json.NewDecoder(resp.Body).Decode(order)

	return order, err
}
