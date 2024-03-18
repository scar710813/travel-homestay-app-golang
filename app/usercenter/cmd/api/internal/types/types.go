// Code generated by goctl. DO NOT EDIT.
package types

type AddWishListReq struct {
	HomestayId int64 `json:"homestayId"`
}

type AddWishListResp struct {
	Homestay Homestay `json:"homestay"`
}

type ClearHistoryReq struct {
}

type ClearHistoryResp struct {
	Success bool `json:"success"`
}

type FileUploadResp struct {
	URL string `json:"url"`
}

type HistoryHomestay struct {
	Id              int64  `json:"id"`
	Title           string `json:"title"`
	Cover           string `json:"cover"`
	Intro           string `json:"intro"`
	Location        string `json:"location"`
	RowState        int64  `json:"rowState"`    //0:下架 1:上架
	PriceBefore     int64  `json:"priceBefore"` //民宿价格
	PriceAfter      int64  `json:"priceAfter"`
	LastBrowingTime int64  `json:"lastBrowingTime"`
}

type HistoryListReq struct {
}

type HistoryListResp struct {
	HistoryList []HistoryHomestay `json:"historyList"`
}

type Homestay struct {
	Id                 int64   `json:"id"`
	Title              string  `json:"title"`
	Cover              string  `json:"cover"`
	Intro              string  `json:"intro"`
	Location           string  `json:"location"`
	HomestayBusinessId int64   `json:"homestayBusinessId"` //店铺id
	UserId             int64   `json:"userId"`             //房东id
	RowState           int64   `json:"rowState"`           //0:下架 1:上架
	RatingStars        float64 `json:"ratingStars"`
	PriceBefore        int64   `json:"priceBefore"` //民宿价格
	PriceAfter         int64   `json:"priceAfter"`
}

type LoginReq struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

type LoginResp struct {
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
	RefreshAfter int64  `json:"refreshAfter"`
}

type RegisterReq struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

type RegisterResp struct {
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
	RefreshAfter int64  `json:"refreshAfter"`
}

type RemoveHistoryReq struct {
	HistoryId int64 `json:"historyId"`
}

type RemoveHistoryResp struct {
	Success bool `json:"success"`
}

type RemoveWishListReq struct {
	Id int64 `json:"id"`
}

type RemoveWishListResp struct {
	Success bool `json:"success"`
}

type UpdateUserInfoReq struct {
	UserInfo User `json:"userInfo"`
}

type UpdateUserInfoResp struct {
	Success bool `json:"success"`
}

type User struct {
	Id          int64   `json:"id"`
	Mobile      string  `json:"mobile"`
	Nickname    string  `json:"nickname"`
	Sex         int64   `json:"sex"`
	Avatar      string  `json:"avatar"`
	Info        string  `json:"info"`
	HomeStayIds []int64 `json:"homestayIds"`
}

type UserInfoReq struct {
}

type UserInfoResp struct {
	UserInfo User `json:"userInfo"`
}

type WXMiniAuthReq struct {
	Code          string `json:"code"`
	IV            string `json:"iv"`
	EncryptedData string `json:"encryptedData"`
}

type WXMiniAuthResp struct {
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
	RefreshAfter int64  `json:"refreshAfter"`
}

type WishListReq struct {
}

type WishListResp struct {
	List []Homestay `json:"list"`
}
