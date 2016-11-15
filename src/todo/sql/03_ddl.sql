-- Run against the todo database
SET SESSION search_path = todo,public;

DROP TABLE IF EXISTS todo.task_tags;
DROP TABLE IF EXISTS todo.tasks;
DROP TABLE IF EXISTS todo.priorities;
DROP TABLE IF EXISTS todo.status;
DROP TABLE IF EXISTS todo.tags;
DROP TABLE IF EXISTS todo.sessions;
DROP TABLE IF EXISTS todo.users;

-- Table: todo.users
CREATE TABLE todo.users
(
  user_id serial NOT NULL,
  name character varying(255),
  email character varying(255) NOT NULL,
  password character varying(255) NOT NULL,
  created_at timestamp without time zone NOT NULL,
  CONSTRAINT users_pkey PRIMARY KEY (user_id),
  CONSTRAINT users_email_key UNIQUE (email),
  CONSTRAINT users_name_key UNIQUE (name)
)
WITH (
  OIDS=FALSE
);
ALTER TABLE todo.users
  OWNER TO todo;

-- Table: todo.sessions
CREATE TABLE todo.sessions
(
  session_id serial NOT NULL,
  uuid  character varying(64) not null unique,
  email character varying(255),
  user_id integer,
  created_at timestamp without time zone NOT NULL,
  CONSTRAINT sessions_pkey PRIMARY KEY (session_id),
  CONSTRAINT sessions_user_id_fkey FOREIGN KEY (user_id)
      REFERENCES todo.users (user_id) MATCH SIMPLE
      ON UPDATE NO ACTION ON DELETE NO ACTION
)
WITH (
  OIDS=FALSE
);
ALTER TABLE todo.sessions
  OWNER TO todo;

-- Table: todo.priorities
CREATE TABLE todo.priorities
(
  priority_id serial NOT NULL,
  value character varying(255),
  CONSTRAINT priorities_pkey PRIMARY KEY (priority_id),
  CONSTRAINT priorities_priority_key UNIQUE (value)
)
WITH (
  OIDS=FALSE
);
ALTER TABLE todo.priorities
  OWNER TO todo;

-- Table: todo.status
CREATE TABLE todo.status
(
  status_id serial NOT NULL,
  value character varying(255),
  CONSTRAINT status_pkey PRIMARY KEY (status_id),
  CONSTRAINT status_status_key UNIQUE (value)
)
WITH (
  OIDS=FALSE
);
ALTER TABLE todo.status
  OWNER TO todo;

-- Table: todo.tasks
CREATE TABLE todo.tasks
(
  task_id serial NOT NULL,
  value character varying(255),
  user_id integer,
  priority_id integer,
  status_id integer,
  created_at timestamp without time zone NOT NULL,
  due_at timestamp without time zone,
  complete_at timestamp without time zone,
  CONSTRAINT tasks_pkey PRIMARY KEY (task_id),
  CONSTRAINT tasks_user_id_fkey FOREIGN KEY (user_id)
      REFERENCES todo.users (user_id) MATCH SIMPLE
      ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT tasks_priority_id_fkey FOREIGN KEY (priority_id)
      REFERENCES todo.priorities (priority_id) MATCH SIMPLE
      ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT tasks_status_id_fkey FOREIGN KEY (status_id)
      REFERENCES todo.status (status_id) MATCH SIMPLE
      ON UPDATE NO ACTION ON DELETE NO ACTION
)
WITH (
  OIDS=FALSE
);
ALTER TABLE todo.tasks
  OWNER TO todo;

-- Table: todo.tags
CREATE TABLE todo.tags
(
  tag_id serial NOT NULL,
  value character varying(255),
  CONSTRAINT tags_pkey PRIMARY KEY (tag_id),
  CONSTRAINT tags_tag UNIQUE (value)
)
WITH (
  OIDS=FALSE
);
ALTER TABLE todo.tags
  OWNER TO todo;

-- Table: todo.task_tags
CREATE TABLE todo.task_tags
(
  task_tag_id serial NOT NULL,
  task_id integer,
  tag_id integer,
  CONSTRAINT task_tags_pkey PRIMARY KEY (task_tag_id),
  CONSTRAINT task_tags_tag_id_fkey FOREIGN KEY (tag_id)
      REFERENCES todo.tags (tag_id) MATCH SIMPLE
      ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT task_tags_task_id_fkey FOREIGN KEY (task_id)
      REFERENCES todo.tasks (task_id) MATCH SIMPLE
      ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT task_tags_unique UNIQUE (task_id, tag_id)
)
WITH (
  OIDS=FALSE
);
ALTER TABLE todo.task_tags
  OWNER TO todo;