create keyspace if not exists attendance_system with replication = {'class': 'NetworkTopologyStrategy', 'replication_factor': 1};

create type if not exists Point(lat double, long double);

create table if not exists attendance_system.device
(
    id       uuid,
    name     text,
    location Point,
    rules    set<text>,
    primary key (id, location),
);

create table if not exists attendance_system.device_startup
(
    device_id            text,
    server_timestamp     timestamp,
    time_after_start_sec int,
    primary key ( device_id, server_timestamp )
);

create table if not exists attendance_system.admin
(
    username        uuid,
    password_hash   text,
    first_name      text,
    last_name       text,
    profile_picture blob,
    rules           set<text>,
    primary key ( username, password_hash)
);
create index if not exists on attendance_system.admin (username) with options = {'unique': 'true'};

create table if not exists attendance_system.employee
(
    cardId       uuid,
    first_name   text,
    last_name    text,
    email        text,
    phone_number text,
    rules        set<text>,
    primary key ( cardId )
);
create index if not exists on attendance_system.employee (cardId) with options = {'unique': 'true'};

create table if not exists attendance_system.attendance_log
(
    timestamp     timestamp,
    cardId        uuid,
    action        text,
    device_id     text,
    seen_by_admin boolean,
    primary key ( timestamp, action )
);

create table attendance_system.admin_commands_log
(
    timestamp      timestamp,
    device_id      text,
    command        text,
    processed      boolean,
    admin_username text,
    primary key (timestamp, command, processed, admin_username)
);

