// Code generated by entc, DO NOT EDIT.

package abilitypatientrights

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/team10/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Abilitypatientrights {
	return predicate.Abilitypatientrights(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Abilitypatientrights {
	return predicate.Abilitypatientrights(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Abilitypatientrights {
	return predicate.Abilitypatientrights(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Abilitypatientrights {
	return predicate.Abilitypatientrights(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Abilitypatientrights {
	return predicate.Abilitypatientrights(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Abilitypatientrights {
	return predicate.Abilitypatientrights(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Abilitypatientrights {
	return predicate.Abilitypatientrights(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Abilitypatientrights {
	return predicate.Abilitypatientrights(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Abilitypatientrights {
	return predicate.Abilitypatientrights(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Operative applies equality check predicate on the "Operative" field. It's identical to OperativeEQ.
func Operative(v string) predicate.Abilitypatientrights {
	return predicate.Abilitypatientrights(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOperative), v))
	})
}

// MedicalSupplies applies equality check predicate on the "MedicalSupplies" field. It's identical to MedicalSuppliesEQ.
func MedicalSupplies(v string) predicate.Abilitypatientrights {
	return predicate.Abilitypatientrights(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMedicalSupplies), v))
	})
}

// Examine applies equality check predicate on the "Examine" field. It's identical to ExamineEQ.
func Examine(v string) predicate.Abilitypatientrights {
	return predicate.Abilitypatientrights(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExamine), v))
	})
}

// OperativeEQ applies the EQ predicate on the "Operative" field.
func OperativeEQ(v string) predicate.Abilitypatientrights {
	return predicate.Abilitypatientrights(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOperative), v))
	})
}

// OperativeNEQ applies the NEQ predicate on the "Operative" field.
func OperativeNEQ(v string) predicate.Abilitypatientrights {
	return predicate.Abilitypatientrights(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOperative), v))
	})
}

// OperativeIn applies the In predicate on the "Operative" field.
func OperativeIn(vs ...string) predicate.Abilitypatientrights {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Abilitypatientrights(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOperative), v...))
	})
}

// OperativeNotIn applies the NotIn predicate on the "Operative" field.
func OperativeNotIn(vs ...string) predicate.Abilitypatientrights {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Abilitypatientrights(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOperative), v...))
	})
}

// OperativeGT applies the GT predicate on the "Operative" field.
func OperativeGT(v string) predicate.Abilitypatientrights {
	return predicate.Abilitypatientrights(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOperative), v))
	})
}

// OperativeGTE applies the GTE predicate on the "Operative" field.
func OperativeGTE(v string) predicate.Abilitypatientrights {
	return predicate.Abilitypatientrights(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOperative), v))
	})
}

// OperativeLT applies the LT predicate on the "Operative" field.
func OperativeLT(v string) predicate.Abilitypatientrights {
	return predicate.Abilitypatientrights(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOperative), v))
	})
}

// OperativeLTE applies the LTE predicate on the "Operative" field.
func OperativeLTE(v string) predicate.Abilitypatientrights {
	return predicate.Abilitypatientrights(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOperative), v))
	})
}

// OperativeContains applies the Contains predicate on the "Operative" field.
func OperativeContains(v string) predicate.Abilitypatientrights {
	return predicate.Abilitypatientrights(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOperative), v))
	})
}

// OperativeHasPrefix applies the HasPrefix predicate on the "Operative" field.
func OperativeHasPrefix(v string) predicate.Abilitypatientrights {
	return predicate.Abilitypatientrights(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOperative), v))
	})
}

// OperativeHasSuffix applies the HasSuffix predicate on the "Operative" field.
func OperativeHasSuffix(v string) predicate.Abilitypatientrights {
	return predicate.Abilitypatientrights(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOperative), v))
	})
}

// OperativeEqualFold applies the EqualFold predicate on the "Operative" field.
func OperativeEqualFold(v string) predicate.Abilitypatientrights {
	return predicate.Abilitypatientrights(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOperative), v))
	})
}

