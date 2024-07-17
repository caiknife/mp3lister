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

func newWtPlayer(db *gorm.DB, opts ...gen.DOOption) wtPlayer {
	_wtPlayer := wtPlayer{}

	_wtPlayer.wtPlayerDo.UseDB(db, opts...)
	_wtPlayer.wtPlayerDo.UseModel(&model.WtPlayer{})

	tableName := _wtPlayer.wtPlayerDo.TableName()
	_wtPlayer.ALL = field.NewAsterisk(tableName)
	_wtPlayer.ID = field.NewUint64(tableName, "id")
	_wtPlayer.VisitorID = field.NewString(tableName, "visitor_id")
	_wtPlayer.Tag = field.NewString(tableName, "tag")
	_wtPlayer.PlayerLvl = field.NewInt32(tableName, "player_lvl")
	_wtPlayer.PlayerExp = field.NewInt32(tableName, "player_exp")
	_wtPlayer.Nickname = field.NewString(tableName, "nickname")
	_wtPlayer.Icons = field.NewField(tableName, "icons")
	_wtPlayer.GoldPool = field.NewInt32(tableName, "gold_pool")
	_wtPlayer.GoldPoolTs = field.NewInt32(tableName, "gold_pool_ts")
	_wtPlayer.Diamond = field.NewInt32(tableName, "diamond")
	_wtPlayer.Gold = field.NewInt64(tableName, "gold")
	_wtPlayer.LegionID = field.NewInt32(tableName, "legion_id")
	_wtPlayer.LegionName = field.NewString(tableName, "legion_name")
	_wtPlayer.LegionPosition = field.NewInt32(tableName, "legion_position")
	_wtPlayer.Tier = field.NewInt32(tableName, "tier")
	_wtPlayer.Trophy = field.NewInt32(tableName, "trophy")
	_wtPlayer.TrophyRoad = field.NewField(tableName, "trophy_road")
	_wtPlayer.Vip = field.NewField(tableName, "vip")
	_wtPlayer.ChestInfo = field.NewField(tableName, "chest_info")
	_wtPlayer.Garage = field.NewField(tableName, "garage")
	_wtPlayer.Inventory = field.NewField(tableName, "inventory")
	_wtPlayer.PathOfValor = field.NewField(tableName, "path_of_valor")
	_wtPlayer.IPRegion = field.NewString(tableName, "ip_region")
	_wtPlayer.DeviceRegion = field.NewString(tableName, "device_region")
	_wtPlayer.RenameTimes = field.NewInt32(tableName, "rename_times")
	_wtPlayer.SettlementTrophy = field.NewInt32(tableName, "settlement_trophy")
	_wtPlayer.StatisticsInfo = field.NewField(tableName, "statistics_info")
	_wtPlayer.GuideInfo = field.NewField(tableName, "guide_info")
	_wtPlayer.Status = field.NewInt32(tableName, "status")
	_wtPlayer.CreateTime = field.NewTime(tableName, "create_time")
	_wtPlayer.UpdateTime = field.NewTime(tableName, "update_time")
	_wtPlayer.LastLoginTime = field.NewTime(tableName, "last_login_time")
	_wtPlayer.Version = field.NewInt64(tableName, "version")
	_wtPlayer.JoinWar = field.NewInt32(tableName, "join_war")
	_wtPlayer.TankTeam = field.NewField(tableName, "tank_team")
	_wtPlayer.CompetitiveRank = field.NewField(tableName, "competitive_rank")

	_wtPlayer.fillFieldMap()

	return _wtPlayer
}

// wtPlayer 玩家表
type wtPlayer struct {
	wtPlayerDo

	ALL              field.Asterisk
	ID               field.Uint64 // 玩家ID
	VisitorID        field.String // 访问ID
	Tag              field.String // 8位简码
	PlayerLvl        field.Int32  // 玩家等级
	PlayerExp        field.Int32  // 玩家经验值
	Nickname         field.String // 玩家名称
	Icons            field.Field  // 头像
	GoldPool         field.Int32  // 金币池
	GoldPoolTs       field.Int32  // 金币池上次更新的秒数
	Diamond          field.Int32  // 钻石
	Gold             field.Int64  // 金币
	LegionID         field.Int32  // 军团ID
	LegionName       field.String // 军团名
	LegionPosition   field.Int32  // 军团职位
	Tier             field.Int32  // 战场级别
	Trophy           field.Int32  // 奖杯数
	TrophyRoad       field.Field  // 荣耀之路
	Vip              field.Field
	ChestInfo        field.Field  // 宝箱信息
	Garage           field.Field  // 车库
	Inventory        field.Field  // 库存信息
	PathOfValor      field.Field  // 英勇之路
	IPRegion         field.String // IP地区
	DeviceRegion     field.String // 设备地区
	RenameTimes      field.Int32  // 改名次数
	SettlementTrophy field.Int32  // 结算奖杯数
	StatisticsInfo   field.Field  // 统计信息
	GuideInfo        field.Field  // 新手信息
	Status           field.Int32  // 账号状态
	CreateTime       field.Time   // 创建时间
	UpdateTime       field.Time   // 更新时间
	LastLoginTime    field.Time   // 最后登录时间
	Version          field.Int64  // 锁版本
	JoinWar          field.Int32  // 是否加入军团战
	TankTeam         field.Field  // 军团战坦克编组
	CompetitiveRank  field.Field  // 竞技模式排行数据

	fieldMap map[string]field.Expr
}

