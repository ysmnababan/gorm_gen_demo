// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package repo

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"gorm_demo/staging/model"
)

func newOutletOp(db *gorm.DB, opts ...gen.DOOption) outletOp {
	_outletOp := outletOp{}

	_outletOp.outletOpDo.UseDB(db, opts...)
	_outletOp.outletOpDo.UseModel(&model.OutletOp{})

	tableName := _outletOp.outletOpDo.TableName()
	_outletOp.ALL = field.NewAsterisk(tableName)
	_outletOp.CreatedAt = field.NewInt64(tableName, "created_at")
	_outletOp.CreatedBy = field.NewString(tableName, "created_by")
	_outletOp.ModifiedAt = field.NewInt64(tableName, "modified_at")
	_outletOp.ModifiedBy = field.NewString(tableName, "modified_by")
	_outletOp.DeletedAt = field.NewInt64(tableName, "deleted_at")
	_outletOp.DeletedBy = field.NewString(tableName, "deleted_by")
	_outletOp.OutletOpID = field.NewString(tableName, "outlet_op_id")
	_outletOp.OutletID = field.NewString(tableName, "outlet_id")
	_outletOp.OpenDay = field.NewString(tableName, "open_day")
	_outletOp.OpenHour = field.NewString(tableName, "open_hour")
	_outletOp.CloseHour = field.NewString(tableName, "close_hour")
	_outletOp.IsOpen = field.NewBool(tableName, "is_open")

	_outletOp.fillFieldMap()

	return _outletOp
}

type outletOp struct {
	outletOpDo

	ALL        field.Asterisk
	CreatedAt  field.Int64
	CreatedBy  field.String
	ModifiedAt field.Int64
	ModifiedBy field.String
	DeletedAt  field.Int64
	DeletedBy  field.String
	OutletOpID field.String
	OutletID   field.String
	OpenDay    field.String
	OpenHour   field.String
	CloseHour  field.String
	IsOpen     field.Bool

	fieldMap map[string]field.Expr
}

func (o outletOp) Table(newTableName string) *outletOp {
	o.outletOpDo.UseTable(newTableName)
	return o.updateTableName(newTableName)
}

func (o outletOp) As(alias string) *outletOp {
	o.outletOpDo.DO = *(o.outletOpDo.As(alias).(*gen.DO))
	return o.updateTableName(alias)
}

func (o *outletOp) updateTableName(table string) *outletOp {
	o.ALL = field.NewAsterisk(table)
	o.CreatedAt = field.NewInt64(table, "created_at")
	o.CreatedBy = field.NewString(table, "created_by")
	o.ModifiedAt = field.NewInt64(table, "modified_at")
	o.ModifiedBy = field.NewString(table, "modified_by")
	o.DeletedAt = field.NewInt64(table, "deleted_at")
	o.DeletedBy = field.NewString(table, "deleted_by")
	o.OutletOpID = field.NewString(table, "outlet_op_id")
	o.OutletID = field.NewString(table, "outlet_id")
	o.OpenDay = field.NewString(table, "open_day")
	o.OpenHour = field.NewString(table, "open_hour")
	o.CloseHour = field.NewString(table, "close_hour")
	o.IsOpen = field.NewBool(table, "is_open")

	o.fillFieldMap()

	return o
}

func (o *outletOp) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := o.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (o *outletOp) fillFieldMap() {
	o.fieldMap = make(map[string]field.Expr, 12)
	o.fieldMap["created_at"] = o.CreatedAt
	o.fieldMap["created_by"] = o.CreatedBy
	o.fieldMap["modified_at"] = o.ModifiedAt
	o.fieldMap["modified_by"] = o.ModifiedBy
	o.fieldMap["deleted_at"] = o.DeletedAt
	o.fieldMap["deleted_by"] = o.DeletedBy
	o.fieldMap["outlet_op_id"] = o.OutletOpID
	o.fieldMap["outlet_id"] = o.OutletID
	o.fieldMap["open_day"] = o.OpenDay
	o.fieldMap["open_hour"] = o.OpenHour
	o.fieldMap["close_hour"] = o.CloseHour
	o.fieldMap["is_open"] = o.IsOpen
}

