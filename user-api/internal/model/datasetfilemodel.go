package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ DatasetFileModel = (*customDatasetFileModel)(nil)

type (
	// DatasetFileModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDatasetFileModel.
	DatasetFileModel interface {
		datasetFileModel
	}

	customDatasetFileModel struct {
		*defaultDatasetFileModel
	}
)

// NewDatasetFileModel returns a model for the database table.
func NewDatasetFileModel(conn sqlx.SqlConn) DatasetFileModel {
	return &customDatasetFileModel{
		defaultDatasetFileModel: newDatasetFileModel(conn),
	}
}
