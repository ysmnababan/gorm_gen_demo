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

func newActorAdmin(db *gorm.DB, opts ...gen.DOOption) actorAdmin {
	_actorAdmin := actorAdmin{}

	_actorAdmin.actorAdminDo.UseDB(db, opts...)
	_actorAdmin.actorAdminDo.UseModel(&model.ActorAdmin{})

	tableName := _actorAdmin.actorAdminDo.TableName()
	_actorAdmin.ALL = field.NewAsterisk(tableName)
	_actorAdmin.CreatedAt = field.NewInt64(tableName, "created_at")
	_actorAdmin.CreatedBy = field.NewString(tableName, "created_by")
	_actorAdmin.ModifiedAt = field.NewInt64(tableName, "modified_at")
	_actorAdmin.ModifiedBy = field.NewString(tableName, "modified_by")
	_actorAdmin.DeletedAt = field.NewInt64(tableName, "deleted_at")
	_actorAdmin.DeletedBy = field.NewString(tableName, "deleted_by")
	_actorAdmin.ActorAdminID = field.NewString(tableName, "actor_admin_id")
	_actorAdmin.ActorName = field.NewString(tableName, "actor_name")
	_actorAdmin.Email = field.NewString(tableName, "email")
	_actorAdmin.Msisdn = field.NewString(tableName, "msisdn")
	_actorAdmin.Password = field.NewString(tableName, "password")
	_actorAdmin.RoleID = field.NewString(tableName, "role_id")
	_actorAdmin.MerchantID = field.NewString(tableName, "merchant_id")
	_actorAdmin.IsActive = field.NewBool(tableName, "is_active")

	_actorAdmin.fillFieldMap()

	return _actorAdmin
}

type actorAdmin struct {
	actorAdminDo

	ALL          field.Asterisk
	CreatedAt    field.Int64
	CreatedBy    field.String
	ModifiedAt   field.Int64
	ModifiedBy   field.String
	DeletedAt    field.Int64
	DeletedBy    field.String
	ActorAdminID field.String
	ActorName    field.String
	Email        field.String
	Msisdn       field.String
	Password     field.String
	RoleID       field.String
	MerchantID   field.String
	IsActive     field.Bool

	fieldMap map[string]field.Expr
}

func (a actorAdmin) Table(newTableName string) *actorAdmin {
	a.actorAdminDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a actorAdmin) As(alias string) *actorAdmin {
	a.actorAdminDo.DO = *(a.actorAdminDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *actorAdmin) updateTableName(table string) *actorAdmin {
	a.ALL = field.NewAsterisk(table)
	a.CreatedAt = field.NewInt64(table, "created_at")
	a.CreatedBy = field.NewString(table, "created_by")
	a.ModifiedAt = field.NewInt64(table, "modified_at")
	a.ModifiedBy = field.NewString(table, "modified_by")
	a.DeletedAt = field.NewInt64(table, "deleted_at")
	a.DeletedBy = field.NewString(table, "deleted_by")
	a.ActorAdminID = field.NewString(table, "actor_admin_id")
	a.ActorName = field.NewString(table, "actor_name")
	a.Email = field.NewString(table, "email")
	a.Msisdn = field.NewString(table, "msisdn")
	a.Password = field.NewString(table, "password")
	a.RoleID = field.NewString(table, "role_id")
	a.MerchantID = field.NewString(table, "merchant_id")
	a.IsActive = field.NewBool(table, "is_active")

	a.fillFieldMap()

	return a
}

func (a *actorAdmin) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *actorAdmin) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 14)
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["created_by"] = a.CreatedBy
	a.fieldMap["modified_at"] = a.ModifiedAt
	a.fieldMap["modified_by"] = a.ModifiedBy
	a.fieldMap["deleted_at"] = a.DeletedAt
	a.fieldMap["deleted_by"] = a.DeletedBy
	a.fieldMap["actor_admin_id"] = a.ActorAdminID
	a.fieldMap["actor_name"] = a.ActorName
	a.fieldMap["email"] = a.Email
	a.fieldMap["msisdn"] = a.Msisdn
	a.fieldMap["password"] = a.Password
	a.fieldMap["role_id"] = a.RoleID
	a.fieldMap["merchant_id"] = a.MerchantID
	a.fieldMap["is_active"] = a.IsActive
}

func (a actorAdmin) clone(db *gorm.DB) actorAdmin {
	a.actorAdminDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a actorAdmin) replaceDB(db *gorm.DB) actorAdmin {
	a.actorAdminDo.ReplaceDB(db)
	return a
}

type actorAdminDo struct{ gen.DO }

