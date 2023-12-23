package handler

import (
    "clinic-appointment-system/pkg/model"
    "clinic-appointment-system/pkg/service"
    "net/http"
    "strconv"
    "time"

    "github.com/gin-gonic/gin"

    "github.com/dgrijalva/jwt-go"
)


type LoginCredentials struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}


type UserHandler struct {
    userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
    return &UserHandler{userService: userService}
}

func (h *UserHandler) RegisterUser(c *gin.Context) {
    var user model.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    newUser, err := h.userService.CreateUser(user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, newUser)
}

func (h *UserHandler) LoginUser(c *gin.Context) {
    var creds LoginCredentials
    if err := c.ShouldBindJSON(&creds); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid login credentials"})
        return
    }

    // Authenticate the user. This part depends on how you store and verify passwords.
    user, err := h.userService.AuthenticateUser(creds.Email, creds.Password)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid login credentials"})
        return
    }

    // Create token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": user.ID,
        "exp":     time.Now().Add(time.Hour * 24).Unix(), // Token expires in 1 day
    })

    // Sign and get the complete encoded token as a string
    tokenString, err := token.SignedString([]byte("YourJWTSecretHere")) // Replace with your secret key
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": tokenString})
}


func (h *UserHandler) GetUser(c *gin.Context) {
    userID, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
        return
    }

    user, err := h.userService.GetUserByID(uint(userID))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
        return
    }

    c.JSON(http.StatusOK, user)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
    userID, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
        return
    }

    var user model.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user.ID = uint(userID) // Set the user ID to the one specified in the route
    updatedUser, err := h.userService.UpdateUser(user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, updatedUser)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
    userID, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
        return
    }

    err = h.userService.DeleteUser(uint(userID))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "user deleted"})
}

