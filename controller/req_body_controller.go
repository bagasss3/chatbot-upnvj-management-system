package controller

import (
	"cbupnvj/constant"
	"cbupnvj/model"
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type reqBodyController struct {
	reqBodyService model.ReqBodyService
}

func NewReqBodyController(reqBodyService model.ReqBodyService) model.ReqBodyController {
	return &reqBodyController{
		reqBodyService: reqBodyService,
	}
}

func (r *reqBodyController) HandleCreateReqBody() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := model.CreateReqBodyRequest{}
		if err := c.Bind(&req); err != nil {
			log.Error(err)
			return constant.ErrInternal
		}

		create, err := r.reqBodyService.CreateReqBody(c.Request().Context(), req)
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, create)
	}
}

func (r *reqBodyController) HandleFindAllReqBodyByActionHttpID() echo.HandlerFunc {
	return func(c echo.Context) error {
		actionHttpId := c.Param("actionHttpId")

		reqBodies, err := r.reqBodyService.FindAllReqBodyByActionHttpID(c.Request().Context(), actionHttpId)
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, reqBodies)
	}
}

func (r *reqBodyController) HandleUpdateReqBody() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := model.UpdateReqBodyRequest{}
		if err := c.Bind(&req); err != nil {
			log.Error(err)
			return constant.ErrInternal
		}

		actionHttpId := c.Param("actionHttpId")
		reqBodyId := c.Param("reqBodyId")

		update, err := r.reqBodyService.UpdateReqBody(c.Request().Context(), reqBodyId, actionHttpId, req)
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, update)
	}
}

func (r *reqBodyController) HandleDeleteReqBody() echo.HandlerFunc {
	return func(c echo.Context) error {
		actionHttpId := c.Param("actionHttpId")
		reqBodyId := c.Param("reqBodyId")

		isDeleted, err := r.reqBodyService.DeleteReqBody(c.Request().Context(), reqBodyId, actionHttpId)
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, isDeleted)
	}
}
