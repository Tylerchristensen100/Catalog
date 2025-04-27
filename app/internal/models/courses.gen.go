package models

const TableNameCourse = "courses"

// Course mapped from table <courses>
type Course struct {
	ID            int32  `gorm:"column:id;type:int;primaryKey;autoIncrement:true;uniqueIndex:id_UNIQUECOURSE,priority:1" json:"id"`
	Name          string `gorm:"column:name;type:varchar(100);not null" json:"name"`
	MajorCode     string `gorm:"column:major_code;type:varchar(10);not null" json:"major_code"`
	Major         Major  `gorm:"foreignKey:MajorCode" json:"major"`
	Code          string `gorm:"column:code;type:varchar(8);not null" json:"code"`
	CreditHours   int32  `gorm:"column:credit_hours;type:int;not null" json:"credit_hours"`
	Prerequisites string `gorm:"column:prerequisites;type:varchar(150);not null" json:"prerequisites"`
	Description   string `gorm:"column:description;type:varchar(1000);not null" json:"description"`
	CreatedBy     int32  `gorm:"column:created_by;type:int;not null;index:created_by_idx,priority:1" json:"created_by"`
}

// TableName Course's table name
func (*Course) TableName() string {
	return TableNameCourse
}
