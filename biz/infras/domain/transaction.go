package domain

import (
	"context"
	"kcers/biz/dal/db/mysql/ent"
)

type TransactionProvider struct {
	db *ent.Client
}

func (p *TransactionProvider) Transactional(fn func() (interface{}, error)) (interface{}, error) {
	tx, err := p.db.Tx(context.Background())
	if err != nil {
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			err := tx.Rollback()
			if err != nil {
				return
			}
		}
	}()
	result, err := fn()
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return nil, err
		}
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return result, err
}
