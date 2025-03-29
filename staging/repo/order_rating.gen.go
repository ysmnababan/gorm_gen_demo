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

func newOrderRating(db *gorm.DB, opts ...gen.DOOption) orderRating {
	_orderRating := orderRating{}

	_orderRating.orderRatingDo.UseDB(db, opts...)
	_orderRating.orderRatingDo.UseModel(&model.OrderRating{})

	tableName := _orderRating.orderRatingDo.TableName()
	_orderRating.ALL = field.NewAsterisk(tableName)
	_orderRating.OrderRatingID = field.NewString(tableName, "order_rating_id")
	_orderRating.OrderID = field.NewString(tableName, "order_id")
	_orderRating.Rating = field.NewInt32(tableName, "rating")
	_orderRating.ReviewAt = field.NewInt64(tableName, "review_at")
	_orderRating.ReviewMsg = field.NewString(tableName, "review_msg")
	_orderRating.ReplyAt = field.NewInt64(tableName, "reply_at")
	_orderRating.ReplyBy = field.NewString(tableName, "reply_by")
	_orderRating.ReplyMsg = field.NewString(tableName, "reply_msg")
	_orderRating.IsAnonymous = field.NewBool(tableName, "is_anonymous")

	_orderRating.fillFieldMap()

	return _orderRating
}

type orderRating struct {
	orderRatingDo

	ALL           field.Asterisk
	OrderRatingID field.String
	OrderID       field.String
	Rating        field.Int32
	ReviewAt      field.Int64
	ReviewMsg     field.String
	ReplyAt       field.Int64
	ReplyBy       field.String
	ReplyMsg      field.String
	IsAnonymous   field.Bool

	fieldMap map[string]field.Expr
}

func (o orderRating) Table(newTableName string) *orderRating {
	o.orderRatingDo.UseTable(newTableName)
	return o.updateTableName(newTableName)
}

func (o orderRating) As(alias string) *orderRating {
	o.orderRatingDo.DO = *(o.orderRatingDo.As(alias).(*gen.DO))
	return o.updateTableName(alias)
}

func (o *orderRating) updateTableName(table string) *orderRating {
	o.ALL = field.NewAsterisk(table)
	o.OrderRatingID = field.NewString(table, "order_rating_id")
	o.OrderID = field.NewString(table, "order_id")
	o.Rating = field.NewInt32(table, "rating")
	o.ReviewAt = field.NewInt64(table, "review_at")
	o.ReviewMsg = field.NewString(table, "review_msg")
	o.ReplyAt = field.NewInt64(table, "reply_at")
	o.ReplyBy = field.NewString(table, "reply_by")
	o.ReplyMsg = field.NewString(table, "reply_msg")
	o.IsAnonymous = field.NewBool(table, "is_anonymous")

	o.fillFieldMap()

	return o
}

func (o *orderRating) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := o.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (o *orderRating) fillFieldMap() {
	o.fieldMap = make(map[string]field.Expr, 9)
	o.fieldMap["order_rating_id"] = o.OrderRatingID
	o.fieldMap["order_id"] = o.OrderID
	o.fieldMap["rating"] = o.Rating
	o.fieldMap["review_at"] = o.ReviewAt
	o.fieldMap["review_msg"] = o.ReviewMsg
	o.fieldMap["reply_at"] = o.ReplyAt
	o.fieldMap["reply_by"] = o.ReplyBy
	o.fieldMap["reply_msg"] = o.ReplyMsg
	o.fieldMap["is_anonymous"] = o.IsAnonymous
}

func (o orderRating) clone(db *gorm.DB) orderRating {
	o.orderRatingDo.ReplaceConnPool(db.Statement.ConnPool)
	return o
}

func (o orderRating) replaceDB(db *gorm.DB) orderRating {
	o.orderRatingDo.ReplaceDB(db)
	return o
}

type orderRatingDo struct{ gen.DO }