// OperativeContainsFold applies the ContainsFold predicate on the "Operative" field.
func OperativeContainsFold(v string) predicate.Abilitypatientrights {
	return predicate.Abilitypatientrights(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOperative), v))
	})
}

// MedicalSuppliesEQ applies the EQ predicate on the "MedicalSupplies" field.
func MedicalSuppliesEQ(v string) predicate.Abilitypatientrights {
	return predicate.Abilitypatientrights(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMedicalSupplies), v))
	})
}

// MedicalSuppliesNEQ applies the NEQ predicate on the "MedicalSupplies" field.
func MedicalSuppliesNEQ(v string) predicate.Abilitypatientrights {
	return predicate.Abilitypatientrights(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMedicalSupplies), v))
	})
}

// MedicalSuppliesIn applies the In predicate on the "MedicalSupplies" field.
func MedicalSuppliesIn(vs ...string) predicate.Abilitypatientrights {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Abilitypatientrights(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMedicalSupplies), v...))
	})
}

// MedicalSuppliesNotIn applies the NotIn predicate on the "MedicalSupplies" field.
func MedicalSuppliesNotIn(vs ...string) predicate.Abilitypatientrights {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Abilitypatientrights(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMedicalSupplies), v...))
	})
}

// MedicalSuppliesGT applies the GT predicate on the "MedicalSupplies" field.
func MedicalSuppliesGT(v string) predicate.Abilitypatientrights {
	return predicate.Abilitypatientrights(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMedicalSupplies), v))
	})
}

// MedicalSuppliesGTE applies the GTE predicate on the "MedicalSupplies" field.
func MedicalSuppliesGTE(v string) predicate.Abilitypatientrights {
	return predicate.Abilitypatientrights(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMedicalSupplies), v))
	})
}

// MedicalSuppliesLT applies the LT predicate on the "MedicalSupplies" field.
func MedicalSuppliesLT(v string) predicate.Abilitypatientrights {
	return predicate.Abilitypatientrights(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMedicalSupplies), v))
	})
}

// MedicalSuppliesLTE applies the LTE predicate on the "MedicalSupplies" field.
func MedicalSuppliesLTE(v string) predicate.Abilitypatientrights {
	return predicate.Abilitypatientrights(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMedicalSupplies), v))
	})
}

// MedicalSuppliesContains applies the Contains predicate on the "MedicalSupplies" field.
func MedicalSuppliesContains(v string) predicate.Abilitypatientrights {
	return predicate.Abilitypatientrights(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMedicalSupplies), v))
	})
}

// MedicalSuppliesHasPrefix applies the HasPrefix predicate on the "MedicalSupplies" field.
func MedicalSuppliesHasPrefix(v string) predicate.Abilitypatientrights {
	return predicate.Abilitypatientrights(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMedicalSupplies), v))
	})
}

// MedicalSuppliesHasSuffix applies the HasSuffix predicate on the "MedicalSupplies" field.
func MedicalSuppliesHasSuffix(v string) predicate.Abilitypatientrights {
	return predicate.Abilitypatientrights(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMedicalSupplies), v))
	})
}

// MedicalSuppliesEqualFold applies the EqualFold predicate on the "MedicalSupplies" field.
func MedicalSuppliesEqualFold(v string) predicate.Abilitypatientrights {
	return predicate.Abilitypatientrights(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMedicalSupplies), v))
	})
}

// MedicalSuppliesContainsFold applies the ContainsFold predicate on the "MedicalSupplies" field.
func MedicalSuppliesContainsFold(v string) predicate.Abilitypatientrights {
	return predicate.Abilitypatientrights(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMedicalSupplies), v))
	})
}

// ExamineEQ applies the EQ predicate on the "Examine" field.
func ExamineEQ(v string) predicate.Abilitypatientrights {
	return predicate.Abilitypatientrights(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExamine), v))
	})
}

