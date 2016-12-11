//Do not edit this file, which is automatically generated by the generator.
package dbschema

import (
	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/factory"
	
	"time"
)

type FtpUserGroup struct {
	param   *factory.Param
	trans	*factory.Transaction
	objects []*FtpUserGroup
	
	Id          	uint    	`db:"id,omitempty,pk" bson:"id,omitempty" comment:"" json:"id" xml:"id"`
	Name        	string  	`db:"name" bson:"name" comment:"组名称" json:"name" xml:"name"`
	Created     	uint    	`db:"created" bson:"created" comment:"创建时间" json:"created" xml:"created"`
	Updated     	uint    	`db:"updated" bson:"updated" comment:"修改时间" json:"updated" xml:"updated"`
	Disabled    	string  	`db:"disabled" bson:"disabled" comment:"是否禁用" json:"disabled" xml:"disabled"`
	Banned      	string  	`db:"banned" bson:"banned" comment:"是否禁止组内用户连接" json:"banned" xml:"banned"`
	Directory   	string  	`db:"directory" bson:"directory" comment:"授权目录" json:"directory" xml:"directory"`
	IpWhitelist 	string  	`db:"ip_whitelist" bson:"ip_whitelist" comment:"IP白名单(一行一个)" json:"ip_whitelist" xml:"ip_whitelist"`
	IpBlacklist 	string  	`db:"ip_blacklist" bson:"ip_blacklist" comment:"IP黑名单(一行一个)" json:"ip_blacklist" xml:"ip_blacklist"`
}

func (this *FtpUserGroup) Trans() *factory.Transaction {
	return this.trans
}

func (this *FtpUserGroup) Use(trans *factory.Transaction) factory.Model {
	this.trans = trans
	return this
}

func (this *FtpUserGroup) Objects() []*FtpUserGroup {
	if this.objects == nil {
		return nil
	}
	return this.objects[:]
}

func (this *FtpUserGroup) NewObjects() *[]*FtpUserGroup {
	this.objects = []*FtpUserGroup{}
	return &this.objects
}

func (this *FtpUserGroup) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetTrans(this.trans).SetCollection("ftp_user_group").SetModel(this)
}

func (this *FtpUserGroup) SetParam(param *factory.Param) factory.Model {
	this.param = param
	return this
}

func (this *FtpUserGroup) Param() *factory.Param {
	if this.param == nil {
		return this.NewParam()
	}
	return this.param
}

func (this *FtpUserGroup) Get(mw func(db.Result) db.Result, args ...interface{}) error {
	return this.Param().SetArgs(args...).SetRecv(this).SetMiddleware(mw).One()
}

func (this *FtpUserGroup) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.NewObjects()
	}
	return this.Param().SetArgs(args...).SetPage(page).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *FtpUserGroup) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.NewObjects()
	}
	return this.Param().SetArgs(args...).SetOffset(offset).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *FtpUserGroup) Add() (pk interface{}, err error) {
	this.Created = uint(time.Now().Unix())
	this.Id = 0
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

func (this *FtpUserGroup) Edit(mw func(db.Result) db.Result, args ...interface{}) error {
	this.Updated = uint(time.Now().Unix())
	return this.Param().SetArgs(args...).SetSend(this).SetMiddleware(mw).Update()
}

func (this *FtpUserGroup) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = this.Param().SetArgs(args...).SetSend(this).SetMiddleware(mw).Upsert(func(){
		this.Updated = uint(time.Now().Unix())
	},func(){
		this.Created = uint(time.Now().Unix())
	this.Id = 0
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

func (this *FtpUserGroup) Delete(mw func(db.Result) db.Result, args ...interface{}) error {
	
	return this.Param().SetArgs(args...).SetMiddleware(mw).Delete()
}

