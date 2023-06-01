package server

import (
	"digibala/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func voucherRoutes(e *echo.Echo) {
	e.GET("/api/v1/voucher", listVoucherHandler)
	e.POST("/api/v1/voucher", createVoucherHandler)
	e.GET("/api/v1/voucher/:id", retreveVoucherHandler)
	e.DELETE("/api/v1/voucher/:id", deleteVoucherHandler)
	e.PUT("/api/v1/voucher", updateVoucherHandler)
}

func createVoucherHandler(c echo.Context) error {
	voucher := &models.Voucher{}
	if err := c.Bind(voucher); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, voucher)
}

func updateVoucherHandler(c echo.Context) error {
	voucher := &models.Voucher{}
	if err := c.Bind(voucher); err != nil {
		return err
	}
	fmt.Println("Updating voucher id:", voucher.ID)
	return c.JSON(http.StatusCreated, voucher)
}

func deleteVoucherHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println("Deleting voucher id:", id)
	return c.JSON(http.StatusOK, models.StatusOK{OK: "OK"})
}

func listVoucherHandler(c echo.Context) error {
	vouchers := []*models.Voucher{}
	return c.JSON(http.StatusOK, vouchers)
}

func retreveVoucherHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	voucher := &models.Voucher{ID: id}
	return c.JSON(http.StatusOK, voucher)
}