// ExamineNEQ applies the NEQ predicate on the "Examine" field.
func ExamineNEQ(v string) predicate.Abilitypatientrights {
	return predicate.Abilitypatientrights(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldExamine), v))
	})
}

// ExamineIn applies the In predicate on the "Examine" field.
func ExamineIn(vs ...string) predicate.Abilitypatientrights {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Abilitypatientrights(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldExamine), v...))
	})
}

// ExamineNotIn applies the NotIn predicate on the "Examine" field.
func ExamineNotIn(vs ...string) predicate.Abilitypatientrights {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Abilitypatientrights(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldExamine), v...))
	})
}

// ExamineGT applies the GT predicate on the "Examine" field.
func ExamineGT(v string) predicate.Abilitypatientrights {
	return predicate.Abilitypatientrights(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldExamine), v))
	})
}

// ExamineGTE applies the GTE predicate on the "Examine" field.
func ExamineGTE(v string) predicate.Abilitypatientrights {
	return predicate.Abilitypatientrights(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldExamine), v))
	})
}

// ExamineLT applies the LT predicate on the "Examine" field.
func ExamineLT(v string) predicate.Abilitypatientrights {
	return predicate.Abilitypatientrights(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldExamine), v))
	})
}

// ExamineLTE applies the LTE predicate on the "Examine" field.
func ExamineLTE(v string) predicate.Abilitypatientrights {
	return predicate.Abilitypatientrights(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldExamine), v))
	})
}

// ExamineContains applies the Contains predicate on the "Examine" field.
func ExamineContains(v string) predicate.Abilitypatientrights {
	return predicate.Abilitypatientrights(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldExamine), v))
	})
}

// ExamineHasPrefix applies the HasPrefix predicate on the "Examine" field.
func ExamineHasPrefix(v string) predicate.Abilitypatientrights {
	return predicate.Abilitypatientrights(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldExamine), v))
	})
}

// ExamineHasSuffix applies the HasSuffix predicate on the "Examine" field.
func ExamineHasSuffix(v string) predicate.Abilitypatientrights {
	return predicate.Abilitypatientrights(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldExamine), v))
	})
}

// ExamineEqualFold applies the EqualFold predicate on the "Examine" field.
func ExamineEqualFold(v string) predicate.Abilitypatientrights {
	return predicate.Abilitypatientrights(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldExamine), v))
	})
}

// ExamineContainsFold applies the ContainsFold predicate on the "Examine" field.
func ExamineContainsFold(v string) predicate.Abilitypatientrights {
	return predicate.Abilitypatientrights(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldExamine), v))
	})
}

// HasAbilitypatientrightsPatientrightstype applies the HasEdge predicate on the "AbilitypatientrightsPatientrightstype" edge.
func HasAbilitypatientrightsPatientrightstype() predicate.Abilitypatientrights {
	return predicate.Abilitypatientrights(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AbilitypatientrightsPatientrightstypeTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AbilitypatientrightsPatientrightstypeTable, AbilitypatientrightsPatientrightstypeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAbilitypatientrightsPatientrightstypeWith applies the HasEdge predicate on the "AbilitypatientrightsPatientrightstype" edge with a given conditions (other predicates).
func HasAbilitypatientrightsPatientrightstypeWith(preds ...predicate.Patientrightstype) predicate.Abilitypatientrights {
	return predicate.Abilitypatientrights(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AbilitypatientrightsPatientrightstypeInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AbilitypatientrightsPatientrightstypeTable, AbilitypatientrightsPatientrightstypeColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Abilitypatientrights) predicate.Abilitypatientrights {
	return predicate.Abilitypatientrights(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Abilitypatientrights) predicate.Abilitypatientrights {
	return predicate.Abilitypatientrights(func(s *sql.Selector) {
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
func Not(p predicate.Abilitypatientrights) predicate.Abilitypatientrights {
	return predicate.Abilitypatientrights(func(s *sql.Selector) {
		p(s.Not())
	})
}
