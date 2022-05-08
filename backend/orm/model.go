package orm

import (
	"mysite/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type LegalEntity struct {
	ID int `gorm:"primaryKey,index"`
}

type Company struct {
	ID            int    `json:"id" gorm:"primaryKey,index"`
	LegalEntityId int    `json:"legal_entity_id,omitempty"`
	Name          string `json:"name"`
}

type Exchange struct {
	ID   int `gorm:"primaryKey,index"`
	Name string
	Code string
}

type Industry struct {
	ID                 int `gorm:"primaryKey,index"`
	Name               string
	IndustryCategoryID int
	Description        string
}
type IndustryCategory struct {
	ID int `gorm:"primaryKey,index"`
}

type Ownership struct {
	ID            int `gorm:"primaryKey,index"`
	LegalEntityID int
	StockID       int
	Amount        int
}

type Person struct {
	ID            int `gorm:"primaryKey,index"`
	Name          string
	Age           int
	LegalEntityID int
	ImageUrl      string
}

func (Person) TableName() string {
	return "persons"
}

type Stock struct {
	ID         int `gorm:"primaryKey,index"`
	Code       string
	ExchangeId int
}

type StockPrice struct {
	ID           int `gorm:"primaryKey,index"`
	Date         int64
	OpeningPrice int
	ClosingPrice int
	Highest      int
	Lowest       int
	StockId      int
}

type IndustryRelationType struct {
	ID          int `gorm:"primaryKey,index"`
	Name        string
	Coefficient float32
}

type IndustryRelation struct {
	ID                     int `gorm:"primaryKey,index"`
	FromCompanyID          int
	ToCompanyID            int
	IndustryRelationTypeID int
}

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open(postgres.Open(utils.Config.ConnStr), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
