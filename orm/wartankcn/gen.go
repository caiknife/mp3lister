// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package wartankcn

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q            = new(Query)
	WtDevice     *wtDevice
	WtGamecenter *wtGamecenter
	WtGift       *wtGift
	WtLegion     *wtLegion
	WtOrder      *wtOrder
	WtPlayer     *wtPlayer
	WtSocial     *wtSocial
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	WtDevice = &Q.WtDevice
	WtGamecenter = &Q.WtGamecenter
	WtGift = &Q.WtGift
	WtLegion = &Q.WtLegion
	WtOrder = &Q.WtOrder
	WtPlayer = &Q.WtPlayer
	WtSocial = &Q.WtSocial
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:           db,
		WtDevice:     newWtDevice(db, opts...),
		WtGamecenter: newWtGamecenter(db, opts...),
		WtGift:       newWtGift(db, opts...),
		WtLegion:     newWtLegion(db, opts...),
		WtOrder:      newWtOrder(db, opts...),
		WtPlayer:     newWtPlayer(db, opts...),
		WtSocial:     newWtSocial(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	WtDevice     wtDevice
	WtGamecenter wtGamecenter
	WtGift       wtGift
	WtLegion     wtLegion
	WtOrder      wtOrder
	WtPlayer     wtPlayer
	WtSocial     wtSocial
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:           db,
		WtDevice:     q.WtDevice.clone(db),
		WtGamecenter: q.WtGamecenter.clone(db),
		WtGift:       q.WtGift.clone(db),
		WtLegion:     q.WtLegion.clone(db),
		WtOrder:      q.WtOrder.clone(db),
		WtPlayer:     q.WtPlayer.clone(db),
		WtSocial:     q.WtSocial.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:           db,
		WtDevice:     q.WtDevice.replaceDB(db),
		WtGamecenter: q.WtGamecenter.replaceDB(db),
		WtGift:       q.WtGift.replaceDB(db),
		WtLegion:     q.WtLegion.replaceDB(db),
		WtOrder:      q.WtOrder.replaceDB(db),
		WtPlayer:     q.WtPlayer.replaceDB(db),
		WtSocial:     q.WtSocial.replaceDB(db),
	}
}

type queryCtx struct {
	WtDevice     *wtDeviceDo
	WtGamecenter *wtGamecenterDo
	WtGift       *wtGiftDo
	WtLegion     *wtLegionDo
	WtOrder      *wtOrderDo
	WtPlayer     *wtPlayerDo
	WtSocial     *wtSocialDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		WtDevice:     q.WtDevice.WithContext(ctx),
		WtGamecenter: q.WtGamecenter.WithContext(ctx),
		WtGift:       q.WtGift.WithContext(ctx),
		WtLegion:     q.WtLegion.WithContext(ctx),
		WtOrder:      q.WtOrder.WithContext(ctx),
		WtPlayer:     q.WtPlayer.WithContext(ctx),
		WtSocial:     q.WtSocial.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}