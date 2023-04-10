package db

import (
	_"context"
	"database/sql"
	_"fmt"
)

type Model interface{
	Querier
}

type SQLModel struct{
	*Queries
	db *sql.DB
}

func NewSQLModel(db *sql.DB)Model{
	return &SQLModel{
		Queries: New(db),
		db:db,
	}
}
// func (model *SQLModel)execTx(ctx context.Context,fn func (q *Queries,txCtx context.Context)error)error{
// 	tx,err:=model.db.BeginTx(ctx,nil)
// 	if err!=nil{
// 		return err
// 	}
// 	q:=New(tx)
// 	err=fn(q,ctx)
// 	if err!=nil{
// 		if rbErr:=tx.Rollback();rbErr!=nil{
// 			return fmt.Errorf("Transaction Err:%s, Rollback Err:%s",err.Error(),rbErr.Error())
// 		}
// 		return err
// 	}
// 	return tx.Commit()
// }
