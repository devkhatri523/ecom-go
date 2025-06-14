package paginator

import (
	"github.com/devkhatri523/ecom-go/cursor-paginator/internal/util"
	"gorm.io/gorm"
	"reflect"
)

// Rule for paginator
type Rule struct {
	Key             string
	Order           Order
	SQLRepr         string
	SQLType         *string
	NULLReplacement interface{}
	CustomType      *CustomType
}

// CustomType for paginator. It provides extra info needed to paginate across custom types (e.g. JSON)
type CustomType struct {
	Meta interface{}
	Type reflect.Type
}

func (r *Rule) validate(db *gorm.DB, dest interface{}) (err error) {
	if schema, err := util.ParseSchema(db, dest); err != nil {
		return ErrInvalidModel
	} else if f := schema.LookUpField(r.Key); f == nil {
		return ErrInvalidModel
	}
	if r.Order != "" {
		if err = r.Order.validate(); err != nil {
			return
		}
	}
	return nil
}
