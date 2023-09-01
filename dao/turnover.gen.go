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

func newTurnover(db *gorm.DB, opts ...gen.DOOption) turnover {
	_turnover := turnover{}

	_turnover.turnoverDo.UseDB(db, opts...)
	_turnover.turnoverDo.UseModel(&model.Turnover{})

	tableName := _turnover.turnoverDo.TableName()
	_turnover.ALL = field.NewAsterisk(tableName)
	_turnover.ID = field.NewInt64(tableName, "id")
	_turnover.UserID = field.NewInt64(tableName, "user_id")
	_turnover.Describe = field.NewString(tableName, "describe")
	_turnover.Value = field.NewString(tableName, "value")
	_turnover.CreateTime = field.NewTime(tableName, "create_time")
	_turnover.UpdateTime = field.NewTime(tableName, "update_time")
	_turnover.IsDelete = field.NewInt32(tableName, "is_delete")

	_turnover.fillFieldMap()

	return _turnover
}

type turnover struct {
	turnoverDo

	ALL        field.Asterisk
	ID         field.Int64
	UserID     field.Int64  // 用户
	Describe   field.String // 描述
	Value      field.String // 值
	CreateTime field.Time
	UpdateTime field.Time
	IsDelete   field.Int32

	fieldMap map[string]field.Expr
}

func (t turnover) Table(newTableName string) *turnover {
	t.turnoverDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t turnover) As(alias string) *turnover {
	t.turnoverDo.DO = *(t.turnoverDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *turnover) updateTableName(table string) *turnover {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewInt64(table, "id")
	t.UserID = field.NewInt64(table, "user_id")
	t.Describe = field.NewString(table, "describe")
	t.Value = field.NewString(table, "value")
	t.CreateTime = field.NewTime(table, "create_time")
	t.UpdateTime = field.NewTime(table, "update_time")
	t.IsDelete = field.NewInt32(table, "is_delete")

	t.fillFieldMap()

	return t
}

func (t *turnover) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *turnover) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 7)
	t.fieldMap["id"] = t.ID
	t.fieldMap["user_id"] = t.UserID
	t.fieldMap["describe"] = t.Describe
	t.fieldMap["value"] = t.Value
	t.fieldMap["create_time"] = t.CreateTime
	t.fieldMap["update_time"] = t.UpdateTime
	t.fieldMap["is_delete"] = t.IsDelete
}

func (t turnover) clone(db *gorm.DB) turnover {
	t.turnoverDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t turnover) replaceDB(db *gorm.DB) turnover {
	t.turnoverDo.ReplaceDB(db)
	return t
}

type turnoverDo struct{ gen.DO }

type ITurnoverDo interface {
	gen.SubQuery
	Debug() ITurnoverDo
	WithContext(ctx context.Context) ITurnoverDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITurnoverDo
	WriteDB() ITurnoverDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITurnoverDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITurnoverDo
	Not(conds ...gen.Condition) ITurnoverDo
	Or(conds ...gen.Condition) ITurnoverDo
	Select(conds ...field.Expr) ITurnoverDo
	Where(conds ...gen.Condition) ITurnoverDo
	Order(conds ...field.Expr) ITurnoverDo
	Distinct(cols ...field.Expr) ITurnoverDo
	Omit(cols ...field.Expr) ITurnoverDo
	Join(table schema.Tabler, on ...field.Expr) ITurnoverDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITurnoverDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITurnoverDo
	Group(cols ...field.Expr) ITurnoverDo
	Having(conds ...gen.Condition) ITurnoverDo
	Limit(limit int) ITurnoverDo
	Offset(offset int) ITurnoverDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITurnoverDo
	Unscoped() ITurnoverDo
	Create(values ...*model.Turnover) error
	CreateInBatches(values []*model.Turnover, batchSize int) error
	Save(values ...*model.Turnover) error
	First() (*model.Turnover, error)
	Take() (*model.Turnover, error)
	Last() (*model.Turnover, error)
	Find() ([]*model.Turnover, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Turnover, err error)
	FindInBatches(result *[]*model.Turnover, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Turnover) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITurnoverDo
	Assign(attrs ...field.AssignExpr) ITurnoverDo
	Joins(fields ...field.RelationField) ITurnoverDo
	Preload(fields ...field.RelationField) ITurnoverDo
	FirstOrInit() (*model.Turnover, error)
	FirstOrCreate() (*model.Turnover, error)
	FindByPage(offset int, limit int) (result []*model.Turnover, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITurnoverDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	FilterWithNameAndRole(name string, role string) (result []model.Turnover, err error)
}

// FilterWithNameAndRole SELECT * FROM @@table WHERE name = @name{{if role !=""}} AND role = @role{{end}}
func (t turnoverDo) FilterWithNameAndRole(name string, role string) (result []model.Turnover, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, name)
	generateSQL.WriteString("SELECT * FROM turnover WHERE name = ? ")
	if role != "" {
		params = append(params, role)
		generateSQL.WriteString("AND role = ? ")
	}

	var executeSQL *gorm.DB
	executeSQL = t.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (t turnoverDo) Debug() ITurnoverDo {
	return t.withDO(t.DO.Debug())
}

func (t turnoverDo) WithContext(ctx context.Context) ITurnoverDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t turnoverDo) ReadDB() ITurnoverDo {
	return t.Clauses(dbresolver.Read)
}

