package loginHandler

import (
	"log"
	"time"

	"github.com/azizanhakim/items-api-fiber/database"
	"github.com/azizanhakim/items-api-fiber/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// PrivateKey is the RSA private key used for signing JWT tokens
var (
	privateKeyPEM = `-----BEGIN RSA PRIVATE KEY-----
MIICWgIBAAKBgEhAvj2GbFMJRIfsLJI1E6tGS2iqTzhBpdFkZaRbD6KVtwE10++1
DeH5+q49jT+GDXBgGoqQ9qaoE2DK3Ap/H/+JZR4WrNEbZw3bHFJo4n/Z0QqwzINu
2BBACJY4I9A44yGjt7cCE4zgHA/yht5OcSWemjDt/DsOxETL/RzlnhX5AgMBAAEC
gYAhFhzH2dWjFLxgLeOfNFUEepUvocXTMiS3xWzSHa0EO+Do2fhqbZOk5q9HuQIE
k+N1kVy2FXoNiSwOh/bJi7tf8agzVigVxlEoM1j9ZKjym/IYbRYmK7nlYzsdfu4W
dLkaD4w+417fzC9r5fVqjemjyYnLaXguEoHDssEvjCOhgQJBAI1t08tyIzjojeaR
oH1b72pIq+HmHx3x6gxGRkdEe8NwLbwRCqxZrCsTfrbL01I2EntK0aLd8Oe8vycZ
juHt+lECQQCCyN5vVX+wtDObCF9rOnsi/DLm1LCCRXKdW9L9dT63hxyJiOHJ1S53
z3CX8IGASzMMQ9b+3Jxpcmu7wmUvok8pAkB6y/xULhMFC26B3rmxfsyexOBwsMUd
0/k6lR3aLU0kgVdEbquMwANsF24zO0CNpiNf57OjrP7JxylVwqw74MwxAkAfnrsQ
xeYibd8QO5z+StxcoAcagg/O30WPwBSqDP/F1ZfTtNGKP82FUBUT1yUoRIYqD/ja
E7KJWA2uDpjyRFbJAkBCLGxcXSjLgLuc9CDwqYMjFijyue4PuLtlOnVTHjMPdc+j
KYbxMfUDcywLgyr+3XRKFtz3ZZwCpuoqr/HnAIi+
-----END RSA PRIVATE KEY-----`

	publicKeyPEM = `-----BEGIN PUBLIC KEY-----
MIGeMA0GCSqGSIb3DQEBAQUAA4GMADCBiAKBgEhAvj2GbFMJRIfsLJI1E6tGS2iq
TzhBpdFkZaRbD6KVtwE10++1DeH5+q49jT+GDXBgGoqQ9qaoE2DK3Ap/H/+JZR4W
rNEbZw3bHFJo4n/Z0QqwzINu2BBACJY4I9A44yGjt7cCE4zgHA/yht5OcSWemjDt
/DsOxETL/RzlnhX5AgMBAAE=
-----END PUBLIC KEY-----`

	// Parse the private key
	PrivateKey, _ = jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKeyPEM))

	// Parse the public key
	PublicKey, _ = jwt.ParseRSAPublicKeyFromPEM([]byte(publicKeyPEM))
)

func Login(c *fiber.Ctx) error {

	username := c.FormValue("username")
	password := c.FormValue("password")

	db := database.DB
	var user model.User

	// Throws Unauthorized error

	db.Find(&user, "username = ?", username)

	if user.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Username not found", "data": nil})
	}

	if !user.CheckPassword(password) {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"name":  user.Fullname,
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString(PrivateKey)
	if err != nil {
		log.Printf("token.SignedString: %v", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}
