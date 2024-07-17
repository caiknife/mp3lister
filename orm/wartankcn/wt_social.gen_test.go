// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package wartankcn

import (
	"context"
	"fmt"
	"testing"

	"github.com/caiknife/mp3lister/orm/wartankcn/model"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm/clause"
)

func init() {
	InitializeDB()
	err := _gen_test_db.AutoMigrate(&model.WtSocial{})
	if err != nil {
		fmt.Printf("Error: AutoMigrate(&model.WtSocial{}) fail: %s", err)
	}
}

func Test_wtSocialQuery(t *testing.T) {
	wtSocial := newWtSocial(_gen_test_db)
	wtSocial = *wtSocial.As(wtSocial.TableName())
	_do := wtSocial.WithContext(context.Background()).Debug()

	primaryKey := field.NewString(wtSocial.TableName(), clause.PrimaryKey)
	_, err := _do.Unscoped().Where(primaryKey.IsNotNull()).Delete()
	if err != nil {
		t.Error("clean table <wt_social> fail:", err)
		return
	}

	_, ok := wtSocial.GetFieldByName("")
	if ok {
		t.Error("GetFieldByName(\"\") from wtSocial success")
	}

	err = _do.Create(&model.WtSocial{})
	if err != nil {
		t.Error("create item in table <wt_social> fail:", err)
	}

	err = _do.Save(&model.WtSocial{})
	if err != nil {
		t.Error("create item in table <wt_social> fail:", err)
	}

	err = _do.CreateInBatches([]*model.WtSocial{{}, {}}, 10)
	if err != nil {
		t.Error("create item in table <wt_social> fail:", err)
	}

	_, err = _do.Select(wtSocial.ALL).Take()
	if err != nil {
		t.Error("Take() on table <wt_social> fail:", err)
	}

	_, err = _do.First()
	if err != nil {
		t.Error("First() on table <wt_social> fail:", err)
	}

	_, err = _do.Last()
	if err != nil {
		t.Error("First() on table <wt_social> fail:", err)
	}

	_, err = _do.Where(primaryKey.IsNotNull()).FindInBatch(10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatch() on table <wt_social> fail:", err)
	}

	err = _do.Where(primaryKey.IsNotNull()).FindInBatches(&[]*model.WtSocial{}, 10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatches() on table <wt_social> fail:", err)
	}

	_, err = _do.Select(wtSocial.ALL).Where(primaryKey.IsNotNull()).Order(primaryKey.Desc()).Find()
	if err != nil {
		t.Error("Find() on table <wt_social> fail:", err)
	}

	_, err = _do.Distinct(primaryKey).Take()
	if err != nil {
		t.Error("select Distinct() on table <wt_social> fail:", err)
	}

	_, err = _do.Select(wtSocial.ALL).Omit(primaryKey).Take()
	if err != nil {
		t.Error("Omit() on table <wt_social> fail:", err)
	}

	_, err = _do.Group(primaryKey).Find()
	if err != nil {
		t.Error("Group() on table <wt_social> fail:", err)
	}

	_, err = _do.Scopes(func(dao gen.Dao) gen.Dao { return dao.Where(primaryKey.IsNotNull()) }).Find()
	if err != nil {
		t.Error("Scopes() on table <wt_social> fail:", err)
	}

	_, _, err = _do.FindByPage(0, 1)
	if err != nil {
		t.Error("FindByPage() on table <wt_social> fail:", err)
	}

	_, err = _do.ScanByPage(&model.WtSocial{}, 0, 1)
	if err != nil {
		t.Error("ScanByPage() on table <wt_social> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrInit()
	if err != nil {
		t.Error("FirstOrInit() on table <wt_social> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrCreate()
	if err != nil {
		t.Error("FirstOrCreate() on table <wt_social> fail:", err)
	}

	var _a _another
	var _aPK = field.NewString(_a.TableName(), "id")

	err = _do.Join(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("Join() on table <wt_social> fail:", err)
	}

	err = _do.LeftJoin(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("LeftJoin() on table <wt_social> fail:", err)
	}

	_, err = _do.Not().Or().Clauses().Take()
	if err != nil {
		t.Error("Not/Or/Clauses on table <wt_social> fail:", err)
	}
}
