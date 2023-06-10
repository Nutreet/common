package common

type Service string

const (
	User      Service = "user"
	Nutrition Service = "nutrition"
)

type ServicesMap struct {
	User      Service
	Nutrition Service
}

var Services ServicesMap = ServicesMap{
	User:      User,
	Nutrition: Nutrition,
}
