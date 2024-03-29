CREATE TABLE "{{.schema}}"."{{.table}}"
(
    "installed_rank" INT           NOT NULL PRIMARY KEY,
    "version"        VARCHAR(50),
    "description"    VARCHAR(200)  NOT NULL,
    "type"           VARCHAR(20)   NOT NULL,
    "script"         VARCHAR(1000) NOT NULL,
    "checksum"       INT,
    "installed_by"   VARCHAR(100)  NOT NULL,
    "installed_on"   TEXT          NOT NULL DEFAULT (strftime('%Y-%m-%d %H:%M:%f', 'now')),
    "execution_time" INT           NOT NULL,
    "success"        BOOLEAN       NOT NULL
);