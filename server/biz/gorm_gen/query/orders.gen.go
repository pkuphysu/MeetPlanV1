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

func newOrder(db *gorm.DB, opts ...gen.DOOption) order {
	_order := order{}

	_order.orderDo.UseDB(db, opts...)
	_order.orderDo.UseModel(&gorm_gen.Order{})

	tableName := _order.orderDo.TableName()
	_order.ALL = field.NewAsterisk(tableName)
	_order.ID = field.NewInt64(tableName, "id")
	_order.Status = field.NewInt8(tableName, "status")
	_order.Message = field.NewString(tableName, "message")
	_order.PlanID = field.NewInt64(tableName, "plan_id")
	_order.StudentID = field.NewInt64(tableName, "student_id")
	_order.Student = orderBelongsToStudent{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Student", "gorm_gen.User"),
	}

	_order.Teacher = orderBelongsToTeacher{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Teacher", "gorm_gen.User"),
	}

	_order.Plan = orderBelongsToPlan{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Plan", "gorm_gen.Plan"),
	}

	_order.fillFieldMap()

	return _order
}

type order struct {
	orderDo

	ALL       field.Asterisk
	ID        field.Int64
	Status    field.Int8
	Message   field.String
	PlanID    field.Int64
	StudentID field.Int64
	Student   orderBelongsToStudent

	Teacher orderBelongsToTeacher

	Plan orderBelongsToPlan

	fieldMap map[string]field.Expr
}

func (o order) Table(newTableName string) *order {
	o.orderDo.UseTable(newTableName)
	return o.updateTableName(newTableName)
}

func (o order) As(alias string) *order {
	o.orderDo.DO = *(o.orderDo.As(alias).(*gen.DO))
	return o.updateTableName(alias)
}

func (o *order) updateTableName(table string) *order {
	o.ALL = field.NewAsterisk(table)
	o.ID = field.NewInt64(table, "id")
	o.Status = field.NewInt8(table, "status")
	o.Message = field.NewString(table, "message")
	o.PlanID = field.NewInt64(table, "plan_id")
	o.StudentID = field.NewInt64(table, "student_id")

	o.fillFieldMap()

	return o
}

func (o *order) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := o.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (o *order) fillFieldMap() {
	o.fieldMap = make(map[string]field.Expr, 8)
	o.fieldMap["id"] = o.ID
	o.fieldMap["status"] = o.Status
	o.fieldMap["message"] = o.Message
	o.fieldMap["plan_id"] = o.PlanID
	o.fieldMap["student_id"] = o.StudentID

}

func (o order) clone(db *gorm.DB) order {
	o.orderDo.ReplaceConnPool(db.Statement.ConnPool)
	return o
}

func (o order) replaceDB(db *gorm.DB) order {
	o.orderDo.ReplaceDB(db)
	return o
}

type orderBelongsToStudent struct {
	db *gorm.DB

	field.RelationField
}

func (a orderBelongsToStudent) Where(conds ...field.Expr) *orderBelongsToStudent {
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

func (a orderBelongsToStudent) WithContext(ctx context.Context) *orderBelongsToStudent {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a orderBelongsToStudent) Session(session *gorm.Session) *orderBelongsToStudent {
	a.db = a.db.Session(session)
	return &a
}

func (a orderBelongsToStudent) Model(m *gorm_gen.Order) *orderBelongsToStudentTx {
	return &orderBelongsToStudentTx{a.db.Model(m).Association(a.Name())}
}

type orderBelongsToStudentTx struct{ tx *gorm.Association }

func (a orderBelongsToStudentTx) Find() (result *gorm_gen.User, err error) {
	return result, a.tx.Find(&result)
}

func (a orderBelongsToStudentTx) Append(values ...*gorm_gen.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a orderBelongsToStudentTx) Replace(values ...*gorm_gen.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a orderBelongsToStudentTx) Delete(values ...*gorm_gen.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a orderBelongsToStudentTx) Clear() error {
	return a.tx.Clear()
}

func (a orderBelongsToStudentTx) Count() int64 {
	return a.tx.Count()
}

type orderBelongsToTeacher struct {
	db *gorm.DB

	field.RelationField
}

func (a orderBelongsToTeacher) Where(conds ...field.Expr) *orderBelongsToTeacher {
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

func (a orderBelongsToTeacher) WithContext(ctx context.Context) *orderBelongsToTeacher {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a orderBelongsToTeacher) Session(session *gorm.Session) *orderBelongsToTeacher {
	a.db = a.db.Session(session)
	return &a
}

func (a orderBelongsToTeacher) Model(m *gorm_gen.Order) *orderBelongsToTeacherTx {
	return &orderBelongsToTeacherTx{a.db.Model(m).Association(a.Name())}
}

type orderBelongsToTeacherTx struct{ tx *gorm.Association }

func (a orderBelongsToTeacherTx) Find() (result *gorm_gen.User, err error) {
	return result, a.tx.Find(&result)
}

func (a orderBelongsToTeacherTx) Append(values ...*gorm_gen.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a orderBelongsToTeacherTx) Replace(values ...*gorm_gen.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a orderBelongsToTeacherTx) Delete(values ...*gorm_gen.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a orderBelongsToTeacherTx) Clear() error {
	return a.tx.Clear()
}

func (a orderBelongsToTeacherTx) Count() int64 {
	return a.tx.Count()
}

type orderBelongsToPlan struct {
	db *gorm.DB

	field.RelationField
}

func (a orderBelongsToPlan) Where(conds ...field.Expr) *orderBelongsToPlan {
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

func (a orderBelongsToPlan) WithContext(ctx context.Context) *orderBelongsToPlan {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a orderBelongsToPlan) Session(session *gorm.Session) *orderBelongsToPlan {
	a.db = a.db.Session(session)
	return &a
}

func (a orderBelongsToPlan) Model(m *gorm_gen.Order) *orderBelongsToPlanTx {
	return &orderBelongsToPlanTx{a.db.Model(m).Association(a.Name())}
}

type orderBelongsToPlanTx struct{ tx *gorm.Association }

func (a orderBelongsToPlanTx) Find() (result *gorm_gen.Plan, err error) {
	return result, a.tx.Find(&result)
}

func (a orderBelongsToPlanTx) Append(values ...*gorm_gen.Plan) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a orderBelongsToPlanTx) Replace(values ...*gorm_gen.Plan) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a orderBelongsToPlanTx) Delete(values ...*gorm_gen.Plan) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a orderBelongsToPlanTx) Clear() error {
	return a.tx.Clear()
}

