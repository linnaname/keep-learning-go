gorm

忍受不了的几点：

1.gorm.Model的有默认字段，如果数据库中没有居然会自动在数据库创建

2.表名默认是结构体名称的复数形式，我醉了，如果是已有表明需要自己实现TableName方法指定。
如果一不小心没自定义表名还会自动创建一个默认的表并增加默认的字段...........这是什么操作?


总体上设计并不是很清晰，有很多潜规则，不太好用