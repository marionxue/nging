// @generated Do not edit this file, which is automatically generated by the generator.

package dbschema

import (
	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/factory"
	
	"time"
)

type DbAccount struct {
	param   *factory.Param
	trans	*factory.Transaction
	objects []*DbAccount
	namer   func(string) string
	connID  int
	
	Id      	uint    	`db:"id,omitempty,pk" bson:"id,omitempty" comment:"ID" json:"id" xml:"id"`
	Title   	string  	`db:"title" bson:"title" comment:"标题" json:"title" xml:"title"`
	Uid     	uint    	`db:"uid" bson:"uid" comment:"UID" json:"uid" xml:"uid"`
	Engine  	string  	`db:"engine" bson:"engine" comment:"数据库引擎" json:"engine" xml:"engine"`
	Host    	string  	`db:"host" bson:"host" comment:"服务器地址" json:"host" xml:"host"`
	User    	string  	`db:"user" bson:"user" comment:"用户名" json:"user" xml:"user"`
	Password	string  	`db:"password" bson:"password" comment:"密码" json:"password" xml:"password"`
	Name    	string  	`db:"name" bson:"name" comment:"数据库名称" json:"name" xml:"name"`
	Options 	string  	`db:"options" bson:"options" comment:"其它选项(JSON)" json:"options" xml:"options"`
	Created 	uint    	`db:"created" bson:"created" comment:"创建时间" json:"created" xml:"created"`
	Updated 	uint    	`db:"updated" bson:"updated" comment:"修改时间" json:"updated" xml:"updated"`
}

func (this *DbAccount) Trans() *factory.Transaction {
	return this.trans
}

func (this *DbAccount) Use(trans *factory.Transaction) factory.Model {
	this.trans = trans
	return this
}

func (this *DbAccount) SetConnID(connID int) factory.Model {
	this.connID = connID
	return this
}

func (this *DbAccount) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName,connID[0]).Use(this.trans)
	}
	return factory.NewModel(structName,this.connID).Use(this.trans)
}

func (this *DbAccount) Objects() []*DbAccount {
	if this.objects == nil {
		return nil
	}
	return this.objects[:]
}

func (this *DbAccount) NewObjects() *[]*DbAccount {
	this.objects = []*DbAccount{}
	return &this.objects
}

func (this *DbAccount) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(this.connID).SetTrans(this.trans).SetCollection(this.Name_()).SetModel(this)
}

func (this *DbAccount) SetNamer(namer func (string) string) factory.Model {
	this.namer = namer
	return this
}

func (this *DbAccount) Name_() string {
	if this.namer != nil {
		return this.namer("db_account")
	}
	return factory.TableNamerGet("db_account")(this)
}

func (this *DbAccount) FullName_(connID ...int) string {
	if len(connID) > 0 {
		return factory.DefaultFactory.Cluster(connID[0]).Table(this.Name_())
	}
	return factory.DefaultFactory.Cluster(this.connID).Table(this.Name_())
}

func (this *DbAccount) SetParam(param *factory.Param) factory.Model {
	this.param = param
	return this
}

func (this *DbAccount) Param() *factory.Param {
	if this.param == nil {
		return this.NewParam()
	}
	return this.param
}

func (this *DbAccount) Get(mw func(db.Result) db.Result, args ...interface{}) error {
	return this.Param().SetArgs(args...).SetRecv(this).SetMiddleware(mw).One()
}

func (this *DbAccount) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.NewObjects()
	}
	return this.Param().SetArgs(args...).SetPage(page).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *DbAccount) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.NewObjects()
	}
	return this.Param().SetArgs(args...).SetOffset(offset).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *DbAccount) Add() (pk interface{}, err error) {
	this.Created = uint(time.Now().Unix())
	this.Id = 0
	if len(this.Engine) == 0 { this.Engine = "mysql" }
	if len(this.Host) == 0 { this.Host = "localhost:3306" }
	if len(this.User) == 0 { this.User = "root" }
	pk, err = this.Param().SetSend(this).Insert()
	if err == nil && pk != nil {
		if v, y := pk.(uint); y {
			this.Id = v
		} else if v, y := pk.(int64); y {
			this.Id = uint(v)
		}
	}
	return
}

