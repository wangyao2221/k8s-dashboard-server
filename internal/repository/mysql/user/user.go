package user

import (
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"

	"k8s-dashboard-server/internal/repository/mysql"
)

// 可以写一个生成器直接生成数据库代码
func NewModel() *User {
	return new(User)
}

func NewQueryBuilder() *userQueryBuilder {
	return new(userQueryBuilder)
}

func (t *User) Create(db *gorm.DB) (id int32, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

type userQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *userQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
	ret := db
	for _, where := range qb.where {
		ret = ret.Where(where.prefix, where.value)
	}
	for _, order := range qb.order {
		ret = ret.Order(order)
	}
	ret = ret.Limit(qb.limit).Offset(qb.offset)
	return ret
}

func (qb *userQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (err error) {
	db = db.Model(&User{})

	for _, where := range qb.where {
		db.Where(where.prefix, where.value)
	}

	if err = db.Updates(m).Error; err != nil {
		return errors.Wrap(err, "updates err")
	}
	return nil
}

func (qb *userQueryBuilder) Delete(db *gorm.DB) (err error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	if err = db.Delete(&User{}).Error; err != nil {
		return errors.Wrap(err, "delete err")
	}
	return nil
}

func (qb *userQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&User{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		c = 0
	}
	return c, res.Error
}

func (qb *userQueryBuilder) First(db *gorm.DB) (*User, error) {
	ret := &User{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		ret = nil
	}
	return ret, res.Error
}

func (qb *userQueryBuilder) QueryOne(db *gorm.DB) (*User, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *userQueryBuilder) QueryAll(db *gorm.DB) ([]*User, error) {
	var ret []*User
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *userQueryBuilder) Limit(limit int) *userQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *userQueryBuilder) Offset(offset int) *userQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *userQueryBuilder) WhereId(p mysql.Predicate, value int32) *userQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *userQueryBuilder) WhereIdIn(value []int32) *userQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *userQueryBuilder) WhereIdNotIn(value []int32) *userQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *userQueryBuilder) OrderById(asc bool) *userQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *userQueryBuilder) WhereUsername(p mysql.Predicate, value string) *userQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "username", p),
		value,
	})
	return qb
}

func (qb *userQueryBuilder) WhereUsernameIn(value []string) *userQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "username", "IN"),
		value,
	})
	return qb
}

func (qb *userQueryBuilder) WhereUsernameNotIn(value []string) *userQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "username", "NOT IN"),
		value,
	})
	return qb
}

func (qb *userQueryBuilder) OrderByUsername(asc bool) *userQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "username "+order)
	return qb
}

func (qb *userQueryBuilder) WherePassword(p mysql.Predicate, value string) *userQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "password", p),
		value,
	})
	return qb
}

func (qb *userQueryBuilder) WherePasswordIn(value []string) *userQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "password", "IN"),
		value,
	})
	return qb
}

func (qb *userQueryBuilder) WherePasswordNotIn(value []string) *userQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "password", "NOT IN"),
		value,
	})
	return qb
}

func (qb *userQueryBuilder) OrderByPassword(asc bool) *userQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "password "+order)
	return qb
}

func (qb *userQueryBuilder) WhereNickname(p mysql.Predicate, value string) *userQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "nickname", p),
		value,
	})
	return qb
}

func (qb *userQueryBuilder) WhereNicknameIn(value []string) *userQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "nickname", "IN"),
		value,
	})
	return qb
}

func (qb *userQueryBuilder) WhereNicknameNotIn(value []string) *userQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "nickname", "NOT IN"),
		value,
	})
	return qb
}

func (qb *userQueryBuilder) OrderByNickname(asc bool) *userQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "nickname "+order)
	return qb
}

func (qb *userQueryBuilder) WhereMobile(p mysql.Predicate, value string) *userQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "mobile", p),
		value,
	})
	return qb
}

