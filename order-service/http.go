package order_service

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/greenbahar/manage-order-process/order"
	"io/ioutil"
	"net/http"
)

type Handler struct {
	Service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) SaveOrder(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "cannot read the body of the request")
		return
	}
	ord := &order.Order{}
	if err = json.Unmarshal(body, ord); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "bad request")
		return
	}
	if err = h.Service.SaveOrder(ord); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "cannot read the body of the request")
		return
	}
	c.JSON(http.StatusOK, ord)
}
