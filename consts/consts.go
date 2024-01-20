package consts

const (
	CMD_NAME_MIGRATE  = "migrate"
	CMD_NAME_VALIDATE = "validate"

	DEFAULT_HISTORY_TABLE = "flyway_schema_history"

	LOCATION_PREFIX_SEQ = "::"
	LOCATION_PREFIX_OS  = "system" + LOCATION_PREFIX_SEQ
	LOCATION_DEFAULT    = "db_migration"

	BASE_LINE_DESC = "<< Flyway Baseline >>"
	BASE_LINE_TYPE = "BASELINE"
	SQL_TYPE       = "SQL"
)
