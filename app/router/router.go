package router

import (
	"fmt"
	"go-fiber-crud/app/router/middleware"
	v1 "go-fiber-crud/app/router/v1"
	"go-fiber-crud/app/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB) *fiber.App {
	router := fiber.New(fiber.Config{
		// Prefork: true,
		CaseSensitive: true,
		StrictRouting: true,
	})
	router.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin,Authorization",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))
	apiV1 := router.Group("/api")
	v1.AuthRouter(apiV1, db)

	apiV1.Post("/example/upload", func(c *fiber.Ctx) error {
		// multiple files upload
		form, err := c.MultipartForm()
		if err != nil {
			utils.HandleValidationError(c, err)
			return err
		}

		if form.File["files"] == nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "File does not exist",
			})
		}

		path, err := utils.CheckFolderExists("example")

		if err != nil {
			utils.HandleValidationError(c, err)
			return err
		}

		// save all file field
		// for fieldName, files := range form.File {
		// 	for _, file := range files {
		// 		// fileName := fmt.Sprintf("%s-%s", fieldName, file.Filename)
		// 		fileName := utils.RandomImageName(file.Filename)
		// 		destination := fmt.Sprintf("%s/%s", path, fileName)
		// 		if err := c.SaveFile(file, destination); err != nil {
		// 			utils.HandleValidationError(c, err)
		// 			return err
		// 		}
		// 	}
		// }

		for _, file := range form.File["files"] {
			fileName := utils.RandomImageName(file.Filename)
			destination := fmt.Sprintf("%s/%s", path, fileName)
			if err := c.SaveFile(file, destination); err != nil {
				utils.HandleValidationError(c, err)
				return err
			}
		}

		// single file upload
		// file, err := c.FormFile("file")

		// if err != nil {
		// 	// utils.HandleValidationError(c, err)
		// 	return statusBadRequest(err, c)
		// }
		// path := filepath.Join(viper.GetString("app.asset"), "upload_file")
		// destination := fmt.Sprintf("./%s/%s", path, file.Filename)

		// if err := os.MkdirAll(path, os.ModePerm); err != nil {
		// 	return statusBadRequest(err, c)
		// }
		// if err := c.SaveFile(file, destination); err != nil {
		// 	return statusBadRequest(err, c)
		// }

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"msg": "File uploaded successfully",
		})
	})

	apiV1.Get("/assets/:path/:filename", func(c *fiber.Ctx) error {
		path := c.Params("path")
		filename := c.Params("filename")
		// filePath := filepath.Join(viper.GetString("app.asset"), path)
		filePath := fmt.Sprintf("%s/%s", viper.GetString("app.asset"), path)
		file := fmt.Sprintf("%s/%s", filePath, filename)
		if utils.CheckFileExists(file) {
		}
		return c.SendFile(fmt.Sprintf("%s/%s", filePath, filename))
	})

	apiV1.Get("/abouts", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "abouts",
		})
	})

	//middleware
	apiV1.Use(middleware.Authorize())

	// authentication middleware
	v1.CustomerRouter(apiV1, db)
	v1.UserRouter(apiV1, db)
	apiV1.Get("/abouts-with-middleware", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "abouts-with-middleware",
		})
	})

	apiV1.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Not found",
		})
	})
	return router
}

func statusBadRequest(err error, c *fiber.Ctx) error {
	println(err.Error())
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"error": err.Error(),
	})
}
