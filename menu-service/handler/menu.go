package handler

import (
	"net/http"

	"github.com/syaripuddin/pengenalan-microservice/utils"
)

func AddMenu(w http.ResponseWriter, r *http.Request) {
	utils.WrapAPISuccess(w, r, "success", http.StatusOK)
}
