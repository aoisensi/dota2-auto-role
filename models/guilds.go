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

// Guild is an object representing the database table.
type Guild struct {
	ID int64 `boil:"id" json:"id" toml:"id" yaml:"id"`

	R *guildR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L guildL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var GuildColumns = struct {
	ID string
}{
	ID: "id",
}

// Generated where

type whereHelperint64 struct{ field string }

func (w whereHelperint64) EQ(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint64) NEQ(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint64) LT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint64) LTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint64) GT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint64) GTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint64) IN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}

var GuildWhere = struct {
	ID whereHelperint64
}{
	ID: whereHelperint64{field: "\"guilds\".\"id\""},
}

// GuildRels is where relationship names are stored.
var GuildRels = struct {
	RegisterdRole string
	RankRoles     string
}{
	RegisterdRole: "RegisterdRole",
	RankRoles:     "RankRoles",
}

// guildR is where relationships are stored.
type guildR struct {
	RegisterdRole *RegisterdRole
	RankRoles     RankRoleSlice
}

// NewStruct creates a new relationship struct
func (*guildR) NewStruct() *guildR {
	return &guildR{}
}

// guildL is where Load methods for each relationship are stored.
type guildL struct{}

var (
	guildAllColumns            = []string{"id"}
	guildColumnsWithoutDefault = []string{"id"}
	guildColumnsWithDefault    = []string{}
	guildPrimaryKeyColumns     = []string{"id"}
)

type (
	// GuildSlice is an alias for a slice of pointers to Guild.
	// This should generally be used opposed to []Guild.
	GuildSlice []*Guild
	// GuildHook is the signature for custom Guild hook methods
	GuildHook func(context.Context, boil.ContextExecutor, *Guild) error

	guildQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	guildType                 = reflect.TypeOf(&Guild{})
	guildMapping              = queries.MakeStructMapping(guildType)
	guildPrimaryKeyMapping, _ = queries.BindMapping(guildType, guildMapping, guildPrimaryKeyColumns)
	guildInsertCacheMut       sync.RWMutex
	guildInsertCache          = make(map[string]insertCache)
	guildUpdateCacheMut       sync.RWMutex
	guildUpdateCache          = make(map[string]updateCache)
	guildUpsertCacheMut       sync.RWMutex
	guildUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var guildBeforeInsertHooks []GuildHook
var guildBeforeUpdateHooks []GuildHook
var guildBeforeDeleteHooks []GuildHook
var guildBeforeUpsertHooks []GuildHook

var guildAfterInsertHooks []GuildHook
var guildAfterSelectHooks []GuildHook
var guildAfterUpdateHooks []GuildHook
var guildAfterDeleteHooks []GuildHook
var guildAfterUpsertHooks []GuildHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Guild) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range guildBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Guild) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range guildBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Guild) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range guildBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Guild) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range guildBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Guild) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range guildAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Guild) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range guildAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Guild) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range guildAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Guild) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range guildAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Guild) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range guildAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddGuildHook registers your hook function for all future operations.
func AddGuildHook(hookPoint boil.HookPoint, guildHook GuildHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		guildBeforeInsertHooks = append(guildBeforeInsertHooks, guildHook)
	case boil.BeforeUpdateHook:
		guildBeforeUpdateHooks = append(guildBeforeUpdateHooks, guildHook)
	case boil.BeforeDeleteHook:
		guildBeforeDeleteHooks = append(guildBeforeDeleteHooks, guildHook)
	case boil.BeforeUpsertHook:
		guildBeforeUpsertHooks = append(guildBeforeUpsertHooks, guildHook)
	case boil.AfterInsertHook:
		guildAfterInsertHooks = append(guildAfterInsertHooks, guildHook)
	case boil.AfterSelectHook:
		guildAfterSelectHooks = append(guildAfterSelectHooks, guildHook)
	case boil.AfterUpdateHook:
		guildAfterUpdateHooks = append(guildAfterUpdateHooks, guildHook)
	case boil.AfterDeleteHook:
		guildAfterDeleteHooks = append(guildAfterDeleteHooks, guildHook)
	case boil.AfterUpsertHook:
		guildAfterUpsertHooks = append(guildAfterUpsertHooks, guildHook)
	}
}

