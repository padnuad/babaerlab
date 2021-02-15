package delivery

import (
	"baberlab/domain"
	"encoding/json"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/trewzaki/gutils"
)

type Handler struct {
	usecase domain.UserUsecase
}

func NewHandler(e *echo.Group, u domain.UserUsecase) *Handler {
	h := Handler{usecase: u}
	e.POST("/users", h.CreateUser)
	return &h
}

// CreateUser ..
func (h *Handler) CreateUser(c echo.Context) error {
	reqMap := map[string]interface{}{}
	user := domain.User{}

	// gutils.EchoBind(c, &reqMap, []string{"user_id"})
	// gutils.EchoGatewayLogger(c, "CreateUser")

	reqByte, _ := json.Marshal(reqMap)
	log.Println("Request payload: ", string(reqByte))

	userID := reqMap["user_id"].(string)

	json.Unmarshal(reqByte, &user)
	user.UserID = userID
	result, _ := h.usecase.CreateUser(user)

	// if err != nil {
	// 	return c.JSONBlob(http.StatusOK, gutils.SendResponse(c, err))
	// }

	// return utils.SuccessResponse(c, result)
	return c.JSONBlob(http.StatusOK, gutils.SendResponse(true, nil, result))
}
