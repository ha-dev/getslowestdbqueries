package controllers

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/ha-dev/getslowestdbqueries/env"
	"github.com/ha-dev/getslowestdbqueries/models"
	"github.com/ha-dev/getslowestdbqueries/platform/database"
)

func GetQuerys(c *fiber.Ctx) error {

	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	listSlowestQuerys, err := db.GetDbQueries()
	if err != nil {
		// Return, if books not found.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":             true,
			"msg":               "listSlowestQuerys were not found",
			"count":             0,
			"listSlowestQuerys": nil,
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error":             false,
		"msg":               nil,
		"count":             len(listSlowestQuerys),
		"listSlowestQuerys": listSlowestQuerys,
	})

}

func AddQuerysForCheckSlowest(c *fiber.Ctx) error {

	//get param post
	p := new(models.Querys)
	if err := c.BodyParser(p); err != nil {
		log.Println(err)
		return c.Status(400).JSON(&fiber.Map{
			"error": true,
			"msg":   err,
		})

	}

	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	//check slowest
	elapsed, err := db.ExecuteDbQueries(p.Query)
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	log.Printf("query took %d", elapsed)

	if elapsed > env.PivotTimetoCheckSlowesQuery {
		p.Time_to_run = fmt.Sprintf("%d", elapsed)
		log.Println("Query Is Slowest")
		log.Println(p)
		//insert to db
		err := db.InsertDbQueries(p)
		if err != nil {
			// Return status 500 and database connection error.
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}

	}

	return nil

}