// One returns a single guild record from the query.
func (q guildQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Guild, error) {
	o := &Guild{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for guilds")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Guild records from the query.
func (q guildQuery) All(ctx context.Context, exec boil.ContextExecutor) (GuildSlice, error) {
	var o []*Guild

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Guild slice")
	}

	if len(guildAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Guild records in the query.
func (q guildQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count guilds rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q guildQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if guilds exists")
	}

	return count > 0, nil
}

// RegisterdRole pointed to by the foreign key.
func (o *Guild) RegisterdRole(mods ...qm.QueryMod) registerdRoleQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"guild_id\" = ?", o.ID),
	}

	queryMods = append(queryMods, mods...)

	query := RegisterdRoles(queryMods...)
	queries.SetFrom(query.Query, "\"registerd_roles\"")

	return query
}

// RankRoles retrieves all the rank_role's RankRoles with an executor.
func (o *Guild) RankRoles(mods ...qm.QueryMod) rankRoleQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"rank_roles\".\"guild_id\"=?", o.ID),
	)

	query := RankRoles(queryMods...)
	queries.SetFrom(query.Query, "\"rank_roles\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"rank_roles\".*"})
	}

	return query
}

// LoadRegisterdRole allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-1 relationship.
func (guildL) LoadRegisterdRole(ctx context.Context, e boil.ContextExecutor, singular bool, maybeGuild interface{}, mods queries.Applicator) error {
	var slice []*Guild
	var object *Guild

	if singular {
		object = maybeGuild.(*Guild)
	} else {
		slice = *maybeGuild.(*[]*Guild)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &guildR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &guildR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`registerd_roles`), qm.WhereIn(`registerd_roles.guild_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load RegisterdRole")
	}

	var resultSlice []*RegisterdRole
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice RegisterdRole")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for registerd_roles")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for registerd_roles")
	}

	if len(guildAfterSelectHooks) != 0 {
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
		object.R.RegisterdRole = foreign
		if foreign.R == nil {
			foreign.R = &registerdRoleR{}
		}
		foreign.R.Guild = object
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ID == foreign.GuildID {
				local.R.RegisterdRole = foreign
				if foreign.R == nil {
					foreign.R = &registerdRoleR{}
				}
				foreign.R.Guild = local
				break
			}
		}
	}

	return nil
}

// LoadRankRoles allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (guildL) LoadRankRoles(ctx context.Context, e boil.ContextExecutor, singular bool, maybeGuild interface{}, mods queries.Applicator) error {
	var slice []*Guild
	var object *Guild

	if singular {
		object = maybeGuild.(*Guild)
	} else {
		slice = *maybeGuild.(*[]*Guild)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &guildR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &guildR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`rank_roles`), qm.WhereIn(`rank_roles.guild_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load rank_roles")
	}

	var resultSlice []*RankRole
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice rank_roles")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on rank_roles")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for rank_roles")
	}

	if len(rankRoleAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.RankRoles = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &rankRoleR{}
			}
			foreign.R.Guild = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.GuildID {
				local.R.RankRoles = append(local.R.RankRoles, foreign)
				if foreign.R == nil {
					foreign.R = &rankRoleR{}
				}
				foreign.R.Guild = local
				break
			}
		}
	}

	return nil
}

