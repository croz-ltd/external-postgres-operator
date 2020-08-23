/*
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PostgresDatabaseSpec struct {
	Name  string `json:"name"`
	Owner string `json:"owner"`
	// +optional
	Comment string `json:"comment,omitempty"`
	// +optional
	Definition PostgresDatabaseSpecDefinition `json:"definition,omitempty"`
	// +optional
	Security PostgresDatabaseSpecSecurity `json:"security,omitempty"`
	// +optional
	Parameters []PostgresDatabaseSpecParameter `json:"parameters,omitempty"`
	// +optional
	Schemas []PostgresDatabaseSpecSchema `json:"schemas,omitempty"`
}

type PostgresDatabaseSpecDefinition struct {
	// +optional
	Encoding Encoding `json:"encoding,omitempty"`
	// +optional
	Template string `json:"template,omitempty"`
	// +optional
	Tablespace string `json:"tablespace,omitempty"`
	// +optional
	Collation Collation `json:"collation,omitempty"`
	// +optional
	CharacterType CharacterType `json:"characterType,omitempty"`
	// +optional
	ConnectionLimit int `json:"connectionLimit,omitempty"`
}

type Encoding string

const (
	BIG5         = "BIG5"
	EUCCN        = "EUCCN"
	EUCJIS2004   = "EUCJIS2004"
	EUCJP        = "EUCJP"
	EUCKR        = "EUCKR"
	EUCTW        = "EUCTW"
	GB18030      = "GB18030"
	GBK          = "GBK"
	ISO88595     = "ISO88595"
	ISO88596     = "ISO88596"
	ISO88597     = "ISO88597"
	ISO88598     = "ISO88598"
	JOHAB        = "JOHAB"
	KOI8R        = "KOI8R"
	KOI8U        = "KOI8U"
	LATIN1       = "LATIN1"
	LATIN10      = "LATIN10"
	LATIN2       = "LATIN2"
	LATIN3       = "LATIN3"
	LATIN4       = "LATIN4"
	LATIN5       = "LATIN5"
	LATIN6       = "LATIN6"
	LATIN7       = "LATIN7"
	LATIN8       = "LATIN8"
	LATIN9       = "LATIN9"
	MULEINTERNAL = "MULEINTERNAL"
	SHIFTJIS2004 = "SHIFTJIS2004"
	SJIS         = "SJIS"
	SQLASCII     = "SQLASCII"
	UHC          = "UHC"
	UTF8         = "UTF8"
	WIN1250      = "WIN1250"
	WIN1251      = "WIN1251"
	WIN1252      = "WIN1252"
	WIN1253      = "WIN1253"
	WIN1254      = "WIN1254"
	WIN1255      = "WIN1255"
	WIN1256      = "WIN1256"
	WIN1257      = "WIN1257"
	WIN1258      = "WIN1258"
	WIN866       = "WIN866"
	WIN874       = "WIN874"
)

type Collation string

const (
	Collation_C         Collation = "C"
	Collation_POSIX               = "POSIX"
	Collation_EnUS_UTF8           = "en_US.utf8"
)

type CharacterType string

const (
	CharacterType_C         CharacterType = "C"
	CharacterType_POSIX                   = "POSIX"
	CharacterType_EnUS_UTF8               = "en_US.utf8"
)

type PostgresDatabaseSpecSecurity struct {
	// +optional
	Privileges []PostgresDatabaseSpecSecurityPrivileges `json:"privileges,omitempty"`
	// +optional
	SecurityLabels []PostgresDatabaseSpecSecurityLabels `json:"securityLabels,omitempty"`
}

type PostgresDatabaseSpecSecurityPrivileges struct {
	Type            DatabaseGrantType `json:"type"`
	WithGrantOption bool              `json:"withGrantOption"`
	// +listType=set
	Grantees []string `json:"grantees"`
}

type DatabaseGrantType string

const (
	DatabaseGrantType_All       DatabaseGrantType = "all"
	DatabaseGrantType_Create                      = "create"
	DatabaseGrantType_Temporary                   = "temporary"
	DatabaseGrantType_Connect                     = "connect"
)

type SchemaGrantType string

const (
	SchemaGrantType_All    SchemaGrantType = "all"
	SchemaGrantType_Create                 = "create"
	SchemaGrantType_Usage                  = "usage"
)

type TablesGrantType string

const (
	TableGrantType_All        TablesGrantType = "all"
	TableGrantType_Select                     = "insert"
	TableGrantType_Insert                     = "select"
	TableGrantType_Update                     = "update"
	TableGrantType_Delete                     = "delete"
	TableGrantType_Truncate                   = "truncate"
	TableGrantType_References                 = "references"
	TableGrantType_Trigger                    = "trigger"
)

type SequenceGrantType string

const (
	SequenceGrantType_All    SequenceGrantType = "all"
	SequenceGrantType_Select                   = "select"
	SequenceGrantType_Update                   = "update"
	SequenceGrantType_Usage                    = "usage"
)

type FunctionsGrantType string

const (
	FunctionsGrantType_Execute FunctionsGrantType = "execute"
)

type TypesGrantType string

const (
	TypesGrantType_Usage TypesGrantType = "usage"
)

type PostgresDatabaseSpecSecurityLabels struct {
	Provider string `json:"provider"`
	Label    string `json:"label"`
}

type PostgresDatabaseSpecParameter struct {
	Name  ParameterName `json:"name"`
	Value string        `json:"value"`
	Role  string        `json:"role"`
}

type ParameterName string

const (
	Role                            ParameterName = "role"
	AllowSystemTableMods                          = "allow_system_table_mods"
	ApplicationName                               = "application_name"
	ArrayNulls                                    = "array_nulls"
	BackendFlushAfter                             = "backend_flush_after"
	BackslashQuote                                = "backslash_quote"
	BacktraceFunctions                            = "backtrace_functions"
	ByteaOutput                                   = "bytea_output"
	CheckFunctionBodies                           = "check_function_bodies"
	ClientEncoding                                = "client_encoding"
	ClientMinMessages                             = "client_min_messages"
	CommitDelay                                   = "commit_delay"
	CommitSiblings                                = "commit_siblings"
	ConstraintExclusion                           = "constraint_exclusion"
	CpuIndexTupleCost                             = "cpu_index_tuple_cost"
	CpuOperatorCost                               = "cpu_operator_cost"
	CpuTupleCost                                  = "cpu_tuple_cost"
	CursorTupleFraction                           = "cursor_tuple_fraction"
	DateStyle                                     = "DateStyle"
	DeadlockTimeout                               = "deadlock_timeout"
	DebugPrettyPrint                              = "debug_pretty_print"
	DebugPrintParse                               = "debug_print_parse"
	DebugPrintPlan                                = "debug_print_plan"
	DebugPrintRewritten                           = "debug_print_rewritten"
	DefaultStatisticsTarget                       = "default_statistics_target"
	DefaultTableAccessMethod                      = "default_table_access_method"
	DefaultTablespace                             = "default_tablespace"
	DefaultTextSearchConfig                       = "default_text_search_config"
	DefaultTransactionDeferrable                  = "default_transaction_deferrable"
	DefaultTransactionIsolation                   = "default_transaction_isolation"
	DefaultTransactionReadOnly                    = "default_transaction_read_only"
	DynamicLibraryPath                            = "dynamic_library_path"
	EffectiveCacheSize                            = "effective_cache_size"
	EffectiveIoConcurrency                        = "effective_io_concurrency"
	EnableBitmapscan                              = "enable_bitmapscan"
	EnableGathermerge                             = "enable_gathermerge"
	EnableHashagg                                 = "enable_hashagg"
	EnableHashjoin                                = "enable_hashjoin"
	EnableIncrementalSort                         = "enable_incremental_sort"
	EnableIndexonlyscan                           = "enable_indexonlyscan"
	EnableIndexscan                               = "enable_indexscan"
	EnableMaterial                                = "enable_material"
	EnableMergejoin                               = "enable_mergejoin"
	EnableNestloop                                = "enable_nestloop"
	EnableParallelAppend                          = "enable_parallel_append"
	EnableParallelHash                            = "enable_parallel_hash"
	EnablePartitionPruning                        = "enable_partition_pruning"
	EnablePartitionwiseAggregate                  = "enable_partitionwise_aggregate"
	EnablePartitionwiseJoin                       = "enable_partitionwise_join"
	EnableSeqscan                                 = "enable_seqscan"
	EnableSort                                    = "enable_sort"
	EnableTidscan                                 = "enable_tidscan"
	EscapeStringWarning                           = "escape_string_warning"
	ExitOnError                                   = "exit_on_error"
	ExtraFloatDigits                              = "extra_float_digits"
	ForceParallelMode                             = "force_parallel_mode"
	FromCollapseLimit                             = "from_collapse_limit"
	Geqo                                          = "geqo"
	GeqoEffort                                    = "geqo_effort"
	GeqoGenerations                               = "geqo_generations"
	GeqoPoolSize                                  = "geqo_pool_size"
	GeqoSeed                                      = "geqo_seed"
	GeqoSelectionBias                             = "geqo_selection_bias"
	GeqoThreshold                                 = "geqo_threshold"
	GinFuzzySearchLimit                           = "gin_fuzzy_search_limit"
	GinPendingListLimit                           = "gin_pending_list_limit"
	HashMemMultiplier                             = "hash_mem_multiplier"
	IdleInTransactionSessionTimeout               = "idle_in_transaction_session_timeout"
	IgnoreChecksumFailure                         = "ignore_checksum_failure"
	IntervalStyle                                 = "IntervalStyle"
	Jit                                           = "jit"
	JitAboveCost                                  = "jit_above_cost"
	JitDumpBitcode                                = "jit_dump_bitcode"
	JitExpressions                                = "jit_expressions"
	JitInlineAboveCost                            = "jit_inline_above_cost"
	JitOptimizeAboveCost                          = "jit_optimize_above_cost"
	JitTupleDeforming                             = "jit_tuple_deforming"
	JoinCollapseLimit                             = "join_collapse_limit"
	LcMessages                                    = "lc_messages"
	LcMonetary                                    = "lc_monetary"
	LcNumeric                                     = "lc_numeric"
	LcTime                                        = "lc_time"
	LoCompatPrivileges                            = "lo_compat_privileges"
	LocalPreloadLibraries                         = "local_preload_libraries"
	LockTimeout                                   = "lock_timeout"
	LogDuration                                   = "log_duration"
	LogErrorVerbosity                             = "log_error_verbosity"
	LogExecutorStats                              = "log_executor_stats"
	LogLockWaits                                  = "log_lock_waits"
	LogMinDurationSample                          = "log_min_duration_sample"
	LogMinDurationStatement                       = "log_min_duration_statement"
	LogMinErrorStatement                          = "log_min_error_statement"
	LogMinMessages                                = "log_min_messages"
	LogParameterMaxLength                         = "log_parameter_max_length"
	LogParameterMaxLengthOnError                  = "log_parameter_max_length_on_error"
	LogParserStats                                = "log_parser_stats"
	LogPlannerStats                               = "log_planner_stats"
	LogReplicationCommands                        = "log_replication_commands"
	LogStatement                                  = "log_statement"
	LogStatementSampleRate                        = "log_statement_sample_rate"
	LogStatementStats                             = "log_statement_stats"
	LogTempFiles                                  = "log_temp_files"
	LogTransactionSampleRate                      = "log_transaction_sample_rate"
	LogicalDecodingWorkMem                        = "logical_decoding_work_mem"
	MaintenanceIoConcurrency                      = "maintenance_io_concurrency"
	MaintenanceWorkMem                            = "maintenance_work_mem"
	MaxParallelMaintenanceWorkers                 = "max_parallel_maintenance_workers"
	MaxParallelWorkers                            = "max_parallel_workers"
	MaxParallelWorkersPerGather                   = "max_parallel_workers_per_gather"
	MaxStackDepth                                 = "max_stack_depth"
	MinParallelIndexScanSize                      = "min_parallel_index_scan_size"
	MinParallelTableScanSize                      = "min_parallel_table_scan_size"
	OperatorPrecedenceWarning                     = "operator_precedence_warning"
	ParallelLeaderParticipation                   = "parallel_leader_participation"
	ParallelSetupCost                             = "parallel_setup_cost"
	ParallelTupleCost                             = "parallel_tuple_cost"
	PasswordEncryption                            = "password_encryption"
	PlanCacheMode                                 = "plan_cache_mode"
	QuoteAllIdentifiers                           = "quote_all_identifiers"
	RandomPageCost                                = "random_page_cost"
	RowSecurity                                   = "row_security"
	SearchPath                                    = "search_path"
	SeqPageCost                                   = "seq_page_cost"
	SessionPreloadLibraries                       = "session_preload_libraries"
	SessionReplicationRole                        = "session_replication_role"
	StandardConformingStrings                     = "standard_conforming_strings"
	StatementTimeout                              = "statement_timeout"
	SynchronizeSeqscans                           = "synchronize_seqscans"
	SynchronousCommit                             = "synchronous_commit"
	TcpKeepalivesCount                            = "tcp_keepalives_count"
	TcpKeepalivesIdle                             = "tcp_keepalives_idle"
	TcpKeepalivesInterval                         = "tcp_keepalives_interval"
	TcpUserTimeout                                = "tcp_user_timeout"
	TempBuffers                                   = "temp_buffers"
	TempFileLimit                                 = "temp_file_limit"
	TempTablespaces                               = "temp_tablespaces"
	TimeZone                                      = "TimeZone"
	TimezoneAbbreviations                         = "timezone_abbreviations"
	TraceNotify                                   = "trace_notify"
	TraceSort                                     = "trace_sort"
	TrackActivities                               = "track_activities"
	TrackCounts                                   = "track_counts"
	TrackFunctions                                = "track_functions"
	TrackIoTiming                                 = "track_io_timing"
	TransactionDeferrable                         = "transaction_deferrable"
	TransactionIsolation                          = "transaction_isolation"
	TransactionReadOnly                           = "transaction_read_only"
	TransformNullEquals                           = "transform_null_equals"
	UpdateProcessTitle                            = "update_process_title"
	VacuumCleanupIndexScaleFactor                 = "vacuum_cleanup_index_scale_factor"
	VacuumCostDelay                               = "vacuum_cost_delay"
	VacuumCostLimit                               = "vacuum_cost_limit"
	VacuumCostPageDirty                           = "vacuum_cost_page_dirty"
	VacuumCostPageHit                             = "vacuum_cost_page_hit"
	VacuumCostPageMiss                            = "vacuum_cost_page_miss"
	VacuumFreezeMinAge                            = "vacuum_freeze_min_age"
	VacuumFreezeTableAge                          = "vacuum_freeze_table_age"
	VacuumMultixactFreezeMinAge                   = "vacuum_multixact_freeze_min_age"
	VacuumMultixactFreezeTableAge                 = "vacuum_multixact_freeze_table_age"
	WalCompression                                = "wal_compression"
	WalConsistencyChecking                        = "wal_consistency_checking"
	WalInitZero                                   = "wal_init_zero"
	WalRecycle                                    = "wal_recycle"
	WalSenderTimeout                              = "wal_sender_timeout"
	WalSkipThreshold                              = "wal_skip_threshold"
	WorkMem                                       = "work_mem"
	Xmlbinary                                     = "xmlbinary"
	Xmloption                                     = "xmloption"
	ZeroDamagedPages                              = "zero_damaged_pages"
)

type PostgresDatabaseSpecSchema struct {
	Name  string `json:"name"`
	Owner string `json:"owner,omitempty"`
	// +optional
	Comment string `json:"comment,omitempty"`
	// +optional
	Security PostgresDatabaseSpecSchemaSecurity `json:"security,omitempty"`
	// +optional
	DefaultPrivileges PostgresDatabaseSpecSchemaDefaultPrivileges `json:"defaultPrivileges,omitempty"`
}

type PostgresDatabaseSpecSchemaSecurity struct {
	// +optional
	Privileges []PostgresDatabaseSpecSchemaSecurityPrivileges `json:"privileges,omitempty"`
	// +optional
	SecurityLabels []PostgresDatabaseSpecSecurityLabels `json:"securityLabels,omitempty"`
}

type PostgresDatabaseSpecSchemaSecurityPrivileges struct {
	Type            SchemaGrantType `json:"type"`
	WithGrantOption bool            `json:"withGrantOption"`
	// +listType=set
	Grantees []string `json:"grantees"`
}

type PostgresDatabaseSpecSchemaDefaultPrivileges struct {
	// +optional
	Tables []PostgresDatabaseSpecTablesSecurityPrivileges `json:"tables,omitempty"`
	// +optional
	Sequences []PostgresDatabaseSpecSequenceSecurityPrivileges `json:"sequences,omitempty"`
	// +optional
	Functions []PostgresDatabaseSpecFunctionsSecurityPrivileges `json:"functions,omitempty"`
	// +optional
	Types []PostgresDatabaseSpecTypesSecurityPrivileges `json:"types,omitempty"`
}

type PostgresDatabaseSpecTablesSecurityPrivileges struct {
	Type            TablesGrantType `json:"type"`
	WithGrantOption bool            `json:"withGrantOption"`
	// +listType=set
	Grantees []string `json:"grantees"`
}

type PostgresDatabaseSpecSequenceSecurityPrivileges struct {
	Type            SequenceGrantType `json:"type"`
	WithGrantOption bool              `json:"withGrantOption"`
	// +listType=set
	Grantees []string `json:"grantees"`
}

type PostgresDatabaseSpecFunctionsSecurityPrivileges struct {
	Type            FunctionsGrantType `json:"type"`
	WithGrantOption bool               `json:"withGrantOption"`
	// +listType=set
	Grantees []string `json:"grantees"`
}

type PostgresDatabaseSpecTypesSecurityPrivileges struct {
	Type            TypesGrantType `json:"type"`
	WithGrantOption bool           `json:"withGrantOption"`
	// +listType=set
	Grantees []string `json:"grantees"`
}

// PostgresDatabaseStatus defines the observed state of PostgresDatabase
type PostgresDatabaseStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// PostgresDatabase is the Schema for the postgresdatabases API
type PostgresDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PostgresDatabaseSpec   `json:"spec,omitempty"`
	Status PostgresDatabaseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PostgresDatabaseList contains a list of PostgresDatabase
type PostgresDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PostgresDatabase `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PostgresDatabase{}, &PostgresDatabaseList{})
}
