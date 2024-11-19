// Code generated by SQLBoiler 4.14.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package master

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
	"github.com/volatiletech/strmangle"
)

// Character is an object representing the database table.
type Character struct {
	ID     uint32 `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name   string `boil:"name" json:"name" toml:"name" yaml:"name"`
	Gender int32  `boil:"gender" json:"gender" toml:"gender" yaml:"gender"`
	Age    int32  `boil:"age" json:"age" toml:"age" yaml:"age"`

	R *characterR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L characterL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CharacterColumns = struct {
	ID     string
	Name   string
	Gender string
	Age    string
}{
	ID:     "id",
	Name:   "name",
	Gender: "gender",
	Age:    "age",
}

var CharacterTableColumns = struct {
	ID     string
	Name   string
	Gender string
	Age    string
}{
	ID:     "character.id",
	Name:   "character.name",
	Gender: "character.gender",
	Age:    "character.age",
}

// Generated where

var CharacterWhere = struct {
	ID     whereHelperuint32
	Name   whereHelperstring
	Gender whereHelperint32
	Age    whereHelperint32
}{
	ID:     whereHelperuint32{field: "`character`.`id`"},
	Name:   whereHelperstring{field: "`character`.`name`"},
	Gender: whereHelperint32{field: "`character`.`gender`"},
	Age:    whereHelperint32{field: "`character`.`age`"},
}

// CharacterRels is where relationship names are stored.
var CharacterRels = struct {
}{}

// characterR is where relationships are stored.
type characterR struct {
}

// NewStruct creates a new relationship struct
func (*characterR) NewStruct() *characterR {
	return &characterR{}
}

// characterL is where Load methods for each relationship are stored.
type characterL struct{}

var (
	characterAllColumns            = []string{"id", "name", "gender", "age"}
	characterColumnsWithoutDefault = []string{"id", "name", "gender", "age"}
	characterColumnsWithDefault    = []string{}
	characterPrimaryKeyColumns     = []string{"id"}
	characterGeneratedColumns      = []string{}
)

type (
	// CharacterSlice is an alias for a slice of pointers to Character.
	// This should almost always be used instead of []Character.
	CharacterSlice []*Character
	// CharacterHook is the signature for custom Character hook methods
	CharacterHook func(context.Context, boil.ContextExecutor, *Character) error

	characterQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	characterType                 = reflect.TypeOf(&Character{})
	characterMapping              = queries.MakeStructMapping(characterType)
	characterPrimaryKeyMapping, _ = queries.BindMapping(characterType, characterMapping, characterPrimaryKeyColumns)
	characterInsertCacheMut       sync.RWMutex
	characterInsertCache          = make(map[string]insertCache)
	characterUpdateCacheMut       sync.RWMutex
	characterUpdateCache          = make(map[string]updateCache)
	characterUpsertCacheMut       sync.RWMutex
	characterUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var characterAfterSelectHooks []CharacterHook

var characterBeforeInsertHooks []CharacterHook
var characterAfterInsertHooks []CharacterHook

var characterBeforeUpdateHooks []CharacterHook
var characterAfterUpdateHooks []CharacterHook

var characterBeforeDeleteHooks []CharacterHook
var characterAfterDeleteHooks []CharacterHook

var characterBeforeUpsertHooks []CharacterHook
var characterAfterUpsertHooks []CharacterHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Character) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range characterAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Character) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range characterBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Character) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range characterAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Character) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range characterBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Character) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range characterAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Character) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range characterBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Character) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range characterAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Character) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range characterBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Character) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range characterAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCharacterHook registers your hook function for all future operations.
func AddCharacterHook(hookPoint boil.HookPoint, characterHook CharacterHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		characterAfterSelectHooks = append(characterAfterSelectHooks, characterHook)
	case boil.BeforeInsertHook:
		characterBeforeInsertHooks = append(characterBeforeInsertHooks, characterHook)
	case boil.AfterInsertHook:
		characterAfterInsertHooks = append(characterAfterInsertHooks, characterHook)
	case boil.BeforeUpdateHook:
		characterBeforeUpdateHooks = append(characterBeforeUpdateHooks, characterHook)
	case boil.AfterUpdateHook:
		characterAfterUpdateHooks = append(characterAfterUpdateHooks, characterHook)
	case boil.BeforeDeleteHook:
		characterBeforeDeleteHooks = append(characterBeforeDeleteHooks, characterHook)
	case boil.AfterDeleteHook:
		characterAfterDeleteHooks = append(characterAfterDeleteHooks, characterHook)
	case boil.BeforeUpsertHook:
		characterBeforeUpsertHooks = append(characterBeforeUpsertHooks, characterHook)
	case boil.AfterUpsertHook:
		characterAfterUpsertHooks = append(characterAfterUpsertHooks, characterHook)
	}
}

// One returns a single character record from the query.
func (q characterQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Character, error) {
	o := &Character{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "master: failed to execute a one query for character")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Character records from the query.
func (q characterQuery) All(ctx context.Context, exec boil.ContextExecutor) (CharacterSlice, error) {
	var o []*Character

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "master: failed to assign all query results to Character slice")
	}

	if len(characterAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Character records in the query.
func (q characterQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "master: failed to count character rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q characterQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "master: failed to check if character exists")
	}

	return count > 0, nil
}

// Characters retrieves all the records using an executor.
func Characters(mods ...qm.QueryMod) characterQuery {
	mods = append(mods, qm.From("`character`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`character`.*"})
	}

	return characterQuery{q}
}

// FindCharacter retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCharacter(ctx context.Context, exec boil.ContextExecutor, iD uint32, selectCols ...string) (*Character, error) {
	characterObj := &Character{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `character` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, characterObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "master: unable to select from character")
	}

	if err = characterObj.doAfterSelectHooks(ctx, exec); err != nil {
		return characterObj, err
	}

	return characterObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Character) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("master: no character provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(characterColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	characterInsertCacheMut.RLock()
	cache, cached := characterInsertCache[key]
	characterInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			characterAllColumns,
			characterColumnsWithDefault,
			characterColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(characterType, characterMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(characterType, characterMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `character` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `character` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `character` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, characterPrimaryKeyColumns))
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
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "master: unable to insert into character")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "master: unable to populate default values for character")
	}

CacheNoHooks:
	if !cached {
		characterInsertCacheMut.Lock()
		characterInsertCache[key] = cache
		characterInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Character.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Character) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	characterUpdateCacheMut.RLock()
	cache, cached := characterUpdateCache[key]
	characterUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			characterAllColumns,
			characterPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("master: unable to update character, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `character` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, characterPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(characterType, characterMapping, append(wl, characterPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "master: unable to update character row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "master: failed to get rows affected by update for character")
	}

	if !cached {
		characterUpdateCacheMut.Lock()
		characterUpdateCache[key] = cache
		characterUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q characterQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "master: unable to update all for character")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "master: unable to retrieve rows affected for character")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CharacterSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("master: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), characterPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `character` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, characterPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "master: unable to update all in character slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "master: unable to retrieve rows affected all in update all character")
	}
	return rowsAff, nil
}

// Delete deletes a single Character record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Character) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("master: no Character provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), characterPrimaryKeyMapping)
	sql := "DELETE FROM `character` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "master: unable to delete from character")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "master: failed to get rows affected by delete for character")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q characterQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("master: no characterQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "master: unable to delete all from character")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "master: failed to get rows affected by deleteall for character")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CharacterSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(characterBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), characterPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `character` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, characterPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "master: unable to delete all from character slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "master: failed to get rows affected by deleteall for character")
	}

	if len(characterAfterDeleteHooks) != 0 {
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
func (o *Character) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCharacter(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CharacterSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CharacterSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), characterPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `character`.* FROM `character` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, characterPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "master: unable to reload all in CharacterSlice")
	}

	*o = slice

	return nil
}

// CharacterExists checks if the Character row exists.
func CharacterExists(ctx context.Context, exec boil.ContextExecutor, iD uint32) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `character` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "master: unable to check if character exists")
	}

	return exists, nil
}

// Exists checks if the Character row exists.
func (o *Character) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return CharacterExists(ctx, exec, o.ID)
}

var mySQLCharacterUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Character) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("master: no character provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(characterColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLCharacterUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
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
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	characterUpsertCacheMut.RLock()
	cache, cached := characterUpsertCache[key]
	characterUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			characterAllColumns,
			characterColumnsWithDefault,
			characterColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			characterAllColumns,
			characterPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("master: unable to upsert character, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`character`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `character` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(characterType, characterMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(characterType, characterMapping, ret)
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
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "master: unable to upsert for character")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(characterType, characterMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "master: unable to retrieve unique values for character")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "master: unable to populate default values for character")
	}

CacheNoHooks:
	if !cached {
		characterUpsertCacheMut.Lock()
		characterUpsertCache[key] = cache
		characterUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}