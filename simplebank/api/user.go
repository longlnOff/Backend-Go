package api

import (
	"log"
	"net/http"
	db "simplebank/db/sqlc"
	"simplebank/db/util"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

// struct to store the create account request.
// use valudator package internally of Gin to perform data validation automatically under the hood.
type CreateUserRequest struct {
	Username    string `json:"username" binding:"required,alphanum"`
	Password	string `json:"password" binding:"required,min=6"`
	FullName	string `json:"full_name" binding:"required"`
	Email		string `json:"email" binding:"required,email"`
}

type createUserResponse struct {
	Username          string    `json:"username"`
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}



// Note that: ctx.JSON function sends request back to the client
func (server *Server) createUser(ctx *gin.Context) {
	var req CreateUserRequest
	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	hashedPasswrd, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedPasswrd,
		FullName:       req.FullName,
		Email:          req.Email,
	}


	
	user, err := server.store.CreateUser(ctx, arg)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				log.Println(pqErr.Code.Name())
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := createUserResponse{
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		PasswordChangedAt: user.PasswordChangedAt,
		CreatedAt:         user.CreatedAt,
	}

	ctx.JSON(http.StatusOK, response)
}
