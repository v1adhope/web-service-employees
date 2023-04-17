// TODO: fields MAY BE NIL
package repository

import (
	"context"
	"fmt"

	"github.com/v1adhope/web-service-employees/internal/entity"
	"github.com/v1adhope/web-service-employees/pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	_colEmployee = "employee"
)

type Repo struct {
	*mongodb.MongoCol
}

func New(col *mongodb.MongoCol) *Repo {
	return &Repo{col}
}

func (r *Repo) Insert(ctx context.Context, emp entity.Employee) (string, error) {
	var dto empolyeeDTO

	dto.FromEntity(emp)

	dto.ID = primitive.NewObjectID()

	res, err := r.Col.InsertOne(ctx, dto)
	if err != nil {
		return "", fmt.Errorf("repository: Insert: InsertOne: %w", err)
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (r *Repo) DeleteByID(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)

	filter := bson.D{{"_id", objID}}

	res, err := r.Col.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("repository: DeleteByID: DeleteOne: %w", err)
	}
	if res.DeletedCount == 0 {
		return fmt.Errorf("repository: DeleteByID: DeleteOne: %w", mongo.ErrNoDocuments)
	}

	return nil
}

func (r *Repo) GetByCompanyID(ctx context.Context, companyID int) ([]entity.Employee, error) {
	filter := bson.D{{"companyID", companyID}}

	cursor, err := r.Col.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("repository: GetByCompany: Find: %w", err)
	}
	defer cursor.Close(ctx)

	if cursor.RemainingBatchLength() == 0 {
		return nil, mongo.ErrNoDocuments
	}

	var res []entity.Employee

	if err := cursor.All(ctx, &res); err != nil {
		return nil, fmt.Errorf("repository: GetByCompany: All: %w", err)
	}

	return res, nil
}

func (r *Repo) GetByDeportamentName(ctx context.Context, deportmentName string) ([]entity.Employee, error) {
	filter := bson.D{{"deportament.name", deportmentName}}

	cursor, err := r.Col.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("repository: GetByDepartament: Find: %w", err)
	}
	defer cursor.Close(ctx)

	if cursor.RemainingBatchLength() == 0 {
		return nil, mongo.ErrNoDocuments
	}

	var res []entity.Employee

	if err := cursor.All(ctx, &res); err != nil {
		return nil, fmt.Errorf("repository: GetByDepartament: All: %w", err)
	}

	return res, nil
}

func (r *Repo) UpdateByID(ctx context.Context, emp entity.Employee) error {
	var dto empolyeeDTO

	dto.FromEntity(emp)

	var field []bson.E

	if dto.Name != "" {
		field = append(field, bson.E{"name", dto.Name})
	}

	if dto.Surname != "" {
		field = append(field, bson.E{"surname", dto.Surname})
	}

	if dto.Phone != "" {
		field = append(field, bson.E{"phone", dto.Phone})
	}

	//NOTE: I think ID 0 not exists
	if dto.CompanyID != 0 {
		field = append(field, bson.E{"companyID", dto.CompanyID})
	}

	if dto.Passport.Type != "" {
		field = append(field, bson.E{"passport.type", dto.Passport.Type})
	}

	if dto.Passport.Number != "" {
		field = append(field, bson.E{"passport.number", dto.Passport.Number})
	}

	if dto.Deportament.Name != "" {
		field = append(field, bson.E{"deportament.name", dto.Deportament.Name})
	}

	if dto.Deportament.Phone != "" {
		field = append(field, bson.E{"deportament.phone", dto.Deportament.Phone})
	}

	filter := bson.D{{"$set", field}}

	var err error

	dto.ID, err = primitive.ObjectIDFromHex(emp.ID)

	res, err := r.Col.UpdateByID(ctx, dto.ID, filter)
	if err != nil {
		return fmt.Errorf("repository: UpdateByID: UpdateByID: %w", err)
	}
	if res.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}
