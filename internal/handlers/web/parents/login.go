package parents

import (
	"net/http"

	"school-portal/internal/dtos"
	apperrors "school-portal/internal/errors"
	"school-portal/internal/utils"
)

// Login handles POST /api/parents/login. There is no OTP step: a valid Parent ID alone
// authenticates the parent, and the response carries the account login information directly.
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req dtos.ParentLoginRequest
	if err := utils.DecodeJSON(r, &req); err != nil {
		utils.WriteError(w, err)
		return
	}

	if !utils.ValidID(req.ParentID) {
		utils.WriteError(w, apperrors.ErrInvalidRequest)
		return
	}

	result, err := h.service.Login(r.Context(), req.ParentID)
	if err != nil {
		utils.WriteError(w, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, dtos.ParentLoginResponse{
		FullName:        result.FullName,
		AccountUsername: result.AccountUsername,
		AccountPassword: result.AccountPassword,
	})
}
