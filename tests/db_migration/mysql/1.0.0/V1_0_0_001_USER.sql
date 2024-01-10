create table catalog.user
(
    id     int NOT NULL AUTO_INCREMENT ,
    name    varchar(30) not null,
    age     int         not null,
    id_card varchar(50) null,
    PRIMARY KEY (id)
)