package utils
import (
	"os"
	"log"
    "time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)
var JwtSecret []byte
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("JWT_SECRET not set")
	}

	JwtSecret = []byte(secret)
}
type Claims struct{
    UserID uuid.UUID `json:"userId"`
	Email  string    `json:"email"`
	RoleID uuid.UUID `json:"roleId"`
	jwt.RegisteredClaims
}
//roleID uuid.UUID in the GenretaeJWT parameter
func GenerateJWT(userID uuid.UUID , email string  )(string ,error){
    claims:=Claims{
		UserID: userID,
		Email: email,
		// RoleID: roleID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtSecret)
}