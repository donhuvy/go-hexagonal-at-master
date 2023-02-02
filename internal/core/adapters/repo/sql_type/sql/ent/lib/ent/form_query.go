// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"hex-base/internal/core/adapters/repo/sql_type/sql/ent/lib/ent/form"
	"hex-base/internal/core/adapters/repo/sql_type/sql/ent/lib/ent/predicate"
	"hex-base/internal/core/adapters/repo/sql_type/sql/ent/lib/ent/todo"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FormQuery is the builder for querying Form entities.
type FormQuery struct {
	config
	ctx        *QueryContext
	order      []OrderFunc
	inters     []Interceptor
	predicates []predicate.Form
	withTodo   *TodoQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FormQuery builder.
func (fq *FormQuery) Where(ps ...predicate.Form) *FormQuery {
	fq.predicates = append(fq.predicates, ps...)
	return fq
}

// Limit the number of records to be returned by this query.
func (fq *FormQuery) Limit(limit int) *FormQuery {
	fq.ctx.Limit = &limit
	return fq
}

// Offset to start from.
func (fq *FormQuery) Offset(offset int) *FormQuery {
	fq.ctx.Offset = &offset
	return fq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (fq *FormQuery) Unique(unique bool) *FormQuery {
	fq.ctx.Unique = &unique
	return fq
}

// Order specifies how the records should be ordered.
func (fq *FormQuery) Order(o ...OrderFunc) *FormQuery {
	fq.order = append(fq.order, o...)
	return fq
}

// QueryTodo chains the current query on the "todo" edge.
func (fq *FormQuery) QueryTodo() *TodoQuery {
	query := (&TodoClient{config: fq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(form.Table, form.FieldID, selector),
			sqlgraph.To(todo.Table, todo.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, form.TodoTable, form.TodoColumn),
		)
		fromU = sqlgraph.SetNeighbors(fq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Form entity from the query.
// Returns a *NotFoundError when no Form was found.
func (fq *FormQuery) First(ctx context.Context) (*Form, error) {
	nodes, err := fq.Limit(1).All(setContextOp(ctx, fq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{form.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fq *FormQuery) FirstX(ctx context.Context) *Form {
	node, err := fq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Form ID from the query.
// Returns a *NotFoundError when no Form ID was found.
func (fq *FormQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = fq.Limit(1).IDs(setContextOp(ctx, fq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{form.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (fq *FormQuery) FirstIDX(ctx context.Context) int64 {
	id, err := fq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Form entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Form entity is found.
// Returns a *NotFoundError when no Form entities are found.
func (fq *FormQuery) Only(ctx context.Context) (*Form, error) {
	nodes, err := fq.Limit(2).All(setContextOp(ctx, fq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{form.Label}
	default:
		return nil, &NotSingularError{form.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fq *FormQuery) OnlyX(ctx context.Context) *Form {
	node, err := fq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Form ID in the query.
// Returns a *NotSingularError when more than one Form ID is found.
// Returns a *NotFoundError when no entities are found.
func (fq *FormQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = fq.Limit(2).IDs(setContextOp(ctx, fq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{form.Label}
	default:
		err = &NotSingularError{form.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (fq *FormQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := fq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Forms.
func (fq *FormQuery) All(ctx context.Context) ([]*Form, error) {
	ctx = setContextOp(ctx, fq.ctx, "All")
	if err := fq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Form, *FormQuery]()
	return withInterceptors[[]*Form](ctx, fq, qr, fq.inters)
}

// AllX is like All, but panics if an error occurs.
func (fq *FormQuery) AllX(ctx context.Context) []*Form {
	nodes, err := fq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Form IDs.
func (fq *FormQuery) IDs(ctx context.Context) ([]int64, error) {
	var ids []int64
	ctx = setContextOp(ctx, fq.ctx, "IDs")
	if err := fq.Select(form.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (fq *FormQuery) IDsX(ctx context.Context) []int64 {
	ids, err := fq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fq *FormQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, fq.ctx, "Count")
	if err := fq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, fq, querierCount[*FormQuery](), fq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (fq *FormQuery) CountX(ctx context.Context) int {
	count, err := fq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fq *FormQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, fq.ctx, "Exist")
	switch _, err := fq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (fq *FormQuery) ExistX(ctx context.Context) bool {
	exist, err := fq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FormQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fq *FormQuery) Clone() *FormQuery {
	if fq == nil {
		return nil
	}
	return &FormQuery{
		config:     fq.config,
		ctx:        fq.ctx.Clone(),
		order:      append([]OrderFunc{}, fq.order...),
		inters:     append([]Interceptor{}, fq.inters...),
		predicates: append([]predicate.Form{}, fq.predicates...),
		withTodo:   fq.withTodo.Clone(),
		// clone intermediate query.
		sql:  fq.sql.Clone(),
		path: fq.path,
	}
}

// WithTodo tells the query-builder to eager-load the nodes that are connected to
// the "todo" edge. The optional arguments are used to configure the query builder of the edge.
func (fq *FormQuery) WithTodo(opts ...func(*TodoQuery)) *FormQuery {
	query := (&TodoClient{config: fq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	fq.withTodo = query
	return fq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Category string `json:"category,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Form.Query().
//		GroupBy(form.FieldCategory).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (fq *FormQuery) GroupBy(field string, fields ...string) *FormGroupBy {
	fq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &FormGroupBy{build: fq}
	grbuild.flds = &fq.ctx.Fields
	grbuild.label = form.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Category string `json:"category,omitempty"`
//	}
//
//	client.Form.Query().
//		Select(form.FieldCategory).
//		Scan(ctx, &v)
func (fq *FormQuery) Select(fields ...string) *FormSelect {
	fq.ctx.Fields = append(fq.ctx.Fields, fields...)
	sbuild := &FormSelect{FormQuery: fq}
	sbuild.label = form.Label
	sbuild.flds, sbuild.scan = &fq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a FormSelect configured with the given aggregations.
func (fq *FormQuery) Aggregate(fns ...AggregateFunc) *FormSelect {
	return fq.Select().Aggregate(fns...)
}

func (fq *FormQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range fq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, fq); err != nil {
				return err
			}
		}
	}
	for _, f := range fq.ctx.Fields {
		if !form.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if fq.path != nil {
		prev, err := fq.path(ctx)
		if err != nil {
			return err
		}
		fq.sql = prev
	}
	return nil
}

func (fq *FormQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Form, error) {
	var (
		nodes       = []*Form{}
		_spec       = fq.querySpec()
		loadedTypes = [1]bool{
			fq.withTodo != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Form).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Form{config: fq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, fq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := fq.withTodo; query != nil {
		if err := fq.loadTodo(ctx, query, nodes, nil,
			func(n *Form, e *Todo) { n.Edges.Todo = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (fq *FormQuery) loadTodo(ctx context.Context, query *TodoQuery, nodes []*Form, init func(*Form), assign func(*Form, *Todo)) error {
	ids := make([]int32, 0, len(nodes))
	nodeids := make(map[int32][]*Form)
	for i := range nodes {
		fk := nodes[i].TodoID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(todo.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "todo_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (fq *FormQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := fq.querySpec()
	_spec.Node.Columns = fq.ctx.Fields
	if len(fq.ctx.Fields) > 0 {
		_spec.Unique = fq.ctx.Unique != nil && *fq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, fq.driver, _spec)
}

func (fq *FormQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   form.Table,
			Columns: form.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: form.FieldID,
			},
		},
		From:   fq.sql,
		Unique: true,
	}
	if unique := fq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := fq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, form.FieldID)
		for i := range fields {
			if fields[i] != form.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := fq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := fq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := fq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := fq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (fq *FormQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(fq.driver.Dialect())
	t1 := builder.Table(form.Table)
	columns := fq.ctx.Fields
	if len(columns) == 0 {
		columns = form.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if fq.sql != nil {
		selector = fq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if fq.ctx.Unique != nil && *fq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range fq.predicates {
		p(selector)
	}
	for _, p := range fq.order {
		p(selector)
	}
	if offset := fq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := fq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// FormGroupBy is the group-by builder for Form entities.
type FormGroupBy struct {
	selector
	build *FormQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fgb *FormGroupBy) Aggregate(fns ...AggregateFunc) *FormGroupBy {
	fgb.fns = append(fgb.fns, fns...)
	return fgb
}

// Scan applies the selector query and scans the result into the given value.
func (fgb *FormGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, fgb.build.ctx, "GroupBy")
	if err := fgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FormQuery, *FormGroupBy](ctx, fgb.build, fgb, fgb.build.inters, v)
}

func (fgb *FormGroupBy) sqlScan(ctx context.Context, root *FormQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(fgb.fns))
	for _, fn := range fgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*fgb.flds)+len(fgb.fns))
		for _, f := range *fgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*fgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// FormSelect is the builder for selecting fields of Form entities.
type FormSelect struct {
	*FormQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (fs *FormSelect) Aggregate(fns ...AggregateFunc) *FormSelect {
	fs.fns = append(fs.fns, fns...)
	return fs
}

// Scan applies the selector query and scans the result into the given value.
func (fs *FormSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, fs.ctx, "Select")
	if err := fs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FormQuery, *FormSelect](ctx, fs.FormQuery, fs, fs.inters, v)
}

func (fs *FormSelect) sqlScan(ctx context.Context, root *FormQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(fs.fns))
	for _, fn := range fs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*fs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
