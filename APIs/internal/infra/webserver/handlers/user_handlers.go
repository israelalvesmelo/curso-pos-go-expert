package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/israelalvesmelo/curso-pos-go-expert/APIs/internal/dto"
	"github.com/israelalvesmelo/curso-pos-go-expert/APIs/internal/entity"

	"github.com/israelalvesmelo/curso-pos-go-expert/APIs/internal/infra/database"
)

type Error struct {
	Message string `json:"message"`
}

type UserHandler struct {
	UserDB database.UserInterface
}

func NewUserHandler(db database.UserInterface) *UserHandler {
	return &UserHandler{
		UserDB: db,
	}
}

// GetJWT godoc
// @Summary      Get a user JWT
// @Description  Get a user JWT
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request   body     dto.GetJWTInput  true  "user credentials"
// @Success      200  {object}  dto.GetJWTOutput
// @Failure      404  {object}  Error
// @Failure      500  {object}  Error
// @Router       /users/generate_token [post]
func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	jwtExpiresIn := r.Context().Value("JwtExperesIn").(int)
	var jwtInput dto.GetJWTInput
	err := json.NewDecoder(r.Body).Decode(&jwtInput)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user, err := h.UserDB.FindByEmail(jwtInput.Email)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		err := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}

	if !user.ValidatePassword(jwtInput.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, tokenString, _ := jwt.Encode(
		map[string]interface{}{
			"sub": user.ID.String(),
			"exp": time.Now().Add(time.Second * time.Duration(jwtExpiresIn)).Unix(),
		})

	accessToken := dto.GetJWTOutput{AccessToken: tokenString}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}

// Create user godoc
// @Summary      Create user
// @Description  Create user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request     body      dto.UserInput  true  "user request"
// @Success      201
// @Failure      500         {object}  Error
// @Router       /users [post]
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userDto dto.UserInput
	err := json.NewDecoder(r.Body).Decode(&userDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := entity.NewUser(userDto.Name, userDto.Email, userDto.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	err = h.UserDB.Create(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
