package delivery

import (
	"baberlab/domain"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/trewzaki/gutils"
)

type Handler struct {
	usecase domain.UserUsecase
}

func NewHandler(e *echo.Group, u domain.UserUsecase) *Handler {
	h := Handler{usecase: u}
	e.POST("/user/add", h.CreateUser)
	e.GET("/user/:user_id", h.GetUserByID)
	e.GET("/user", h.GetUser)
	return &h
}

// CreateUser ..
func (h *Handler) CreateUser(c echo.Context) error {
	log.Println("[X] CreateUser")
	reqMap := map[string]interface{}{}
	user := domain.User{}
	if err := c.Bind(&reqMap); err != nil {
		log.Fatal(err)
	}
	// gutils.EchoBind(c, &reqMap, []string{"user_id"})
	// gutils.EchoGatewayLogger(c, "CreateUser")
	fmt.Println(reqMap)
	reqByte, _ := json.Marshal(reqMap)
	log.Println("Request payload: ", string(reqByte))

	// userID := reqMap["user_id"].(string)

	json.Unmarshal(reqByte, &user)
	// user.UserID = userID
	err := h.usecase.CreateUser(user)
	// if err := h.usecase.CreateUser(user); err != nil {
	// 	return c.JSON(http.StatusOK, map[string]interface{}{"success": false, "message": "GetUserByID error"})
	// }
	if err != nil {
		// return c.JSONBlob(http.StatusOK, gutils.SendResponse(c, err))
		return c.JSON(http.StatusOK, map[string]interface{}{"success": false, "message": "GetUserByID error"})
	}

	// return utils.SuccessResponse(c, result)
	// return c.JSONBlob(http.StatusOK, gutils.SendResponse(true, nil, nil))
	return c.JSON(http.StatusOK, map[string]interface{}{"success": true, "message": "GetUserByID"})
}

// GetUserByID ..
func (h *Handler) GetUserByID(c echo.Context) error {
	log.Println("[X] CreateUser")
	reqMap := map[string]interface{}{}

	gutils.EchoBind(c, &reqMap, []string{})
	gutils.EchoGatewayLogger(c, "GetUserByID")
	// if err := c.Bind(&reqMap); err != nil {
	// 	log.Fatal(err)
	// }
	fmt.Println(reqMap)
	reqByte, _ := json.Marshal(reqMap)
	log.Println("Request payload: ", string(reqByte))

	userID := reqMap["user_id"].(string)

	result, err := h.usecase.GetUserByID(userID)

	if err != nil {
		// return utils.ErrorResponse(c, err)
		return c.JSON(http.StatusOK, map[string]interface{}{"success": false, "message": "GetUserByID error"})
	}

	// return utils.SuccessResponse(c, result)
	return c.JSON(http.StatusOK, result)
}

// GetUser ..
func (h *Handler) GetUser(c echo.Context) error {
	log.Println("[X] GetUser")
	reqMap := map[string]interface{}{}

	gutils.EchoBind(c, &reqMap, []string{})
	gutils.EchoGatewayLogger(c, "GetUser")
	// if err := c.Bind(&reqMap); err != nil {
	// 	log.Fatal(err)
	// }
	fmt.Println(reqMap)
	reqByte, _ := json.Marshal(reqMap)
	log.Println("Request payload: ", string(reqByte))

	result, err := h.usecase.GetUser()

	if err != nil {
		// return utils.ErrorResponse(c, err)
		return c.JSON(http.StatusOK, map[string]interface{}{"success": false, "message": "GetUser error"})
	}

	// return utils.SuccessResponse(c, result)
	return c.JSON(http.StatusOK, result)
}
