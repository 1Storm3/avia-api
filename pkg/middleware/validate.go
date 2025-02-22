package middleware

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

type ValidationSource int

const (
	FromBody ValidationSource = 1 << iota
	FromQuery
	FromBoth = FromBody | FromQuery
)

func ValidateMiddleware[T any](source ValidationSource) fiber.Handler {
	return func(c *fiber.Ctx) error {
		reqData := new(T)
		if source&FromBody != 0 {
			if err := c.BodyParser(reqData); err != nil {
				return c.Status(http.StatusBadRequest).JSON(fiber.Map{
					"statusCode": http.StatusBadRequest,
					"message":    "Некорректные данные запроса (Body)",
				})
			}
		}

		if source&FromQuery != 0 {
			if err := c.QueryParser(reqData); err != nil {
				return c.Status(http.StatusBadRequest).JSON(fiber.Map{
					"statusCode": http.StatusBadRequest,
					"message":    "Некорректные параметры запроса (Query)",
				})
			}
		}

		if err := validate.Struct(reqData); err != nil {
			var ve validator.ValidationErrors
			if errors.As(err, &ve) {
				var errorMessages []string
				for _, fe := range ve {
					switch fe.Tag() {
					case "min":
						errorMessages = append(errorMessages, fmt.Sprintf("Поле '%s' должно содержать минимум %s символов", fe.Field(), fe.Param()))
					case "required":
						errorMessages = append(errorMessages, fmt.Sprintf("Поле '%s' обязательно для заполнения", fe.Field()))
					case "email":
						errorMessages = append(errorMessages, "Некорректный формат email")
					default:
						errorMessages = append(errorMessages, fmt.Sprintf("Поле '%s' некорректно", fe.Field()))
					}
				}

				return c.Status(http.StatusBadRequest).JSON(fiber.Map{
					"statusCode": http.StatusBadRequest,
					"message":    errorMessages[0],
				})
			}
		}
		return c.Next()
	}
}
