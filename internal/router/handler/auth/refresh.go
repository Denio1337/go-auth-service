package auth

import (
	cerror "app/internal/router/types/error"
	"app/internal/router/types/response"
	"app/internal/service/auth"
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// PATCH: Refresh tokens
func Refresh(c *fiber.Ctx) error {
	// Get user info from cookie
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	// Route to service
	result, err := auth.Refresh(&auth.RefreshParams{
		Username:  claims["username"].(string),
		ID:        uint(claims["id"].(float64)),
		Identity:  c.Locals("identity").(string),
		PairID:    claims["pair_id"].(string),
		UserAgent: c.Locals("userAgent").(string),
		IP:        c.Locals("ip").(string),
	})

	// Handle error from service
	if err != nil {
		// User agent changed
		if errors.Is(err, auth.ErrUserAgentChanged) {
			setAuthCookies(c, nil)

			cerr := *cerror.ErrForbidden
			cerr.Message = err.Error()
			return &cerr
		}

		// IP changed
		webhookURL, ok := c.Locals("webhook").(string)
		if errors.Is(err, auth.ErrIPChanged) && ok {
			go sendWebhookNotification(webhookURL, uint(claims["id"].(float64)))
		}
	}

	// Update tokens
	setAuthCookies(c, result)

	return c.JSON(response.SuccessResponse(struct{}{}))
}

func sendWebhookNotification(webhookURL string, userID uint) {
	payload := WebhookPayload{
		ID:        userID,
		Message:   "Attempt to enter from new IP",
		Timestamp: time.Now(),
	}

	body, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", webhookURL, bytes.NewBuffer(body))
	if err != nil {
		log.Printf("failed to create request: %v", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("webhook request error: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		log.Printf("webhook failed with status code: %d", resp.StatusCode)
	}
}
