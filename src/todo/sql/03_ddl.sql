-- Run against the todo database

SET SESSION search_path = todo,public;

-- Table: todo.users
DROP TABLE IF EXISTS todo.users;
CREATE TABLE todo.users
(
  id serial NOT NULL,
  name character varying(255),
  email character varying(255) NOT NULL,
  password character varying(255) NOT NULL,
  created_at timestamp without time zone NOT NULL,
  CONSTRAINT users_pkey PRIMARY KEY (id),
  CONSTRAINT users_email_key UNIQUE (email)
)
WITH (
  OIDS=FALSE
);
ALTER TABLE todo.users
  OWNER TO todo;

-- Table: todo.sessions
DROP TABLE IF EXISTS todo.sessions;
CREATE TABLE todo.sessions
(
  id serial NOT NULL,
  email character varying(255),
  user_id integer,
  created_at timestamp without time zone NOT NULL,
  CONSTRAINT sessions_pkey PRIMARY KEY (id),
  CONSTRAINT sessions_user_id_fkey FOREIGN KEY (user_id)
      REFERENCES todo.users (id) MATCH SIMPLE
      ON UPDATE NO ACTION ON DELETE NO ACTION
)
WITH (
  OIDS=FALSE
);
ALTER TABLE todo.sessions
  OWNER TO todo;

-- Table: todo.priorities
DROP TABLE IF EXISTS todo.priorities;
CREATE TABLE todo.priorities
(
  id serial NOT NULL,
  value character varying(255),
  CONSTRAINT priorities_pkey PRIMARY KEY (id)
)
WITH (
  OIDS=FALSE
);
ALTER TABLE todo.priorities
  OWNER TO todo;

-- Table: todo.status
DROP TABLE IF EXISTS todo.status;
CREATE TABLE todo.status
(
  id serial NOT NULL,
  value character varying(255),
  CONSTRAINT status_pkey PRIMARY KEY (id)
)
WITH (
  OIDS=FALSE
);
ALTER TABLE todo.status
  OWNER TO todo;

-- Table: todo.tasks
DROP TABLE IF EXISTS todo.tasks;
CREATE TABLE todo.tasks
(
  id serial NOT NULL,
  value character varying(255),
  priority_id integer,
  status_id integer,
  created_at timestamp without time zone NOT NULL,
  due_at timestamp without time zone,
  complete_at timestamp without time zone,
  CONSTRAINT tasks_pkey PRIMARY KEY (id),
  CONSTRAINT tasks_priority_id_fkey FOREIGN KEY (priority_id)
      REFERENCES todo.priorities (id) MATCH SIMPLE
      ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT tasks_status_id_fkey FOREIGN KEY (status_id)
      REFERENCES todo.status (id) MATCH SIMPLE
      ON UPDATE NO ACTION ON DELETE NO ACTION
)
WITH (
  OIDS=FALSE
);
ALTER TABLE todo.tasks
  OWNER TO todo;

-- Table: todo.user_tasks
DROP TABLE IF EXISTS todo.user_tasks;
CREATE TABLE todo.user_tasks
(
  id serial NOT NULL,
  task_id integer,
  user_id integer,
  CONSTRAINT user_tasks_pkey PRIMARY KEY (id),
  CONSTRAINT user_tasks_task_id_fkey FOREIGN KEY (task_id)
      REFERENCES todo.tasks (id) MATCH SIMPLE
      ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT user_tasks_user_id_fkey FOREIGN KEY (user_id)
      REFERENCES todo.users (id) MATCH SIMPLE
      ON UPDATE NO ACTION ON DELETE NO ACTION
)
WITH (
  OIDS=FALSE
);
ALTER TABLE todo.user_tasks
  OWNER TO todo;

-- Table: todo.tags
DROP TABLE IF EXISTS todo.tags;
CREATE TABLE todo.tags
(
  id serial NOT NULL,
  tag character varying(255),
  CONSTRAINT tags_pkey PRIMARY KEY (id)
)
WITH (
  OIDS=FALSE
);
ALTER TABLE todo.tags
  OWNER TO todo;

-- Table: todo.task_tags
DROP TABLE IF EXISTS todo.task_tags;
CREATE TABLE todo.task_tags
(
  id serial NOT NULL,
  task_id integer,
  tag_id integer,
  CONSTRAINT task_tags_pkey PRIMARY KEY (id),
  CONSTRAINT task_tags_tag_id_fkey FOREIGN KEY (tag_id)
      REFERENCES todo.tags (id) MATCH SIMPLE
      ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT task_tags_task_id_fkey FOREIGN KEY (task_id)
      REFERENCES todo.tasks (id) MATCH SIMPLE
      ON UPDATE NO ACTION ON DELETE NO ACTION
)
WITH (
  OIDS=FALSE
);
ALTER TABLE todo.task_tags
  OWNER TO todo;