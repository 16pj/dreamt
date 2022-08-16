package api

import (
	"dreamt/pkg/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (a API) FGetDreams(c *fiber.Ctx) error {
	// get dreams from controller
	dreams, err := a.controller.GetDreams()
	if err != nil {
		return c.Status(http.StatusInternalServerError).Send([]byte(err.Error()))
	}

	// write response
	return c.Status(200).JSON(dreams)
}

func (a API) FGetDream(c *fiber.Ctx) error {
	// get id from url
	id := c.Params("id")
	fmt.Println(id)

	// get dreams from controller
	dream, err := a.controller.GetDream(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).Send([]byte(err.Error()))
	}

	// write response
	return c.Status(200).JSON(dream)
}

func (a API) FGetInterpretation(c *fiber.Ctx) error {
	// get keyword from url
	keyword := c.Params("keyword")

	// get interpretation from controller
	interpret, err := a.controller.GetInterpret(keyword)
	if err != nil {
		return c.Status(http.StatusInternalServerError).Send([]byte(err.Error()))
	}

	// write response
	return c.Status(200).JSON(interpret)
}

func (a API) FGetKeywords(c *fiber.Ctx) error {
	// get query param limit
	limit := c.Query("limit")
	if limit == "" {
		limit = "10"
	}

	top, err := strconv.Atoi(limit)
	if err != nil {
		return c.Status(http.StatusBadRequest).Send([]byte(err.Error()))
	}

	// get keywords from controller
	keywords, err := a.controller.GetKeywords(top)
	if err != nil {
		return c.Status(http.StatusInternalServerError).Send([]byte(err.Error()))
	}

	// write response
	return c.Status(200).JSON(keywords)
}

func (a API) FCreateDream(c *fiber.Ctx) error {
	// get dream from body
	var dream models.Dream
	if err := c.BodyParser(&dream); err != nil {
		return c.Status(http.StatusBadRequest).Send([]byte(err.Error()))
	}

	// create dream in controller
	id, err := a.controller.WriteDreams(dream)
	if err != nil {
		return c.Status(http.StatusInternalServerError).Send([]byte(err.Error()))
	}

	// write response
	return c.Status(200).JSON(id)
}

func (a API) FDeleteDream(c *fiber.Ctx) error {
	// get id from url
	id := c.Params("id")

	// delete dream in controller
	if err := a.controller.DeleteDream(id); err != nil {
		return c.Status(http.StatusInternalServerError).Send([]byte(err.Error()))
	}

	// write response
	return c.Status(200).JSON("Dream deleted")
}
