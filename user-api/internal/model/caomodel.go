package model

import (
	"gorm.io/gorm"
)

var _ CaoModel = (*customCaoModel)(nil)

type (
	// CaoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCaoModel.
	CaoModel interface {
		caoModel
		customCaoLogicModel
	}

	customCaoModel struct {
		*defaultCaoModel
	}

	customCaoLogicModel interface {
	}
)

// NewCaoModel returns a model for the database table.
func NewCaoModel(conn *gorm.DB) CaoModel {
	return &customCaoModel{
		defaultCaoModel: newCaoModel(conn),
	}
}

func (m *defaultCaoModel) customCacheKeys(data *Cao) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
