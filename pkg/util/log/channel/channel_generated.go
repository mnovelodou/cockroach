// Code generated by gen/main.go. DO NOT EDIT.

package channel

import "github.com/cockroachdb/cockroach/pkg/util/log/logpb"

// DEV is the channel used during development, to collect log
// details useful for troubleshooting when it is unclear which other
// channel to use. It is also the default logging channel in
// CockroachDB, when the caller does not indicate a channel.
//
// This channel is special in that there are no constraints as to
// what may or may not be logged on it. Conversely, users in
// production deployments are invited to not collect DEV logs in
// centralized logging facilities, because they likely contain
// sensitive operational data.
const DEV = logpb.Channel_DEV

// OPS is the channel used to report "point" operational events,
// initiated by user operators or automation:
//
// - operator or system actions on server processes: process starts,
//   stops, shutdowns, crashes (if they can be logged),
//   including each time: command-line parameters, current version being run.
// - actions that impact the topology of a cluster: node additions,
//   removals, decommissions, etc.
// - job-related initiation or termination.
// - cluster setting changes.
// - zone configuration changes.
const OPS = logpb.Channel_OPS

// HEALTH is the channel used to report "background" operational
// events, initiated by CockroachDB or reporting on automatic processes:
//
// - current resource usage, including critical resource usage.
// - node-node connection events, including connection errors and
//   gossip details.
// - range and table leasing events.
// - up-, down-replication; range unavailability.
const HEALTH = logpb.Channel_HEALTH

// STORAGE is the channel used to report low-level storage
// layer events (RocksDB/Pebble).
const STORAGE = logpb.Channel_STORAGE

// SESSIONS is the channel used to report client network activity:
//
// - connections opened/closed.
// - authentication events: logins, failed attempts.
// - session and query cancellation.
//
// This is typically configured in "audit" mode, with event
// numbering and synchronous writes.
const SESSIONS = logpb.Channel_SESSIONS

// SQL_SCHEMA is the channel used to report changes to the
// SQL logical schema, excluding privilege and ownership changes
// (which are reported on the separate channel PRIVILEGES) and
// zone config changes (which go to OPS).
//
// This includes:
//
// - database/schema/table/sequence/view/type creation
// - adding/removing/changing table columns
// - changing sequence parameters
//
// etc., more generally changes to the schema that affect the
// functional behavior of client apps using stored objects.
const SQL_SCHEMA = logpb.Channel_SQL_SCHEMA

// USER_ADMIN is the channel used to report changes
// in users and roles, including:
//
// - users added/dropped.
// - changes to authentication credentials, incl passwords, validity etc.
// - role grants/revocations.
// - role option grants/revocations.
//
// This is typically configured in "audit" mode, with event
// numbering and synchronous writes.
const USER_ADMIN = logpb.Channel_USER_ADMIN

// PRIVILEGES is the channel used to report data
// authorization changes, including:
//
// - privilege grants/revocations on database, objects etc.
// - object ownership changes.
//
// This is typically configured in "audit" mode, with event
// numbering and synchronous writes.
const PRIVILEGES = logpb.Channel_PRIVILEGES

// SENSITIVE_ACCESS is the channel used to report SQL
// data access to sensitive data (when enabled):
//
// - data access audit events (when table audit is enabled).
// - SQL statements executed by users with the ADMIN bit.
// - operations that write to `system` tables.
//
// This is typically configured in "audit" mode, with event
// numbering and synchronous writes.
const SENSITIVE_ACCESS = logpb.Channel_SENSITIVE_ACCESS

// SQL_EXEC is the channel used to report SQL execution on
// behalf of client connections:
//
// - logical SQL statement executions (if enabled)
// - pgwire events (if enabled)
const SQL_EXEC = logpb.Channel_SQL_EXEC

// SQL_PERF is the channel used to report SQL executions
// that are marked to be highlighted as "out of the ordinary"
// to facilitate performance investigations.
// This includes the "SQL slow query log".
//
// Arguably, this channel overlaps with SQL_EXEC defined above.
// However, we keep them separate for backward-compatibility
// with previous versions, where the corresponding events
// were redirected to separate files.
const SQL_PERF = logpb.Channel_SQL_PERF

// SQL_INTERNAL_PERF is like the SQL perf channel above but aimed at
// helping developers of CockroachDB itself. It exists as a separate
// channel so as to not pollute the SQL perf logging output with
// internal troubleshooting details.
const SQL_INTERNAL_PERF = logpb.Channel_SQL_INTERNAL_PERF