func (a orderBelongsToPlanTx) Count() int64 {
	return a.tx.Count()
}

type orderDo struct{ gen.DO }

type IOrderDo interface {
	gen.SubQuery
	Debug() IOrderDo
	WithContext(ctx context.Context) IOrderDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IOrderDo
	WriteDB() IOrderDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IOrderDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IOrderDo
	Not(conds ...gen.Condition) IOrderDo
	Or(conds ...gen.Condition) IOrderDo
	Select(conds ...field.Expr) IOrderDo
	Where(conds ...gen.Condition) IOrderDo
	Order(conds ...field.Expr) IOrderDo
	Distinct(cols ...field.Expr) IOrderDo
	Omit(cols ...field.Expr) IOrderDo
	Join(table schema.Tabler, on ...field.Expr) IOrderDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IOrderDo
	RightJoin(table schema.Tabler, on ...field.Expr) IOrderDo
	Group(cols ...field.Expr) IOrderDo
	Having(conds ...gen.Condition) IOrderDo
	Limit(limit int) IOrderDo
	Offset(offset int) IOrderDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IOrderDo
	Unscoped() IOrderDo
	Create(values ...*gorm_gen.Order) error
	CreateInBatches(values []*gorm_gen.Order, batchSize int) error
	Save(values ...*gorm_gen.Order) error
	First() (*gorm_gen.Order, error)
	Take() (*gorm_gen.Order, error)
	Last() (*gorm_gen.Order, error)
	Find() ([]*gorm_gen.Order, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*gorm_gen.Order, err error)
	FindInBatches(result *[]*gorm_gen.Order, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*gorm_gen.Order) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IOrderDo
	Assign(attrs ...field.AssignExpr) IOrderDo
	Joins(fields ...field.RelationField) IOrderDo
	Preload(fields ...field.RelationField) IOrderDo
	FirstOrInit() (*gorm_gen.Order, error)
	FirstOrCreate() (*gorm_gen.Order, error)
	FindByPage(offset int, limit int) (result []*gorm_gen.Order, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IOrderDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (o orderDo) Debug() IOrderDo {
	return o.withDO(o.DO.Debug())
}

func (o orderDo) WithContext(ctx context.Context) IOrderDo {
	return o.withDO(o.DO.WithContext(ctx))
}

func (o orderDo) ReadDB() IOrderDo {
	return o.Clauses(dbresolver.Read)
}

func (o orderDo) WriteDB() IOrderDo {
	return o.Clauses(dbresolver.Write)
}

func (o orderDo) Session(config *gorm.Session) IOrderDo {
	return o.withDO(o.DO.Session(config))
}

func (o orderDo) Clauses(conds ...clause.Expression) IOrderDo {
	return o.withDO(o.DO.Clauses(conds...))
}

func (o orderDo) Returning(value interface{}, columns ...string) IOrderDo {
	return o.withDO(o.DO.Returning(value, columns...))
}

func (o orderDo) Not(conds ...gen.Condition) IOrderDo {
	return o.withDO(o.DO.Not(conds...))
}

func (o orderDo) Or(conds ...gen.Condition) IOrderDo {
	return o.withDO(o.DO.Or(conds...))
}

func (o orderDo) Select(conds ...field.Expr) IOrderDo {
	return o.withDO(o.DO.Select(conds...))
}

func (o orderDo) Where(conds ...gen.Condition) IOrderDo {
	return o.withDO(o.DO.Where(conds...))
}

func (o orderDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IOrderDo {
	return o.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (o orderDo) Order(conds ...field.Expr) IOrderDo {
	return o.withDO(o.DO.Order(conds...))
}

func (o orderDo) Distinct(cols ...field.Expr) IOrderDo {
	return o.withDO(o.DO.Distinct(cols...))
}

func (o orderDo) Omit(cols ...field.Expr) IOrderDo {
	return o.withDO(o.DO.Omit(cols...))
}

func (o orderDo) Join(table schema.Tabler, on ...field.Expr) IOrderDo {
	return o.withDO(o.DO.Join(table, on...))
}

func (o orderDo) LeftJoin(table schema.Tabler, on ...field.Expr) IOrderDo {
	return o.withDO(o.DO.LeftJoin(table, on...))
}

func (o orderDo) RightJoin(table schema.Tabler, on ...field.Expr) IOrderDo {
	return o.withDO(o.DO.RightJoin(table, on...))
}

func (o orderDo) Group(cols ...field.Expr) IOrderDo {
	return o.withDO(o.DO.Group(cols...))
}

func (o orderDo) Having(conds ...gen.Condition) IOrderDo {
	return o.withDO(o.DO.Having(conds...))
}

func (o orderDo) Limit(limit int) IOrderDo {
	return o.withDO(o.DO.Limit(limit))
}

func (o orderDo) Offset(offset int) IOrderDo {
	return o.withDO(o.DO.Offset(offset))
}

func (o orderDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IOrderDo {
	return o.withDO(o.DO.Scopes(funcs...))
}

func (o orderDo) Unscoped() IOrderDo {
	return o.withDO(o.DO.Unscoped())
}

func (o orderDo) Create(values ...*gorm_gen.Order) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Create(values)
}

func (o orderDo) CreateInBatches(values []*gorm_gen.Order, batchSize int) error {
	return o.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (o orderDo) Save(values ...*gorm_gen.Order) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Save(values)
}

func (o orderDo) First() (*gorm_gen.Order, error) {
	if result, err := o.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*gorm_gen.Order), nil
	}
}

func (o orderDo) Take() (*gorm_gen.Order, error) {
	if result, err := o.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*gorm_gen.Order), nil
	}
}

