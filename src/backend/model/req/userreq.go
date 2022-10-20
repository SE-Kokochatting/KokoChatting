package req

type UserReq struct {
	Name string `json:"name"`
	Password string `json:"password"`
}