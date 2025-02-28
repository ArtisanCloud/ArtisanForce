package product

import (
	"PowerX/internal/model/powermodel"
)

// PriceBook 数据表结构
type PriceBook struct {
	Products []*Product `gorm:"many2many:price_book_entries;foreignKey:Id;joinForeignKey:price_book_id;References:Id;JoinReferences:price_book_id"`
	//Resellers []*Reseller `gorm:"foreignKey:PriceBookId;references:Id" json:"resellers"`

	powermodel.PowerModel
	IsStandard  bool   `gorm:"column:is_standard; comment:是否是标准手册,标准手册只能有一条" json:"isStandard"`
	Name        string `gorm:"column:name; comment:价格手册名字" json:"name"`
	Description string `gorm:"column:description; comment:手册描述" json:"description"`
	StoreId     string `gorm:"column:storeId; comment:门店Id" json:"storeId"`
}

const TableNamePriceBook = "price_books"
const PriceBookUniqueId = powermodel.UniqueId
