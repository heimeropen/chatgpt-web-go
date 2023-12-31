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

func newAmountDetail(db *gorm.DB, opts ...gen.DOOption) amountDetail {
	_amountDetail := amountDetail{}

	_amountDetail.amountDetailDo.UseDB(db, opts...)
	_amountDetail.amountDetailDo.UseModel(&model.AmountDetail{})

	tableName := _amountDetail.amountDetailDo.TableName()
	_amountDetail.ALL = field.NewAsterisk(tableName)
	_amountDetail.ID = field.NewInt64(tableName, "id")
	_amountDetail.UserID = field.NewInt64(tableName, "user_id")
	_amountDetail.Type = field.NewString(tableName, "type")
	_amountDetail.CorrelationID = field.NewInt64(tableName, "correlation_id")
	_amountDetail.OriginalAmount = field.NewString(tableName, "original_amount")
	_amountDetail.OperateAmount = field.NewString(tableName, "operate_amount")
	_amountDetail.CurrentAmount = field.NewString(tableName, "current_amount")
	_amountDetail.Remarks = field.NewString(tableName, "remarks")
	_amountDetail.Status = field.NewInt32(tableName, "status")
	_amountDetail.CreateTime = field.NewTime(tableName, "create_time")
	_amountDetail.UpdateTime = field.NewTime(tableName, "update_time")
	_amountDetail.IsDelete = field.NewInt32(tableName, "is_delete")

	_amountDetail.fillFieldMap()

	return _amountDetail
}

type amountDetail struct {
	amountDetailDo

	ALL            field.Asterisk
	ID             field.Int64
	UserID         field.Int64
	Type           field.String // 提现 or 提成
	CorrelationID  field.Int64  // 关联ID
	OriginalAmount field.String // 原始金额 分
	OperateAmount  field.String // 操作金额
	CurrentAmount  field.String // 当前金额
	Remarks        field.String // 备注
	Status         field.Int32  // 1-正常
	CreateTime     field.Time
	UpdateTime     field.Time
	IsDelete       field.Int32

	fieldMap map[string]field.Expr
}

func (a amountDetail) Table(newTableName string) *amountDetail {
	a.amountDetailDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a amountDetail) As(alias string) *amountDetail {
	a.amountDetailDo.DO = *(a.amountDetailDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *amountDetail) updateTableName(table string) *amountDetail {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt64(table, "id")
	a.UserID = field.NewInt64(table, "user_id")
	a.Type = field.NewString(table, "type")
	a.CorrelationID = field.NewInt64(table, "correlation_id")
	a.OriginalAmount = field.NewString(table, "original_amount")
	a.OperateAmount = field.NewString(table, "operate_amount")
	a.CurrentAmount = field.NewString(table, "current_amount")
	a.Remarks = field.NewString(table, "remarks")
	a.Status = field.NewInt32(table, "status")
	a.CreateTime = field.NewTime(table, "create_time")
	a.UpdateTime = field.NewTime(table, "update_time")
	a.IsDelete = field.NewInt32(table, "is_delete")

	a.fillFieldMap()

	return a
}

func (a *amountDetail) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *amountDetail) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 12)
	a.fieldMap["id"] = a.ID
	a.fieldMap["user_id"] = a.UserID
	a.fieldMap["type"] = a.Type
	a.fieldMap["correlation_id"] = a.CorrelationID
	a.fieldMap["original_amount"] = a.OriginalAmount
	a.fieldMap["operate_amount"] = a.OperateAmount
	a.fieldMap["current_amount"] = a.CurrentAmount
	a.fieldMap["remarks"] = a.Remarks
	a.fieldMap["status"] = a.Status
	a.fieldMap["create_time"] = a.CreateTime
	a.fieldMap["update_time"] = a.UpdateTime
	a.fieldMap["is_delete"] = a.IsDelete
}

func (a amountDetail) clone(db *gorm.DB) amountDetail {
	a.amountDetailDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a amountDetail) replaceDB(db *gorm.DB) amountDetail {
	a.amountDetailDo.ReplaceDB(db)
	return a
}

type amountDetailDo struct{ gen.DO }

