package basemigrate

import (
	"errors"
)

const (
	sqlChangelog = `CREATE TABLE IF NOT EXISTS databasechangelog (
	id varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
	author varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
	filename varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
	dateexecuted datetime NOT NULL,
	orderexecuted int(11) NOT NULL,
	md5sum varchar(35) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
	description varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
	tag varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
	version varchar(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;`
)

const (
	appVersion       = "1.0"
	elementChangeset = "--changeset "
	elementRollback  = "--rollback "
	elementInclude   = "--include "
)

var (
	// ErrInvalidHeader is when the changeset header is invalid.
	ErrInvalidHeader = errors.New("invalid changeset header")
	// ErrInvalidFormat is when a changeset is not found.
	ErrInvalidFormat = errors.New("invalid changeset format")
)