type IOrderRatingDo interface {
	gen.SubQuery
	Debug() IOrderRatingDo
	WithContext(ctx context.Context) IOrderRatingDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IOrderRatingDo
	WriteDB() IOrderRatingDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IOrderRatingDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IOrderRatingDo
	Not(conds ...gen.Condition) IOrderRatingDo
	Or(conds ...gen.Condition) IOrderRatingDo
	Select(conds ...field.Expr) IOrderRatingDo
	Where(conds ...gen.Condition) IOrderRatingDo
	Order(conds ...field.Expr) IOrderRatingDo
	Distinct(cols ...field.Expr) IOrderRatingDo
	Omit(cols ...field.Expr) IOrderRatingDo
	Join(table schema.Tabler, on ...field.Expr) IOrderRatingDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IOrderRatingDo
	RightJoin(table schema.Tabler, on ...field.Expr) IOrderRatingDo
	Group(cols ...field.Expr) IOrderRatingDo
	Having(conds ...gen.Condition) IOrderRatingDo
	Limit(limit int) IOrderRatingDo
	Offset(offset int) IOrderRatingDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IOrderRatingDo
	Unscoped() IOrderRatingDo
	Create(values ...*model.OrderRating) error
	CreateInBatches(values []*model.OrderRating, batchSize int) error
	Save(values ...*model.OrderRating) error
	First() (*model.OrderRating, error)
	Take() (*model.OrderRating, error)
	Last() (*model.OrderRating, error)
	Find() ([]*model.OrderRating, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.OrderRating, err error)
	FindInBatches(result *[]*model.OrderRating, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.OrderRating) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IOrderRatingDo
	Assign(attrs ...field.AssignExpr) IOrderRatingDo
	Joins(fields ...field.RelationField) IOrderRatingDo
	Preload(fields ...field.RelationField) IOrderRatingDo
	FirstOrInit() (*model.OrderRating, error)
	FirstOrCreate() (*model.OrderRating, error)
	FindByPage(offset int, limit int) (result []*model.OrderRating, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IOrderRatingDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (o orderRatingDo) Debug() IOrderRatingDo {
	return o.withDO(o.DO.Debug())
}

func (o orderRatingDo) WithContext(ctx context.Context) IOrderRatingDo {
	return o.withDO(o.DO.WithContext(ctx))
}

func (o orderRatingDo) ReadDB() IOrderRatingDo {
	return o.Clauses(dbresolver.Read)
}

func (o orderRatingDo) WriteDB() IOrderRatingDo {
	return o.Clauses(dbresolver.Write)
}

func (o orderRatingDo) Session(config *gorm.Session) IOrderRatingDo {
	return o.withDO(o.DO.Session(config))
}

func (o orderRatingDo) Clauses(conds ...clause.Expression) IOrderRatingDo {
	return o.withDO(o.DO.Clauses(conds...))
}

func (o orderRatingDo) Returning(value interface{}, columns ...string) IOrderRatingDo {
	return o.withDO(o.DO.Returning(value, columns...))
}

func (o orderRatingDo) Not(conds ...gen.Condition) IOrderRatingDo {
	return o.withDO(o.DO.Not(conds...))
}

func (o orderRatingDo) Or(conds ...gen.Condition) IOrderRatingDo {
	return o.withDO(o.DO.Or(conds...))
}

func (o orderRatingDo) Select(conds ...field.Expr) IOrderRatingDo {
	return o.withDO(o.DO.Select(conds...))
}

func (o orderRatingDo) Where(conds ...gen.Condition) IOrderRatingDo {
	return o.withDO(o.DO.Where(conds...))
}

func (o orderRatingDo) Order(conds ...field.Expr) IOrderRatingDo {
	return o.withDO(o.DO.Order(conds...))
}

func (o orderRatingDo) Distinct(cols ...field.Expr) IOrderRatingDo {
	return o.withDO(o.DO.Distinct(cols...))
}

func (o orderRatingDo) Omit(cols ...field.Expr) IOrderRatingDo {
	return o.withDO(o.DO.Omit(cols...))
}

func (o orderRatingDo) Join(table schema.Tabler, on ...field.Expr) IOrderRatingDo {
	return o.withDO(o.DO.Join(table, on...))
}

func (o orderRatingDo) LeftJoin(table schema.Tabler, on ...field.Expr) IOrderRatingDo {
	return o.withDO(o.DO.LeftJoin(table, on...))
}

func (o orderRatingDo) RightJoin(table schema.Tabler, on ...field.Expr) IOrderRatingDo {
	return o.withDO(o.DO.RightJoin(table, on...))
}

func (o orderRatingDo) Group(cols ...field.Expr) IOrderRatingDo {
	return o.withDO(o.DO.Group(cols...))
}

func (o orderRatingDo) Having(conds ...gen.Condition) IOrderRatingDo {
	return o.withDO(o.DO.Having(conds...))
}

func (o orderRatingDo) Limit(limit int) IOrderRatingDo {
	return o.withDO(o.DO.Limit(limit))
}

func (o orderRatingDo) Offset(offset int) IOrderRatingDo {
	return o.withDO(o.DO.Offset(offset))
}

func (o orderRatingDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IOrderRatingDo {
	return o.withDO(o.DO.Scopes(funcs...))
}

func (o orderRatingDo) Unscoped() IOrderRatingDo {
	return o.withDO(o.DO.Unscoped())
}

func (o orderRatingDo) Create(values ...*model.OrderRating) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Create(values)
}

func (o orderRatingDo) CreateInBatches(values []*model.OrderRating, batchSize int) error {
	return o.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (o orderRatingDo) Save(values ...*model.OrderRating) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Save(values)
}

func (o orderRatingDo) First() (*model.OrderRating, error) {
	if result, err := o.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.OrderRating), nil
	}
}

