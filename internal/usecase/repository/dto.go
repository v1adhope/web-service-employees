package repository

import (
	"github.com/v1adhope/web-service-employees/internal/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type empolyeeDTO struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        string             `bson:"name"`
	Surname     string             `bson:"surname"`
	Phone       string             `bson:"phone"`
	CompanyID   string             `bson:"companyID"`
	Passport    entity.Passport    `bson:"passport"`
	Deportament entity.Deportament `bson:"deportament"`
}

func (dto *empolyeeDTO) FromEntity(emp entity.Employee) {
	dto.Name = emp.Name
	dto.Surname = emp.Surname
	dto.Phone = emp.Phone
	dto.CompanyID = emp.CompanyID
	dto.Passport = emp.Passport
	dto.Deportament = emp.Deportament
}
