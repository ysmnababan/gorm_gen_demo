// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"gorm_demo/src/model"
)

func newNavigation(db *gorm.DB, opts ...gen.DOOption) navigation {
	_navigation := navigation{}

	_navigation.navigationDo.UseDB(db, opts...)
	_navigation.navigationDo.UseModel(&model.Navigation{})

	tableName := _navigation.navigationDo.TableName()
	_navigation.ALL = field.NewAsterisk(tableName)
	_navigation.CreatedAt = field.NewInt64(tableName, "created_at")
	_navigation.CreatedBy = field.NewString(tableName, "created_by")
	_navigation.ModifiedAt = field.NewInt64(tableName, "modified_at")
	_navigation.ModifiedBy = field.NewString(tableName, "modified_by")
	_navigation.DeletedAt = field.NewInt64(tableName, "deleted_at")
	_navigation.DeletedBy = field.NewString(tableName, "deleted_by")
	_navigation.NavigationID = field.NewString(tableName, "navigation_id")
	_navigation.NavigationName = field.NewString(tableName, "navigation_name")
	_navigation.ParentNavigationID = field.NewString(tableName, "parent_navigation_id")
	_navigation.SortOrder = field.NewInt32(tableName, "sort_order")
	_navigation.URLPath = field.NewString(tableName, "url_path")
	_navigation.IsActive = field.NewBool(tableName, "is_active")
	_navigation.RolesNav = navigationHasManyRolesNav{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("RolesNav", "model.RolesNavigation"),
	}

	_navigation.fillFieldMap()

	return _navigation
}

type navigation struct {
	navigationDo

	ALL                field.Asterisk
	CreatedAt          field.Int64
	CreatedBy          field.String
	ModifiedAt         field.Int64
	ModifiedBy         field.String
	DeletedAt          field.Int64
	DeletedBy          field.String
	NavigationID       field.String
	NavigationName     field.String
	ParentNavigationID field.String
	SortOrder          field.Int32
	URLPath            field.String
	IsActive           field.Bool
	RolesNav           navigationHasManyRolesNav

	fieldMap map[string]field.Expr
}

func (n navigation) Table(newTableName string) *navigation {
	n.navigationDo.UseTable(newTableName)
	return n.updateTableName(newTableName)
}

func (n navigation) As(alias string) *navigation {
	n.navigationDo.DO = *(n.navigationDo.As(alias).(*gen.DO))
	return n.updateTableName(alias)
}

func (n *navigation) updateTableName(table string) *navigation {
	n.ALL = field.NewAsterisk(table)
	n.CreatedAt = field.NewInt64(table, "created_at")
	n.CreatedBy = field.NewString(table, "created_by")
	n.ModifiedAt = field.NewInt64(table, "modified_at")
	n.ModifiedBy = field.NewString(table, "modified_by")
	n.DeletedAt = field.NewInt64(table, "deleted_at")
	n.DeletedBy = field.NewString(table, "deleted_by")
	n.NavigationID = field.NewString(table, "navigation_id")
	n.NavigationName = field.NewString(table, "navigation_name")
	n.ParentNavigationID = field.NewString(table, "parent_navigation_id")
	n.SortOrder = field.NewInt32(table, "sort_order")
	n.URLPath = field.NewString(table, "url_path")
	n.IsActive = field.NewBool(table, "is_active")

	n.fillFieldMap()

	return n
}

func (n *navigation) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := n.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (n *navigation) fillFieldMap() {
	n.fieldMap = make(map[string]field.Expr, 13)
	n.fieldMap["created_at"] = n.CreatedAt
	n.fieldMap["created_by"] = n.CreatedBy
	n.fieldMap["modified_at"] = n.ModifiedAt
	n.fieldMap["modified_by"] = n.ModifiedBy
	n.fieldMap["deleted_at"] = n.DeletedAt
	n.fieldMap["deleted_by"] = n.DeletedBy
	n.fieldMap["navigation_id"] = n.NavigationID
	n.fieldMap["navigation_name"] = n.NavigationName
	n.fieldMap["parent_navigation_id"] = n.ParentNavigationID
	n.fieldMap["sort_order"] = n.SortOrder
	n.fieldMap["url_path"] = n.URLPath
	n.fieldMap["is_active"] = n.IsActive

}

func (n navigation) clone(db *gorm.DB) navigation {
	n.navigationDo.ReplaceConnPool(db.Statement.ConnPool)
	return n
}

func (n navigation) replaceDB(db *gorm.DB) navigation {
	n.navigationDo.ReplaceDB(db)
	return n
}

type navigationHasManyRolesNav struct {
	db *gorm.DB

	field.RelationField
}

