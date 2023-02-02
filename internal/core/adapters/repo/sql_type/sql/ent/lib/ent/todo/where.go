// Code generated by ent, DO NOT EDIT.

package todo

import (
	"hex-base/internal/core/adapters/repo/sql_type/sql/ent/lib/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int32) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int32) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int32) predicate.Todo {
	return predicate.Todo(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int32) predicate.Todo {
	return predicate.Todo(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int32) predicate.Todo {
	return predicate.Todo(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int32) predicate.Todo {
	return predicate.Todo(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int32) predicate.Todo {
	return predicate.Todo(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int32) predicate.Todo {
	return predicate.Todo(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int32) predicate.Todo {
	return predicate.Todo(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldName, v))
}

// Task applies equality check predicate on the "task" field. It's identical to TaskEQ.
func Task(v string) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldTask, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Todo {
	return predicate.Todo(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Todo {
	return predicate.Todo(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Todo {
	return predicate.Todo(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Todo {
	return predicate.Todo(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Todo {
	return predicate.Todo(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Todo {
	return predicate.Todo(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Todo {
	return predicate.Todo(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Todo {
	return predicate.Todo(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Todo {
	return predicate.Todo(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Todo {
	return predicate.Todo(sql.FieldHasSuffix(FieldName, v))
}

// NameIsNil applies the IsNil predicate on the "name" field.
func NameIsNil() predicate.Todo {
	return predicate.Todo(sql.FieldIsNull(FieldName))
}

// NameNotNil applies the NotNil predicate on the "name" field.
func NameNotNil() predicate.Todo {
	return predicate.Todo(sql.FieldNotNull(FieldName))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Todo {
	return predicate.Todo(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Todo {
	return predicate.Todo(sql.FieldContainsFold(FieldName, v))
}

// TaskEQ applies the EQ predicate on the "task" field.
func TaskEQ(v string) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldTask, v))
}

// TaskNEQ applies the NEQ predicate on the "task" field.
func TaskNEQ(v string) predicate.Todo {
	return predicate.Todo(sql.FieldNEQ(FieldTask, v))
}

// TaskIn applies the In predicate on the "task" field.
func TaskIn(vs ...string) predicate.Todo {
	return predicate.Todo(sql.FieldIn(FieldTask, vs...))
}

// TaskNotIn applies the NotIn predicate on the "task" field.
func TaskNotIn(vs ...string) predicate.Todo {
	return predicate.Todo(sql.FieldNotIn(FieldTask, vs...))
}

// TaskGT applies the GT predicate on the "task" field.
func TaskGT(v string) predicate.Todo {
	return predicate.Todo(sql.FieldGT(FieldTask, v))
}

// TaskGTE applies the GTE predicate on the "task" field.
func TaskGTE(v string) predicate.Todo {
	return predicate.Todo(sql.FieldGTE(FieldTask, v))
}

// TaskLT applies the LT predicate on the "task" field.
func TaskLT(v string) predicate.Todo {
	return predicate.Todo(sql.FieldLT(FieldTask, v))
}

// TaskLTE applies the LTE predicate on the "task" field.
func TaskLTE(v string) predicate.Todo {
	return predicate.Todo(sql.FieldLTE(FieldTask, v))
}

// TaskContains applies the Contains predicate on the "task" field.
func TaskContains(v string) predicate.Todo {
	return predicate.Todo(sql.FieldContains(FieldTask, v))
}

// TaskHasPrefix applies the HasPrefix predicate on the "task" field.
func TaskHasPrefix(v string) predicate.Todo {
	return predicate.Todo(sql.FieldHasPrefix(FieldTask, v))
}

// TaskHasSuffix applies the HasSuffix predicate on the "task" field.
func TaskHasSuffix(v string) predicate.Todo {
	return predicate.Todo(sql.FieldHasSuffix(FieldTask, v))
}

// TaskIsNil applies the IsNil predicate on the "task" field.
func TaskIsNil() predicate.Todo {
	return predicate.Todo(sql.FieldIsNull(FieldTask))
}

// TaskNotNil applies the NotNil predicate on the "task" field.
func TaskNotNil() predicate.Todo {
	return predicate.Todo(sql.FieldNotNull(FieldTask))
}

// TaskEqualFold applies the EqualFold predicate on the "task" field.
func TaskEqualFold(v string) predicate.Todo {
	return predicate.Todo(sql.FieldEqualFold(FieldTask, v))
}

// TaskContainsFold applies the ContainsFold predicate on the "task" field.
func TaskContainsFold(v string) predicate.Todo {
	return predicate.Todo(sql.FieldContainsFold(FieldTask, v))
}

// HasForm applies the HasEdge predicate on the "form" edge.
func HasForm() predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FormTable, FormColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFormWith applies the HasEdge predicate on the "form" edge with a given conditions (other predicates).
func HasFormWith(preds ...predicate.Form) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FormInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FormTable, FormColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Todo) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Todo) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Todo) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		p(s.Not())
	})
}
