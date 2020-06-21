// Code generated by SQLBoiler 3.7.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// RegisterdRole is an object representing the database table.
type RegisterdRole struct {
	ID      int64 `boil:"id" json:"id" toml:"id" yaml:"id"`
	GuildID int64 `boil:"guild_id" json:"guild_id" toml:"guild_id" yaml:"guild_id"`

	R *registerdRoleR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L registerdRoleL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var RegisterdRoleColumns = struct {
	ID      string
	GuildID string
}{
	ID:      "id",
	GuildID: "guild_id",
}

// Generated where

var RegisterdRoleWhere = struct {
	ID      whereHelperint64
	GuildID whereHelperint64
}{
	ID:      whereHelperint64{field: "\"registerd_roles\".\"id\""},
	GuildID: whereHelperint64{field: "\"registerd_roles\".\"guild_id\""},
}

// RegisterdRoleRels is where relationship names are stored.
var RegisterdRoleRels = struct {
	Guild string
}{
	Guild: "Guild",
}

// registerdRoleR is where relationships are stored.
type registerdRoleR struct {
	Guild *Guild
}

// NewStruct creates a new relationship struct
func (*registerdRoleR) NewStruct() *registerdRoleR {
	return &registerdRoleR{}
}

// registerdRoleL is where Load methods for each relationship are stored.
type registerdRoleL struct{}

var (
	registerdRoleAllColumns            = []string{"id", "guild_id"}
	registerdRoleColumnsWithoutDefault = []string{"id", "guild_id"}
	registerdRoleColumnsWithDefault    = []string{}
	registerdRolePrimaryKeyColumns     = []string{"id"}
)

type (
	// RegisterdRoleSlice is an alias for a slice of pointers to RegisterdRole.
	// This should generally be used opposed to []RegisterdRole.
	RegisterdRoleSlice []*RegisterdRole
	// RegisterdRoleHook is the signature for custom RegisterdRole hook methods
	RegisterdRoleHook func(context.Context, boil.ContextExecutor, *RegisterdRole) error

	registerdRoleQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	registerdRoleType                 = reflect.TypeOf(&RegisterdRole{})
	registerdRoleMapping              = queries.MakeStructMapping(registerdRoleType)
	registerdRolePrimaryKeyMapping, _ = queries.BindMapping(registerdRoleType, registerdRoleMapping, registerdRolePrimaryKeyColumns)
	registerdRoleInsertCacheMut       sync.RWMutex
	registerdRoleInsertCache          = make(map[string]insertCache)
	registerdRoleUpdateCacheMut       sync.RWMutex
	registerdRoleUpdateCache          = make(map[string]updateCache)
	registerdRoleUpsertCacheMut       sync.RWMutex
	registerdRoleUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var registerdRoleBeforeInsertHooks []RegisterdRoleHook
var registerdRoleBeforeUpdateHooks []RegisterdRoleHook
var registerdRoleBeforeDeleteHooks []RegisterdRoleHook
var registerdRoleBeforeUpsertHooks []RegisterdRoleHook

var registerdRoleAfterInsertHooks []RegisterdRoleHook
var registerdRoleAfterSelectHooks []RegisterdRoleHook
var registerdRoleAfterUpdateHooks []RegisterdRoleHook
var registerdRoleAfterDeleteHooks []RegisterdRoleHook
var registerdRoleAfterUpsertHooks []RegisterdRoleHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *RegisterdRole) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range registerdRoleBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *RegisterdRole) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range registerdRoleBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *RegisterdRole) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range registerdRoleBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *RegisterdRole) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range registerdRoleBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *RegisterdRole) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range registerdRoleAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *RegisterdRole) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range registerdRoleAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *RegisterdRole) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range registerdRoleAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *RegisterdRole) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range registerdRoleAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *RegisterdRole) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range registerdRoleAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddRegisterdRoleHook registers your hook function for all future operations.
func AddRegisterdRoleHook(hookPoint boil.HookPoint, registerdRoleHook RegisterdRoleHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		registerdRoleBeforeInsertHooks = append(registerdRoleBeforeInsertHooks, registerdRoleHook)
	case boil.BeforeUpdateHook:
		registerdRoleBeforeUpdateHooks = append(registerdRoleBeforeUpdateHooks, registerdRoleHook)
	case boil.BeforeDeleteHook:
		registerdRoleBeforeDeleteHooks = append(registerdRoleBeforeDeleteHooks, registerdRoleHook)
	case boil.BeforeUpsertHook:
		registerdRoleBeforeUpsertHooks = append(registerdRoleBeforeUpsertHooks, registerdRoleHook)
	case boil.AfterInsertHook:
		registerdRoleAfterInsertHooks = append(registerdRoleAfterInsertHooks, registerdRoleHook)
	case boil.AfterSelectHook:
		registerdRoleAfterSelectHooks = append(registerdRoleAfterSelectHooks, registerdRoleHook)
	case boil.AfterUpdateHook:
		registerdRoleAfterUpdateHooks = append(registerdRoleAfterUpdateHooks, registerdRoleHook)
	case boil.AfterDeleteHook:
		registerdRoleAfterDeleteHooks = append(registerdRoleAfterDeleteHooks, registerdRoleHook)
	case boil.AfterUpsertHook:
		registerdRoleAfterUpsertHooks = append(registerdRoleAfterUpsertHooks, registerdRoleHook)
	}
}

