// Code generated by ent, DO NOT EDIT.

package internal

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/pulkitbhardwaj/omega/pkgs/ats/internal/candidate"
	"github.com/pulkitbhardwaj/omega/pkgs/ats/internal/position"
	"github.com/pulkitbhardwaj/omega/pkgs/ats/internal/predicate"
	"github.com/pulkitbhardwaj/omega/pkgs/ats/internal/skill"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeCandidate = "Candidate"
	TypePosition  = "Position"
	TypeSkill     = "Skill"
)

// CandidateMutation represents an operation that mutates the Candidate nodes in the graph.
type CandidateMutation struct {
	config
	op            Op
	typ           string
	id            *uuid.UUID
	created_at    *time.Time
	updated_at    *time.Time
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Candidate, error)
	predicates    []predicate.Candidate
}

var _ ent.Mutation = (*CandidateMutation)(nil)

// candidateOption allows management of the mutation configuration using functional options.
type candidateOption func(*CandidateMutation)

// newCandidateMutation creates new mutation for the Candidate entity.
func newCandidateMutation(c config, op Op, opts ...candidateOption) *CandidateMutation {
	m := &CandidateMutation{
		config:        c,
		op:            op,
		typ:           TypeCandidate,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withCandidateID sets the ID field of the mutation.
func withCandidateID(id uuid.UUID) candidateOption {
	return func(m *CandidateMutation) {
		var (
			err   error
			once  sync.Once
			value *Candidate
		)
		m.oldValue = func(ctx context.Context) (*Candidate, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Candidate.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withCandidate sets the old Candidate of the mutation.
func withCandidate(node *Candidate) candidateOption {
	return func(m *CandidateMutation) {
		m.oldValue = func(context.Context) (*Candidate, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m CandidateMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m CandidateMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("internal: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Candidate entities.
func (m *CandidateMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *CandidateMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *CandidateMutation) IDs(ctx context.Context) ([]uuid.UUID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uuid.UUID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Candidate.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetCreatedAt sets the "created_at" field.
func (m *CandidateMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *CandidateMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Candidate entity.
// If the Candidate object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CandidateMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *CandidateMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *CandidateMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *CandidateMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Candidate entity.
// If the Candidate object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CandidateMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *CandidateMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// Where appends a list predicates to the CandidateMutation builder.
func (m *CandidateMutation) Where(ps ...predicate.Candidate) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the CandidateMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *CandidateMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Candidate, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *CandidateMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *CandidateMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Candidate).
func (m *CandidateMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *CandidateMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.created_at != nil {
		fields = append(fields, candidate.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, candidate.FieldUpdatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *CandidateMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case candidate.FieldCreatedAt:
		return m.CreatedAt()
	case candidate.FieldUpdatedAt:
		return m.UpdatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *CandidateMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case candidate.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case candidate.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Candidate field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CandidateMutation) SetField(name string, value ent.Value) error {
	switch name {
	case candidate.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case candidate.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Candidate field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *CandidateMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *CandidateMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CandidateMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Candidate numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *CandidateMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *CandidateMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *CandidateMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Candidate nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *CandidateMutation) ResetField(name string) error {
	switch name {
	case candidate.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case candidate.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	}
	return fmt.Errorf("unknown Candidate field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *CandidateMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *CandidateMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *CandidateMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *CandidateMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *CandidateMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *CandidateMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *CandidateMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Candidate unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *CandidateMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Candidate edge %s", name)
}

// PositionMutation represents an operation that mutates the Position nodes in the graph.
type PositionMutation struct {
	config
	op            Op
	typ           string
	id            *uuid.UUID
	created_at    *time.Time
	updated_at    *time.Time
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Position, error)
	predicates    []predicate.Position
}

var _ ent.Mutation = (*PositionMutation)(nil)

// positionOption allows management of the mutation configuration using functional options.
type positionOption func(*PositionMutation)

// newPositionMutation creates new mutation for the Position entity.
func newPositionMutation(c config, op Op, opts ...positionOption) *PositionMutation {
	m := &PositionMutation{
		config:        c,
		op:            op,
		typ:           TypePosition,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withPositionID sets the ID field of the mutation.
func withPositionID(id uuid.UUID) positionOption {
	return func(m *PositionMutation) {
		var (
			err   error
			once  sync.Once
			value *Position
		)
		m.oldValue = func(ctx context.Context) (*Position, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Position.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withPosition sets the old Position of the mutation.
func withPosition(node *Position) positionOption {
	return func(m *PositionMutation) {
		m.oldValue = func(context.Context) (*Position, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m PositionMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m PositionMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("internal: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Position entities.
func (m *PositionMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *PositionMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *PositionMutation) IDs(ctx context.Context) ([]uuid.UUID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uuid.UUID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Position.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetCreatedAt sets the "created_at" field.
func (m *PositionMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *PositionMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Position entity.
// If the Position object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PositionMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *PositionMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *PositionMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *PositionMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Position entity.
// If the Position object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PositionMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *PositionMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// Where appends a list predicates to the PositionMutation builder.
func (m *PositionMutation) Where(ps ...predicate.Position) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the PositionMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *PositionMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Position, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *PositionMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *PositionMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Position).
func (m *PositionMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *PositionMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.created_at != nil {
		fields = append(fields, position.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, position.FieldUpdatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *PositionMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case position.FieldCreatedAt:
		return m.CreatedAt()
	case position.FieldUpdatedAt:
		return m.UpdatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *PositionMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case position.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case position.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Position field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *PositionMutation) SetField(name string, value ent.Value) error {
	switch name {
	case position.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case position.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Position field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *PositionMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *PositionMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *PositionMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Position numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *PositionMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *PositionMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *PositionMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Position nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *PositionMutation) ResetField(name string) error {
	switch name {
	case position.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case position.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	}
	return fmt.Errorf("unknown Position field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *PositionMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *PositionMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *PositionMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *PositionMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *PositionMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *PositionMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *PositionMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Position unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *PositionMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Position edge %s", name)
}

// SkillMutation represents an operation that mutates the Skill nodes in the graph.
type SkillMutation struct {
	config
	op            Op
	typ           string
	id            *uuid.UUID
	created_at    *time.Time
	updated_at    *time.Time
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Skill, error)
	predicates    []predicate.Skill
}

var _ ent.Mutation = (*SkillMutation)(nil)

// skillOption allows management of the mutation configuration using functional options.
type skillOption func(*SkillMutation)

// newSkillMutation creates new mutation for the Skill entity.
func newSkillMutation(c config, op Op, opts ...skillOption) *SkillMutation {
	m := &SkillMutation{
		config:        c,
		op:            op,
		typ:           TypeSkill,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withSkillID sets the ID field of the mutation.
func withSkillID(id uuid.UUID) skillOption {
	return func(m *SkillMutation) {
		var (
			err   error
			once  sync.Once
			value *Skill
		)
		m.oldValue = func(ctx context.Context) (*Skill, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Skill.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withSkill sets the old Skill of the mutation.
func withSkill(node *Skill) skillOption {
	return func(m *SkillMutation) {
		m.oldValue = func(context.Context) (*Skill, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m SkillMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m SkillMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("internal: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Skill entities.
func (m *SkillMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *SkillMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *SkillMutation) IDs(ctx context.Context) ([]uuid.UUID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uuid.UUID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Skill.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetCreatedAt sets the "created_at" field.
func (m *SkillMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *SkillMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Skill entity.
// If the Skill object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *SkillMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *SkillMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *SkillMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *SkillMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Skill entity.
// If the Skill object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *SkillMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *SkillMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// Where appends a list predicates to the SkillMutation builder.
func (m *SkillMutation) Where(ps ...predicate.Skill) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the SkillMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *SkillMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Skill, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *SkillMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *SkillMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Skill).
func (m *SkillMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *SkillMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.created_at != nil {
		fields = append(fields, skill.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, skill.FieldUpdatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *SkillMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case skill.FieldCreatedAt:
		return m.CreatedAt()
	case skill.FieldUpdatedAt:
		return m.UpdatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *SkillMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case skill.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case skill.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Skill field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *SkillMutation) SetField(name string, value ent.Value) error {
	switch name {
	case skill.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case skill.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Skill field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *SkillMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *SkillMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *SkillMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Skill numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *SkillMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *SkillMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *SkillMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Skill nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *SkillMutation) ResetField(name string) error {
	switch name {
	case skill.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case skill.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	}
	return fmt.Errorf("unknown Skill field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *SkillMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *SkillMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *SkillMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *SkillMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *SkillMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *SkillMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *SkillMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Skill unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *SkillMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Skill edge %s", name)
}
