package controllers

type AddPointForm struct {
	PathName string  `json:"name"`
	Lat      float32 `json:"lat"`
	Long     float32 `json:"long"`
}
