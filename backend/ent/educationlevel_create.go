// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team10/app/ent/doctorinfo"
	"github.com/team10/app/ent/educationlevel"
)

// EducationlevelCreate is the builder for creating a Educationlevel entity.
type EducationlevelCreate struct {
	config
	mutation *EducationlevelMutation
	hooks    []Hook
}

// SetLevel sets the level field.
func (ec *EducationlevelCreate) SetLevel(s string) *EducationlevelCreate {
	ec.mutation.SetLevel(s)
	return ec
}

// AddEducationlevel2doctorinfoIDs adds the educationlevel2doctorinfo edge to Doctorinfo by ids.
func (ec *EducationlevelCreate) AddEducationlevel2doctorinfoIDs(ids ...int) *EducationlevelCreate {
	ec.mutation.AddEducationlevel2doctorinfoIDs(ids...)
	return ec
}

// AddEducationlevel2doctorinfo adds the educationlevel2doctorinfo edges to Doctorinfo.
func (ec *EducationlevelCreate) AddEducationlevel2doctorinfo(d ...*Doctorinfo) *EducationlevelCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return ec.AddEducationlevel2doctorinfoIDs(ids...)
}

// Mutation returns the EducationlevelMutation object of the builder.
func (ec *EducationlevelCreate) Mutation() *EducationlevelMutation {
	return ec.mutation
}

// Save creates the Educationlevel in the database.
func (ec *EducationlevelCreate) Save(ctx context.Context) (*Educationlevel, error) {
	if _, ok := ec.mutation.Level(); !ok {
		return nil, &ValidationError{Name: "level", err: errors.New("ent: missing required field \"level\"")}
	}
	if v, ok := ec.mutation.Level(); ok {
		if err := educationlevel.LevelValidator(v); err != nil {
			return nil, &ValidationError{Name: "level", err: fmt.Errorf("ent: validator failed for field \"level\": %w", err)}
		}
	}
	var (
		err  error
		node *Educationlevel
	)
	if len(ec.hooks) == 0 {
		node, err = ec.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EducationlevelMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ec.mutation = mutation
			node, err = ec.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ec.hooks) - 1; i >= 0; i-- {
			mut = ec.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ec.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ec *EducationlevelCreate) SaveX(ctx context.Context) *Educationlevel {
	v, err := ec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ec *EducationlevelCreate) sqlSave(ctx context.Context) (*Educationlevel, error) {
	e, _spec := ec.createSpec()
	if err := sqlgraph.CreateNode(ctx, ec.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	e.ID = int(id)
	return e, nil
}

func (ec *EducationlevelCreate) createSpec() (*Educationlevel, *sqlgraph.CreateSpec) {
	var (
		e     = &Educationlevel{config: ec.config}
		_spec = &sqlgraph.CreateSpec{
			Table: educationlevel.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: educationlevel.FieldID,
			},
		}
	)
	if value, ok := ec.mutation.Level(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: educationlevel.FieldLevel,
		})
		e.Level = value
	}
	if nodes := ec.mutation.Educationlevel2doctorinfoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   educationlevel.Educationlevel2doctorinfoTable,
			Columns: []string{educationlevel.Educationlevel2doctorinfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: doctorinfo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return e, _spec
}
