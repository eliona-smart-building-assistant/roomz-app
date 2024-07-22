// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package appdb

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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/v4/types"
	"github.com/volatiletech/strmangle"
)

// Configuration is an object representing the database table.
type Configuration struct {
	ID         int64             `boil:"id" json:"id" toml:"id" yaml:"id"`
	Active     bool              `boil:"active" json:"active" toml:"active" yaml:"active"`
	Enable     bool              `boil:"enable" json:"enable" toml:"enable" yaml:"enable"`
	ProjectIds types.StringArray `boil:"project_ids" json:"project_ids" toml:"project_ids" yaml:"project_ids"`
	UserID     string            `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`

	R *configurationR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L configurationL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ConfigurationColumns = struct {
	ID         string
	Active     string
	Enable     string
	ProjectIds string
	UserID     string
}{
	ID:         "id",
	Active:     "active",
	Enable:     "enable",
	ProjectIds: "project_ids",
	UserID:     "user_id",
}

var ConfigurationTableColumns = struct {
	ID         string
	Active     string
	Enable     string
	ProjectIds string
	UserID     string
}{
	ID:         "configuration.id",
	Active:     "configuration.active",
	Enable:     "configuration.enable",
	ProjectIds: "configuration.project_ids",
	UserID:     "configuration.user_id",
}

// Generated where

type whereHelperbool struct{ field string }

func (w whereHelperbool) EQ(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperbool) NEQ(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperbool) LT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperbool) LTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperbool) GT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperbool) GTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

type whereHelpertypes_StringArray struct{ field string }

func (w whereHelpertypes_StringArray) EQ(x types.StringArray) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertypes_StringArray) NEQ(x types.StringArray) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertypes_StringArray) LT(x types.StringArray) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertypes_StringArray) LTE(x types.StringArray) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertypes_StringArray) GT(x types.StringArray) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertypes_StringArray) GTE(x types.StringArray) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var ConfigurationWhere = struct {
	ID         whereHelperint64
	Active     whereHelperbool
	Enable     whereHelperbool
	ProjectIds whereHelpertypes_StringArray
	UserID     whereHelperstring
}{
	ID:         whereHelperint64{field: "\"roomz\".\"configuration\".\"id\""},
	Active:     whereHelperbool{field: "\"roomz\".\"configuration\".\"active\""},
	Enable:     whereHelperbool{field: "\"roomz\".\"configuration\".\"enable\""},
	ProjectIds: whereHelpertypes_StringArray{field: "\"roomz\".\"configuration\".\"project_ids\""},
	UserID:     whereHelperstring{field: "\"roomz\".\"configuration\".\"user_id\""},
}

// ConfigurationRels is where relationship names are stored.
var ConfigurationRels = struct {
	Assets string
}{
	Assets: "Assets",
}

// configurationR is where relationships are stored.
type configurationR struct {
	Assets AssetSlice `boil:"Assets" json:"Assets" toml:"Assets" yaml:"Assets"`
}

// NewStruct creates a new relationship struct
func (*configurationR) NewStruct() *configurationR {
	return &configurationR{}
}

func (r *configurationR) GetAssets() AssetSlice {
	if r == nil {
		return nil
	}
	return r.Assets
}

// configurationL is where Load methods for each relationship are stored.
type configurationL struct{}

var (
	configurationAllColumns            = []string{"id", "active", "enable", "project_ids", "user_id"}
	configurationColumnsWithoutDefault = []string{"project_ids", "user_id"}
	configurationColumnsWithDefault    = []string{"id", "active", "enable"}
	configurationPrimaryKeyColumns     = []string{"id"}
	configurationGeneratedColumns      = []string{}
)

