package model

import (
	"gorm.io/gorm"
)

var _ TestModel = (*customTestModel)(nil)

type (
	// TestModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTestModel.
	TestModel interface {
		testModel
		customTestLogicModel
	}

	customTestModel struct {
		*defaultTestModel
	}

	customTestLogicModel interface {
	}
)

// NewTestModel returns a model for the database table.
func NewTestModel(conn *gorm.DB) TestModel {
	return &customTestModel{
		defaultTestModel: newTestModel(conn),
	}
}

func (m *defaultTestModel) customCacheKeys(data *Test) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
