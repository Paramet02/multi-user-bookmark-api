package handlers

import (
	"github.com/Paramet02/multi-user-bookmark-api/internal/port/inbound"
	"github.com/gofiber/fiber/v2"
)

type LoginRequset struct {
	Identifier string `json:"identifier"` // username OR email
	Password string `json:"password"`
}

type authHandler struct {
	authHand inbound.AuthService
}

func NewAuthHandler(authHand inbound.AuthService) *authHandler {
	return &authHandler{authHand: authHand}
}

func (s *authHandler) Login(c *fiber.Ctx) error {
	var req LoginRequset
	ctx := c.UserContext()
	
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid input"})
	}

	token , err := s.authHand.Login(ctx , req.Identifier , req.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid credentials"})
	}

	c.Cookie(&fiber.Cookie{
		Name: "access_token",
		Value: token,
		Path: "/",
		Secure: false, 
		HTTPOnly: true, // protection xss
		MaxAge: 86400, // 1 day
		SameSite: "none",
	})

	return c.JSON(fiber.Map{
		"message": "login successful",
		"token": token,
	})
}
