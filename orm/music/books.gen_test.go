// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package music

import (
	"context"
	"fmt"
	"testing"

	"github.com/caiknife/mp3lister/orm/music/model"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm/clause"
)

func init() {
	InitializeDB()
	err := db.AutoMigrate(&model.Book{})
	if err != nil {
		fmt.Printf("Error: AutoMigrate(&model.Book{}) fail: %s", err)
	}
}

func Test_bookQuery(t *testing.T) {
	book := newBook(db)
	book = *book.As(book.TableName())
	_do := book.WithContext(context.Background()).Debug()

	primaryKey := field.NewString(book.TableName(), clause.PrimaryKey)
	_, err := _do.Unscoped().Where(primaryKey.IsNotNull()).Delete()
	if err != nil {
		t.Error("clean table <books> fail:", err)
		return
	}

	_, ok := book.GetFieldByName("")
	if ok {
		t.Error("GetFieldByName(\"\") from book success")
	}

	err = _do.Create(&model.Book{})
	if err != nil {
		t.Error("create item in table <books> fail:", err)
	}

	err = _do.Save(&model.Book{})
	if err != nil {
		t.Error("create item in table <books> fail:", err)
	}

	err = _do.CreateInBatches([]*model.Book{{}, {}}, 10)
	if err != nil {
		t.Error("create item in table <books> fail:", err)
	}

	_, err = _do.Select(book.ALL).Take()
	if err != nil {
		t.Error("Take() on table <books> fail:", err)
	}

	_, err = _do.First()
	if err != nil {
		t.Error("First() on table <books> fail:", err)
	}

	_, err = _do.Last()
	if err != nil {
		t.Error("First() on table <books> fail:", err)
	}

	_, err = _do.Where(primaryKey.IsNotNull()).FindInBatch(10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatch() on table <books> fail:", err)
	}

	err = _do.Where(primaryKey.IsNotNull()).FindInBatches(&[]*model.Book{}, 10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatches() on table <books> fail:", err)
	}

	_, err = _do.Select(book.ALL).Where(primaryKey.IsNotNull()).Order(primaryKey.Desc()).Find()
	if err != nil {
		t.Error("Find() on table <books> fail:", err)
	}

	_, err = _do.Distinct(primaryKey).Take()
	if err != nil {
		t.Error("select Distinct() on table <books> fail:", err)
	}

	_, err = _do.Select(book.ALL).Omit(primaryKey).Take()
	if err != nil {
		t.Error("Omit() on table <books> fail:", err)
	}

	_, err = _do.Group(primaryKey).Find()
	if err != nil {
		t.Error("Group() on table <books> fail:", err)
	}

	_, err = _do.Scopes(func(dao gen.Dao) gen.Dao { return dao.Where(primaryKey.IsNotNull()) }).Find()
	if err != nil {
		t.Error("Scopes() on table <books> fail:", err)
	}

	_, _, err = _do.FindByPage(0, 1)
	if err != nil {
		t.Error("FindByPage() on table <books> fail:", err)
	}

	_, err = _do.ScanByPage(&model.Book{}, 0, 1)
	if err != nil {
		t.Error("ScanByPage() on table <books> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrInit()
	if err != nil {
		t.Error("FirstOrInit() on table <books> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrCreate()
	if err != nil {
		t.Error("FirstOrCreate() on table <books> fail:", err)
	}

	var _a _another
	var _aPK = field.NewString(_a.TableName(), "id")

	err = _do.Join(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("Join() on table <books> fail:", err)
	}

	err = _do.LeftJoin(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("LeftJoin() on table <books> fail:", err)
	}

	_, err = _do.Not().Or().Clauses().Take()
	if err != nil {
		t.Error("Not/Or/Clauses on table <books> fail:", err)
	}
}
