package models

type Resource struct {
	Value int64
	Max   int64
}

type User struct {
	Id        string
	Name      string
	Email     string
	Resources struct {
		Metal     Resource
		Crystal   Resource
		Deuterium Resource
		Energy    Resource
	}
}
