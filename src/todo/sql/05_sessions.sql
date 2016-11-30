---------------------------------------------------------------------------------------------
-- SESSIONS
---------------------------------------------------------------------------------------------
set session search_path = todo;

-- get sessions
create or replace function sessions_get_all() 
returns table
(
    session_id      int
    , uuid          varchar(64)
    , email         varchar(255)
    , user_id       int
    , created_at    timestamptz
)
as $$
begin

    return query 
    select session_id
            , uuid
            , email
            , user_id
            , created_at
    from todo.sessions;

end;
$$ language plpgsql stable;

-- check session
create or replace function session_check(_uuid varchar(64))
returns table 
(
    session_id      int
    , uuid          varchar(64)
    , email         varchar(255)
    , user_id       int
    , created_at    timestamptz
)
as $$
begin

    return query
    select session_id
            , uuid
            , email
            , user_id
            , created_at
    from todo.sessions
    where uuid = _uuid;

end;
$$ language plpgsql stable;

-- add status
create or replace function session_add
(
    _uuid       varchar(64)
    , _email    varchar(255)
    , _user_id  int
)
returns table
(
    session_id      int
    , uuid          varchar(64)
    , email         varchar(255)
    , user_id       int
    , created_at    timestamptz
)
as $$

    insert into todo.sessions 
    (
        uuid
        , email
        , user_id
        , created_at
    ) 
    values 
    (
        _uuid
        , _email
        , _user_id
        , now()
    ) 
    returning session_id, uuid, email, user_id, created_at;

$$ language sql volatile;

-- delete session by uuid
create or replace function session_delete(_uuid varchar(64)) 
returns void
as $$

    delete from todo.sessions where uuid = _uuid;

$$ language sql volatile;

-- delete all sessions
create or replace function sessions_delete_all() 
returns void
as $$

    delete from todo.sessions;

$$ language sql volatile;