package services

import (
	infraestructure "github.com/WelintonJunior/billing-and-subscription-service/infraestructure/postgres"
	"github.com/WelintonJunior/billing-and-subscription-service/types"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func LogAction(c *fiber.Ctx, userID uuid.UUID, action string) error {
	ip := c.IP()
	userAgent := c.Get("User-Agent")

	auditLog := types.AuditLog{
		UserID:    userID,
		Action:    action,
		IpAddress: ip,
		UserAgent: userAgent,
	}

	return infraestructure.Db.WithContext(c.Context()).Create(&auditLog).Error
}
