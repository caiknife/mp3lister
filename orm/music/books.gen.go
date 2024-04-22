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

func newBook(db *gorm.DB, opts ...gen.DOOption) book {
	_book := book{}

	_book.bookDo.UseDB(db, opts...)
	_book.bookDo.UseModel(&model.Book{})

	tableName := _book.bookDo.TableName()
	_book.ALL = field.NewAsterisk(tableName)
	_book.ID = field.NewUint64(tableName, "id")
	_book.CreatedAt = field.NewTime(tableName, "created_at")
	_book.UpdatedAt = field.NewTime(tableName, "updated_at")
	_book.DeletedAt = field.NewField(tableName, "deleted_at")
	_book.Title = field.NewString(tableName, "title")
	_book.Author = field.NewString(tableName, "author")
	_book.Genre = field.NewString(tableName, "genre")

	_book.fillFieldMap()

	return _book
}

type book struct {
	bookDo

	ALL       field.Asterisk
	ID        field.Uint64
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field
	Title     field.String // 书籍名称
	Author    field.String // 作者
	Genre     field.String // 分类

	fieldMap map[string]field.Expr
}

func (b book) Table(newTableName string) *book {
	b.bookDo.UseTable(newTableName)
	return b.updateTableName(newTableName)
}

func (b book) As(alias string) *book {
	b.bookDo.DO = *(b.bookDo.As(alias).(*gen.DO))
	return b.updateTableName(alias)
}

func (b *book) updateTableName(table string) *book {
	b.ALL = field.NewAsterisk(table)
	b.ID = field.NewUint64(table, "id")
	b.CreatedAt = field.NewTime(table, "created_at")
	b.UpdatedAt = field.NewTime(table, "updated_at")
	b.DeletedAt = field.NewField(table, "deleted_at")
	b.Title = field.NewString(table, "title")
	b.Author = field.NewString(table, "author")
	b.Genre = field.NewString(table, "genre")

	b.fillFieldMap()

	return b
}

func (b *book) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := b.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (b *book) fillFieldMap() {
	b.fieldMap = make(map[string]field.Expr, 7)
	b.fieldMap["id"] = b.ID
	b.fieldMap["created_at"] = b.CreatedAt
	b.fieldMap["updated_at"] = b.UpdatedAt
	b.fieldMap["deleted_at"] = b.DeletedAt
	b.fieldMap["title"] = b.Title
	b.fieldMap["author"] = b.Author
	b.fieldMap["genre"] = b.Genre
}

func (b book) clone(db *gorm.DB) book {
	b.bookDo.ReplaceConnPool(db.Statement.ConnPool)
	return b
}

func (b book) replaceDB(db *gorm.DB) book {
	b.bookDo.ReplaceDB(db)
	return b
}

type bookDo struct{ gen.DO }

func (b bookDo) Debug() *bookDo {
	return b.withDO(b.DO.Debug())
}

func (b bookDo) WithContext(ctx context.Context) *bookDo {
	return b.withDO(b.DO.WithContext(ctx))
}

func (b bookDo) ReadDB() *bookDo {
	return b.Clauses(dbresolver.Read)
}

func (b bookDo) WriteDB() *bookDo {
	return b.Clauses(dbresolver.Write)
}

func (b bookDo) Session(config *gorm.Session) *bookDo {
	return b.withDO(b.DO.Session(config))
}

func (b bookDo) Clauses(conds ...clause.Expression) *bookDo {
	return b.withDO(b.DO.Clauses(conds...))
}

func (b bookDo) Returning(value interface{}, columns ...string) *bookDo {
	return b.withDO(b.DO.Returning(value, columns...))
}

func (b bookDo) Not(conds ...gen.Condition) *bookDo {
	return b.withDO(b.DO.Not(conds...))
}

func (b bookDo) Or(conds ...gen.Condition) *bookDo {
	return b.withDO(b.DO.Or(conds...))
}

func (b bookDo) Select(conds ...field.Expr) *bookDo {
	return b.withDO(b.DO.Select(conds...))
}