func (w wtPlayer) Table(newTableName string) *wtPlayer {
	w.wtPlayerDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w wtPlayer) As(alias string) *wtPlayer {
	w.wtPlayerDo.DO = *(w.wtPlayerDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *wtPlayer) updateTableName(table string) *wtPlayer {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewUint64(table, "id")
	w.VisitorID = field.NewString(table, "visitor_id")
	w.Tag = field.NewString(table, "tag")
	w.PlayerLvl = field.NewInt32(table, "player_lvl")
	w.PlayerExp = field.NewInt32(table, "player_exp")
	w.Nickname = field.NewString(table, "nickname")
	w.Icons = field.NewField(table, "icons")
	w.GoldPool = field.NewInt32(table, "gold_pool")
	w.GoldPoolTs = field.NewInt32(table, "gold_pool_ts")
	w.Diamond = field.NewInt32(table, "diamond")
	w.Gold = field.NewInt64(table, "gold")
	w.LegionID = field.NewInt32(table, "legion_id")
	w.LegionName = field.NewString(table, "legion_name")
	w.LegionPosition = field.NewInt32(table, "legion_position")
	w.Tier = field.NewInt32(table, "tier")
	w.Trophy = field.NewInt32(table, "trophy")
	w.TrophyRoad = field.NewField(table, "trophy_road")
	w.Vip = field.NewField(table, "vip")
	w.ChestInfo = field.NewField(table, "chest_info")
	w.Garage = field.NewField(table, "garage")
	w.Inventory = field.NewField(table, "inventory")
	w.PathOfValor = field.NewField(table, "path_of_valor")
	w.IPRegion = field.NewString(table, "ip_region")
	w.DeviceRegion = field.NewString(table, "device_region")
	w.RenameTimes = field.NewInt32(table, "rename_times")
	w.SettlementTrophy = field.NewInt32(table, "settlement_trophy")
	w.StatisticsInfo = field.NewField(table, "statistics_info")
	w.GuideInfo = field.NewField(table, "guide_info")
	w.Status = field.NewInt32(table, "status")
	w.CreateTime = field.NewTime(table, "create_time")
	w.UpdateTime = field.NewTime(table, "update_time")
	w.LastLoginTime = field.NewTime(table, "last_login_time")
	w.Version = field.NewInt64(table, "version")
	w.JoinWar = field.NewInt32(table, "join_war")
	w.TankTeam = field.NewField(table, "tank_team")
	w.CompetitiveRank = field.NewField(table, "competitive_rank")

	w.fillFieldMap()

	return w
}

func (w *wtPlayer) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *wtPlayer) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 36)
	w.fieldMap["id"] = w.ID
	w.fieldMap["visitor_id"] = w.VisitorID
	w.fieldMap["tag"] = w.Tag
	w.fieldMap["player_lvl"] = w.PlayerLvl
	w.fieldMap["player_exp"] = w.PlayerExp
	w.fieldMap["nickname"] = w.Nickname
	w.fieldMap["icons"] = w.Icons
	w.fieldMap["gold_pool"] = w.GoldPool
	w.fieldMap["gold_pool_ts"] = w.GoldPoolTs
	w.fieldMap["diamond"] = w.Diamond
	w.fieldMap["gold"] = w.Gold
	w.fieldMap["legion_id"] = w.LegionID
	w.fieldMap["legion_name"] = w.LegionName
	w.fieldMap["legion_position"] = w.LegionPosition
	w.fieldMap["tier"] = w.Tier
	w.fieldMap["trophy"] = w.Trophy
	w.fieldMap["trophy_road"] = w.TrophyRoad
	w.fieldMap["vip"] = w.Vip
	w.fieldMap["chest_info"] = w.ChestInfo
	w.fieldMap["garage"] = w.Garage
	w.fieldMap["inventory"] = w.Inventory
	w.fieldMap["path_of_valor"] = w.PathOfValor
	w.fieldMap["ip_region"] = w.IPRegion
	w.fieldMap["device_region"] = w.DeviceRegion
	w.fieldMap["rename_times"] = w.RenameTimes
	w.fieldMap["settlement_trophy"] = w.SettlementTrophy
	w.fieldMap["statistics_info"] = w.StatisticsInfo
	w.fieldMap["guide_info"] = w.GuideInfo
	w.fieldMap["status"] = w.Status
	w.fieldMap["create_time"] = w.CreateTime
	w.fieldMap["update_time"] = w.UpdateTime
	w.fieldMap["last_login_time"] = w.LastLoginTime
	w.fieldMap["version"] = w.Version
	w.fieldMap["join_war"] = w.JoinWar
	w.fieldMap["tank_team"] = w.TankTeam
	w.fieldMap["competitive_rank"] = w.CompetitiveRank
}

