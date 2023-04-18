package entity

type (
	Employee struct {
		ID          string      `bson:"_id" json:"id"`
		Name        string      `bson:"name" json:"name"`
		Surname     string      `bson:"surname" json:"surname"`
		Phone       string      `bson:"phone" json:"phone"`
		CompanyID   string      `bson:"companyID" json:"companyID"`
		Passport    Passport    `bson:"passport" json:"passport"`
		Deportament Deportament `bson:"deportament" json:"deportament"`
	}

	Passport struct {
		Type   string `bson:"type" json:"type"`
		Number string `bson:"number" json:"number"`
	}

	Deportament struct {
		Name  string `bson:"name" json:"name"`
		Phone string `bson:"phone" json:"phone"`
	}
)