type (
	// ConfigurationSlice is an alias for a slice of pointers to Configuration.
	// This should almost always be used instead of []Configuration.
	ConfigurationSlice []*Configuration
	// ConfigurationHook is the signature for custom Configuration hook methods
	ConfigurationHook func(context.Context, boil.ContextExecutor, *Configuration) error

	configurationQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	configurationType                 = reflect.TypeOf(&Configuration{})
	configurationMapping              = queries.MakeStructMapping(configurationType)
	configurationPrimaryKeyMapping, _ = queries.BindMapping(configurationType, configurationMapping, configurationPrimaryKeyColumns)
	configurationInsertCacheMut       sync.RWMutex
	configurationInsertCache          = make(map[string]insertCache)
	configurationUpdateCacheMut       sync.RWMutex
	configurationUpdateCache          = make(map[string]updateCache)
	configurationUpsertCacheMut       sync.RWMutex
	configurationUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var configurationAfterSelectMu sync.Mutex
var configurationAfterSelectHooks []ConfigurationHook

var configurationBeforeInsertMu sync.Mutex
var configurationBeforeInsertHooks []ConfigurationHook
var configurationAfterInsertMu sync.Mutex
var configurationAfterInsertHooks []ConfigurationHook

var configurationBeforeUpdateMu sync.Mutex
var configurationBeforeUpdateHooks []ConfigurationHook
var configurationAfterUpdateMu sync.Mutex
var configurationAfterUpdateHooks []ConfigurationHook

var configurationBeforeDeleteMu sync.Mutex
var configurationBeforeDeleteHooks []ConfigurationHook
var configurationAfterDeleteMu sync.Mutex
var configurationAfterDeleteHooks []ConfigurationHook

var configurationBeforeUpsertMu sync.Mutex
var configurationBeforeUpsertHooks []ConfigurationHook
var configurationAfterUpsertMu sync.Mutex
var configurationAfterUpsertHooks []ConfigurationHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Configuration) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range configurationAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Configuration) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range configurationBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Configuration) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range configurationAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Configuration) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range configurationBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Configuration) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range configurationAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Configuration) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range configurationBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Configuration) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range configurationAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Configuration) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range configurationBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Configuration) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range configurationAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddConfigurationHook registers your hook function for all future operations.
func AddConfigurationHook(hookPoint boil.HookPoint, configurationHook ConfigurationHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		configurationAfterSelectMu.Lock()
		configurationAfterSelectHooks = append(configurationAfterSelectHooks, configurationHook)
		configurationAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		configurationBeforeInsertMu.Lock()
		configurationBeforeInsertHooks = append(configurationBeforeInsertHooks, configurationHook)
		configurationBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		configurationAfterInsertMu.Lock()
		configurationAfterInsertHooks = append(configurationAfterInsertHooks, configurationHook)
		configurationAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		configurationBeforeUpdateMu.Lock()
		configurationBeforeUpdateHooks = append(configurationBeforeUpdateHooks, configurationHook)
		configurationBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		configurationAfterUpdateMu.Lock()
		configurationAfterUpdateHooks = append(configurationAfterUpdateHooks, configurationHook)
		configurationAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		configurationBeforeDeleteMu.Lock()
		configurationBeforeDeleteHooks = append(configurationBeforeDeleteHooks, configurationHook)
		configurationBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		configurationAfterDeleteMu.Lock()
		configurationAfterDeleteHooks = append(configurationAfterDeleteHooks, configurationHook)
		configurationAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		configurationBeforeUpsertMu.Lock()
		configurationBeforeUpsertHooks = append(configurationBeforeUpsertHooks, configurationHook)
		configurationBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		configurationAfterUpsertMu.Lock()
		configurationAfterUpsertHooks = append(configurationAfterUpsertHooks, configurationHook)
		configurationAfterUpsertMu.Unlock()
	}
}

// OneG returns a single configuration record from the query using the global executor.
func (q configurationQuery) OneG(ctx context.Context) (*Configuration, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single configuration record from the query.
func (q configurationQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Configuration, error) {
	o := &Configuration{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "appdb: failed to execute a one query for configuration")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all Configuration records from the query using the global executor.
func (q configurationQuery) AllG(ctx context.Context) (ConfigurationSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all Configuration records from the query.
func (q configurationQuery) All(ctx context.Context, exec boil.ContextExecutor) (ConfigurationSlice, error) {
	var o []*Configuration

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "appdb: failed to assign all query results to Configuration slice")
	}

	if len(configurationAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all Configuration records in the query using the global executor
func (q configurationQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all Configuration records in the query.
func (q configurationQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "appdb: failed to count configuration rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q configurationQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q configurationQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "appdb: failed to check if configuration exists")
	}

	return count > 0, nil
}

// Assets retrieves all the asset's Assets with an executor.
func (o *Configuration) Assets(mods ...qm.QueryMod) assetQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"roomz\".\"asset\".\"configuration_id\"=?", o.ID),
	)

	return Assets(queryMods...)
}

