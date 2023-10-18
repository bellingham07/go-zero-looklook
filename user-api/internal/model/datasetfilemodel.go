package model

import (
	"gorm.io/gorm"
)

var _ DatasetFileModel = (*customDatasetFileModel)(nil)

type (
	// DatasetFileModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDatasetFileModel.
	DatasetFileModel interface {
		datasetFileModel
		customDatasetFileLogicModel
	}

	customDatasetFileModel struct {
		*defaultDatasetFileModel
	}

	customDatasetFileLogicModel interface {
	}
)

// NewDatasetFileModel returns a model for the database table.
func NewDatasetFileModel(conn *gorm.DB) DatasetFileModel {
	return &customDatasetFileModel{
		defaultDatasetFileModel: newDatasetFileModel(conn),
	}
}

func (m *defaultDatasetFileModel) customCacheKeys(data *DatasetFile) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
