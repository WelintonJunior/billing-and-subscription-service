package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/stripe/stripe-go/v82"
)

func CreateSubscription(sc *stripe.Client, customerID, priceID string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		customerID := c.Query("customer_id", "")
		priceID := c.Query("price_id", "")

		if customerID == "" || priceID == "" {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"error":   ErrMalformedRequest,
			})
		}

		params := &stripe.SubscriptionCreateParams{
			Customer: stripe.String(customerID),
			Items: []*stripe.SubscriptionCreateItemParams{
				{
					Price: stripe.String(priceID),
				},
			},
			PaymentBehavior: stripe.String("default_incomplete"),
			Expand:          []*string{stripe.String("latest_invoice.payment_intent")},
			TrialPeriodDays: stripe.Int64(7),
		}

		subscription, err := sc.V1Subscriptions.Create(c.Context(), params)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"error":   err,
			})
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"success":      true,
			"subscription": subscription,
		})
	}
}