type IAmountDetailDo interface {
	gen.SubQuery
	Debug() IAmountDetailDo
	WithContext(ctx context.Context) IAmountDetailDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAmountDetailDo
	WriteDB() IAmountDetailDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAmountDetailDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAmountDetailDo
	Not(conds ...gen.Condition) IAmountDetailDo
	Or(conds ...gen.Condition) IAmountDetailDo
	Select(conds ...field.Expr) IAmountDetailDo
	Where(conds ...gen.Condition) IAmountDetailDo
	Order(conds ...field.Expr) IAmountDetailDo
	Distinct(cols ...field.Expr) IAmountDetailDo
	Omit(cols ...field.Expr) IAmountDetailDo
	Join(table schema.Tabler, on ...field.Expr) IAmountDetailDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAmountDetailDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAmountDetailDo
	Group(cols ...field.Expr) IAmountDetailDo
	Having(conds ...gen.Condition) IAmountDetailDo
	Limit(limit int) IAmountDetailDo
	Offset(offset int) IAmountDetailDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAmountDetailDo
	Unscoped() IAmountDetailDo
	Create(values ...*model.AmountDetail) error
	CreateInBatches(values []*model.AmountDetail, batchSize int) error
	Save(values ...*model.AmountDetail) error
	First() (*model.AmountDetail, error)
	Take() (*model.AmountDetail, error)
	Last() (*model.AmountDetail, error)
	Find() ([]*model.AmountDetail, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AmountDetail, err error)
	FindInBatches(result *[]*model.AmountDetail, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.AmountDetail) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAmountDetailDo
	Assign(attrs ...field.AssignExpr) IAmountDetailDo
	Joins(fields ...field.RelationField) IAmountDetailDo
	Preload(fields ...field.RelationField) IAmountDetailDo
	FirstOrInit() (*model.AmountDetail, error)
	FirstOrCreate() (*model.AmountDetail, error)
	FindByPage(offset int, limit int) (result []*model.AmountDetail, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAmountDetailDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	FilterWithNameAndRole(name string, role string) (result []model.AmountDetail, err error)
}

// FilterWithNameAndRole SELECT * FROM @@table WHERE name = @name{{if role !=""}} AND role = @role{{end}}
func (a amountDetailDo) FilterWithNameAndRole(name string, role string) (result []model.AmountDetail, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, name)
	generateSQL.WriteString("SELECT * FROM amount_details WHERE name = ? ")
	if role != "" {
		params = append(params, role)
		generateSQL.WriteString("AND role = ? ")
	}

	var executeSQL *gorm.DB
	executeSQL = a.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (a amountDetailDo) Debug() IAmountDetailDo {
	return a.withDO(a.DO.Debug())
}

func (a amountDetailDo) WithContext(ctx context.Context) IAmountDetailDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a amountDetailDo) ReadDB() IAmountDetailDo {
	return a.Clauses(dbresolver.Read)
}

func (a amountDetailDo) WriteDB() IAmountDetailDo {
	return a.Clauses(dbresolver.Write)
}

func (a amountDetailDo) Session(config *gorm.Session) IAmountDetailDo {
	return a.withDO(a.DO.Session(config))
}

func (a amountDetailDo) Clauses(conds ...clause.Expression) IAmountDetailDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a amountDetailDo) Returning(value interface{}, columns ...string) IAmountDetailDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a amountDetailDo) Not(conds ...gen.Condition) IAmountDetailDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a amountDetailDo) Or(conds ...gen.Condition) IAmountDetailDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a amountDetailDo) Select(conds ...field.Expr) IAmountDetailDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a amountDetailDo) Where(conds ...gen.Condition) IAmountDetailDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a amountDetailDo) Order(conds ...field.Expr) IAmountDetailDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a amountDetailDo) Distinct(cols ...field.Expr) IAmountDetailDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a amountDetailDo) Omit(cols ...field.Expr) IAmountDetailDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a amountDetailDo) Join(table schema.Tabler, on ...field.Expr) IAmountDetailDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a amountDetailDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAmountDetailDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a amountDetailDo) RightJoin(table schema.Tabler, on ...field.Expr) IAmountDetailDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a amountDetailDo) Group(cols ...field.Expr) IAmountDetailDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a amountDetailDo) Having(conds ...gen.Condition) IAmountDetailDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a amountDetailDo) Limit(limit int) IAmountDetailDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a amountDetailDo) Offset(offset int) IAmountDetailDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a amountDetailDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAmountDetailDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a amountDetailDo) Unscoped() IAmountDetailDo {
	return a.withDO(a.DO.Unscoped())
}

func (a amountDetailDo) Create(values ...*model.AmountDetail) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a amountDetailDo) CreateInBatches(values []*model.AmountDetail, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a amountDetailDo) Save(values ...*model.AmountDetail) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a amountDetailDo) First() (*model.AmountDetail, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.AmountDetail), nil
	}
}

func (a amountDetailDo) Take() (*model.AmountDetail, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.AmountDetail), nil
	}
}

func (a amountDetailDo) Last() (*model.AmountDetail, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.AmountDetail), nil
	}
}

func (a amountDetailDo) Find() ([]*model.AmountDetail, error) {
	result, err := a.DO.Find()
	return result.([]*model.AmountDetail), err
}

func (a amountDetailDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AmountDetail, err error) {
	buf := make([]*model.AmountDetail, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a amountDetailDo) FindInBatches(result *[]*model.AmountDetail, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a amountDetailDo) Attrs(attrs ...field.AssignExpr) IAmountDetailDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a amountDetailDo) Assign(attrs ...field.AssignExpr) IAmountDetailDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a amountDetailDo) Joins(fields ...field.RelationField) IAmountDetailDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a amountDetailDo) Preload(fields ...field.RelationField) IAmountDetailDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a amountDetailDo) FirstOrInit() (*model.AmountDetail, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.AmountDetail), nil
	}
}

func (a amountDetailDo) FirstOrCreate() (*model.AmountDetail, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.AmountDetail), nil
	}
}

func (a amountDetailDo) FindByPage(offset int, limit int) (result []*model.AmountDetail, count int64, err error) {
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

func (a amountDetailDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a amountDetailDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a amountDetailDo) Delete(models ...*model.AmountDetail) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *amountDetailDo) withDO(do gen.Dao) *amountDetailDo {
	a.DO = *do.(*gen.DO)
	return a
}
