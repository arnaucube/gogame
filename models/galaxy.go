package models

type Galaxy struct {
	SolarSystem []string
}

type SolarSystem struct {
	Id      string
	Planets []string // array with ids of the planets
}
