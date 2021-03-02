package controllers

import (
	"github.com/gofiber/fiber/v2"
)

// GetPersons is to get all Persons data
func GetPersons(c *fiber.Ctx) error {
	return c.SendString("All Persons")
}

// GetPerson is to get one Person data
func GetPerson(c *fiber.Ctx) error {
	return c.SendString("Single Person")
}

// NewPerson is create new Person data
func NewPerson(c *fiber.Ctx) error {
	return c.SendString("New Person")
}

// DeletePerson is to delete Person data
func DeletePerson(c *fiber.Ctx) error {
	return c.SendString("Delete Person")
}
