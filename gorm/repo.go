package main

import (
	"errors"
	"gorm.io/gorm"
)

// DomainRepository defines methods for interacting with domain blacklists
type DomainRepository interface {
	Create(domain *DomainBlackList) error
	Delete(domain string) error
}

// domainRepository is the implementation of DomainRepository
type domainRepository struct {
	db *gorm.DB
}

func NewDomainRepository(db *gorm.DB) DomainRepository {
	return &domainRepository{db: db}
}

func (r *domainRepository) Create(domain *DomainBlackList) error {
	var existing DomainBlackList
	result := r.db.Where("domain = ?", domain.Domain).First(&existing)
	if result.Error == nil {
		return gorm.ErrDuplicatedKey
	}
	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}

	result = r.db.Create(domain)
	return result.Error
}

func (r *domainRepository) Delete(domain string) error {
	result := r.db.Where("domain = ?", domain).Delete(&DomainBlackList{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