func (a navigationHasManyRolesNav) Where(conds ...field.Expr) *navigationHasManyRolesNav {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a navigationHasManyRolesNav) WithContext(ctx context.Context) *navigationHasManyRolesNav {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a navigationHasManyRolesNav) Session(session *gorm.Session) *navigationHasManyRolesNav {
	a.db = a.db.Session(session)
	return &a
}

func (a navigationHasManyRolesNav) Model(m *model.Navigation) *navigationHasManyRolesNavTx {
	return &navigationHasManyRolesNavTx{a.db.Model(m).Association(a.Name())}
}

type navigationHasManyRolesNavTx struct{ tx *gorm.Association }

func (a navigationHasManyRolesNavTx) Find() (result []*model.RolesNavigation, err error) {
	return result, a.tx.Find(&result)
}

func (a navigationHasManyRolesNavTx) Append(values ...*model.RolesNavigation) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a navigationHasManyRolesNavTx) Replace(values ...*model.RolesNavigation) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a navigationHasManyRolesNavTx) Delete(values ...*model.RolesNavigation) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a navigationHasManyRolesNavTx) Clear() error {
	return a.tx.Clear()
}

func (a navigationHasManyRolesNavTx) Count() int64 {
	return a.tx.Count()
}

type navigationDo struct{ gen.DO }

type INavigationDo interface {
	gen.SubQuery
	Debug() INavigationDo
	WithContext(ctx context.Context) INavigationDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() INavigationDo
	WriteDB() INavigationDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) INavigationDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) INavigationDo
	Not(conds ...gen.Condition) INavigationDo
	Or(conds ...gen.Condition) INavigationDo
	Select(conds ...field.Expr) INavigationDo
	Where(conds ...gen.Condition) INavigationDo
	Order(conds ...field.Expr) INavigationDo
	Distinct(cols ...field.Expr) INavigationDo
	Omit(cols ...field.Expr) INavigationDo
	Join(table schema.Tabler, on ...field.Expr) INavigationDo
	LeftJoin(table schema.Tabler, on ...field.Expr) INavigationDo
	RightJoin(table schema.Tabler, on ...field.Expr) INavigationDo
	Group(cols ...field.Expr) INavigationDo
	Having(conds ...gen.Condition) INavigationDo
	Limit(limit int) INavigationDo
	Offset(offset int) INavigationDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) INavigationDo
	Unscoped() INavigationDo
	Create(values ...*model.Navigation) error
	CreateInBatches(values []*model.Navigation, batchSize int) error
	Save(values ...*model.Navigation) error
	First() (*model.Navigation, error)
	Take() (*model.Navigation, error)
	Last() (*model.Navigation, error)
	Find() ([]*model.Navigation, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Navigation, err error)
	FindInBatches(result *[]*model.Navigation, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Navigation) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) INavigationDo
	Assign(attrs ...field.AssignExpr) INavigationDo
	Joins(fields ...field.RelationField) INavigationDo
	Preload(fields ...field.RelationField) INavigationDo
	FirstOrInit() (*model.Navigation, error)
	FirstOrCreate() (*model.Navigation, error)
	FindByPage(offset int, limit int) (result []*model.Navigation, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) INavigationDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (n navigationDo) Debug() INavigationDo {
	return n.withDO(n.DO.Debug())
}

func (n navigationDo) WithContext(ctx context.Context) INavigationDo {
	return n.withDO(n.DO.WithContext(ctx))
}

func (n navigationDo) ReadDB() INavigationDo {
	return n.Clauses(dbresolver.Read)
}

func (n navigationDo) WriteDB() INavigationDo {
	return n.Clauses(dbresolver.Write)
}

func (n navigationDo) Session(config *gorm.Session) INavigationDo {
	return n.withDO(n.DO.Session(config))
}

func (n navigationDo) Clauses(conds ...clause.Expression) INavigationDo {
	return n.withDO(n.DO.Clauses(conds...))
}

func (n navigationDo) Returning(value interface{}, columns ...string) INavigationDo {
	return n.withDO(n.DO.Returning(value, columns...))
}

func (n navigationDo) Not(conds ...gen.Condition) INavigationDo {
	return n.withDO(n.DO.Not(conds...))
}

func (n navigationDo) Or(conds ...gen.Condition) INavigationDo {
	return n.withDO(n.DO.Or(conds...))
}

func (n navigationDo) Select(conds ...field.Expr) INavigationDo {
	return n.withDO(n.DO.Select(conds...))
}

func (n navigationDo) Where(conds ...gen.Condition) INavigationDo {
	return n.withDO(n.DO.Where(conds...))
}

