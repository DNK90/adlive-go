package models

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strings"
)

type Params struct {
	FromField   	int
	FromFieldValue 	int
	Limit 			int
	Fields 			[]string
	OrderBy 		string
	OrderType 		string
	Preloads        map[string][]interface{}
	Joins           map[string][]interface{}
}

func (p *Params) Process(db *gorm.DB) *gorm.DB {
	if p.Fields != nil {
		db = db.Select(strings.Join(p.Fields, ","))
	}
	if p.Preloads != nil {
		for field, conditions := range p.Preloads {
			db = db.Preload(field, conditions...)
		}
	}
	if p.OrderBy != "" {
		desc := true
		if p.OrderBy == "ASC" {
			desc = false
		}
		db = db.Order(clause.OrderByColumn{Column: clause.Column{Name: p.OrderBy}, Desc: desc})
	}
	return db
}