type IActorAdminDo interface {
	gen.SubQuery
	Debug() IActorAdminDo
	WithContext(ctx context.Context) IActorAdminDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IActorAdminDo
	WriteDB() IActorAdminDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IActorAdminDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IActorAdminDo
	Not(conds ...gen.Condition) IActorAdminDo
	Or(conds ...gen.Condition) IActorAdminDo
	Select(conds ...field.Expr) IActorAdminDo
	Where(conds ...gen.Condition) IActorAdminDo
	Order(conds ...field.Expr) IActorAdminDo
	Distinct(cols ...field.Expr) IActorAdminDo
	Omit(cols ...field.Expr) IActorAdminDo
	Join(table schema.Tabler, on ...field.Expr) IActorAdminDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IActorAdminDo
	RightJoin(table schema.Tabler, on ...field.Expr) IActorAdminDo
	Group(cols ...field.Expr) IActorAdminDo
	Having(conds ...gen.Condition) IActorAdminDo
	Limit(limit int) IActorAdminDo
	Offset(offset int) IActorAdminDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IActorAdminDo
	Unscoped() IActorAdminDo
	Create(values ...*model.ActorAdmin) error
	CreateInBatches(values []*model.ActorAdmin, batchSize int) error
	Save(values ...*model.ActorAdmin) error
	First() (*model.ActorAdmin, error)
	Take() (*model.ActorAdmin, error)
	Last() (*model.ActorAdmin, error)
	Find() ([]*model.ActorAdmin, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ActorAdmin, err error)
	FindInBatches(result *[]*model.ActorAdmin, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.ActorAdmin) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IActorAdminDo
	Assign(attrs ...field.AssignExpr) IActorAdminDo
	Joins(fields ...field.RelationField) IActorAdminDo
	Preload(fields ...field.RelationField) IActorAdminDo
	FirstOrInit() (*model.ActorAdmin, error)
	FirstOrCreate() (*model.ActorAdmin, error)
	FindByPage(offset int, limit int) (result []*model.ActorAdmin, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IActorAdminDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a actorAdminDo) Debug() IActorAdminDo {
	return a.withDO(a.DO.Debug())
}

func (a actorAdminDo) WithContext(ctx context.Context) IActorAdminDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a actorAdminDo) ReadDB() IActorAdminDo {
	return a.Clauses(dbresolver.Read)
}

func (a actorAdminDo) WriteDB() IActorAdminDo {
	return a.Clauses(dbresolver.Write)
}

func (a actorAdminDo) Session(config *gorm.Session) IActorAdminDo {
	return a.withDO(a.DO.Session(config))
}

func (a actorAdminDo) Clauses(conds ...clause.Expression) IActorAdminDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a actorAdminDo) Returning(value interface{}, columns ...string) IActorAdminDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a actorAdminDo) Not(conds ...gen.Condition) IActorAdminDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a actorAdminDo) Or(conds ...gen.Condition) IActorAdminDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a actorAdminDo) Select(conds ...field.Expr) IActorAdminDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a actorAdminDo) Where(conds ...gen.Condition) IActorAdminDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a actorAdminDo) Order(conds ...field.Expr) IActorAdminDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a actorAdminDo) Distinct(cols ...field.Expr) IActorAdminDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a actorAdminDo) Omit(cols ...field.Expr) IActorAdminDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a actorAdminDo) Join(table schema.Tabler, on ...field.Expr) IActorAdminDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a actorAdminDo) LeftJoin(table schema.Tabler, on ...field.Expr) IActorAdminDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a actorAdminDo) RightJoin(table schema.Tabler, on ...field.Expr) IActorAdminDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a actorAdminDo) Group(cols ...field.Expr) IActorAdminDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a actorAdminDo) Having(conds ...gen.Condition) IActorAdminDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a actorAdminDo) Limit(limit int) IActorAdminDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a actorAdminDo) Offset(offset int) IActorAdminDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a actorAdminDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IActorAdminDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a actorAdminDo) Unscoped() IActorAdminDo {
	return a.withDO(a.DO.Unscoped())
}

func (a actorAdminDo) Create(values ...*model.ActorAdmin) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a actorAdminDo) CreateInBatches(values []*model.ActorAdmin, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a actorAdminDo) Save(values ...*model.ActorAdmin) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a actorAdminDo) First() (*model.ActorAdmin, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ActorAdmin), nil
	}
}

func (a actorAdminDo) Take() (*model.ActorAdmin, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ActorAdmin), nil
	}
}

func (a actorAdminDo) Last() (*model.ActorAdmin, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ActorAdmin), nil
	}
}

func (a actorAdminDo) Find() ([]*model.ActorAdmin, error) {
	result, err := a.DO.Find()
	return result.([]*model.ActorAdmin), err
}

func (a actorAdminDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ActorAdmin, err error) {
	buf := make([]*model.ActorAdmin, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a actorAdminDo) FindInBatches(result *[]*model.ActorAdmin, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a actorAdminDo) Attrs(attrs ...field.AssignExpr) IActorAdminDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a actorAdminDo) Assign(attrs ...field.AssignExpr) IActorAdminDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a actorAdminDo) Joins(fields ...field.RelationField) IActorAdminDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a actorAdminDo) Preload(fields ...field.RelationField) IActorAdminDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a actorAdminDo) FirstOrInit() (*model.ActorAdmin, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ActorAdmin), nil
	}
}

func (a actorAdminDo) FirstOrCreate() (*model.ActorAdmin, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ActorAdmin), nil
	}
}

func (a actorAdminDo) FindByPage(offset int, limit int) (result []*model.ActorAdmin, count int64, err error) {
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

func (a actorAdminDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a actorAdminDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a actorAdminDo) Delete(models ...*model.ActorAdmin) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *actorAdminDo) withDO(do gen.Dao) *actorAdminDo {
	a.DO = *do.(*gen.DO)
	return a
}
