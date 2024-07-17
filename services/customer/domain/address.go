package domain

type Address struct {
	Street      string `bson:"street" validate:"required"`
	HouseNumber string `bson:"house-number" validate:"required"`
	ZipCode     string `bson:"zip-code" validate:"required"`
}