func (o outletOp) clone(db *gorm.DB) outletOp {
	o.outletOpDo.ReplaceConnPool(db.Statement.ConnPool)
	return o
}

func (o outletOp) replaceDB(db *gorm.DB) outletOp {
	o.outletOpDo.ReplaceDB(db)
	return o
}

type outletOpDo struct{ gen.DO }

type IOutletOpDo interface {
	gen.SubQuery
	Debug() IOutletOpDo
	WithContext(ctx context.Context) IOutletOpDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IOutletOpDo
	WriteDB() IOutletOpDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IOutletOpDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IOutletOpDo
	Not(conds ...gen.Condition) IOutletOpDo
	Or(conds ...gen.Condition) IOutletOpDo
	Select(conds ...field.Expr) IOutletOpDo
	Where(conds ...gen.Condition) IOutletOpDo
	Order(conds ...field.Expr) IOutletOpDo
	Distinct(cols ...field.Expr) IOutletOpDo
	Omit(cols ...field.Expr) IOutletOpDo
	Join(table schema.Tabler, on ...field.Expr) IOutletOpDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IOutletOpDo
	RightJoin(table schema.Tabler, on ...field.Expr) IOutletOpDo
	Group(cols ...field.Expr) IOutletOpDo
	Having(conds ...gen.Condition) IOutletOpDo
	Limit(limit int) IOutletOpDo
	Offset(offset int) IOutletOpDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IOutletOpDo
	Unscoped() IOutletOpDo
	Create(values ...*model.OutletOp) error
	CreateInBatches(values []*model.OutletOp, batchSize int) error
	Save(values ...*model.OutletOp) error
	First() (*model.OutletOp, error)
	Take() (*model.OutletOp, error)
	Last() (*model.OutletOp, error)
	Find() ([]*model.OutletOp, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.OutletOp, err error)
	FindInBatches(result *[]*model.OutletOp, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.OutletOp) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IOutletOpDo
	Assign(attrs ...field.AssignExpr) IOutletOpDo
	Joins(fields ...field.RelationField) IOutletOpDo
	Preload(fields ...field.RelationField) IOutletOpDo
	FirstOrInit() (*model.OutletOp, error)
	FirstOrCreate() (*model.OutletOp, error)
	FindByPage(offset int, limit int) (result []*model.OutletOp, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IOutletOpDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (o outletOpDo) Debug() IOutletOpDo {
	return o.withDO(o.DO.Debug())
}

func (o outletOpDo) WithContext(ctx context.Context) IOutletOpDo {
	return o.withDO(o.DO.WithContext(ctx))
}

func (o outletOpDo) ReadDB() IOutletOpDo {
	return o.Clauses(dbresolver.Read)
}

func (o outletOpDo) WriteDB() IOutletOpDo {
	return o.Clauses(dbresolver.Write)
}

func (o outletOpDo) Session(config *gorm.Session) IOutletOpDo {
	return o.withDO(o.DO.Session(config))
}

func (o outletOpDo) Clauses(conds ...clause.Expression) IOutletOpDo {
	return o.withDO(o.DO.Clauses(conds...))
}

func (o outletOpDo) Returning(value interface{}, columns ...string) IOutletOpDo {
	return o.withDO(o.DO.Returning(value, columns...))
}

func (o outletOpDo) Not(conds ...gen.Condition) IOutletOpDo {
	return o.withDO(o.DO.Not(conds...))
}

func (o outletOpDo) Or(conds ...gen.Condition) IOutletOpDo {
	return o.withDO(o.DO.Or(conds...))
}

func (o outletOpDo) Select(conds ...field.Expr) IOutletOpDo {
	return o.withDO(o.DO.Select(conds...))
}

func (o outletOpDo) Where(conds ...gen.Condition) IOutletOpDo {
	return o.withDO(o.DO.Where(conds...))
}

func (o outletOpDo) Order(conds ...field.Expr) IOutletOpDo {
	return o.withDO(o.DO.Order(conds...))
}

func (o outletOpDo) Distinct(cols ...field.Expr) IOutletOpDo {
	return o.withDO(o.DO.Distinct(cols...))
}

func (o outletOpDo) Omit(cols ...field.Expr) IOutletOpDo {
	return o.withDO(o.DO.Omit(cols...))
}

func (o outletOpDo) Join(table schema.Tabler, on ...field.Expr) IOutletOpDo {
	return o.withDO(o.DO.Join(table, on...))
}

func (o outletOpDo) LeftJoin(table schema.Tabler, on ...field.Expr) IOutletOpDo {
	return o.withDO(o.DO.LeftJoin(table, on...))
}

func (o outletOpDo) RightJoin(table schema.Tabler, on ...field.Expr) IOutletOpDo {
	return o.withDO(o.DO.RightJoin(table, on...))
}

func (o outletOpDo) Group(cols ...field.Expr) IOutletOpDo {
	return o.withDO(o.DO.Group(cols...))
}

func (o outletOpDo) Having(conds ...gen.Condition) IOutletOpDo {
	return o.withDO(o.DO.Having(conds...))
}

func (o outletOpDo) Limit(limit int) IOutletOpDo {
	return o.withDO(o.DO.Limit(limit))
}

func (o outletOpDo) Offset(offset int) IOutletOpDo {
	return o.withDO(o.DO.Offset(offset))
}

func (o outletOpDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IOutletOpDo {
	return o.withDO(o.DO.Scopes(funcs...))
}

func (o outletOpDo) Unscoped() IOutletOpDo {
	return o.withDO(o.DO.Unscoped())
}

func (o outletOpDo) Create(values ...*model.OutletOp) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Create(values)
}

func (o outletOpDo) CreateInBatches(values []*model.OutletOp, batchSize int) error {
	return o.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (o outletOpDo) Save(values ...*model.OutletOp) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Save(values)
}

func (o outletOpDo) First() (*model.OutletOp, error) {
	if result, err := o.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.OutletOp), nil
	}
}

