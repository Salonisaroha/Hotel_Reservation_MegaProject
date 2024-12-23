package api

import (
	"context"

	"github.com/Salonisaroha/db"
	"github.com/Salonisaroha/types"
	"github.com/gofiber/fiber/v2"
)
type UserHandler struct{
	userStore db.UserStore
}
func NewUserHandler(userStore db.UserStore) *UserHandler{
	return &UserHandler{
            userStore: userStore,
	}
}
func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	var(
		id = c.Params("id")
		ctx = context.Background()
	)
	
	user, err :=h.userStore.GetUserByID(ctx,id)
	if err != nil{
		return err
	}
	return c.JSON(user)
}
func(h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "Saloni",
		LastName:  "Saroha",
	}
	return c.JSON(u)
}

