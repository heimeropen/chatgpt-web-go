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

func newAction(db *gorm.DB, opts ...gen.DOOption) action {
	_action := action{}

	_action.actionDo.UseDB(db, opts...)
	_action.actionDo.UseModel(&model.Action{})

	tableName := _action.actionDo.TableName()
	_action.ALL = field.NewAsterisk(tableName)
	_action.ID = field.NewInt64(tableName, "id")
	_action.UserID = field.NewInt64(tableName, "user_id")
	_action.Type = field.NewString(tableName, "type")
	_action.Describe = field.NewString(tableName, "describe")
	_action.IP = field.NewString(tableName, "ip")
	_action.CreateTime = field.NewTime(tableName, "create_time")
	_action.UpdateTime = field.NewTime(tableName, "update_time")
	_action.IsDelete = field.NewInt32(tableName, "is_delete")

	_action.fillFieldMap()

	return _action
}

type action struct {
	actionDo

	ALL        field.Asterisk
	ID         field.Int64
	UserID     field.Int64
	Type       field.String
	Describe   field.String
	IP         field.String
	CreateTime field.Time
	UpdateTime field.Time
	IsDelete   field.Int32

	fieldMap map[string]field.Expr
}

func (a action) Table(newTableName string) *action {
	a.actionDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a action) As(alias string) *action {
	a.actionDo.DO = *(a.actionDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *action) updateTableName(table string) *action {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt64(table, "id")
	a.UserID = field.NewInt64(table, "user_id")
	a.Type = field.NewString(table, "type")
	a.Describe = field.NewString(table, "describe")
	a.IP = field.NewString(table, "ip")
	a.CreateTime = field.NewTime(table, "create_time")
	a.UpdateTime = field.NewTime(table, "update_time")
	a.IsDelete = field.NewInt32(table, "is_delete")

	a.fillFieldMap()

	return a
}

func (a *action) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *action) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 8)
	a.fieldMap["id"] = a.ID
	a.fieldMap["user_id"] = a.UserID
	a.fieldMap["type"] = a.Type
	a.fieldMap["describe"] = a.Describe
	a.fieldMap["ip"] = a.IP
	a.fieldMap["create_time"] = a.CreateTime
	a.fieldMap["update_time"] = a.UpdateTime
	a.fieldMap["is_delete"] = a.IsDelete
}

func (a action) clone(db *gorm.DB) action {
	a.actionDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a action) replaceDB(db *gorm.DB) action {
	a.actionDo.ReplaceDB(db)
	return a
}

type actionDo struct{ gen.DO }

type IActionDo interface {
	gen.SubQuery
	Debug() IActionDo
	WithContext(ctx context.Context) IActionDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IActionDo
	WriteDB() IActionDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IActionDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IActionDo
	Not(conds ...gen.Condition) IActionDo
	Or(conds ...gen.Condition) IActionDo
	Select(conds ...field.Expr) IActionDo
	Where(conds ...gen.Condition) IActionDo
	Order(conds ...field.Expr) IActionDo
	Distinct(cols ...field.Expr) IActionDo
	Omit(cols ...field.Expr) IActionDo
	Join(table schema.Tabler, on ...field.Expr) IActionDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IActionDo
	RightJoin(table schema.Tabler, on ...field.Expr) IActionDo
	Group(cols ...field.Expr) IActionDo
	Having(conds ...gen.Condition) IActionDo
	Limit(limit int) IActionDo
	Offset(offset int) IActionDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IActionDo
	Unscoped() IActionDo
	Create(values ...*model.Action) error
	CreateInBatches(values []*model.Action, batchSize int) error
	Save(values ...*model.Action) error
	First() (*model.Action, error)
	Take() (*model.Action, error)
	Last() (*model.Action, error)
	Find() ([]*model.Action, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Action, err error)
	FindInBatches(result *[]*model.Action, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Action) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IActionDo
	Assign(attrs ...field.AssignExpr) IActionDo
	Joins(fields ...field.RelationField) IActionDo
	Preload(fields ...field.RelationField) IActionDo
	FirstOrInit() (*model.Action, error)
	FirstOrCreate() (*model.Action, error)
	FindByPage(offset int, limit int) (result []*model.Action, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IActionDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	FilterWithNameAndRole(name string, role string) (result []model.Action, err error)
}

// FilterWithNameAndRole SELECT * FROM @@table WHERE name = @name{{if role !=""}} AND role = @role{{end}}
func (a actionDo) FilterWithNameAndRole(name string, role string) (result []model.Action, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, name)
	generateSQL.WriteString("SELECT * FROM action WHERE name = ? ")
	if role != "" {
		params = append(params, role)
		generateSQL.WriteString("AND role = ? ")
	}

	var executeSQL *gorm.DB
	executeSQL = a.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (a actionDo) Debug() IActionDo {
	return a.withDO(a.DO.Debug())
}

func (a actionDo) WithContext(ctx context.Context) IActionDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a actionDo) ReadDB() IActionDo {
	return a.Clauses(dbresolver.Read)
}

