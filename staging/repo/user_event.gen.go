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

func newUserEvent(db *gorm.DB, opts ...gen.DOOption) userEvent {
	_userEvent := userEvent{}

	_userEvent.userEventDo.UseDB(db, opts...)
	_userEvent.userEventDo.UseModel(&model.UserEvent{})

	tableName := _userEvent.userEventDo.TableName()
	_userEvent.ALL = field.NewAsterisk(tableName)
	_userEvent.ActorEventID = field.NewString(tableName, "actor_event_id")
	_userEvent.CustomerID = field.NewString(tableName, "customer_id")
	_userEvent.MerchantID = field.NewString(tableName, "merchant_id")
	_userEvent.EventName = field.NewString(tableName, "event_name")
	_userEvent.SecretCode = field.NewString(tableName, "secret_code")
	_userEvent.ExpiredAt = field.NewInt64(tableName, "expired_at")

	_userEvent.fillFieldMap()

	return _userEvent
}

type userEvent struct {
	userEventDo

	ALL          field.Asterisk
	ActorEventID field.String
	CustomerID   field.String
	MerchantID   field.String
	EventName    field.String
	SecretCode   field.String
	ExpiredAt    field.Int64

	fieldMap map[string]field.Expr
}

func (u userEvent) Table(newTableName string) *userEvent {
	u.userEventDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u userEvent) As(alias string) *userEvent {
	u.userEventDo.DO = *(u.userEventDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *userEvent) updateTableName(table string) *userEvent {
	u.ALL = field.NewAsterisk(table)
	u.ActorEventID = field.NewString(table, "actor_event_id")
	u.CustomerID = field.NewString(table, "customer_id")
	u.MerchantID = field.NewString(table, "merchant_id")
	u.EventName = field.NewString(table, "event_name")
	u.SecretCode = field.NewString(table, "secret_code")
	u.ExpiredAt = field.NewInt64(table, "expired_at")

	u.fillFieldMap()

	return u
}

func (u *userEvent) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *userEvent) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 6)
	u.fieldMap["actor_event_id"] = u.ActorEventID
	u.fieldMap["customer_id"] = u.CustomerID
	u.fieldMap["merchant_id"] = u.MerchantID
	u.fieldMap["event_name"] = u.EventName
	u.fieldMap["secret_code"] = u.SecretCode
	u.fieldMap["expired_at"] = u.ExpiredAt
}

func (u userEvent) clone(db *gorm.DB) userEvent {
	u.userEventDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u userEvent) replaceDB(db *gorm.DB) userEvent {
	u.userEventDo.ReplaceDB(db)
	return u
}

type userEventDo struct{ gen.DO }

type IUserEventDo interface {
	gen.SubQuery
	Debug() IUserEventDo
	WithContext(ctx context.Context) IUserEventDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUserEventDo
	WriteDB() IUserEventDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUserEventDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUserEventDo
	Not(conds ...gen.Condition) IUserEventDo
	Or(conds ...gen.Condition) IUserEventDo
	Select(conds ...field.Expr) IUserEventDo
	Where(conds ...gen.Condition) IUserEventDo
	Order(conds ...field.Expr) IUserEventDo
	Distinct(cols ...field.Expr) IUserEventDo
	Omit(cols ...field.Expr) IUserEventDo
	Join(table schema.Tabler, on ...field.Expr) IUserEventDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUserEventDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUserEventDo
	Group(cols ...field.Expr) IUserEventDo
	Having(conds ...gen.Condition) IUserEventDo
	Limit(limit int) IUserEventDo
	Offset(offset int) IUserEventDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUserEventDo
	Unscoped() IUserEventDo
	Create(values ...*model.UserEvent) error
	CreateInBatches(values []*model.UserEvent, batchSize int) error
	Save(values ...*model.UserEvent) error
	First() (*model.UserEvent, error)
	Take() (*model.UserEvent, error)
	Last() (*model.UserEvent, error)
	Find() ([]*model.UserEvent, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserEvent, err error)
	FindInBatches(result *[]*model.UserEvent, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.UserEvent) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUserEventDo
	Assign(attrs ...field.AssignExpr) IUserEventDo
	Joins(fields ...field.RelationField) IUserEventDo
	Preload(fields ...field.RelationField) IUserEventDo
	FirstOrInit() (*model.UserEvent, error)
	FirstOrCreate() (*model.UserEvent, error)
	FindByPage(offset int, limit int) (result []*model.UserEvent, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUserEventDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u userEventDo) Debug() IUserEventDo {
	return u.withDO(u.DO.Debug())
}

func (u userEventDo) WithContext(ctx context.Context) IUserEventDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userEventDo) ReadDB() IUserEventDo {
	return u.Clauses(dbresolver.Read)
}

func (u userEventDo) WriteDB() IUserEventDo {
	return u.Clauses(dbresolver.Write)
}

func (u userEventDo) Session(config *gorm.Session) IUserEventDo {
	return u.withDO(u.DO.Session(config))
}

func (u userEventDo) Clauses(conds ...clause.Expression) IUserEventDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userEventDo) Returning(value interface{}, columns ...string) IUserEventDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userEventDo) Not(conds ...gen.Condition) IUserEventDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userEventDo) Or(conds ...gen.Condition) IUserEventDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userEventDo) Select(conds ...field.Expr) IUserEventDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userEventDo) Where(conds ...gen.Condition) IUserEventDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userEventDo) Order(conds ...field.Expr) IUserEventDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userEventDo) Distinct(cols ...field.Expr) IUserEventDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userEventDo) Omit(cols ...field.Expr) IUserEventDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userEventDo) Join(table schema.Tabler, on ...field.Expr) IUserEventDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userEventDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUserEventDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userEventDo) RightJoin(table schema.Tabler, on ...field.Expr) IUserEventDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userEventDo) Group(cols ...field.Expr) IUserEventDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userEventDo) Having(conds ...gen.Condition) IUserEventDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userEventDo) Limit(limit int) IUserEventDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userEventDo) Offset(offset int) IUserEventDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userEventDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUserEventDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userEventDo) Unscoped() IUserEventDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userEventDo) Create(values ...*model.UserEvent) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userEventDo) CreateInBatches(values []*model.UserEvent, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userEventDo) Save(values ...*model.UserEvent) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userEventDo) First() (*model.UserEvent, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserEvent), nil
	}
}

func (u userEventDo) Take() (*model.UserEvent, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserEvent), nil
	}
}

func (u userEventDo) Last() (*model.UserEvent, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserEvent), nil
	}
}

func (u userEventDo) Find() ([]*model.UserEvent, error) {
	result, err := u.DO.Find()
	return result.([]*model.UserEvent), err
}

func (u userEventDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserEvent, err error) {
	buf := make([]*model.UserEvent, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userEventDo) FindInBatches(result *[]*model.UserEvent, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userEventDo) Attrs(attrs ...field.AssignExpr) IUserEventDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userEventDo) Assign(attrs ...field.AssignExpr) IUserEventDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userEventDo) Joins(fields ...field.RelationField) IUserEventDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userEventDo) Preload(fields ...field.RelationField) IUserEventDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userEventDo) FirstOrInit() (*model.UserEvent, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserEvent), nil
	}
}

func (u userEventDo) FirstOrCreate() (*model.UserEvent, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserEvent), nil
	}
}

func (u userEventDo) FindByPage(offset int, limit int) (result []*model.UserEvent, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u userEventDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userEventDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u userEventDo) Delete(models ...*model.UserEvent) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *userEventDo) withDO(do gen.Dao) *userEventDo {
	u.DO = *do.(*gen.DO)
	return u
}
