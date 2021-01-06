// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team10/app/ent/gender"
	"github.com/team10/app/ent/patientrecord"
)

// GenderCreate is the builder for creating a Gender entity.
type GenderCreate struct {
	config
	mutation *GenderMutation
	hooks    []Hook
}

// SetGenderstatus sets the Genderstatus field.
func (gc *GenderCreate) SetGenderstatus(s string) *GenderCreate {
	gc.mutation.SetGenderstatus(s)
	return gc
}

// AddPatientrecordIDs adds the patientrecord edge to Patientrecord by ids.
func (gc *GenderCreate) AddPatientrecordIDs(ids ...int) *GenderCreate {
	gc.mutation.AddPatientrecordIDs(ids...)
	return gc
}

// AddPatientrecord adds the patientrecord edges to Patientrecord.
func (gc *GenderCreate) AddPatientrecord(p ...*Patientrecord) *GenderCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return gc.AddPatientrecordIDs(ids...)
}

// Mutation returns the GenderMutation object of the builder.
func (gc *GenderCreate) Mutation() *GenderMutation {
	return gc.mutation
}

// Save creates the Gender in the database.
func (gc *GenderCreate) Save(ctx context.Context) (*Gender, error) {
	if _, ok := gc.mutation.Genderstatus(); !ok {
		return nil, &ValidationError{Name: "Genderstatus", err: errors.New("ent: missing required field \"Genderstatus\"")}
	}
	var (
		err  error
		node *Gender
	)
	if len(gc.hooks) == 0 {
		node, err = gc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GenderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			gc.mutation = mutation
			node, err = gc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(gc.hooks) - 1; i >= 0; i-- {
			mut = gc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (gc *GenderCreate) SaveX(ctx context.Context) *Gender {
	v, err := gc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (gc *GenderCreate) sqlSave(ctx context.Context) (*Gender, error) {
	ge, _spec := gc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	ge.ID = int(id)
	return ge, nil
}

func (gc *GenderCreate) createSpec() (*Gender, *sqlgraph.CreateSpec) {
	var (
		ge    = &Gender{config: gc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: gender.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: gender.FieldID,
			},
		}
	)
	if value, ok := gc.mutation.Genderstatus(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: gender.FieldGenderstatus,
		})
		ge.Genderstatus = value
	}
	if nodes := gc.mutation.PatientrecordIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   gender.PatientrecordTable,
			Columns: []string{gender.PatientrecordColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patientrecord.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return ge, _spec
}