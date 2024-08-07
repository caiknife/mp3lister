// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package wartankcn

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/caiknife/mp3lister/orm/wartankcn/model"
)

func newWtDevice(db *gorm.DB, opts ...gen.DOOption) wtDevice {
	_wtDevice := wtDevice{}

	_wtDevice.wtDeviceDo.UseDB(db, opts...)
	_wtDevice.wtDeviceDo.UseModel(&model.WtDevice{})

	tableName := _wtDevice.wtDeviceDo.TableName()
	_wtDevice.ALL = field.NewAsterisk(tableName)
	_wtDevice.ID = field.NewString(tableName, "id")
	_wtDevice.PlayerID = field.NewInt64(tableName, "player_id")
	_wtDevice.DeviceRegion = field.NewString(tableName, "device_region")
	_wtDevice.CreateTime = field.NewTime(tableName, "create_time")
	_wtDevice.UpdateTime = field.NewTime(tableName, "update_time")

	_wtDevice.fillFieldMap()

	return _wtDevice
}

// wtDevice 玩家设备关系表
type wtDevice struct {
	wtDeviceDo

	ALL          field.Asterisk
	ID           field.String // 设备ID
	PlayerID     field.Int64  // 玩家ID
	DeviceRegion field.String // 设备地区
	CreateTime   field.Time   // 创建时间
	UpdateTime   field.Time   // 更新时间

	fieldMap map[string]field.Expr
}

func (w wtDevice) Table(newTableName string) *wtDevice {
	w.wtDeviceDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w wtDevice) As(alias string) *wtDevice {
	w.wtDeviceDo.DO = *(w.wtDeviceDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *wtDevice) updateTableName(table string) *wtDevice {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewString(table, "id")
	w.PlayerID = field.NewInt64(table, "player_id")
	w.DeviceRegion = field.NewString(table, "device_region")
	w.CreateTime = field.NewTime(table, "create_time")
	w.UpdateTime = field.NewTime(table, "update_time")

	w.fillFieldMap()

	return w
}

func (w *wtDevice) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *wtDevice) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 5)
	w.fieldMap["id"] = w.ID
	w.fieldMap["player_id"] = w.PlayerID
	w.fieldMap["device_region"] = w.DeviceRegion
	w.fieldMap["create_time"] = w.CreateTime
	w.fieldMap["update_time"] = w.UpdateTime
}

func (w wtDevice) clone(db *gorm.DB) wtDevice {
	w.wtDeviceDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w wtDevice) replaceDB(db *gorm.DB) wtDevice {
	w.wtDeviceDo.ReplaceDB(db)
	return w
}

type wtDeviceDo struct{ gen.DO }

func (w wtDeviceDo) Debug() *wtDeviceDo {
	return w.withDO(w.DO.Debug())
}

func (w wtDeviceDo) WithContext(ctx context.Context) *wtDeviceDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w wtDeviceDo) ReadDB() *wtDeviceDo {
	return w.Clauses(dbresolver.Read)
}

func (w wtDeviceDo) WriteDB() *wtDeviceDo {
	return w.Clauses(dbresolver.Write)
}

func (w wtDeviceDo) Session(config *gorm.Session) *wtDeviceDo {
	return w.withDO(w.DO.Session(config))
}

func (w wtDeviceDo) Clauses(conds ...clause.Expression) *wtDeviceDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w wtDeviceDo) Returning(value interface{}, columns ...string) *wtDeviceDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w wtDeviceDo) Not(conds ...gen.Condition) *wtDeviceDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w wtDeviceDo) Or(conds ...gen.Condition) *wtDeviceDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w wtDeviceDo) Select(conds ...field.Expr) *wtDeviceDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w wtDeviceDo) Where(conds ...gen.Condition) *wtDeviceDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w wtDeviceDo) Order(conds ...field.Expr) *wtDeviceDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w wtDeviceDo) Distinct(cols ...field.Expr) *wtDeviceDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w wtDeviceDo) Omit(cols ...field.Expr) *wtDeviceDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w wtDeviceDo) Join(table schema.Tabler, on ...field.Expr) *wtDeviceDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w wtDeviceDo) LeftJoin(table schema.Tabler, on ...field.Expr) *wtDeviceDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w wtDeviceDo) RightJoin(table schema.Tabler, on ...field.Expr) *wtDeviceDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w wtDeviceDo) Group(cols ...field.Expr) *wtDeviceDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w wtDeviceDo) Having(conds ...gen.Condition) *wtDeviceDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w wtDeviceDo) Limit(limit int) *wtDeviceDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w wtDeviceDo) Offset(offset int) *wtDeviceDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w wtDeviceDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *wtDeviceDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w wtDeviceDo) Unscoped() *wtDeviceDo {
	return w.withDO(w.DO.Unscoped())
}

func (w wtDeviceDo) Create(values ...*model.WtDevice) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w wtDeviceDo) CreateInBatches(values []*model.WtDevice, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w wtDeviceDo) Save(values ...*model.WtDevice) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w wtDeviceDo) First() (*model.WtDevice, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WtDevice), nil
	}
}

func (w wtDeviceDo) Take() (*model.WtDevice, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WtDevice), nil
	}
}

func (w wtDeviceDo) Last() (*model.WtDevice, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WtDevice), nil
	}
}

func (w wtDeviceDo) Find() ([]*model.WtDevice, error) {
	result, err := w.DO.Find()
	return result.([]*model.WtDevice), err
}

func (w wtDeviceDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WtDevice, err error) {
	buf := make([]*model.WtDevice, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w wtDeviceDo) FindInBatches(result *[]*model.WtDevice, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w wtDeviceDo) Attrs(attrs ...field.AssignExpr) *wtDeviceDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w wtDeviceDo) Assign(attrs ...field.AssignExpr) *wtDeviceDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w wtDeviceDo) Joins(fields ...field.RelationField) *wtDeviceDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w wtDeviceDo) Preload(fields ...field.RelationField) *wtDeviceDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w wtDeviceDo) FirstOrInit() (*model.WtDevice, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WtDevice), nil
	}
}

func (w wtDeviceDo) FirstOrCreate() (*model.WtDevice, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WtDevice), nil
	}
}

func (w wtDeviceDo) FindByPage(offset int, limit int) (result []*model.WtDevice, count int64, err error) {
	result, err = w.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = w.Offset(-1).Limit(-1).Count()
	return
}

func (w wtDeviceDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w wtDeviceDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w wtDeviceDo) Delete(models ...*model.WtDevice) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *wtDeviceDo) withDO(do gen.Dao) *wtDeviceDo {
	w.DO = *do.(*gen.DO)
	return w
}
