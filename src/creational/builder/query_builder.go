package builder

type QueryBuilder interface {
	SetFields(props ...string) QueryBuilder
	SetFrom(from string) QueryBuilder
	SetWhere(condition string) QueryBuilder
	SetAnd(condition string) QueryBuilder
	SetOr(condition string) QueryBuilder
	SetOrderBy(prop string, desc bool) QueryBuilder
	Build() (string, error)
}
