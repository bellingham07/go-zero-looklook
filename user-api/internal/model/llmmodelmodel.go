package model

import (
	"context"
	"gorm.io/gorm"
)

var _ LlmModelModel = (*customLlmModelModel)(nil)

type (
	// LlmModelModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLlmModelModel.
	LlmModelModel interface {
		llmModelModel
		customLlmModelLogicModel
	}

	customLlmModelModel struct {
		*defaultLlmModelModel
	}

	customLlmModelLogicModel interface {
		UpdateLlm(ctx context.Context) error
	}
)

func (c *customLlmModelModel) UpdateLlm(ctx context.Context) error {
	err := c.conn.WithContext(ctx).Updates(12).Error
	if err != nil {
		return err
	}
	return nil
}

// NewLlmModelModel returns a model for the database table.
func NewLlmModelModel(conn *gorm.DB) LlmModelModel {
	return &customLlmModelModel{
		defaultLlmModelModel: newLlmModelModel(conn),
	}
}

func (m *defaultLlmModelModel) customCacheKeys(data *LlmModel) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
