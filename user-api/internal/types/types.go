// Code generated by goctl. DO NOT EDIT.
package types

type LoginReq struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type LoginResp struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Time string `json:"time"`
}