func (o outletOpDo) Take() (*model.OutletOp, error) {
	if result, err := o.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.OutletOp), nil
	}
}

func (o outletOpDo) Last() (*model.OutletOp, error) {
	if result, err := o.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.OutletOp), nil
	}
}

func (o outletOpDo) Find() ([]*model.OutletOp, error) {
	result, err := o.DO.Find()
	return result.([]*model.OutletOp), err
}

func (o outletOpDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.OutletOp, err error) {
	buf := make([]*model.OutletOp, 0, batchSize)
	err = o.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (o outletOpDo) FindInBatches(result *[]*model.OutletOp, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return o.DO.FindInBatches(result, batchSize, fc)
}

func (o outletOpDo) Attrs(attrs ...field.AssignExpr) IOutletOpDo {
	return o.withDO(o.DO.Attrs(attrs...))
}

func (o outletOpDo) Assign(attrs ...field.AssignExpr) IOutletOpDo {
	return o.withDO(o.DO.Assign(attrs...))
}

func (o outletOpDo) Joins(fields ...field.RelationField) IOutletOpDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Joins(_f))
	}
	return &o
}

func (o outletOpDo) Preload(fields ...field.RelationField) IOutletOpDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Preload(_f))
	}
	return &o
}

func (o outletOpDo) FirstOrInit() (*model.OutletOp, error) {
	if result, err := o.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.OutletOp), nil
	}
}

func (o outletOpDo) FirstOrCreate() (*model.OutletOp, error) {
	if result, err := o.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.OutletOp), nil
	}
}

func (o outletOpDo) FindByPage(offset int, limit int) (result []*model.OutletOp, count int64, err error) {
	result, err = o.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = o.Offset(-1).Limit(-1).Count()
	return
}

func (o outletOpDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = o.Count()
	if err != nil {
		return
	}

	err = o.Offset(offset).Limit(limit).Scan(result)
	return
}

func (o outletOpDo) Scan(result interface{}) (err error) {
	return o.DO.Scan(result)
}

func (o outletOpDo) Delete(models ...*model.OutletOp) (result gen.ResultInfo, err error) {
	return o.DO.Delete(models)
}

func (o *outletOpDo) withDO(do gen.Dao) *outletOpDo {
	o.DO = *do.(*gen.DO)
	return o
}
