DROP TABLE IF EXISTS segments;
CREATE TABLE segments
(
    id serial not null unique,
    title varchar(255) not null  unique
);

DROP TABLE IF EXISTS users;
CREATE TABLE users
(
    id serial not null unique,
    name varchar(255) not null
);

DROP TABLE IF EXISTS user_segment;
CREATE TABLE user_segment
(
    id serial not null unique,
    user_id int not null,
    segment_id int references segments (id) on delete cascade not null,
    added_at timestamptz not null  DEFAULT  CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS user_history;
CREATE TABLE user_history
(
    id serial not null unique,
    user_id int not null,
    segment_id int references segments (id) on delete cascade not null,
    operation_type varchar(255) not null,
    execution_time timestamptz not null DEFAULT  CURRENT_TIMESTAMP
);



DROP EXTENSION IF EXISTS pg_cron;
CREATE EXTENSION pg_cron;

