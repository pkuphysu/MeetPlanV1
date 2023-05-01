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

	"meetplan/biz/gorm_gen"
)

func newFriendLink(db *gorm.DB, opts ...gen.DOOption) friendLink {
	_friendLink := friendLink{}

	_friendLink.friendLinkDo.UseDB(db, opts...)
	_friendLink.friendLinkDo.UseModel(&gorm_gen.FriendLink{})

	tableName := _friendLink.friendLinkDo.TableName()
	_friendLink.ALL = field.NewAsterisk(tableName)
	_friendLink.ID = field.NewInt32(tableName, "id")
	_friendLink.Name = field.NewString(tableName, "name")
	_friendLink.URL = field.NewString(tableName, "url")
	_friendLink.Description = field.NewString(tableName, "description")

	_friendLink.fillFieldMap()

	return _friendLink
}

type friendLink struct {
	friendLinkDo

	ALL         field.Asterisk
	ID          field.Int32
	Name        field.String
	URL         field.String
	Description field.String

	fieldMap map[string]field.Expr
}

func (f friendLink) Table(newTableName string) *friendLink {
	f.friendLinkDo.UseTable(newTableName)
	return f.updateTableName(newTableName)
}

func (f friendLink) As(alias string) *friendLink {
	f.friendLinkDo.DO = *(f.friendLinkDo.As(alias).(*gen.DO))
	return f.updateTableName(alias)
}

func (f *friendLink) updateTableName(table string) *friendLink {
	f.ALL = field.NewAsterisk(table)
	f.ID = field.NewInt32(table, "id")
	f.Name = field.NewString(table, "name")
	f.URL = field.NewString(table, "url")
	f.Description = field.NewString(table, "description")

	f.fillFieldMap()

	return f
}

func (f *friendLink) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := f.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (f *friendLink) fillFieldMap() {
	f.fieldMap = make(map[string]field.Expr, 4)
	f.fieldMap["id"] = f.ID
	f.fieldMap["name"] = f.Name
	f.fieldMap["url"] = f.URL
	f.fieldMap["description"] = f.Description
}

func (f friendLink) clone(db *gorm.DB) friendLink {
	f.friendLinkDo.ReplaceConnPool(db.Statement.ConnPool)
	return f
}

func (f friendLink) replaceDB(db *gorm.DB) friendLink {
	f.friendLinkDo.ReplaceDB(db)
	return f
}

type friendLinkDo struct{ gen.DO }

