package controllers

type AddPointForm struct {
	PathName string  `json:"name"`
	Lat      float64 `json:"lat"`
	Long     float64 `json:"long"`
}

type ResultPoint struct {
	Index int `json:"index"`
	Latitude float64 `json:"lat"`
	Longitude float64 `json:"lng"`
}