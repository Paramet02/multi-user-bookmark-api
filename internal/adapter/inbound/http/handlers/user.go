package handlers

import (
	"strconv"

	"github.com/Paramet02/multi-user-bookmark-api/internal/adapter/inbound/http/response/mapper"
	"github.com/Paramet02/multi-user-bookmark-api/internal/adapter/inbound/http/response/model"
	"github.com/Paramet02/multi-user-bookmark-api/internal/port/inbound"
	"github.com/Paramet02/multi-user-bookmark-api/internal/port/outbound/security"
	"github.com/gofiber/fiber/v2"

	"time"
)

// userHandler struct implements the HTTP handler for user-related operations.
type userHandler struct {
	userService inbound.UserService 
	jwtSecurity security.JWTManager
}

// NewUserHandler creates a new instance of userHandler with the provided userService and userSecurity.
// This function is used to initialize the user handler with the necessary dependencies.
func NewUserHandler(userService inbound.UserService , jwtSecurity security.JWTManager) *userHandler {
	return &userHandler{userService: userService , jwtSecurity: jwtSecurity}
}

func (h *userHandler) GetUserID(c *fiber.Ctx) error {
	uid := c.Params("id")
	if uid == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "User ID is required",
		})
	}

	id, err := strconv.Atoi(uid)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID format",
		})
	}

	ctx := c.UserContext()

	user, err := h.userService.GetUserByID(ctx, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if user == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	return c.JSON(user)
}

func (h *userHandler) GetUserByEmail(c *fiber.Ctx) error {
	email := c.Query("email")

	if email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Email is required",
		})
	}

	ctx := c.UserContext()

	userEmail, err := h.userService.GetUserByEmail(ctx, email)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if userEmail == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	return c.JSON(userEmail)
}

func (h *userHandler) Register(c *fiber.Ctx) error {
	var req model.RegisterRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request format",
		})
	}

	if req.Email == "" || req.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Email and password are required",
		})
	}

	ctx := c.UserContext()

	user , err := h.userService.RegisterUser(ctx , req.Email , req.Username , req.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	publicUser := mapper.ToUserPublicResponse(user)

	return c.JSON(publicUser)
}

func (h *userHandler) UpdateUser(c *fiber.Ctx) error {
    ctx := c.UserContext()

    // ดึง user id จาก JWT (เหมือน DeleteUser)
    userID, _, err := h.jwtSecurity.ExtractUserID(ctx)
    if err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "error": "unauthorized",
        })
    }

    var req model.UpdateUserRequest
    if err := c.BodyParser(&req); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "invalid request format",
        })
    }

    // Validation ขั้นต้น
    if req.Username == "" && req.Password == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "at least one field (username or password) must be provided",
        })
    }

    // เรียก service
	updatedUser, err := h.userService.UpdateUser(ctx, userID, req.Email, req.Username, req.Password)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "failed to update user",
        })
    }

    // map กลับไปเป็น public response
    publicUser := mapper.ToUserPublicResponse(updatedUser)

    return c.JSON(publicUser)
}


func (h *userHandler) DeleteUser(c *fiber.Ctx) error {
	ctx := c.UserContext()
	userid , _ , err := h.jwtSecurity.ExtractUserID(ctx)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error" : err,
		})
	}

	if err := h.userService.DeleteUser(ctx , userid); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error" : err,
		})
	}

	c.Cookie(&fiber.Cookie{
        Name:     "access_token",
        Value:    "",
        Path:     "/",
        Expires:  time.Unix(0, 0),
        HTTPOnly: true,
        Secure:   false,
        SameSite: "none",
    })

	return c.JSON(fiber.Map{"message": "user deleted"})
}
