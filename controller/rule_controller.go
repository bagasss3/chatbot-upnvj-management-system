package controller

import (
	"cbupnvj/constant"
	"cbupnvj/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type ruleController struct {
	ruleService model.RuleService
}

func NewRuleController(ruleService model.RuleService) model.RuleController {
	return &ruleController{
		ruleService: ruleService,
	}
}

func (r *ruleController) HandleCreateRule() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := model.CreateUpdateRuleRequest{}
		if err := c.Bind(&req); err != nil {
			log.Error(err)
			return constant.ErrInternal
		}

		create, err := r.ruleService.CreateRule(c.Request().Context(), req)
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, create)
	}
}

func (r *ruleController) HandleFindAllRule() echo.HandlerFunc {
	return func(c echo.Context) error {
		rules, err := r.ruleService.FindAllRule(c.Request().Context())
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, rules)
	}
}

func (r *ruleController) HandleFindRuleByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		idParam := c.Param("id")

		id, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			log.Error(err)
			return err
		}

		rule, err := r.ruleService.FindRuleByID(c.Request().Context(), id)
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, rule)
	}
}

func (r *ruleController) HandleUpdateRule() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := model.CreateUpdateRuleRequest{}
		if err := c.Bind(&req); err != nil {
			log.Error(err)
			return constant.ErrInternal
		}

		idParam := c.Param("id")
		id, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			log.Error(err)
			return err
		}

		update, err := r.ruleService.UpdateRule(c.Request().Context(), id, req)
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, update)
	}
}

func (r *ruleController) HandleDeleteRule() echo.HandlerFunc {
	return func(c echo.Context) error {
		idParam := c.Param("id")

		id, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			log.Error(err)
			return err
		}

		isDeleted, err := r.ruleService.DeleteRule(c.Request().Context(), id)
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, isDeleted)
	}
}
