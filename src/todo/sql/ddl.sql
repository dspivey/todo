drop table if exists task_tags;
drop table if exists tags;
drop table if exists user_tasks;
drop table if exists tasks;
drop table if exists priorities;
drop table if exists status;
drop table if exists sessions;
drop table if exists users;

create table users (
    id           serial primary key,
    name         varchar(255),
    email        varchar(255) not null unique,
    password     varchar(255) not null,
    created_at   timestamp not null   
);

create table sessions (
    id           serial primary key,
    email        varchar(255),
    user_id      integer references users(id),
    created_at   timestamp not null   
);

create table priorities (
    id          serial primary key,
    value       varchar(255)
);

create table status (
    id          serial primary key,
    value       varchar(255)
);

create table tasks (
    id          serial primary key,
    value       varchar(255),
    priority_id integer references priorities(id),
    status_id   integer references status(id),
    created_at  timestamp not null,
    due_at      timestamp,
    complete_at timestamp
);

create table user_tasks (
    id          serial primary key,
    task_id     integer references tasks(id),
    user_id     integer references users(id)
);

create table tags (
    id          serial primary key,
    tag         varchar(255)
);

create table task_tags (
    id          serial primary key,
    task_id     integer references tasks(id),
    tag_id     integer references tags(id)
);