func (t turnoverDo) WriteDB() ITurnoverDo {
	return t.Clauses(dbresolver.Write)
}

func (t turnoverDo) Session(config *gorm.Session) ITurnoverDo {
	return t.withDO(t.DO.Session(config))
}

func (t turnoverDo) Clauses(conds ...clause.Expression) ITurnoverDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t turnoverDo) Returning(value interface{}, columns ...string) ITurnoverDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t turnoverDo) Not(conds ...gen.Condition) ITurnoverDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t turnoverDo) Or(conds ...gen.Condition) ITurnoverDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t turnoverDo) Select(conds ...field.Expr) ITurnoverDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t turnoverDo) Where(conds ...gen.Condition) ITurnoverDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t turnoverDo) Order(conds ...field.Expr) ITurnoverDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t turnoverDo) Distinct(cols ...field.Expr) ITurnoverDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t turnoverDo) Omit(cols ...field.Expr) ITurnoverDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t turnoverDo) Join(table schema.Tabler, on ...field.Expr) ITurnoverDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t turnoverDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITurnoverDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t turnoverDo) RightJoin(table schema.Tabler, on ...field.Expr) ITurnoverDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t turnoverDo) Group(cols ...field.Expr) ITurnoverDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t turnoverDo) Having(conds ...gen.Condition) ITurnoverDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t turnoverDo) Limit(limit int) ITurnoverDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t turnoverDo) Offset(offset int) ITurnoverDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t turnoverDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITurnoverDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t turnoverDo) Unscoped() ITurnoverDo {
	return t.withDO(t.DO.Unscoped())
}

func (t turnoverDo) Create(values ...*model.Turnover) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t turnoverDo) CreateInBatches(values []*model.Turnover, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t turnoverDo) Save(values ...*model.Turnover) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t turnoverDo) First() (*model.Turnover, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Turnover), nil
	}
}

func (t turnoverDo) Take() (*model.Turnover, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Turnover), nil
	}
}

func (t turnoverDo) Last() (*model.Turnover, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Turnover), nil
	}
}

func (t turnoverDo) Find() ([]*model.Turnover, error) {
	result, err := t.DO.Find()
	return result.([]*model.Turnover), err
}

func (t turnoverDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Turnover, err error) {
	buf := make([]*model.Turnover, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t turnoverDo) FindInBatches(result *[]*model.Turnover, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t turnoverDo) Attrs(attrs ...field.AssignExpr) ITurnoverDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t turnoverDo) Assign(attrs ...field.AssignExpr) ITurnoverDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t turnoverDo) Joins(fields ...field.RelationField) ITurnoverDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t turnoverDo) Preload(fields ...field.RelationField) ITurnoverDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t turnoverDo) FirstOrInit() (*model.Turnover, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Turnover), nil
	}
}

func (t turnoverDo) FirstOrCreate() (*model.Turnover, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Turnover), nil
	}
}

func (t turnoverDo) FindByPage(offset int, limit int) (result []*model.Turnover, count int64, err error) {
	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t turnoverDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t turnoverDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t turnoverDo) Delete(models ...*model.Turnover) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *turnoverDo) withDO(do gen.Dao) *turnoverDo {
	t.DO = *do.(*gen.DO)
	return t
}