// SetRegisterdRole of the guild to the related item.
// Sets o.R.RegisterdRole to related.
// Adds o to related.R.Guild.
func (o *Guild) SetRegisterdRole(ctx context.Context, exec boil.ContextExecutor, insert bool, related *RegisterdRole) error {
	var err error

	if insert {
		related.GuildID = o.ID

		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"registerd_roles\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"guild_id"}),
			strmangle.WhereClause("\"", "\"", 2, registerdRolePrimaryKeyColumns),
		)
		values := []interface{}{o.ID, related.ID}

		if boil.IsDebug(ctx) {
			writer := boil.DebugWriterFrom(ctx)
			fmt.Fprintln(writer, updateQuery)
			fmt.Fprintln(writer, values)
		}
		if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.GuildID = o.ID

	}

	if o.R == nil {
		o.R = &guildR{
			RegisterdRole: related,
		}
	} else {
		o.R.RegisterdRole = related
	}

	if related.R == nil {
		related.R = &registerdRoleR{
			Guild: o,
		}
	} else {
		related.R.Guild = o
	}
	return nil
}

// AddRankRoles adds the given related objects to the existing relationships
// of the guild, optionally inserting them as new records.
// Appends related to o.R.RankRoles.
// Sets related.R.Guild appropriately.
func (o *Guild) AddRankRoles(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*RankRole) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.GuildID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"rank_roles\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"guild_id"}),
				strmangle.WhereClause("\"", "\"", 2, rankRolePrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.GuildID = o.ID
		}
	}

	if o.R == nil {
		o.R = &guildR{
			RankRoles: related,
		}
	} else {
		o.R.RankRoles = append(o.R.RankRoles, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &rankRoleR{
				Guild: o,
			}
		} else {
			rel.R.Guild = o
		}
	}
	return nil
}

// Guilds retrieves all the records using an executor.
func Guilds(mods ...qm.QueryMod) guildQuery {
	mods = append(mods, qm.From("\"guilds\""))
	return guildQuery{NewQuery(mods...)}
}

// FindGuild retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindGuild(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*Guild, error) {
	guildObj := &Guild{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"guilds\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, guildObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from guilds")
	}

	return guildObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Guild) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no guilds provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(guildColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	guildInsertCacheMut.RLock()
	cache, cached := guildInsertCache[key]
	guildInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			guildAllColumns,
			guildColumnsWithDefault,
			guildColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(guildType, guildMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(guildType, guildMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"guilds\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"guilds\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into guilds")
	}

	if !cached {
		guildInsertCacheMut.Lock()
		guildInsertCache[key] = cache
		guildInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Guild.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Guild) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	guildUpdateCacheMut.RLock()
	cache, cached := guildUpdateCache[key]
	guildUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			guildAllColumns,
			guildPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update guilds, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"guilds\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, guildPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(guildType, guildMapping, append(wl, guildPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update guilds row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for guilds")
	}

	if !cached {
		guildUpdateCacheMut.Lock()
		guildUpdateCache[key] = cache
		guildUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q guildQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for guilds")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for guilds")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o GuildSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), guildPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"guilds\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, guildPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in guild slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all guild")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Guild) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no guilds provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(guildColumnsWithDefault, o)

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

	guildUpsertCacheMut.RLock()
	cache, cached := guildUpsertCache[key]
	guildUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			guildAllColumns,
			guildColumnsWithDefault,
			guildColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			guildAllColumns,
			guildPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert guilds, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(guildPrimaryKeyColumns))
			copy(conflict, guildPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"guilds\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(guildType, guildMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(guildType, guildMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert guilds")
	}

	if !cached {
		guildUpsertCacheMut.Lock()
		guildUpsertCache[key] = cache
		guildUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Guild record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Guild) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Guild provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), guildPrimaryKeyMapping)
	sql := "DELETE FROM \"guilds\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from guilds")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for guilds")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q guildQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no guildQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from guilds")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for guilds")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o GuildSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(guildBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), guildPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"guilds\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, guildPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from guild slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for guilds")
	}

	if len(guildAfterDeleteHooks) != 0 {
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
func (o *Guild) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindGuild(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *GuildSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := GuildSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), guildPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"guilds\".* FROM \"guilds\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, guildPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in GuildSlice")
	}

	*o = slice

	return nil
}

// GuildExists checks if the Guild row exists.
func GuildExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"guilds\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if guilds exists")
	}

	return exists, nil
}