func (qb *userQueryBuilder) WhereMobileIn(value []string) *userQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "mobile", "IN"),
		value,
	})
	return qb
}

func (qb *userQueryBuilder) WhereMobileNotIn(value []string) *userQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "mobile", "NOT IN"),
		value,
	})
	return qb
}

func (qb *userQueryBuilder) OrderByMobile(asc bool) *userQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "mobile "+order)
	return qb
}

func (qb *userQueryBuilder) WhereIsUsed(p mysql.Predicate, value int32) *userQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_used", p),
		value,
	})
	return qb
}

func (qb *userQueryBuilder) WhereIsUsedIn(value []int32) *userQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_used", "IN"),
		value,
	})
	return qb
}

func (qb *userQueryBuilder) WhereIsUsedNotIn(value []int32) *userQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_used", "NOT IN"),
		value,
	})
	return qb
}

func (qb *userQueryBuilder) OrderByIsUsed(asc bool) *userQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "is_used "+order)
	return qb
}

func (qb *userQueryBuilder) WhereIsDeleted(p mysql.Predicate, value int32) *userQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_deleted", p),
		value,
	})
	return qb
}

func (qb *userQueryBuilder) WhereIsDeletedIn(value []int32) *userQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_deleted", "IN"),
		value,
	})
	return qb
}

func (qb *userQueryBuilder) WhereIsDeletedNotIn(value []int32) *userQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_deleted", "NOT IN"),
		value,
	})
	return qb
}

func (qb *userQueryBuilder) OrderByIsDeleted(asc bool) *userQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "is_deleted "+order)
	return qb
}

func (qb *userQueryBuilder) WhereCreatedAt(p mysql.Predicate, value time.Time) *userQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", p),
		value,
	})
	return qb
}

func (qb *userQueryBuilder) WhereCreatedAtIn(value []time.Time) *userQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", "IN"),
		value,
	})
	return qb
}

func (qb *userQueryBuilder) WhereCreatedAtNotIn(value []time.Time) *userQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", "NOT IN"),
		value,
	})
	return qb
}

func (qb *userQueryBuilder) OrderByCreatedAt(asc bool) *userQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "created_at "+order)
	return qb
}

func (qb *userQueryBuilder) WhereCreatedUser(p mysql.Predicate, value string) *userQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_user", p),
		value,
	})
	return qb
}

func (qb *userQueryBuilder) WhereCreatedUserIn(value []string) *userQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_user", "IN"),
		value,
	})
	return qb
}

func (qb *userQueryBuilder) WhereCreatedUserNotIn(value []string) *userQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_user", "NOT IN"),
		value,
	})
	return qb
}

func (qb *userQueryBuilder) OrderByCreatedUser(asc bool) *userQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "created_user "+order)
	return qb
}

func (qb *userQueryBuilder) WhereUpdatedAt(p mysql.Predicate, value time.Time) *userQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", p),
		value,
	})
	return qb
}

func (qb *userQueryBuilder) WhereUpdatedAtIn(value []time.Time) *userQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", "IN"),
		value,
	})
	return qb
}

func (qb *userQueryBuilder) WhereUpdatedAtNotIn(value []time.Time) *userQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", "NOT IN"),
		value,
	})
	return qb
}

func (qb *userQueryBuilder) OrderByUpdatedAt(asc bool) *userQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "updated_at "+order)
	return qb
}

func (qb *userQueryBuilder) WhereUpdatedUser(p mysql.Predicate, value string) *userQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_user", p),
		value,
	})
	return qb
}

func (qb *userQueryBuilder) WhereUpdatedUserIn(value []string) *userQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_user", "IN"),
		value,
	})
	return qb
}

func (qb *userQueryBuilder) WhereUpdatedUserNotIn(value []string) *userQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_user", "NOT IN"),
		value,
	})
	return qb
}

func (qb *userQueryBuilder) OrderByUpdatedUser(asc bool) *userQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "updated_user "+order)
	return qb
}
