-- Run against the todo database

-- Role: todo
DROP ROLE IF EXISTS todo;
CREATE ROLE todo LOGIN
  ENCRYPTED PASSWORD 'md50169c4dbcc9bb28f4e4e9616a6051ca4'
  SUPERUSER INHERIT CREATEDB CREATEROLE NOREPLICATION;

-- Change todo database owner to todo
ALTER DATABASE todo OWNER TO todo;

-- Schema: todo
DROP SCHEMA IF EXISTS todo;
CREATE SCHEMA todo AUTHORIZATION todo;