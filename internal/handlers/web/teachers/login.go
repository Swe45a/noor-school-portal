package teachers

import (
	"net/http"

	"school-portal/internal/dtos"
	apperrors "school-portal/internal/errors"
	"school-portal/internal/utils"
)

// Login handles POST /api/teachers/login. Both the ID Number and Employee Number must belong
// to the same teacher record; the service verifies this as a single lookup so the response
// never reveals which of the two values was wrong.
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req dtos.TeacherLoginRequest
	if err := utils.DecodeJSON(r, &req); err != nil {
		utils.WriteError(w, err)
		return
	}

	if !utils.ValidID(req.IDNumber) || !utils.ValidID(req.EmployeeNumber) {
		utils.WriteError(w, apperrors.ErrInvalidRequest)
		return
	}

	result, err := h.service.Login(r.Context(), req.IDNumber, req.EmployeeNumber)
	if err != nil {
		utils.WriteError(w, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, dtos.TeacherLoginResponse{
		FullName:        result.FullName,
		AccountUsername: result.AccountUsername,
		AccountPassword: result.AccountPassword,
	})
}
