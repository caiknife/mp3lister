// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package music

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/caiknife/mp3lister/orm/music/model"
)

func newPlayer(db *gorm.DB, opts ...gen.DOOption) player {
	_player := player{}

	_player.playerDo.UseDB(db, opts...)
	_player.playerDo.UseModel(&model.Player{})

	tableName := _player.playerDo.TableName()
	_player.ALL = field.NewAsterisk(tableName)
	_player.ID = field.NewUint64(tableName, "id")
	_player.CreatedAt = field.NewTime(tableName, "created_at")
	_player.UpdatedAt = field.NewTime(tableName, "updated_at")
	_player.DeletedAt = field.NewField(tableName, "deleted_at")
	_player.Name = field.NewString(tableName, "name")
	_player.Phone = field.NewString(tableName, "phone")
	_player.Email = field.NewString(tableName, "email")
	_player.Gold = field.NewInt64(tableName, "gold")
	_player.Extra = field.NewField(tableName, "extra")

	_player.fillFieldMap()

	return _player
}

type player struct {
	playerDo

	ALL       field.Asterisk
	ID        field.Uint64
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field
	Name      field.String // 姓名
	Phone     field.String // 电话
	Email     field.String // 邮件地址
	Gold      field.Int64  // 金币数量
	Extra     field.Field  // 扩展信息

	fieldMap map[string]field.Expr
}

func (p player) Table(newTableName string) *player {
	p.playerDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p player) As(alias string) *player {
	p.playerDo.DO = *(p.playerDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *player) updateTableName(table string) *player {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewUint64(table, "id")
	p.CreatedAt = field.NewTime(table, "created_at")
	p.UpdatedAt = field.NewTime(table, "updated_at")
	p.DeletedAt = field.NewField(table, "deleted_at")
	p.Name = field.NewString(table, "name")
	p.Phone = field.NewString(table, "phone")
	p.Email = field.NewString(table, "email")
	p.Gold = field.NewInt64(table, "gold")
	p.Extra = field.NewField(table, "extra")

	p.fillFieldMap()

	return p
}

func (p *player) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *player) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 9)
	p.fieldMap["id"] = p.ID
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
	p.fieldMap["deleted_at"] = p.DeletedAt
	p.fieldMap["name"] = p.Name
	p.fieldMap["phone"] = p.Phone
	p.fieldMap["email"] = p.Email
	p.fieldMap["gold"] = p.Gold
	p.fieldMap["extra"] = p.Extra
}

func (p player) clone(db *gorm.DB) player {
	p.playerDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p player) replaceDB(db *gorm.DB) player {
	p.playerDo.ReplaceDB(db)
	return p
}

type playerDo struct{ gen.DO }

func (p playerDo) Debug() *playerDo {
	return p.withDO(p.DO.Debug())
}

func (p playerDo) WithContext(ctx context.Context) *playerDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p playerDo) ReadDB() *playerDo {
	return p.Clauses(dbresolver.Read)
}

func (p playerDo) WriteDB() *playerDo {
	return p.Clauses(dbresolver.Write)
}

func (p playerDo) Session(config *gorm.Session) *playerDo {
	return p.withDO(p.DO.Session(config))
}

func (p playerDo) Clauses(conds ...clause.Expression) *playerDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p playerDo) Returning(value interface{}, columns ...string) *playerDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p playerDo) Not(conds ...gen.Condition) *playerDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p playerDo) Or(conds ...gen.Condition) *playerDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p playerDo) Select(conds ...field.Expr) *playerDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p playerDo) Where(conds ...gen.Condition) *playerDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p playerDo) Order(conds ...field.Expr) *playerDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p playerDo) Distinct(cols ...field.Expr) *playerDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p playerDo) Omit(cols ...field.Expr) *playerDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p playerDo) Join(table schema.Tabler, on ...field.Expr) *playerDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p playerDo) LeftJoin(table schema.Tabler, on ...field.Expr) *playerDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p playerDo) RightJoin(table schema.Tabler, on ...field.Expr) *playerDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p playerDo) Group(cols ...field.Expr) *playerDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p playerDo) Having(conds ...gen.Condition) *playerDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p playerDo) Limit(limit int) *playerDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p playerDo) Offset(offset int) *playerDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p playerDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *playerDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p playerDo) Unscoped() *playerDo {
	return p.withDO(p.DO.Unscoped())
}

func (p playerDo) Create(values ...*model.Player) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p playerDo) CreateInBatches(values []*model.Player, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p playerDo) Save(values ...*model.Player) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p playerDo) First() (*model.Player, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Player), nil
	}
}

func (p playerDo) Take() (*model.Player, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Player), nil
	}
}

func (p playerDo) Last() (*model.Player, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Player), nil
	}
}

func (p playerDo) Find() ([]*model.Player, error) {
	result, err := p.DO.Find()
	return result.([]*model.Player), err
}

func (p playerDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Player, err error) {
	buf := make([]*model.Player, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p playerDo) FindInBatches(result *[]*model.Player, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p playerDo) Attrs(attrs ...field.AssignExpr) *playerDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p playerDo) Assign(attrs ...field.AssignExpr) *playerDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p playerDo) Joins(fields ...field.RelationField) *playerDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p playerDo) Preload(fields ...field.RelationField) *playerDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p playerDo) FirstOrInit() (*model.Player, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Player), nil
	}
}

func (p playerDo) FirstOrCreate() (*model.Player, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Player), nil
	}
}

func (p playerDo) FindByPage(offset int, limit int) (result []*model.Player, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p playerDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p playerDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p playerDo) Delete(models ...*model.Player) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *playerDo) withDO(do gen.Dao) *playerDo {
	p.DO = *do.(*gen.DO)
	return p
}