type IFriendLinkDo interface {
	gen.SubQuery
	Debug() IFriendLinkDo
	WithContext(ctx context.Context) IFriendLinkDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IFriendLinkDo
	WriteDB() IFriendLinkDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IFriendLinkDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IFriendLinkDo
	Not(conds ...gen.Condition) IFriendLinkDo
	Or(conds ...gen.Condition) IFriendLinkDo
	Select(conds ...field.Expr) IFriendLinkDo
	Where(conds ...gen.Condition) IFriendLinkDo
	Order(conds ...field.Expr) IFriendLinkDo
	Distinct(cols ...field.Expr) IFriendLinkDo
	Omit(cols ...field.Expr) IFriendLinkDo
	Join(table schema.Tabler, on ...field.Expr) IFriendLinkDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IFriendLinkDo
	RightJoin(table schema.Tabler, on ...field.Expr) IFriendLinkDo
	Group(cols ...field.Expr) IFriendLinkDo
	Having(conds ...gen.Condition) IFriendLinkDo
	Limit(limit int) IFriendLinkDo
	Offset(offset int) IFriendLinkDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IFriendLinkDo
	Unscoped() IFriendLinkDo
	Create(values ...*gorm_gen.FriendLink) error
	CreateInBatches(values []*gorm_gen.FriendLink, batchSize int) error
	Save(values ...*gorm_gen.FriendLink) error
	First() (*gorm_gen.FriendLink, error)
	Take() (*gorm_gen.FriendLink, error)
	Last() (*gorm_gen.FriendLink, error)
	Find() ([]*gorm_gen.FriendLink, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*gorm_gen.FriendLink, err error)
	FindInBatches(result *[]*gorm_gen.FriendLink, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*gorm_gen.FriendLink) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IFriendLinkDo
	Assign(attrs ...field.AssignExpr) IFriendLinkDo
	Joins(fields ...field.RelationField) IFriendLinkDo
	Preload(fields ...field.RelationField) IFriendLinkDo
	FirstOrInit() (*gorm_gen.FriendLink, error)
	FirstOrCreate() (*gorm_gen.FriendLink, error)
	FindByPage(offset int, limit int) (result []*gorm_gen.FriendLink, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IFriendLinkDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (f friendLinkDo) Debug() IFriendLinkDo {
	return f.withDO(f.DO.Debug())
}

func (f friendLinkDo) WithContext(ctx context.Context) IFriendLinkDo {
	return f.withDO(f.DO.WithContext(ctx))
}

func (f friendLinkDo) ReadDB() IFriendLinkDo {
	return f.Clauses(dbresolver.Read)
}

func (f friendLinkDo) WriteDB() IFriendLinkDo {
	return f.Clauses(dbresolver.Write)
}

func (f friendLinkDo) Session(config *gorm.Session) IFriendLinkDo {
	return f.withDO(f.DO.Session(config))
}

func (f friendLinkDo) Clauses(conds ...clause.Expression) IFriendLinkDo {
	return f.withDO(f.DO.Clauses(conds...))
}

func (f friendLinkDo) Returning(value interface{}, columns ...string) IFriendLinkDo {
	return f.withDO(f.DO.Returning(value, columns...))
}

func (f friendLinkDo) Not(conds ...gen.Condition) IFriendLinkDo {
	return f.withDO(f.DO.Not(conds...))
}

func (f friendLinkDo) Or(conds ...gen.Condition) IFriendLinkDo {
	return f.withDO(f.DO.Or(conds...))
}

func (f friendLinkDo) Select(conds ...field.Expr) IFriendLinkDo {
	return f.withDO(f.DO.Select(conds...))
}

func (f friendLinkDo) Where(conds ...gen.Condition) IFriendLinkDo {
	return f.withDO(f.DO.Where(conds...))
}

func (f friendLinkDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IFriendLinkDo {
	return f.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (f friendLinkDo) Order(conds ...field.Expr) IFriendLinkDo {
	return f.withDO(f.DO.Order(conds...))
}

func (f friendLinkDo) Distinct(cols ...field.Expr) IFriendLinkDo {
	return f.withDO(f.DO.Distinct(cols...))
}

func (f friendLinkDo) Omit(cols ...field.Expr) IFriendLinkDo {
	return f.withDO(f.DO.Omit(cols...))
}

func (f friendLinkDo) Join(table schema.Tabler, on ...field.Expr) IFriendLinkDo {
	return f.withDO(f.DO.Join(table, on...))
}

func (f friendLinkDo) LeftJoin(table schema.Tabler, on ...field.Expr) IFriendLinkDo {
	return f.withDO(f.DO.LeftJoin(table, on...))
}

func (f friendLinkDo) RightJoin(table schema.Tabler, on ...field.Expr) IFriendLinkDo {
	return f.withDO(f.DO.RightJoin(table, on...))
}

func (f friendLinkDo) Group(cols ...field.Expr) IFriendLinkDo {
	return f.withDO(f.DO.Group(cols...))
}

func (f friendLinkDo) Having(conds ...gen.Condition) IFriendLinkDo {
	return f.withDO(f.DO.Having(conds...))
}

func (f friendLinkDo) Limit(limit int) IFriendLinkDo {
	return f.withDO(f.DO.Limit(limit))
}

func (f friendLinkDo) Offset(offset int) IFriendLinkDo {
	return f.withDO(f.DO.Offset(offset))
}

func (f friendLinkDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IFriendLinkDo {
	return f.withDO(f.DO.Scopes(funcs...))
}

func (f friendLinkDo) Unscoped() IFriendLinkDo {
	return f.withDO(f.DO.Unscoped())
}

func (f friendLinkDo) Create(values ...*gorm_gen.FriendLink) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Create(values)
}

func (f friendLinkDo) CreateInBatches(values []*gorm_gen.FriendLink, batchSize int) error {
	return f.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (f friendLinkDo) Save(values ...*gorm_gen.FriendLink) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Save(values)
}

func (f friendLinkDo) First() (*gorm_gen.FriendLink, error) {
	if result, err := f.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*gorm_gen.FriendLink), nil
	}
}

func (f friendLinkDo) Take() (*gorm_gen.FriendLink, error) {
	if result, err := f.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*gorm_gen.FriendLink), nil
	}
}

func (f friendLinkDo) Last() (*gorm_gen.FriendLink, error) {
	if result, err := f.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*gorm_gen.FriendLink), nil
	}
}

func (f friendLinkDo) Find() ([]*gorm_gen.FriendLink, error) {
	result, err := f.DO.Find()
	return result.([]*gorm_gen.FriendLink), err
}

func (f friendLinkDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*gorm_gen.FriendLink, err error) {
	buf := make([]*gorm_gen.FriendLink, 0, batchSize)
	err = f.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (f friendLinkDo) FindInBatches(result *[]*gorm_gen.FriendLink, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return f.DO.FindInBatches(result, batchSize, fc)
}

func (f friendLinkDo) Attrs(attrs ...field.AssignExpr) IFriendLinkDo {
	return f.withDO(f.DO.Attrs(attrs...))
}

func (f friendLinkDo) Assign(attrs ...field.AssignExpr) IFriendLinkDo {
	return f.withDO(f.DO.Assign(attrs...))
}

func (f friendLinkDo) Joins(fields ...field.RelationField) IFriendLinkDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Joins(_f))
	}
	return &f
}

func (f friendLinkDo) Preload(fields ...field.RelationField) IFriendLinkDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Preload(_f))
	}
	return &f
}

func (f friendLinkDo) FirstOrInit() (*gorm_gen.FriendLink, error) {
	if result, err := f.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*gorm_gen.FriendLink), nil
	}
}

func (f friendLinkDo) FirstOrCreate() (*gorm_gen.FriendLink, error) {
	if result, err := f.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*gorm_gen.FriendLink), nil
	}
}

func (f friendLinkDo) FindByPage(offset int, limit int) (result []*gorm_gen.FriendLink, count int64, err error) {
	result, err = f.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = f.Offset(-1).Limit(-1).Count()
	return
}

func (f friendLinkDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = f.Count()
	if err != nil {
		return
	}

	err = f.Offset(offset).Limit(limit).Scan(result)
	return
}

func (f friendLinkDo) Scan(result interface{}) (err error) {
	return f.DO.Scan(result)
}

func (f friendLinkDo) Delete(models ...*gorm_gen.FriendLink) (result gen.ResultInfo, err error) {
	return f.DO.Delete(models)
}

func (f *friendLinkDo) withDO(do gen.Dao) *friendLinkDo {
	f.DO = *do.(*gen.DO)
	return f
}