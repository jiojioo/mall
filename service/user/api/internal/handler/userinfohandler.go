package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"mall/service/user/api/internal/errcode"
	"mall/service/user/api/internal/logic"
	"mall/service/user/api/internal/svc"
	"mall/service/user/api/internal/types"

	"mall/service/user/api/internal/validate"
)

func UserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfoRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		if err := validate.Validate(req); err != nil {
			httpx.OkJson(w, errcode.ValidateErr(err))
			return
		}

		l := logic.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, errcode.OK(resp))
		}
	}
}
