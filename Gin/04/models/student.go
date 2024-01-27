package models

type Student struct {
	Id       int
	Number   int
	Password string
	ClassId  int
	Name     string
}

// 表示配置操作数据库的表名称
// 表示把Student结构体默认操作的表改为students表
func (Student) TableName() string {
	return "students"
}