func (o orderRatingDo) Take() (*model.OrderRating, error) {
	if result, err := o.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.OrderRating), nil
	}
}

func (o orderRatingDo) Last() (*model.OrderRating, error) {
	if result, err := o.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.OrderRating), nil
	}
}

func (o orderRatingDo) Find() ([]*model.OrderRating, error) {
	result, err := o.DO.Find()
	return result.([]*model.OrderRating), err
}

func (o orderRatingDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.OrderRating, err error) {
	buf := make([]*model.OrderRating, 0, batchSize)
	err = o.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (o orderRatingDo) FindInBatches(result *[]*model.OrderRating, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return o.DO.FindInBatches(result, batchSize, fc)
}

func (o orderRatingDo) Attrs(attrs ...field.AssignExpr) IOrderRatingDo {
	return o.withDO(o.DO.Attrs(attrs...))
}

func (o orderRatingDo) Assign(attrs ...field.AssignExpr) IOrderRatingDo {
	return o.withDO(o.DO.Assign(attrs...))
}

func (o orderRatingDo) Joins(fields ...field.RelationField) IOrderRatingDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Joins(_f))
	}
	return &o
}

func (o orderRatingDo) Preload(fields ...field.RelationField) IOrderRatingDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Preload(_f))
	}
	return &o
}

func (o orderRatingDo) FirstOrInit() (*model.OrderRating, error) {
	if result, err := o.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.OrderRating), nil
	}
}

func (o orderRatingDo) FirstOrCreate() (*model.OrderRating, error) {
	if result, err := o.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.OrderRating), nil
	}
}

func (o orderRatingDo) FindByPage(offset int, limit int) (result []*model.OrderRating, count int64, err error) {
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

func (o orderRatingDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = o.Count()
	if err != nil {
		return
	}

	err = o.Offset(offset).Limit(limit).Scan(result)
	return
}

func (o orderRatingDo) Scan(result interface{}) (err error) {
	return o.DO.Scan(result)
}

func (o orderRatingDo) Delete(models ...*model.OrderRating) (result gen.ResultInfo, err error) {
	return o.DO.Delete(models)
}

func (o *orderRatingDo) withDO(do gen.Dao) *orderRatingDo {
	o.DO = *do.(*gen.DO)
	return o
}
