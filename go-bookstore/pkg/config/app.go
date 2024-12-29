package config

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db * gorm.DB
)
func Connect(){
d, err := gorm.Open("mysql", "mrfoor:BaderKousas@12@/simplerest?charst=utf8&parseTime=True&loc=Local")
if err != nil{
	panic(err)

}
db = d
}
func GetDB() *gorm.DB{
	return db
}