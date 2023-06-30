package handler

import (
	"net/http"
	"os"
	"time"

	"github.com/KhoirulAziz99/mnc/internal/models"
	"github.com/KhoirulAziz99/mnc/internal/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type LogHandler interface {
	Login(c *gin.Context)
	Logout(c *gin.Context)
}

type logHandler struct {
	customerService services.CustomerServices
}

func NewLogHandler(customerServices services.CustomerServices) *logHandler {
	return &logHandler{customerService: customerServices}
}

func (l logHandler) Login(c *gin.Context) {
	var customer models.LoginCustomer

	err := c.ShouldBindJSON(&customer) //binding data
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	customerExist, err := l.customerService.GetByEmail(customer.Email) //mencocokkan dengan email yang terdaftar di database
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error :": "Invalid email"})
		return
	}
	if customerExist.Password != customer.Password { //mencocokkan dengan password yang ada di databse
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password"})
		return
	}

	token := jwt.New(jwt.SigningMethodHS256) //Membuat objek token baru 

	claims := token.Claims.(jwt.MapClaims) //Menginisialisasi objek claims sebagai MapClaims dari token yang dibuat sebelumnya
	claims["email"] = customer.Email //Menetapkan email customer ke claim "email" dalam token.
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() //Menetapkan waktu kadaluarsa token dengan menambahkan 24 jam ke waktu saat ini dan mengonversikannya ke Unix timestamp.

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET"))) //Menandatangani token menggunakan secret key yang diambil dari environment variable.
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{ //Kalau berhasil akan dikirimi token
		"message": "Your token will expire in 24 hours",
		"token":   tokenString,
	})

}

func (l logHandler) Logout(c *gin.Context) {
    // Hapus header Authorization
    c.Header("Authorization", "") //tergantung tokennya disimpan di mana.....bisa di cookie atau di header autorization

    c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}

