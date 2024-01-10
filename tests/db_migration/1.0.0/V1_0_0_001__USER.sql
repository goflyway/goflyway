create table user
(
    id      INTEGER PRIMARY KEY autoincrement,
    name    varchar(30) not null,
    age     int         not null,
    id_card varchar(50) null
)