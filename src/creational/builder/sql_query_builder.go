package builder

import (
	"fmt"
	"strings"
)

type SQLQueryBuilder struct {
	slt   string
	frm   string
	whr   string
	order string
}

func NewSQLQueryBuilder() *SQLQueryBuilder {
	return &SQLQueryBuilder{}
}

func (bld *SQLQueryBuilder) SetFields(props ...string) QueryBuilder {
	bld.slt = fmt.Sprintf("SELECT %s ", strings.Join(props, ", "))
	return bld
}

func (bld *SQLQueryBuilder) SetFrom(from string) QueryBuilder {
	bld.frm = fmt.Sprintf("FROM %s ", from)
	return bld
}

func (bld *SQLQueryBuilder) SetWhere(condition string) QueryBuilder {
	bld.whr = fmt.Sprintf("WHERE %s ", condition)
	return bld
}

func (bld *SQLQueryBuilder) SetAnd(condition string) QueryBuilder {
	bld.whr += fmt.Sprintf("AND %s ", condition)
	return bld
}

func (bld *SQLQueryBuilder) SetOr(condition string) QueryBuilder {
	bld.whr += fmt.Sprintf("OR %s ", condition)
	return bld
}

func (bld *SQLQueryBuilder) SetOrderBy(prop string, desc bool) QueryBuilder {
	bld.order = "ORDER BY "
	if desc {
		bld.order += "DESC"
	} else {
		bld.order += "ASC"
	}
	return bld
}

func (bld *SQLQueryBuilder) Build() (string, error) {
	if bld.slt == "" {
		return "", fmt.Errorf("'select' is not set")
	}
	if bld.frm == "" {
		return "", fmt.Errorf("'from' is not set")
	}
	if bld.whr != "" && strings.Index(bld.whr, "WHERE") != 0 {
		return "", fmt.Errorf("'where' has wrong signature:'%s'", bld.whr)
	}
	return bld.slt + bld.frm + bld.whr + bld.order, nil
}
