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
	e.GET("/user/:user_id", h.GetUserById)
	return &h
}

// CreateUser ..
func (h *Handler) CreateUser(c echo.Context) error {
	reqMap := map[string]interface{}{}
	user := domain.User{}
	if err := c.Bind(&reqMap); err != nil {
		log.Fatal(err)
	}
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

// GetUserById ..
func (h *Handler) GetUserById(c echo.Context) error {
	reqMap := map[string]interface{}{}

	// gutils.EchoBind(c, &reqMap, []string{"user_id"})
	// gutils.EchoGatewayLogger(c, "GetUser")
	if err := c.Bind(&reqMap); err != nil {
		log.Fatal(err)
	}
	reqByte, _ := json.Marshal(reqMap)
	log.Println("Request payload: ", string(reqByte))

	userID := reqMap["user_id"].(string)

	result, err := h.usecase.GetUserById(userID)

	if err != nil {
		// return utils.ErrorResponse(c, err)
		return c.JSON(http.StatusOK, map[string]interface{}{"success": false, "message": "GetUserById error"})
	}

	// return utils.SuccessResponse(c, result)
	return c.JSON(http.StatusOK, map[string]interface{}{"success": true, "data": result})
}
