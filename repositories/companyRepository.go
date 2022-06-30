package repositories

import (
	"context"
	"yas/database/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CompanyRepository interface {
	CreateCompany(ctx context.Context, company *models.Company) (*models.Company, error)
}

type companyRepository struct {
	Database   *mongo.Database
	Collection string
}

func NewCompanyRepository(d *mongo.Database) CompanyRepository {
	return &companyRepository{
		Database:   d,
		Collection: "company",
	}
}

func (r *companyRepository) CreateCompany(ctx context.Context, company *models.Company) (*models.Company, error) {
	company.Id = primitive.NewObjectID()
	collection := r.Database.Collection(r.Collection)
	_, err := collection.InsertOne(ctx, company)
	if err != nil {
		return nil, err
	}
	return company, nil
}