func (a actionDo) WriteDB() IActionDo {
	return a.Clauses(dbresolver.Write)
}

func (a actionDo) Session(config *gorm.Session) IActionDo {
	return a.withDO(a.DO.Session(config))
}

func (a actionDo) Clauses(conds ...clause.Expression) IActionDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a actionDo) Returning(value interface{}, columns ...string) IActionDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a actionDo) Not(conds ...gen.Condition) IActionDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a actionDo) Or(conds ...gen.Condition) IActionDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a actionDo) Select(conds ...field.Expr) IActionDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a actionDo) Where(conds ...gen.Condition) IActionDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a actionDo) Order(conds ...field.Expr) IActionDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a actionDo) Distinct(cols ...field.Expr) IActionDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a actionDo) Omit(cols ...field.Expr) IActionDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a actionDo) Join(table schema.Tabler, on ...field.Expr) IActionDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a actionDo) LeftJoin(table schema.Tabler, on ...field.Expr) IActionDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a actionDo) RightJoin(table schema.Tabler, on ...field.Expr) IActionDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a actionDo) Group(cols ...field.Expr) IActionDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a actionDo) Having(conds ...gen.Condition) IActionDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a actionDo) Limit(limit int) IActionDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a actionDo) Offset(offset int) IActionDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a actionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IActionDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a actionDo) Unscoped() IActionDo {
	return a.withDO(a.DO.Unscoped())
}

func (a actionDo) Create(values ...*model.Action) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a actionDo) CreateInBatches(values []*model.Action, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a actionDo) Save(values ...*model.Action) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a actionDo) First() (*model.Action, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Action), nil
	}
}

func (a actionDo) Take() (*model.Action, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Action), nil
	}
}

func (a actionDo) Last() (*model.Action, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Action), nil
	}
}

func (a actionDo) Find() ([]*model.Action, error) {
	result, err := a.DO.Find()
	return result.([]*model.Action), err
}

func (a actionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Action, err error) {
	buf := make([]*model.Action, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a actionDo) FindInBatches(result *[]*model.Action, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a actionDo) Attrs(attrs ...field.AssignExpr) IActionDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a actionDo) Assign(attrs ...field.AssignExpr) IActionDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a actionDo) Joins(fields ...field.RelationField) IActionDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a actionDo) Preload(fields ...field.RelationField) IActionDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a actionDo) FirstOrInit() (*model.Action, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Action), nil
	}
}

func (a actionDo) FirstOrCreate() (*model.Action, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Action), nil
	}
}

func (a actionDo) FindByPage(offset int, limit int) (result []*model.Action, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a actionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a actionDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a actionDo) Delete(models ...*model.Action) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *actionDo) withDO(do gen.Dao) *actionDo {
	a.DO = *do.(*gen.DO)
	return a
}
