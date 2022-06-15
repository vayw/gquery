package gquery

import (
	"fmt"
	"strings"
)

type GQuery struct {
	QFrom  string
	QWhere map[string]Filter
	QGet   []string
}

type Filter struct {
	Condition string
	Value     string
}

func (q *GQuery) From(name string) {
	q.QFrom = name
}

func (q *GQuery) Where(name, condition, value string) {
	if q.QWhere == nil {
		q.QWhere = map[string]Filter{}
	}
	q.QWhere[name] = Filter{condition, value}
}

func (q *GQuery) Get(name string) {
	q.QGet = append(q.QGet, name)
}

func (q GQuery) Build() (qstring string) {
	var qfilter string
	var qfields []string

	if len(q.QWhere) > 0 {
		var filters []string

		for k, v := range q.QWhere {
			q := strings.Split(k, ".")
			if len(q) > 1 {
				qs := strings.Join(q, ": {")
				filters = append(filters, fmt.Sprintf("%s: {%s: %s}%s", qs, v.Condition, v.Value, strings.Repeat("}", len(q)-1)))
			} else {
				filters = append(filters, fmt.Sprintf("%s: {%s: %s}", k, v.Condition, v.Value))
			}
		}

		qfilter = fmt.Sprintf("(where: {%s} )", strings.Join(filters, ","))
	}

	for _, f := range q.QGet {
		qfields = append(qfields, Bracer(f))
	}
	fields := strings.Join(qfields, " ")

	qstring = fmt.Sprintf("{ %s %s { %s }}", q.QFrom, qfilter, fields)
	return
}

func Bracer(st string) (res string) {
	tf := strings.Split(st, ".")
	res = strings.Join(tf, " { ")
	res = fmt.Sprintf("%s %s", res, strings.Repeat(" } ", len(tf)-1))
	return
}
