package mistake

import (
	"net/http"
)

var (
	ErrInvalidCredentials = "Неверный логин или пароль"

	ErrTicketNotFound = "Билет не найден"

	ErrPassengerNotFound = "Билет не найден"

	ErrDatabaseFailure = "Ошибка базы данных"

	ErrRateLimitExceeded = "Превышен лимит запросов"

	ErrInvalidRequestData = "Некорректные данные запроса"

	ErrEntityNotFound = "Сущность не найдена"

	ErrTokenGeneration = "Ошибка генерации токена"

	ErrUnknown = "Неизвестная ошибка"

	ErrDocumentNotFound = "Документ не найден"
)

var ErrorMap = map[string]int{
	ErrInvalidCredentials: http.StatusUnauthorized,
	ErrTicketNotFound:     http.StatusNotFound,
	ErrDatabaseFailure:    http.StatusInternalServerError,
	ErrRateLimitExceeded:  http.StatusTooManyRequests,
	ErrInvalidRequestData: http.StatusBadRequest,
	ErrEntityNotFound:     http.StatusNotFound,
	ErrTokenGeneration:    http.StatusInternalServerError,
	ErrUnknown:            http.StatusInternalServerError,
	ErrDocumentNotFound:   http.StatusNotFound,
	ErrPassengerNotFound:  http.StatusNotFound,
}