func (this *DbAccount) Edit(mw func(db.Result) db.Result, args ...interface{}) error {
	this.Updated = uint(time.Now().Unix())
	if len(this.Engine) == 0 { this.Engine = "mysql" }
	if len(this.Host) == 0 { this.Host = "localhost:3306" }
	if len(this.User) == 0 { this.User = "root" }
	return this.Setter(mw, args...).SetSend(this).Update()
}

func (this *DbAccount) Setter(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	return this.Param().SetArgs(args...).SetMiddleware(mw)
}

func (this *DbAccount) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) error {
	return this.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (this *DbAccount) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) error {
	
	if val, ok := kvset["engine"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["engine"] = "mysql" } }
	if val, ok := kvset["host"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["host"] = "localhost:3306" } }
	if val, ok := kvset["user"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["user"] = "root" } }
	return this.Setter(mw, args...).SetSend(kvset).Update()
}

func (this *DbAccount) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = this.Param().SetArgs(args...).SetSend(this).SetMiddleware(mw).Upsert(func(){
		this.Updated = uint(time.Now().Unix())
	if len(this.Engine) == 0 { this.Engine = "mysql" }
	if len(this.Host) == 0 { this.Host = "localhost:3306" }
	if len(this.User) == 0 { this.User = "root" }
	},func(){
		this.Created = uint(time.Now().Unix())
	this.Id = 0
	if len(this.Engine) == 0 { this.Engine = "mysql" }
	if len(this.Host) == 0 { this.Host = "localhost:3306" }
	if len(this.User) == 0 { this.User = "root" }
	})
	if err == nil && pk != nil {
		if v, y := pk.(uint); y {
			this.Id = v
		} else if v, y := pk.(int64); y {
			this.Id = uint(v)
		}
	}
	return 
}

func (this *DbAccount) Delete(mw func(db.Result) db.Result, args ...interface{}) error {
	
	return this.Param().SetArgs(args...).SetMiddleware(mw).Delete()
}

func (this *DbAccount) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return this.Param().SetArgs(args...).SetMiddleware(mw).Count()
}

func (this *DbAccount) Reset() *DbAccount {
	this.Id = 0
	this.Title = ``
	this.Uid = 0
	this.Engine = ``
	this.Host = ``
	this.User = ``
	this.Password = ``
	this.Name = ``
	this.Options = ``
	this.Created = 0
	this.Updated = 0
	return this
}

func (this *DbAccount) AsMap() map[string]interface{} {
	r := map[string]interface{}{}
	r["Id"] = this.Id
	r["Title"] = this.Title
	r["Uid"] = this.Uid
	r["Engine"] = this.Engine
	r["Host"] = this.Host
	r["User"] = this.User
	r["Password"] = this.Password
	r["Name"] = this.Name
	r["Options"] = this.Options
	r["Created"] = this.Created
	r["Updated"] = this.Updated
	return r
}

func (this *DbAccount) AsRow() map[string]interface{} {
	r := map[string]interface{}{}
	r["id"] = this.Id
	r["title"] = this.Title
	r["uid"] = this.Uid
	r["engine"] = this.Engine
	r["host"] = this.Host
	r["user"] = this.User
	r["password"] = this.Password
	r["name"] = this.Name
	r["options"] = this.Options
	r["created"] = this.Created
	r["updated"] = this.Updated
	return r
}

func (this *DbAccount) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = this.AsRow()
	}
	return factory.BatchValidate("db_account", kvset)
}

func (this *DbAccount) Validate(field string, value interface{}) error {
	return factory.Validate("db_account", field, value)
}

