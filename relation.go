package gocassa

import (
	"fmt"
	"strings"
)

const (
	equality = iota
	in
	greaterThan
	greaterThanOrEquals
	lesserThan
	lesserThanOrEquals
)

type Relation struct {
	op    int
	key   string
	terms []interface{}
}

func (r Relation) cql() (string, []interface{}) {
	ret := ""
	key := strings.ToLower(r.key)
	switch r.op {
	case equality:
		ret = key + " = ?"
	case in:
		// Ideally the above code should work.
		qs := []string{}
		for i := 0; i < len(r.terms); i++ {
			qs = append(qs, "?")
		}
		return fmt.Sprintf(key+" IN (%s)", strings.Join(qs, ", ")), r.terms
	case greaterThan:
		ret = key + " > ?"
	case greaterThanOrEquals:
		ret = key + " >= ?"
	case lesserThan:
		ret = key + " < ?"
	case lesserThanOrEquals:
		ret = key + " <= ?"
	}
	return ret, r.terms
}

func toI(i interface{}) []interface{} {
	return []interface{}{i}
}

func Eq(key string, term interface{}) Relation {
	return Relation{
		op:    equality,
		key:   key,
		terms: toI(term),
	}
}

func In(key string, terms ...interface{}) Relation {
	return Relation{
		op:    in,
		key:   key,
		terms: terms,
	}
}

func GT(key string, term interface{}) Relation {
	return Relation{
		op:    greaterThan,
		key:   key,
		terms: toI(term),
	}
}

func GTE(key string, term interface{}) Relation {
	return Relation{
		op:    greaterThanOrEquals,
		key:   key,
		terms: toI(term),
	}
}

func LT(key string, term interface{}) Relation {
	return Relation{
		op:    lesserThan,
		key:   key,
		terms: toI(term),
	}
}

func LTE(key string, term interface{}) Relation {
	return Relation{
		op:    lesserThanOrEquals,
		key:   key,
		terms: toI(term),
	}
}
