package alur

import (
    "github.com/eaciit/dbox"
    "github.com/eaciit/orm/v1"
)

var db *orm.DataContext

func SetDb(conn dbox.IConnection){
    if db!=nil{
        db.Close()
    }
    db=orm.New(conn)
}

func CloseDb(){
    if db!=nil{
        db.Close()
    }
}

func Db() *orm.DataContext{
    return db
}