func (n navigationDo) Order(conds ...field.Expr) INavigationDo {
	return n.withDO(n.DO.Order(conds...))
}

func (n navigationDo) Distinct(cols ...field.Expr) INavigationDo {
	return n.withDO(n.DO.Distinct(cols...))
}

func (n navigationDo) Omit(cols ...field.Expr) INavigationDo {
	return n.withDO(n.DO.Omit(cols...))
}

func (n navigationDo) Join(table schema.Tabler, on ...field.Expr) INavigationDo {
	return n.withDO(n.DO.Join(table, on...))
}

func (n navigationDo) LeftJoin(table schema.Tabler, on ...field.Expr) INavigationDo {
	return n.withDO(n.DO.LeftJoin(table, on...))
}

func (n navigationDo) RightJoin(table schema.Tabler, on ...field.Expr) INavigationDo {
	return n.withDO(n.DO.RightJoin(table, on...))
}

func (n navigationDo) Group(cols ...field.Expr) INavigationDo {
	return n.withDO(n.DO.Group(cols...))
}

func (n navigationDo) Having(conds ...gen.Condition) INavigationDo {
	return n.withDO(n.DO.Having(conds...))
}

func (n navigationDo) Limit(limit int) INavigationDo {
	return n.withDO(n.DO.Limit(limit))
}

func (n navigationDo) Offset(offset int) INavigationDo {
	return n.withDO(n.DO.Offset(offset))
}

func (n navigationDo) Scopes(funcs ...func(gen.Dao) gen.Dao) INavigationDo {
	return n.withDO(n.DO.Scopes(funcs...))
}

func (n navigationDo) Unscoped() INavigationDo {
	return n.withDO(n.DO.Unscoped())
}

func (n navigationDo) Create(values ...*model.Navigation) error {
	if len(values) == 0 {
		return nil
	}
	return n.DO.Create(values)
}

func (n navigationDo) CreateInBatches(values []*model.Navigation, batchSize int) error {
	return n.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (n navigationDo) Save(values ...*model.Navigation) error {
	if len(values) == 0 {
		return nil
	}
	return n.DO.Save(values)
}

func (n navigationDo) First() (*model.Navigation, error) {
	if result, err := n.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Navigation), nil
	}
}

func (n navigationDo) Take() (*model.Navigation, error) {
	if result, err := n.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Navigation), nil
	}
}

func (n navigationDo) Last() (*model.Navigation, error) {
	if result, err := n.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Navigation), nil
	}
}

func (n navigationDo) Find() ([]*model.Navigation, error) {
	result, err := n.DO.Find()
	return result.([]*model.Navigation), err
}

func (n navigationDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Navigation, err error) {
	buf := make([]*model.Navigation, 0, batchSize)
	err = n.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (n navigationDo) FindInBatches(result *[]*model.Navigation, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return n.DO.FindInBatches(result, batchSize, fc)
}

func (n navigationDo) Attrs(attrs ...field.AssignExpr) INavigationDo {
	return n.withDO(n.DO.Attrs(attrs...))
}

func (n navigationDo) Assign(attrs ...field.AssignExpr) INavigationDo {
	return n.withDO(n.DO.Assign(attrs...))
}

func (n navigationDo) Joins(fields ...field.RelationField) INavigationDo {
	for _, _f := range fields {
		n = *n.withDO(n.DO.Joins(_f))
	}
	return &n
}

func (n navigationDo) Preload(fields ...field.RelationField) INavigationDo {
	for _, _f := range fields {
		n = *n.withDO(n.DO.Preload(_f))
	}
	return &n
}

func (n navigationDo) FirstOrInit() (*model.Navigation, error) {
	if result, err := n.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Navigation), nil
	}
}

func (n navigationDo) FirstOrCreate() (*model.Navigation, error) {
	if result, err := n.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Navigation), nil
	}
}

func (n navigationDo) FindByPage(offset int, limit int) (result []*model.Navigation, count int64, err error) {
	result, err = n.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = n.Offset(-1).Limit(-1).Count()
	return
}

func (n navigationDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = n.Count()
	if err != nil {
		return
	}

	err = n.Offset(offset).Limit(limit).Scan(result)
	return
}

func (n navigationDo) Scan(result interface{}) (err error) {
	return n.DO.Scan(result)
}

func (n navigationDo) Delete(models ...*model.Navigation) (result gen.ResultInfo, err error) {
	return n.DO.Delete(models)
}

func (n *navigationDo) withDO(do gen.Dao) *navigationDo {
	n.DO = *do.(*gen.DO)
	return n
}
