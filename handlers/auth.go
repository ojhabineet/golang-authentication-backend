package handlers

import (
	"backend-auth/models"
	"backend-auth/storage"
	"backend-auth/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type SignupInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Signup(c *gin.Context) {

	var input SignupInput

	c.ShouldBindJSON(&input)

	// check existing user
	for _, user := range storage.Users {
		if user.Email == input.Email {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Email already exists",
			})
			return
		}
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword(
		[]byte(input.Password),
		bcrypt.DefaultCost,
	)

	user := models.User{
		ID:       len(storage.Users) + 1,
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
		Role:     "user",
	}

	storage.Users = append(storage.Users, user)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Signup successful",
	})
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {

	var input LoginInput

	c.ShouldBindJSON(&input)

	for _, user := range storage.Users {

		if user.Email == input.Email {

			err := bcrypt.CompareHashAndPassword(
				[]byte(user.Password),
				[]byte(input.Password),
			)

			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{
					"error": "Wrong password",
				})
				return
			}

			token, _ := utils.GenerateJWT(user.ID, user.Role)

			c.JSON(http.StatusOK, gin.H{
				"token": token,
			})

			return
		}
	}

	c.JSON(http.StatusUnauthorized, gin.H{
		"error": "User not found",
	})
}

func GetProfile(c *gin.Context) {

	userID := int(c.MustGet("user_id").(float64))

	for _, user := range storage.Users {

		if user.ID == userID {

			c.JSON(http.StatusOK, gin.H{
				"id":    user.ID,
				"name":  user.Name,
				"email": user.Email,
				"role":  user.Role,
			})

			return
		}
	}
}

func GetUsers(c *gin.Context) {

	role := c.MustGet("role")

	if role != "admin" {

		c.JSON(http.StatusForbidden, gin.H{
			"error": "Admin only",
		})

		return
	}

	c.JSON(http.StatusOK, storage.Users)
}