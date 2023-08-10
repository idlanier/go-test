package repository

import "github.com/idlanier/go-test.git/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
