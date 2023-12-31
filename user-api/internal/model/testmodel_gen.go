// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"github.com/SpectatorNan/gorm-zero/gormc"

	"gorm.io/gorm"
)

type (
	testModel interface {
		Insert(ctx context.Context, tx *gorm.DB, data *Test) error

		FindOne(ctx context.Context, id int64) (*Test, error)
		Update(ctx context.Context, tx *gorm.DB, data *Test) error

		Delete(ctx context.Context, tx *gorm.DB, id int64) error
		Transaction(ctx context.Context, fn func(db *gorm.DB) error) error
	}

	defaultTestModel struct {
		conn  *gorm.DB
		table string
	}

	Test struct {
		Id   int64          `gorm:"column:id"`
		Name sql.NullString `gorm:"column:name"`
		Time sql.NullInt64  `gorm:"column:time"`
	}
)

func (Test) TableName() string {
	return `"public"."test"`
}

func newTestModel(conn *gorm.DB) *defaultTestModel {
	return &defaultTestModel{
		conn:  conn,
		table: `"public"."test"`,
	}
}

func (m *defaultTestModel) Insert(ctx context.Context, tx *gorm.DB, data *Test) error {
	db := m.conn
	if tx != nil {
		db = tx
	}
	err := db.WithContext(ctx).Save(&data).Error
	return err
}

func (m *defaultTestModel) FindOne(ctx context.Context, id int64) (*Test, error) {
	var resp Test
	err := m.conn.WithContext(ctx).Model(&Test{}).Where("id = ?", id).Take(&resp).Error
	switch err {
	case nil:
		return &resp, nil
	case gormc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTestModel) Update(ctx context.Context, tx *gorm.DB, data *Test) error {
	db := m.conn
	if tx != nil {
		db = tx
	}
	err := db.WithContext(ctx).Save(data).Error
	return err
}

func (m *defaultTestModel) Delete(ctx context.Context, tx *gorm.DB, id int64) error {
	db := m.conn
	if tx != nil {
		db = tx
	}
	err := db.WithContext(ctx).Delete(&Test{}, id).Error

	return err
}

func (m *defaultTestModel) Transaction(ctx context.Context, fn func(db *gorm.DB) error) error {
	return m.conn.WithContext(ctx).Transaction(fn)
}
