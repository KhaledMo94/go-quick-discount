package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckErrorAndReturnData(
    ctx *gin.Context,
    data any,
    err *error,
    successMsg string,
    customResponseCode int,
) {
    if err != nil && *err != nil {
        ctx.JSON(
            http.StatusInternalServerError,
            gin.H{
                "status":  "error",
                "message": "internal server error",
                "data":    nil,
            },
        )
        return
    }

    if customResponseCode >= 500 || customResponseCode < 1 {
        customResponseCode = http.StatusOK
    }

    ctx.JSON(
        customResponseCode,
        gin.H{
            "status":  "success",
            "message": successMsg,
            "data":    data,
        },
    )
}

func NormalizeJSONByLocale(jsonEscaped string , locale string)(string , error){
	var fields map[string]string

	err := json.Unmarshal([]byte(jsonEscaped) , &fields)

	if err != nil {
		return "",err
	}

	return fields[locale] , nil
}