func (w wtPlayer) clone(db *gorm.DB) wtPlayer {
	w.wtPlayerDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w wtPlayer) replaceDB(db *gorm.DB) wtPlayer {
	w.wtPlayerDo.ReplaceDB(db)
	return w
}

type wtPlayerDo struct{ gen.DO }

func (w wtPlayerDo) Debug() *wtPlayerDo {
	return w.withDO(w.DO.Debug())
}

func (w wtPlayerDo) WithContext(ctx context.Context) *wtPlayerDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w wtPlayerDo) ReadDB() *wtPlayerDo {
	return w.Clauses(dbresolver.Read)
}

func (w wtPlayerDo) WriteDB() *wtPlayerDo {
	return w.Clauses(dbresolver.Write)
}

func (w wtPlayerDo) Session(config *gorm.Session) *wtPlayerDo {
	return w.withDO(w.DO.Session(config))
}

func (w wtPlayerDo) Clauses(conds ...clause.Expression) *wtPlayerDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w wtPlayerDo) Returning(value interface{}, columns ...string) *wtPlayerDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w wtPlayerDo) Not(conds ...gen.Condition) *wtPlayerDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w wtPlayerDo) Or(conds ...gen.Condition) *wtPlayerDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w wtPlayerDo) Select(conds ...field.Expr) *wtPlayerDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w wtPlayerDo) Where(conds ...gen.Condition) *wtPlayerDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w wtPlayerDo) Order(conds ...field.Expr) *wtPlayerDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w wtPlayerDo) Distinct(cols ...field.Expr) *wtPlayerDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w wtPlayerDo) Omit(cols ...field.Expr) *wtPlayerDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w wtPlayerDo) Join(table schema.Tabler, on ...field.Expr) *wtPlayerDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w wtPlayerDo) LeftJoin(table schema.Tabler, on ...field.Expr) *wtPlayerDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w wtPlayerDo) RightJoin(table schema.Tabler, on ...field.Expr) *wtPlayerDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w wtPlayerDo) Group(cols ...field.Expr) *wtPlayerDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w wtPlayerDo) Having(conds ...gen.Condition) *wtPlayerDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w wtPlayerDo) Limit(limit int) *wtPlayerDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w wtPlayerDo) Offset(offset int) *wtPlayerDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w wtPlayerDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *wtPlayerDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w wtPlayerDo) Unscoped() *wtPlayerDo {
	return w.withDO(w.DO.Unscoped())
}

func (w wtPlayerDo) Create(values ...*model.WtPlayer) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w wtPlayerDo) CreateInBatches(values []*model.WtPlayer, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w wtPlayerDo) Save(values ...*model.WtPlayer) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w wtPlayerDo) First() (*model.WtPlayer, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WtPlayer), nil
	}
}

func (w wtPlayerDo) Take() (*model.WtPlayer, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WtPlayer), nil
	}
}

func (w wtPlayerDo) Last() (*model.WtPlayer, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WtPlayer), nil
	}
}

func (w wtPlayerDo) Find() ([]*model.WtPlayer, error) {
	result, err := w.DO.Find()
	return result.([]*model.WtPlayer), err
}

func (w wtPlayerDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WtPlayer, err error) {
	buf := make([]*model.WtPlayer, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w wtPlayerDo) FindInBatches(result *[]*model.WtPlayer, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w wtPlayerDo) Attrs(attrs ...field.AssignExpr) *wtPlayerDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w wtPlayerDo) Assign(attrs ...field.AssignExpr) *wtPlayerDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w wtPlayerDo) Joins(fields ...field.RelationField) *wtPlayerDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w wtPlayerDo) Preload(fields ...field.RelationField) *wtPlayerDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w wtPlayerDo) FirstOrInit() (*model.WtPlayer, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WtPlayer), nil
	}
}

func (w wtPlayerDo) FirstOrCreate() (*model.WtPlayer, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WtPlayer), nil
	}
}

func (w wtPlayerDo) FindByPage(offset int, limit int) (result []*model.WtPlayer, count int64, err error) {
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

func (w wtPlayerDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w wtPlayerDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w wtPlayerDo) Delete(models ...*model.WtPlayer) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *wtPlayerDo) withDO(do gen.Dao) *wtPlayerDo {
	w.DO = *do.(*gen.DO)
	return w
}
