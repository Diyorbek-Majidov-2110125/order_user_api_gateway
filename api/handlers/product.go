package handlers

import (
	"practice1/order_user_api_gateway/api/http"
	"practice1/order_user_api_gateway/genproto/order_service"

	"github.com/gin-gonic/gin"
)


// Create product godoc
// @ID create_product
// @Router /product [POST]
// @Summary Create product
// @Description Create product
// @Tags Product
// @Accept json
// @Produce json
// @Success 200 {object} http.Response{data=corporate_service.CreateAgentResponse} "Response body"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateProduct(c *gin.Context) {
	var product order_service.CreateProductRequest

	err := c.ShouldBindJSON(&product)

	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.OrderService().CreateProduct(
		c.Request.Context(),
		&product,
	)

	h.handleResponse(c, http.OK, resp)
}


// Get product by id
// @ID get_product_by_id
// @Router /product/{id} [GET]
// @Summary Get product by id
// @Description Get product by id
// @Tags Product
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=corporate_service.GetAgentResponse} "Response body"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetProductById(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.services.OrderService().GetProductById(
		c.Request.Context(),
		&order_service.Primarykey{Id: id},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}


// Get product list
// @ID get_product_list
// @Router /product [GET]
// @Summary Get product list
// @Description Get product list
// @Tags Product
// @Accept json
// @Produce json
// @Param page query integer false "page"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Success 200 {object} http.Response{data=corporate_service.GetAllAgentsResponse} "Response body"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetAllProducts(c *gin.Context) {
	resp, err := h.services.OrderService().GetAllProducts(
		c.Request.Context(),
		&order_service.GetAllProductsRequest{},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}


// Delete product by id
// @ID delete_product_by_id
// @Router /product/{id} [DELETE]
// @Summary Delete product by id
// @Description Delete product by id
// @Tags Product
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Empty
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.services.OrderService().DeleteProduct(
		c.Request.Context(),
		&order_service.Primarykey{Id: id},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}