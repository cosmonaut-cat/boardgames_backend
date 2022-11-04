create database if not exists event_store;

use event_store;

create table events(
    event_id varchar(64) not null,
    event_type varchar(64) not null,
    event_entity varchar(64) not null,
    event_version int not null,
    event_timestamp varchar(64) not null ,
    event_payload_type varchar(64) not null,
    event_payload blob not null,
    primary key (event_id, event_version, event_entity, event_type),
    index (event_type, event_entity, event_version, event_timestamp)
);