// LoadAssets allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (configurationL) LoadAssets(ctx context.Context, e boil.ContextExecutor, singular bool, maybeConfiguration interface{}, mods queries.Applicator) error {
	var slice []*Configuration
	var object *Configuration

	if singular {
		var ok bool
		object, ok = maybeConfiguration.(*Configuration)
		if !ok {
			object = new(Configuration)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeConfiguration)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeConfiguration))
			}
		}
	} else {
		s, ok := maybeConfiguration.(*[]*Configuration)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeConfiguration)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeConfiguration))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &configurationR{}
		}
		args[object.ID] = struct{}{}
	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &configurationR{}
			}
			args[obj.ID] = struct{}{}
		}
	}

	if len(args) == 0 {
		return nil
	}

	argsSlice := make([]interface{}, len(args))
	i := 0
	for arg := range args {
		argsSlice[i] = arg
		i++
	}

	query := NewQuery(
		qm.From(`roomz.asset`),
		qm.WhereIn(`roomz.asset.configuration_id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load asset")
	}

	var resultSlice []*Asset
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice asset")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on asset")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for asset")
	}

	if len(assetAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Assets = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &assetR{}
			}
			foreign.R.Configuration = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.ConfigurationID {
				local.R.Assets = append(local.R.Assets, foreign)
				if foreign.R == nil {
					foreign.R = &assetR{}
				}
				foreign.R.Configuration = local
				break
			}
		}
	}

	return nil
}

// AddAssetsG adds the given related objects to the existing relationships
// of the configuration, optionally inserting them as new records.
// Appends related to o.R.Assets.
// Sets related.R.Configuration appropriately.
// Uses the global database handle.
func (o *Configuration) AddAssetsG(ctx context.Context, insert bool, related ...*Asset) error {
	return o.AddAssets(ctx, boil.GetContextDB(), insert, related...)
}

// AddAssets adds the given related objects to the existing relationships
// of the configuration, optionally inserting them as new records.
// Appends related to o.R.Assets.
// Sets related.R.Configuration appropriately.
func (o *Configuration) AddAssets(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Asset) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.ConfigurationID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"roomz\".\"asset\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"configuration_id"}),
				strmangle.WhereClause("\"", "\"", 2, assetPrimaryKeyColumns),
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

			rel.ConfigurationID = o.ID
		}
	}

	if o.R == nil {
		o.R = &configurationR{
			Assets: related,
		}
	} else {
		o.R.Assets = append(o.R.Assets, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &assetR{
				Configuration: o,
			}
		} else {
			rel.R.Configuration = o
		}
	}
	return nil
}

// Configurations retrieves all the records using an executor.
func Configurations(mods ...qm.QueryMod) configurationQuery {
	mods = append(mods, qm.From("\"roomz\".\"configuration\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"roomz\".\"configuration\".*"})
	}

	return configurationQuery{q}
}

// FindConfigurationG retrieves a single record by ID.
func FindConfigurationG(ctx context.Context, iD int64, selectCols ...string) (*Configuration, error) {
	return FindConfiguration(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindConfiguration retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindConfiguration(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*Configuration, error) {
	configurationObj := &Configuration{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"roomz\".\"configuration\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, configurationObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "appdb: unable to select from configuration")
	}

	if err = configurationObj.doAfterSelectHooks(ctx, exec); err != nil {
		return configurationObj, err
	}

	return configurationObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Configuration) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Configuration) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("appdb: no configuration provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(configurationColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	configurationInsertCacheMut.RLock()
	cache, cached := configurationInsertCache[key]
	configurationInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			configurationAllColumns,
			configurationColumnsWithDefault,
			configurationColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(configurationType, configurationMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(configurationType, configurationMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"roomz\".\"configuration\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"roomz\".\"configuration\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "appdb: unable to insert into configuration")
	}

	if !cached {
		configurationInsertCacheMut.Lock()
		configurationInsertCache[key] = cache
		configurationInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single Configuration record using the global executor.
// See Update for more documentation.
func (o *Configuration) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the Configuration.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Configuration) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	configurationUpdateCacheMut.RLock()
	cache, cached := configurationUpdateCache[key]
	configurationUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			configurationAllColumns,
			configurationPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("appdb: unable to update configuration, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"roomz\".\"configuration\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, configurationPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(configurationType, configurationMapping, append(wl, configurationPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "appdb: unable to update configuration row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "appdb: failed to get rows affected by update for configuration")
	}

	if !cached {
		configurationUpdateCacheMut.Lock()
		configurationUpdateCache[key] = cache
		configurationUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q configurationQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q configurationQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "appdb: unable to update all for configuration")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "appdb: unable to retrieve rows affected for configuration")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o ConfigurationSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ConfigurationSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("appdb: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), configurationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"roomz\".\"configuration\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, configurationPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "appdb: unable to update all in configuration slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "appdb: unable to retrieve rows affected all in update all configuration")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Configuration) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns, opts...)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Configuration) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("appdb: no configuration provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(configurationColumnsWithDefault, o)

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

	configurationUpsertCacheMut.RLock()
	cache, cached := configurationUpsertCache[key]
	configurationUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			configurationAllColumns,
			configurationColumnsWithDefault,
			configurationColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			configurationAllColumns,
			configurationPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("appdb: unable to upsert configuration, could not build update column list")
		}

		ret := strmangle.SetComplement(configurationAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(configurationPrimaryKeyColumns) == 0 {
				return errors.New("appdb: unable to upsert configuration, could not build conflict column list")
			}

			conflict = make([]string, len(configurationPrimaryKeyColumns))
			copy(conflict, configurationPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"roomz\".\"configuration\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(configurationType, configurationMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(configurationType, configurationMapping, ret)
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
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "appdb: unable to upsert configuration")
	}

	if !cached {
		configurationUpsertCacheMut.Lock()
		configurationUpsertCache[key] = cache
		configurationUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single Configuration record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Configuration) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single Configuration record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Configuration) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("appdb: no Configuration provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), configurationPrimaryKeyMapping)
	sql := "DELETE FROM \"roomz\".\"configuration\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "appdb: unable to delete from configuration")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "appdb: failed to get rows affected by delete for configuration")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q configurationQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q configurationQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("appdb: no configurationQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "appdb: unable to delete all from configuration")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "appdb: failed to get rows affected by deleteall for configuration")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o ConfigurationSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ConfigurationSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(configurationBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), configurationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"roomz\".\"configuration\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, configurationPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "appdb: unable to delete all from configuration slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "appdb: failed to get rows affected by deleteall for configuration")
	}

	if len(configurationAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Configuration) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("appdb: no Configuration provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Configuration) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindConfiguration(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ConfigurationSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("appdb: empty ConfigurationSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ConfigurationSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ConfigurationSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), configurationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"roomz\".\"configuration\".* FROM \"roomz\".\"configuration\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, configurationPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "appdb: unable to reload all in ConfigurationSlice")
	}

	*o = slice

	return nil
}

// ConfigurationExistsG checks if the Configuration row exists.
func ConfigurationExistsG(ctx context.Context, iD int64) (bool, error) {
	return ConfigurationExists(ctx, boil.GetContextDB(), iD)
}

// ConfigurationExists checks if the Configuration row exists.
func ConfigurationExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"roomz\".\"configuration\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "appdb: unable to check if configuration exists")
	}

	return exists, nil
}

// Exists checks if the Configuration row exists.
func (o *Configuration) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return ConfigurationExists(ctx, exec, o.ID)
}
