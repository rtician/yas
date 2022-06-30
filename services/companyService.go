package services

import (
	"context"

	"yas/database/models"
	"yas/repositories"

	"github.com/rs/zerolog/log"
)

type CompanyService interface {
	CreateCompany(ctx context.Context, company *models.Company) (*models.Company, error)
}

type defaultCompanyService struct {
	companyRepository repositories.CompanyRepository
}

func NewCompanyService(repo repositories.CompanyRepository) CompanyService {
	return &defaultCompanyService{
		companyRepository: repo,
	}
}

func (d *defaultCompanyService) CreateCompany(ctx context.Context, company *models.Company) (*models.Company, error) {
	c, err := d.companyRepository.CreateCompany(ctx, company)
	if err != nil {
		log.Error().Err(err).Msg("error creating company")
		return nil, err
	}
	return c, nil
}
