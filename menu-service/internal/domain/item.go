package domain

import (
	"time"
)

type Item struct {
	ID          int         `json:"id" db:"id"`
	Name        string      `json:"name" db:"name"`
	Description string      `json:"description" db:"description"`
	Price       float64     `json:"price" db:"price"`
	Categories  []*Category `json:"categories,omitempty" db:"categories"`
	CreatedAt   time.Time   `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at" db:"updated_at"`
	DeletedAt   *time.Time  `json:"deleted_at,omitempty" db:"deleted_at"`
}

//
//type Categories []*Category
//
//func (c *Categories) Scan(val interface{}) error {
//	var raw []byte
//
//	switch v := val.(type) {
//	case []byte: // []uint8
//		raw = v
//	case string:
//		raw = []byte(v)
//	default:
//		return errors.New(fmt.Sprintf("unsupported type: %T", v))
//	}
//
//	err := json.Unmarshal(raw, &c)
//	if err != nil {
//		log.Printf("fail to unmarshall as byte (%T): %v", v, err)
//		return err
//	}
//	return nil
//}
//
//func (c *Categories) Value() (driver.Value, error) {
//	return json.Marshal(&c)
//}
