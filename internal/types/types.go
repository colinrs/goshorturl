// Code generated by goctl. DO NOT EDIT.
package types

type AccessShortUrlRequest struct {
	Url string `path:"url"`
}

type AccessShortUrlResponse struct {
}

type CreatShortUrlResponse struct {
	Id  uint   `json:"id"`
	Url string `json:"url"`
}

type CreateShortUrlRequest struct {
	Origin      string `json:"origin"`
	Description string `json:"description"`
	ExpireAt    string `json:"expire_at"` // ISO 8601 格式 格式: "2023-08-15T14:30:00+08:00"
}

type DetailShortUrlRequest struct {
	Id  uint   `form:"id,optional"`
	Url string `form:"url,optional"`
}

type DetailShortUrlResponse struct {
	Id          uint   `json:"id"`
	Url         string `json:"url"`
	Origin      string `json:"origin"`
	Description string `json:"description"`
	ExpireAt    string `json:"expire_at"`
	CreatedAt   string `json:"created_at"`
}

type UpdateShortUrlRequest struct {
	Id          uint   `json:"id"`
	Description string `json:"description,optional"`
	ExpireAt    string `json:"expire_at,optional"` // ISO 8601 格式 格式: "2023-08-15T14:30:00+08:00"
}

type UpdateShortUrlResponse struct {
}
