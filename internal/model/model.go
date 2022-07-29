package model

import (
	"fmt"

	"BlogService/global"
	"BlogService/pkg/setting"

	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Model struct {
	gorm.Model
	CreateBy string `json:"create_by" gorm:"type:varchar(100); default:''; comment:'创建人'"`
	UpdateBy string `json:"update_by" gorm:"type:varchar(100); default:''; comment:'修改人'"`
}

// NewDBEngine 连接数据库软件并创建一个数据库引擎
func NewDBEngine(database setting.DB) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		database.Username,
		database.Password,
		database.Address,
		database.Database,
		database.Charset,
		database.ParseTime,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.Wrap(err, "database: unable to connect to the database")
	}
	sqlDB, _ := db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(database.MaxIdleConns)
	// SetMaxOpenConns 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(database.MaxOpenConns)

	//_ = db.Use(&OpentracingPlugin{})
	return db, nil
}

// MigrateDB 迁移数据表
func MigrateDB() error {
	err := global.DBEngine.AutoMigrate(&Tag{}, &Article{}, &ArticleTag{}, &Auth{})
	return errors.Wrap(err, "database: failed to migration schema")
}

//const (
//	parentSpanGormKey  = "opentracing:parent.span"
//	gormSpanKey        = "opentracing:span"
//	callBackBeforeName = "opentracing:before"
//	callBackAfterName  = "opentracing:after"
//)
//
//func WithContext(ctx context.Context, db *gorm.DB) *gorm.DB {
//	if ctx == nil {
//		return db
//	}
//	parentSpan := opentracing.SpanFromContext(ctx)
//	if parentSpan == nil {
//		return db
//	}
//	return db.Set(parentSpanGormKey, parentSpan)
//}
//
//// AddGormCallbacks adds OpentracingPlugin for tracing, you should call SetSpanToGorm to make them work
//func AddGormCallbacks(db *gorm.DB) {
//	callbacks := newCallbacks()
//	registerCallbacks(db, "create", callbacks)
//	registerCallbacks(db, "query", callbacks)
//	registerCallbacks(db, "update", callbacks)
//	registerCallbacks(db, "delete", callbacks)
//	registerCallbacks(db, "row_query", callbacks)
//}
//
//func newCallbacks() *OpentracingPlugin {
//	return &OpentracingPlugin{}
//}
//
//func (op *OpentracingPlugin) beforeCreate(scope *gorm.Scope)   { op.before(scope) }
//func (op *OpentracingPlugin) afterCreate(scope *gorm.Scope)    { op.after(scope, "INSERT") }
//func (op *OpentracingPlugin) beforeQuery(scope *gorm.Scope)    { op.before(scope) }
//func (op *OpentracingPlugin) afterQuery(scope *gorm.Scope)     { op.after(scope, "SELECT") }
//func (op *OpentracingPlugin) beforeUpdate(scope *gorm.Scope)   { op.before(scope) }
//func (op *OpentracingPlugin) afterUpdate(scope *gorm.Scope)    { op.after(scope, "UPDATE") }
//func (op *OpentracingPlugin) beforeDelete(scope *gorm.Scope)   { op.before(scope) }
//func (op *OpentracingPlugin) afterDelete(scope *gorm.Scope)    { op.after(scope, "DELETE") }
//func (op *OpentracingPlugin) beforeRowQuery(scope *gorm.Scope) { op.before(scope) }
//func (op *OpentracingPlugin) afterRowQuery(scope *gorm.Scope)  { op.after(scope, "") }
//
//func before(db *gorm.DB) {
//	span, _ := opentracing.StartSpanFromContext(db.Statement.Context, "gorm")
//	db.InstanceSet(gormSpanKey, span)
//	return
//}
//
//func after(db *gorm.DB) {
//	gormSpan, isExist := db.InstanceGet(gormSpanKey)
//	if !isExist {
//		return
//	}
//
//	span, ok := gormSpan.(opentracing.Span)
//	if !ok {
//		return
//	}
//	defer span.Finish()
//	if db.Error != nil {
//		span.LogFields(tracerLog.Error(db.Error))
//	}
//	span.LogFields(tracerLog.String("sql", db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...)))
//	return
//}
//
//type OpentracingPlugin struct{}
//
//func (op *OpentracingPlugin) Name() string {
//	return "opentracingPlugin"
//}
//
//func (op *OpentracingPlugin) Initialize(db *gorm.DB) error {
//	// 开始前 - 并不是都用相同的方法，可以自己自定义
//	db.Callback().Create().Before("gorm:create").Register(callBackBeforeName, before)
//	db.Callback().Query().Before("gorm:query").Register(callBackBeforeName, before)
//	db.Callback().Delete().Before("gorm:delete").Register(callBackBeforeName, before)
//	db.Callback().Update().Before("gorm:update").Register(callBackBeforeName, before)
//	db.Callback().Row().Before("gorm:row_query").Register(callBackBeforeName, before)
//
//	// 结束后 - 并不是都用相同的方法，可以自己自定义
//	db.Callback().Create().After("gorm:create").Register(callBackAfterName, after)
//	db.Callback().Query().After("gorm:query").Register(callBackAfterName, after)
//	db.Callback().Delete().After("gorm:delete").Register(callBackAfterName, after)
//	db.Callback().Update().After("gorm:update").Register(callBackAfterName, after)
//	db.Callback().Row().After("gorm:row_query").Register(callBackAfterName, after)
//
//	return nil
//}
//
//var _ gorm.Plugin = &OpentracingPlugin{}