func (o orderDo) Last() (*gorm_gen.Order, error) {
	if result, err := o.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*gorm_gen.Order), nil
	}
}

func (o orderDo) Find() ([]*gorm_gen.Order, error) {
	result, err := o.DO.Find()
	return result.([]*gorm_gen.Order), err
}

func (o orderDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*gorm_gen.Order, err error) {
	buf := make([]*gorm_gen.Order, 0, batchSize)
	err = o.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (o orderDo) FindInBatches(result *[]*gorm_gen.Order, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return o.DO.FindInBatches(result, batchSize, fc)
}

func (o orderDo) Attrs(attrs ...field.AssignExpr) IOrderDo {
	return o.withDO(o.DO.Attrs(attrs...))
}

func (o orderDo) Assign(attrs ...field.AssignExpr) IOrderDo {
	return o.withDO(o.DO.Assign(attrs...))
}

func (o orderDo) Joins(fields ...field.RelationField) IOrderDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Joins(_f))
	}
	return &o
}

func (o orderDo) Preload(fields ...field.RelationField) IOrderDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Preload(_f))
	}
	return &o
}

func (o orderDo) FirstOrInit() (*gorm_gen.Order, error) {
	if result, err := o.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*gorm_gen.Order), nil
	}
}

func (o orderDo) FirstOrCreate() (*gorm_gen.Order, error) {
	if result, err := o.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*gorm_gen.Order), nil
	}
}

func (o orderDo) FindByPage(offset int, limit int) (result []*gorm_gen.Order, count int64, err error) {
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

func (o orderDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = o.Count()
	if err != nil {
		return
	}

	err = o.Offset(offset).Limit(limit).Scan(result)
	return
}

func (o orderDo) Scan(result interface{}) (err error) {
	return o.DO.Scan(result)
}

func (o orderDo) Delete(models ...*gorm_gen.Order) (result gen.ResultInfo, err error) {
	return o.DO.Delete(models)
}

func (o *orderDo) withDO(do gen.Dao) *orderDo {
	o.DO = *do.(*gen.DO)
	return o
}
