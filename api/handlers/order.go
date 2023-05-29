package handlers

import (
	"practice1/order_user_api_gateway/api/http"
	"practice1/order_user_api_gateway/genproto/order_service"

	"github.com/gin-gonic/gin"
)

// Create user godoc
// @ID create_user
// @Router /user [POST]
// @Summary Create user
// @Description Create user
// @Tags User
// @Accept json
// @Produce json
// @Param agent body corporate_service.CreateAgentRequest true "Request body"
// @Success 200 {object} http.Response{data=corporate_service.CreateAgentResponse} "Response body"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateOrder(c *gin.Context) {
	var order order_service.CreateOrderRequest

	err := c.ShouldBindJSON(&order)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.OrderService().Create(
		c.Request.Context(),
		&order,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// Get order by id
// @ID get_order_by_id
// @Router /order/{id} [GET]
// @Summary Get order by id
// @Description Get order by id
// @Tags Order
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=corporate_service.GetAgentResponse} "Response body"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetOrderById(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.services.OrderService().GetById(
		c.Request.Context(),
		&order_service.Pkey{Id: id},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error)
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// Get order list
// @ID get_order_list
// @Router /order [GET]
// @Summary Get order list
// @Description Get order list
// @Tags Order
// @Accept json
// @Produce json
// @Param page query integer false "page"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Success 200 {object} http.Response{data=corporate_service.GetAllAgentsResponse} "Response body"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetAllOrders(c *gin.Context) {
	resp, err := h.services.OrderService().GetAll(
		c.Request.Context(),
		&order_service.GetAllOrdersRequest{},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// Delete order by id
// @ID delete_order_by_id
// @Router /order/{id} [DELETE]
// @Summary Delete order by id
// @Description Delete order by id
// @Tags Order
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Empty
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteOrder(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.services.OrderService().Delete(
		c.Request.Context(),
		&order_service.Pkey{Id: id},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)

}
