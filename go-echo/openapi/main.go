package main

import (
	"github.com/GIT_USER_ID/GIT_REPO_ID/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

    //todo: handle the error!
	c, _ := handlers.NewContainer()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())


	// AddPet - Add a new pet to the store
	e.POST("/api/v3/api/d/pet", c.AddPet)

	// DeletePet - Deletes a pet
	e.DELETE("/api/v3/pet/:petId", c.DeletePet)

	// FindPetsByStatus - Finds Pets by status
	e.GET("/api/v3/pet/findByStatus", c.FindPetsByStatus)

	// FindPetsByTags - Finds Pets by tags
	e.GET("/api/v3/pet/findByTags", c.FindPetsByTags)

	// GetPetById - Find pet by ID
	e.GET("/api/v3/pet/:petId", c.GetPetById)

	// UpdatePet - Update an existing pet
	e.PUT("/api/v3/api/d/pet", c.UpdatePet)

	// UpdatePetWithForm - Updates a pet in the store with form data
	e.POST("/api/v3/pet/:petId", c.UpdatePetWithForm)

	// UploadFile - uploads an image
	e.POST("/api/v3/pet/:petId/uploadImage", c.UploadFile)

	// DeleteOrder - Delete purchase order by ID
	e.DELETE("/api/v3/store/order/:orderId", c.DeleteOrder)

	// GetInventory - Returns pet inventories by status
	e.GET("/api/v3/store/inventory", c.GetInventory)

	// GetOrderById - Find purchase order by ID
	e.GET("/api/v3/store/order/:orderId", c.GetOrderById)

	// PlaceOrder - Place an order for a pet
	e.POST("/api/v3/store/order", c.PlaceOrder)

	// CreateUser - Create user
	e.POST("/api/v3/user", c.CreateUser)

	// CreateUsersWithListInput - Creates list of users with given input array
	e.POST("/api/v3/user/createWithList", c.CreateUsersWithListInput)

	// DeleteUser - Delete user
	e.DELETE("/api/v3/user/:username", c.DeleteUser)

	// GetUserByName - Get user by user name
	e.GET("/api/v3/user/:username", c.GetUserByName)

	// LoginUser - Logs user into the system
	e.GET("/api/v3/user/login", c.LoginUser)

	// LogoutUser - Logs out current logged in user session
	e.GET("/api/v3/user/logout", c.LogoutUser)

	// UpdateUser - Update user
	e.PUT("/api/v3/user/:username", c.UpdateUser)


	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}