// One returns a single registerdRole record from the query.
func (q registerdRoleQuery) One(ctx context.Context, exec boil.ContextExecutor) (*RegisterdRole, error) {
	o := &RegisterdRole{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for registerd_roles")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all RegisterdRole records from the query.
func (q registerdRoleQuery) All(ctx context.Context, exec boil.ContextExecutor) (RegisterdRoleSlice, error) {
	var o []*RegisterdRole

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to RegisterdRole slice")
	}

	if len(registerdRoleAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all RegisterdRole records in the query.
func (q registerdRoleQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count registerd_roles rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q registerdRoleQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if registerd_roles exists")
	}

	return count > 0, nil
}

// Guild pointed to by the foreign key.
func (o *RegisterdRole) Guild(mods ...qm.QueryMod) guildQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.GuildID),
	}

	queryMods = append(queryMods, mods...)

	query := Guilds(queryMods...)
	queries.SetFrom(query.Query, "\"guilds\"")

	return query
}

// LoadGuild allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (registerdRoleL) LoadGuild(ctx context.Context, e boil.ContextExecutor, singular bool, maybeRegisterdRole interface{}, mods queries.Applicator) error {
	var slice []*RegisterdRole
	var object *RegisterdRole

	if singular {
		object = maybeRegisterdRole.(*RegisterdRole)
	} else {
		slice = *maybeRegisterdRole.(*[]*RegisterdRole)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &registerdRoleR{}
		}
		args = append(args, object.GuildID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &registerdRoleR{}
			}

			for _, a := range args {
				if a == obj.GuildID {
					continue Outer
				}
			}

			args = append(args, obj.GuildID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`guilds`), qm.WhereIn(`guilds.id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Guild")
	}

	var resultSlice []*Guild
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Guild")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for guilds")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for guilds")
	}

	if len(registerdRoleAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Guild = foreign
		if foreign.R == nil {
			foreign.R = &guildR{}
		}
		foreign.R.RegisterdRole = object
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.GuildID == foreign.ID {
				local.R.Guild = foreign
				if foreign.R == nil {
					foreign.R = &guildR{}
				}
				foreign.R.RegisterdRole = local
				break
			}
		}
	}

	return nil
}

// SetGuild of the registerdRole to the related item.
// Sets o.R.Guild to related.
// Adds o to related.R.RegisterdRole.
func (o *RegisterdRole) SetGuild(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Guild) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"registerd_roles\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"guild_id"}),
		strmangle.WhereClause("\"", "\"", 2, registerdRolePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.GuildID = related.ID
	if o.R == nil {
		o.R = &registerdRoleR{
			Guild: related,
		}
	} else {
		o.R.Guild = related
	}

	if related.R == nil {
		related.R = &guildR{
			RegisterdRole: o,
		}
	} else {
		related.R.RegisterdRole = o
	}

	return nil
}

// RegisterdRoles retrieves all the records using an executor.
func RegisterdRoles(mods ...qm.QueryMod) registerdRoleQuery {
	mods = append(mods, qm.From("\"registerd_roles\""))
	return registerdRoleQuery{NewQuery(mods...)}
}

// FindRegisterdRole retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindRegisterdRole(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*RegisterdRole, error) {
	registerdRoleObj := &RegisterdRole{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"registerd_roles\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, registerdRoleObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from registerd_roles")
	}

	return registerdRoleObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *RegisterdRole) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no registerd_roles provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(registerdRoleColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	registerdRoleInsertCacheMut.RLock()
	cache, cached := registerdRoleInsertCache[key]
	registerdRoleInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			registerdRoleAllColumns,
			registerdRoleColumnsWithDefault,
			registerdRoleColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(registerdRoleType, registerdRoleMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(registerdRoleType, registerdRoleMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"registerd_roles\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"registerd_roles\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into registerd_roles")
	}

	if !cached {
		registerdRoleInsertCacheMut.Lock()
		registerdRoleInsertCache[key] = cache
		registerdRoleInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the RegisterdRole.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *RegisterdRole) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	registerdRoleUpdateCacheMut.RLock()
	cache, cached := registerdRoleUpdateCache[key]
	registerdRoleUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			registerdRoleAllColumns,
			registerdRolePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update registerd_roles, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"registerd_roles\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, registerdRolePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(registerdRoleType, registerdRoleMapping, append(wl, registerdRolePrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update registerd_roles row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for registerd_roles")
	}

	if !cached {
		registerdRoleUpdateCacheMut.Lock()
		registerdRoleUpdateCache[key] = cache
		registerdRoleUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q registerdRoleQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for registerd_roles")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for registerd_roles")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o RegisterdRoleSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), registerdRolePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"registerd_roles\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, registerdRolePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in registerdRole slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all registerdRole")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *RegisterdRole) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no registerd_roles provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(registerdRoleColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	registerdRoleUpsertCacheMut.RLock()
	cache, cached := registerdRoleUpsertCache[key]
	registerdRoleUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			registerdRoleAllColumns,
			registerdRoleColumnsWithDefault,
			registerdRoleColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			registerdRoleAllColumns,
			registerdRolePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert registerd_roles, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(registerdRolePrimaryKeyColumns))
			copy(conflict, registerdRolePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"registerd_roles\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(registerdRoleType, registerdRoleMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(registerdRoleType, registerdRoleMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert registerd_roles")
	}

	if !cached {
		registerdRoleUpsertCacheMut.Lock()
		registerdRoleUpsertCache[key] = cache
		registerdRoleUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single RegisterdRole record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *RegisterdRole) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no RegisterdRole provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), registerdRolePrimaryKeyMapping)
	sql := "DELETE FROM \"registerd_roles\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from registerd_roles")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for registerd_roles")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q registerdRoleQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no registerdRoleQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from registerd_roles")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for registerd_roles")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o RegisterdRoleSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(registerdRoleBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), registerdRolePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"registerd_roles\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, registerdRolePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from registerdRole slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for registerd_roles")
	}

	if len(registerdRoleAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *RegisterdRole) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindRegisterdRole(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *RegisterdRoleSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := RegisterdRoleSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), registerdRolePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"registerd_roles\".* FROM \"registerd_roles\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, registerdRolePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in RegisterdRoleSlice")
	}

	*o = slice

	return nil
}

// RegisterdRoleExists checks if the RegisterdRole row exists.
func RegisterdRoleExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"registerd_roles\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if registerd_roles exists")
	}

	return exists, nil
}
