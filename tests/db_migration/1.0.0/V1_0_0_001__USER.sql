create table user
(
    id      int         not null primary key autoincrement,
    name    varchar(30) not null,
    age     int         not null,
    id_card varchar(50) null
)