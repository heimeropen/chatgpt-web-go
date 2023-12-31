// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"chatgpt-web-new-go/model"
)

func newCarmi(db *gorm.DB, opts ...gen.DOOption) carmi {
	_carmi := carmi{}

	_carmi.carmiDo.UseDB(db, opts...)
	_carmi.carmiDo.UseModel(&model.Carmi{})

	tableName := _carmi.carmiDo.TableName()
	_carmi.ALL = field.NewAsterisk(tableName)
	_carmi.ID = field.NewInt64(tableName, "id")
	_carmi.IP = field.NewString(tableName, "ip")
	_carmi.UserID = field.NewInt64(tableName, "user_id")
	_carmi.Key = field.NewString(tableName, "key")
	_carmi.Value = field.NewInt32(tableName, "value")
	_carmi.Status = field.NewInt32(tableName, "status")
	_carmi.Type = field.NewString(tableName, "type")
	_carmi.EndTime = field.NewString(tableName, "end_time")
	_carmi.Level = field.NewInt32(tableName, "level")
	_carmi.CreateTime = field.NewTime(tableName, "create_time")
	_carmi.UpdateTime = field.NewTime(tableName, "update_time")
	_carmi.IsDelete = field.NewInt32(tableName, "is_delete")

	_carmi.fillFieldMap()

	return _carmi
}

type carmi struct {
	carmiDo

	ALL        field.Asterisk
	ID         field.Int64
	IP         field.String // 使用时候的ip
	UserID     field.Int64  // 使用者
	Key        field.String // 卡密
	Value      field.Int32  // 积分/天数
	Status     field.Int32  // 0有效 1使用 2过期
	Type       field.String // 类型 integral/vip_days/svip_days
	EndTime    field.String // 截止时间
	Level      field.Int32  // 卡密充值等级
	CreateTime field.Time
	UpdateTime field.Time
	IsDelete   field.Int32

	fieldMap map[string]field.Expr
}

func (c carmi) Table(newTableName string) *carmi {
	c.carmiDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c carmi) As(alias string) *carmi {
	c.carmiDo.DO = *(c.carmiDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *carmi) updateTableName(table string) *carmi {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewInt64(table, "id")
	c.IP = field.NewString(table, "ip")
	c.UserID = field.NewInt64(table, "user_id")
	c.Key = field.NewString(table, "key")
	c.Value = field.NewInt32(table, "value")
	c.Status = field.NewInt32(table, "status")
	c.Type = field.NewString(table, "type")
	c.EndTime = field.NewString(table, "end_time")
	c.Level = field.NewInt32(table, "level")
	c.CreateTime = field.NewTime(table, "create_time")
	c.UpdateTime = field.NewTime(table, "update_time")
	c.IsDelete = field.NewInt32(table, "is_delete")

	c.fillFieldMap()

	return c
}

func (c *carmi) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *carmi) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 12)
	c.fieldMap["id"] = c.ID
	c.fieldMap["ip"] = c.IP
	c.fieldMap["user_id"] = c.UserID
	c.fieldMap["key"] = c.Key
	c.fieldMap["value"] = c.Value
	c.fieldMap["status"] = c.Status
	c.fieldMap["type"] = c.Type
	c.fieldMap["end_time"] = c.EndTime
	c.fieldMap["level"] = c.Level
	c.fieldMap["create_time"] = c.CreateTime
	c.fieldMap["update_time"] = c.UpdateTime
	c.fieldMap["is_delete"] = c.IsDelete
}

func (c carmi) clone(db *gorm.DB) carmi {
	c.carmiDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c carmi) replaceDB(db *gorm.DB) carmi {
	c.carmiDo.ReplaceDB(db)
	return c
}

type carmiDo struct{ gen.DO }

