package mahasiswa

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"vps_server_playground/util"
)

type Controller interface {
	CreateMahasiswa(ctx *fiber.Ctx) error
	GetMahasiswaByID(ctx *fiber.Ctx) error
}

type controllerImpl struct {
	service Service
}

func NewController(service Service) Controller {
	return &controllerImpl{service}
}

func (c *controllerImpl) CreateMahasiswa(req *fiber.Ctx) error {
	createReq := new(createRequest)
	if err := req.BodyParser(createReq); err != nil {
		return err
	}
	log.Println(createReq)
	ctx := context.Background()
	mhs, err := c.service.Create(ctx, createReq)
	if err != nil {
		return err
	}
	return req.Status(http.StatusCreated).JSON(&util.WebResponse{
		StatusCode: http.StatusCreated,
		Status:     "created",
		Data:       mhs,
	})
}

func (c *controllerImpl) GetMahasiswaByID(req *fiber.Ctx) error {
	id := req.Params("id")
	ctx := context.Background()
	mhs, err := c.service.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if mhs == nil {
		return req.Status(http.StatusNotFound).JSON(&util.WebResponse{
			StatusCode: http.StatusNotFound,
			Status:     "not found",
			Data:       nil,
		})
	}
	return req.Status(http.StatusOK).JSON(&util.WebResponse{
		StatusCode: http.StatusOK,
		Status:     "ok",
		Data:       mhs,
	})
}