func (b bookDo) Where(conds ...gen.Condition) *bookDo {
	return b.withDO(b.DO.Where(conds...))
}

func (b bookDo) Order(conds ...field.Expr) *bookDo {
	return b.withDO(b.DO.Order(conds...))
}

func (b bookDo) Distinct(cols ...field.Expr) *bookDo {
	return b.withDO(b.DO.Distinct(cols...))
}

func (b bookDo) Omit(cols ...field.Expr) *bookDo {
	return b.withDO(b.DO.Omit(cols...))
}

func (b bookDo) Join(table schema.Tabler, on ...field.Expr) *bookDo {
	return b.withDO(b.DO.Join(table, on...))
}

func (b bookDo) LeftJoin(table schema.Tabler, on ...field.Expr) *bookDo {
	return b.withDO(b.DO.LeftJoin(table, on...))
}

func (b bookDo) RightJoin(table schema.Tabler, on ...field.Expr) *bookDo {
	return b.withDO(b.DO.RightJoin(table, on...))
}

func (b bookDo) Group(cols ...field.Expr) *bookDo {
	return b.withDO(b.DO.Group(cols...))
}

func (b bookDo) Having(conds ...gen.Condition) *bookDo {
	return b.withDO(b.DO.Having(conds...))
}

func (b bookDo) Limit(limit int) *bookDo {
	return b.withDO(b.DO.Limit(limit))
}

func (b bookDo) Offset(offset int) *bookDo {
	return b.withDO(b.DO.Offset(offset))
}

func (b bookDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *bookDo {
	return b.withDO(b.DO.Scopes(funcs...))
}

func (b bookDo) Unscoped() *bookDo {
	return b.withDO(b.DO.Unscoped())
}

func (b bookDo) Create(values ...*model.Book) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Create(values)
}

func (b bookDo) CreateInBatches(values []*model.Book, batchSize int) error {
	return b.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (b bookDo) Save(values ...*model.Book) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Save(values)
}

func (b bookDo) First() (*model.Book, error) {
	if result, err := b.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Book), nil
	}
}

func (b bookDo) Take() (*model.Book, error) {
	if result, err := b.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Book), nil
	}
}

func (b bookDo) Last() (*model.Book, error) {
	if result, err := b.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Book), nil
	}
}

func (b bookDo) Find() ([]*model.Book, error) {
	result, err := b.DO.Find()
	return result.([]*model.Book), err
}

func (b bookDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Book, err error) {
	buf := make([]*model.Book, 0, batchSize)
	err = b.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (b bookDo) FindInBatches(result *[]*model.Book, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return b.DO.FindInBatches(result, batchSize, fc)
}

func (b bookDo) Attrs(attrs ...field.AssignExpr) *bookDo {
	return b.withDO(b.DO.Attrs(attrs...))
}

func (b bookDo) Assign(attrs ...field.AssignExpr) *bookDo {
	return b.withDO(b.DO.Assign(attrs...))
}

func (b bookDo) Joins(fields ...field.RelationField) *bookDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Joins(_f))
	}
	return &b
}

func (b bookDo) Preload(fields ...field.RelationField) *bookDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Preload(_f))
	}
	return &b
}

func (b bookDo) FirstOrInit() (*model.Book, error) {
	if result, err := b.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Book), nil
	}
}

func (b bookDo) FirstOrCreate() (*model.Book, error) {
	if result, err := b.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Book), nil
	}
}

func (b bookDo) FindByPage(offset int, limit int) (result []*model.Book, count int64, err error) {
	result, err = b.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = b.Offset(-1).Limit(-1).Count()
	return
}

func (b bookDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = b.Count()
	if err != nil {
		return
	}

	err = b.Offset(offset).Limit(limit).Scan(result)
	return
}

func (b bookDo) Scan(result interface{}) (err error) {
	return b.DO.Scan(result)
}

func (b bookDo) Delete(models ...*model.Book) (result gen.ResultInfo, err error) {
	return b.DO.Delete(models)
}

func (b *bookDo) withDO(do gen.Dao) *bookDo {
	b.DO = *do.(*gen.DO)
	return b
}
