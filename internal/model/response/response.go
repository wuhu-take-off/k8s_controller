package response

import "github.com/gin-gonic/gin"

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewResponse(code int, message string, data interface{}, c *gin.Context) {
	c.JSON(code, &Response{
		Code:    code,
		Message: message,
		Data:    data,
	})

}
func NewErrorResponse(code int, message string, c *gin.Context) {
	c.JSON(code, &ErrorResponse{
		Code:    code,
		Message: message,
	})
}

func NewResponseWithError(err error, c *gin.Context) {
	NewErrorResponse(500, err.Error(), c)
}
func NewResponseWithErrorString(message string, c *gin.Context) {
	NewErrorResponse(500, message, c)
}
func NewResponseWithData(data interface{}, c *gin.Context) {
	NewResponse(200, "success", data, c)
}
