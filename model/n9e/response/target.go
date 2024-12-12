package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/n9e"

type TargetResponse struct {
	Target model.Target `json:"target"`
}

type TargetListResponse struct {
	Code  int64           `json:"code"`
	List  []*model.Target `json:"list"`
	Total int64           `json:"total"`
}
