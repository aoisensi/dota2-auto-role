// Code generated by SQLBoiler 3.7.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Guilds", testGuilds)
	t.Run("RankRoles", testRankRoles)
	t.Run("RegisterdRoles", testRegisterdRoles)
	t.Run("Users", testUsers)
}

func TestDelete(t *testing.T) {
	t.Run("Guilds", testGuildsDelete)
	t.Run("RankRoles", testRankRolesDelete)
	t.Run("RegisterdRoles", testRegisterdRolesDelete)
	t.Run("Users", testUsersDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Guilds", testGuildsQueryDeleteAll)
	t.Run("RankRoles", testRankRolesQueryDeleteAll)
	t.Run("RegisterdRoles", testRegisterdRolesQueryDeleteAll)
	t.Run("Users", testUsersQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Guilds", testGuildsSliceDeleteAll)
	t.Run("RankRoles", testRankRolesSliceDeleteAll)
	t.Run("RegisterdRoles", testRegisterdRolesSliceDeleteAll)
	t.Run("Users", testUsersSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Guilds", testGuildsExists)
	t.Run("RankRoles", testRankRolesExists)
	t.Run("RegisterdRoles", testRegisterdRolesExists)
	t.Run("Users", testUsersExists)
}

func TestFind(t *testing.T) {
	t.Run("Guilds", testGuildsFind)
	t.Run("RankRoles", testRankRolesFind)
	t.Run("RegisterdRoles", testRegisterdRolesFind)
	t.Run("Users", testUsersFind)
}

func TestBind(t *testing.T) {
	t.Run("Guilds", testGuildsBind)
	t.Run("RankRoles", testRankRolesBind)
	t.Run("RegisterdRoles", testRegisterdRolesBind)
	t.Run("Users", testUsersBind)
}

func TestOne(t *testing.T) {
	t.Run("Guilds", testGuildsOne)
	t.Run("RankRoles", testRankRolesOne)
	t.Run("RegisterdRoles", testRegisterdRolesOne)
	t.Run("Users", testUsersOne)
}

func TestAll(t *testing.T) {
	t.Run("Guilds", testGuildsAll)
	t.Run("RankRoles", testRankRolesAll)
	t.Run("RegisterdRoles", testRegisterdRolesAll)
	t.Run("Users", testUsersAll)
}

func TestCount(t *testing.T) {
	t.Run("Guilds", testGuildsCount)
	t.Run("RankRoles", testRankRolesCount)
	t.Run("RegisterdRoles", testRegisterdRolesCount)
	t.Run("Users", testUsersCount)
}

func TestHooks(t *testing.T) {
	t.Run("Guilds", testGuildsHooks)
	t.Run("RankRoles", testRankRolesHooks)
	t.Run("RegisterdRoles", testRegisterdRolesHooks)
	t.Run("Users", testUsersHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Guilds", testGuildsInsert)
	t.Run("Guilds", testGuildsInsertWhitelist)
	t.Run("RankRoles", testRankRolesInsert)
	t.Run("RankRoles", testRankRolesInsertWhitelist)
	t.Run("RegisterdRoles", testRegisterdRolesInsert)
	t.Run("RegisterdRoles", testRegisterdRolesInsertWhitelist)
	t.Run("Users", testUsersInsert)
	t.Run("Users", testUsersInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("RankRoleToGuildUsingGuild", testRankRoleToOneGuildUsingGuild)
	t.Run("RegisterdRoleToGuildUsingGuild", testRegisterdRoleToOneGuildUsingGuild)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {
	t.Run("GuildToRegisterdRoleUsingRegisterdRole", testGuildOneToOneRegisterdRoleUsingRegisterdRole)
}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("GuildToRankRoles", testGuildToManyRankRoles)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("RankRoleToGuildUsingRankRoles", testRankRoleToOneSetOpGuildUsingGuild)
	t.Run("RegisterdRoleToGuildUsingRegisterdRole", testRegisterdRoleToOneSetOpGuildUsingGuild)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {
	t.Run("GuildToRegisterdRoleUsingRegisterdRole", testGuildOneToOneSetOpRegisterdRoleUsingRegisterdRole)
}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("GuildToRankRoles", testGuildToManyAddOpRankRoles)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("Guilds", testGuildsReload)
	t.Run("RankRoles", testRankRolesReload)
	t.Run("RegisterdRoles", testRegisterdRolesReload)
	t.Run("Users", testUsersReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Guilds", testGuildsReloadAll)
	t.Run("RankRoles", testRankRolesReloadAll)
	t.Run("RegisterdRoles", testRegisterdRolesReloadAll)
	t.Run("Users", testUsersReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Guilds", testGuildsSelect)
	t.Run("RankRoles", testRankRolesSelect)
	t.Run("RegisterdRoles", testRegisterdRolesSelect)
	t.Run("Users", testUsersSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Guilds", testGuildsUpdate)
	t.Run("RankRoles", testRankRolesUpdate)
	t.Run("RegisterdRoles", testRegisterdRolesUpdate)
	t.Run("Users", testUsersUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Guilds", testGuildsSliceUpdateAll)
	t.Run("RankRoles", testRankRolesSliceUpdateAll)
	t.Run("RegisterdRoles", testRegisterdRolesSliceUpdateAll)
	t.Run("Users", testUsersSliceUpdateAll)
}
