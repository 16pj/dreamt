package fiber

import (
	"dreamt/pkg/api"
	rmodels "dreamt/pkg/api/models"
	"dreamt/pkg/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type FiberAPI struct {
	*fiber.App
	*api.API
}

func sendFiberResp(c *fiber.Ctx, resp rmodels.APIResponse) error {
	if resp.Err != nil {
		return c.Status(resp.Status).JSON(resp.Err)
	}

	// write response
	return c.Status(resp.Status).JSON(resp.Body)
}

func NewFiberAPI(api *api.API) *FiberAPI {
	return &FiberAPI{
		fiber.New(),
		api,
	}
}

func (f FiberAPI) FGetDreams(c *fiber.Ctx) error {
	// get dreams from controller
	return sendFiberResp(c, f.GetDreams())
}

func (f FiberAPI) FGetDream(c *fiber.Ctx) error {
	// get id from url
	id := c.Params("id")
	// get dreams from controller
	return sendFiberResp(c, f.GetDream(rmodels.GetDreamRequest{ID: id}))
}

func (f FiberAPI) FGetInterpretation(c *fiber.Ctx) error {
	// get keyword from url
	keyword := c.Params("keyword")

	// get interpretation from controller
	return sendFiberResp(c, f.GetInterpret(rmodels.GetInterpretationRequest{Keyword: keyword}))
}

func (f FiberAPI) FGetKeywords(c *fiber.Ctx) error {
	// get query param limit
	limit := c.Query("limit")

	// get keywords from controller
	return sendFiberResp(c, f.GetKeywords(rmodels.GetKeywordsRequest{Limit: limit}))
}

func (f FiberAPI) FCreateDream(c *fiber.Ctx) error {
	// get dream from body
	var dream models.Dream
	if err := c.BodyParser(&dream); err != nil {
		return c.Status(http.StatusBadRequest).Send([]byte(err.Error()))
	}

	// create dream in controller
	return sendFiberResp(c, f.CreateDream(rmodels.CreateDreamRequest{Dream: dream}))
}

func (f FiberAPI) FDeleteDream(c *fiber.Ctx) error {
	// get id from url
	id := c.Params("id")

	// delete dream in controller
	return sendFiberResp(c, f.DeleteDream(rmodels.DeleteDreamRequest{ID: id}))
}
