package backend

import (
	"github.com/chanzuckerberg/happy/api/pkg/dbutil"
	"github.com/chanzuckerberg/happy/shared/model"
	"github.com/pkg/errors"
	"gorm.io/gorm/clause"
)

type StacklistBackendDB struct {
	DB *dbutil.DB
}

func MakeStacklistBackendDB(db *dbutil.DB) *StacklistBackendDB {
	return &StacklistBackendDB{
		DB: db,
	}
}

func (s *StacklistBackendDB) CreateOrUpdateAppStack(payload model.AppStackPayload) (*model.AppStack, error) {
	db := s.DB.GetDB()
	stack := &model.AppStack{AppStackPayload: payload}
	res := db.Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "app_name"},
			{Name: "environment"},
			{Name: "stack"},
		},
		UpdateAll: true,
	}).Create(&stack)

	return stack, errors.Wrapf(res.Error, "unable to create app stack %s", payload.AppMetadata)
}

func (s *StacklistBackendDB) GetAppStacks(payload model.AppStackPayload) ([]*model.AppStack, error) {
	db := s.DB.GetDB()
	stack := &model.AppStack{AppStackPayload: payload}
	stacks := []*model.AppStack{}
	res := db.Where(stack).Find(&stacks)
	return stacks, errors.Wrapf(res.Error, "unable to get app stacks for %s", stack.AppMetadata)
}

func (s *StacklistBackendDB) DeleteAppStack(payload model.AppStackPayload) (*model.AppStack, error) {
	db := s.DB.GetDB()
	record := &model.AppStack{}
	res := db.Clauses(clause.Returning{}).
		Where("app_name = ? AND environment = ? AND stack = ?",
			payload.AppName,
			payload.Environment,
			payload.Stack,
		).Delete(record)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}
	return record, nil
}
