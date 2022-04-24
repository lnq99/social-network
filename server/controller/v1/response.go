package v1

import (
	"app/internal/model"
	"app/pkg/utils"
)

type Response = utils.Response
type ErrResponse = utils.ErrResponse
type Msg = utils.Msg

var jsonResponse = utils.JsonResponse
var statusResponse = utils.StatusResponse

type loginResponse struct {
	Token string          `json:"token"`
	User  ProfileResponse `json:"user"`
}

type dataResponse struct {
	Data interface{} `json:"data"`
}

type GetMutualAndTypeResponse struct {
	T      string  `json:"type"`
	Mutual []int64 `json:"mutual"`
}

type SearchResponse struct {
	Id     int    `json:"id"`
	Mutual int    `json:"mutual"`
	T      string `json:"type"`
}

type FriendResponse struct {
	Id      int    `json:"id"`
	Name    int    `json:"name"`
	AvatarS string `json:"avatars"`
}

type ProfileResponse struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Gender     string `json:"gender"`
	Birthdate  string `json:"birthdate"`
	Created    string `json:"created"`
	Intro      string `json:"intro"`
	AvatarS    string `json:"avatars"`
	AvatarL    string `json:"avatarl"`
	PostCount  string `json:"postCount"`
	PhotoCount string `json:"photoCount"`
}

func toProfileResponse(profile model.Profile) ProfileResponse {
	return ProfileResponse{
		Id:         profile.Id,
		Name:       profile.Name,
		Gender:     profile.Gender,
		Birthdate:  profile.Birthdate,
		Created:    profile.Created,
		Intro:      profile.Intro,
		AvatarS:    profile.AvatarS,
		AvatarL:    profile.AvatarL,
		PostCount:  profile.PostCount,
		PhotoCount: profile.PhotoCount,
	}
}