type ICarmiDo interface {
	gen.SubQuery
	Debug() ICarmiDo
	WithContext(ctx context.Context) ICarmiDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICarmiDo
	WriteDB() ICarmiDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICarmiDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICarmiDo
	Not(conds ...gen.Condition) ICarmiDo
	Or(conds ...gen.Condition) ICarmiDo
	Select(conds ...field.Expr) ICarmiDo
	Where(conds ...gen.Condition) ICarmiDo
	Order(conds ...field.Expr) ICarmiDo
	Distinct(cols ...field.Expr) ICarmiDo
	Omit(cols ...field.Expr) ICarmiDo
	Join(table schema.Tabler, on ...field.Expr) ICarmiDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICarmiDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICarmiDo
	Group(cols ...field.Expr) ICarmiDo
	Having(conds ...gen.Condition) ICarmiDo
	Limit(limit int) ICarmiDo
	Offset(offset int) ICarmiDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICarmiDo
	Unscoped() ICarmiDo
	Create(values ...*model.Carmi) error
	CreateInBatches(values []*model.Carmi, batchSize int) error
	Save(values ...*model.Carmi) error
	First() (*model.Carmi, error)
	Take() (*model.Carmi, error)
	Last() (*model.Carmi, error)
	Find() ([]*model.Carmi, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Carmi, err error)
	FindInBatches(result *[]*model.Carmi, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Carmi) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICarmiDo
	Assign(attrs ...field.AssignExpr) ICarmiDo
	Joins(fields ...field.RelationField) ICarmiDo
	Preload(fields ...field.RelationField) ICarmiDo
	FirstOrInit() (*model.Carmi, error)
	FirstOrCreate() (*model.Carmi, error)
	FindByPage(offset int, limit int) (result []*model.Carmi, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICarmiDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	FilterWithNameAndRole(name string, role string) (result []model.Carmi, err error)
}

// FilterWithNameAndRole SELECT * FROM @@table WHERE name = @name{{if role !=""}} AND role = @role{{end}}
func (c carmiDo) FilterWithNameAndRole(name string, role string) (result []model.Carmi, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, name)
	generateSQL.WriteString("SELECT * FROM carmi WHERE name = ? ")
	if role != "" {
		params = append(params, role)
		generateSQL.WriteString("AND role = ? ")
	}

	var executeSQL *gorm.DB
	executeSQL = c.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (c carmiDo) Debug() ICarmiDo {
	return c.withDO(c.DO.Debug())
}

func (c carmiDo) WithContext(ctx context.Context) ICarmiDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c carmiDo) ReadDB() ICarmiDo {
	return c.Clauses(dbresolver.Read)
}

func (c carmiDo) WriteDB() ICarmiDo {
	return c.Clauses(dbresolver.Write)
}

func (c carmiDo) Session(config *gorm.Session) ICarmiDo {
	return c.withDO(c.DO.Session(config))
}

func (c carmiDo) Clauses(conds ...clause.Expression) ICarmiDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c carmiDo) Returning(value interface{}, columns ...string) ICarmiDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c carmiDo) Not(conds ...gen.Condition) ICarmiDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c carmiDo) Or(conds ...gen.Condition) ICarmiDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c carmiDo) Select(conds ...field.Expr) ICarmiDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c carmiDo) Where(conds ...gen.Condition) ICarmiDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c carmiDo) Order(conds ...field.Expr) ICarmiDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c carmiDo) Distinct(cols ...field.Expr) ICarmiDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c carmiDo) Omit(cols ...field.Expr) ICarmiDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c carmiDo) Join(table schema.Tabler, on ...field.Expr) ICarmiDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c carmiDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICarmiDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c carmiDo) RightJoin(table schema.Tabler, on ...field.Expr) ICarmiDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c carmiDo) Group(cols ...field.Expr) ICarmiDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c carmiDo) Having(conds ...gen.Condition) ICarmiDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c carmiDo) Limit(limit int) ICarmiDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c carmiDo) Offset(offset int) ICarmiDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c carmiDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICarmiDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c carmiDo) Unscoped() ICarmiDo {
	return c.withDO(c.DO.Unscoped())
}

func (c carmiDo) Create(values ...*model.Carmi) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c carmiDo) CreateInBatches(values []*model.Carmi, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c carmiDo) Save(values ...*model.Carmi) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c carmiDo) First() (*model.Carmi, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Carmi), nil
	}
}

func (c carmiDo) Take() (*model.Carmi, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Carmi), nil
	}
}

func (c carmiDo) Last() (*model.Carmi, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Carmi), nil
	}
}

func (c carmiDo) Find() ([]*model.Carmi, error) {
	result, err := c.DO.Find()
	return result.([]*model.Carmi), err
}

func (c carmiDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Carmi, err error) {
	buf := make([]*model.Carmi, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c carmiDo) FindInBatches(result *[]*model.Carmi, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c carmiDo) Attrs(attrs ...field.AssignExpr) ICarmiDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c carmiDo) Assign(attrs ...field.AssignExpr) ICarmiDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c carmiDo) Joins(fields ...field.RelationField) ICarmiDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c carmiDo) Preload(fields ...field.RelationField) ICarmiDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c carmiDo) FirstOrInit() (*model.Carmi, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Carmi), nil
	}
}

func (c carmiDo) FirstOrCreate() (*model.Carmi, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Carmi), nil
	}
}

func (c carmiDo) FindByPage(offset int, limit int) (result []*model.Carmi, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c carmiDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c carmiDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c carmiDo) Delete(models ...*model.Carmi) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *carmiDo) withDO(do gen.Dao) *carmiDo {
	c.DO = *do.(*gen.DO)
	return c
}
