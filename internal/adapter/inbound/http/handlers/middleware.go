package handlers

import (
	"github.com/Paramet02/multi-user-bookmark-api/internal/port/outbound/security"
	"github.com/gofiber/fiber/v2"
)

type authMiddleware struct {
	jwtManager security.JWTManager
}

func NewAuthMiddleware(jwtManager security.JWTManager) *authMiddleware {
	return &authMiddleware{jwtManager: jwtManager}
}

func (a *authMiddleware) Middleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Cookies("access_token")
		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "missing token",
			})
		}
		
		userID, role, err := a.jwtManager.ValidateToken(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid token",
			})
		}

		ctx := a.jwtManager.InjectUserID(c.Context(), userID, role)
		c.SetUserContext(ctx)

		return c.Next()
	}
}

func (a *authMiddleware) RequestRole(requiredRoles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		_ , role, err := a.jwtManager.ExtractUserID(c.UserContext())
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "unauthorized",
			})
		}

		for _, r := range requiredRoles {
			if r == role {
				return c.Next()
			}
		}
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "forbidden: insufficient role",
		})
	}
}
