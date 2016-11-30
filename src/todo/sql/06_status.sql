---------------------------------------------------------------------------------------------
-- STATUS
---------------------------------------------------------------------------------------------
set session search_path = todo;

-- get status
create or replace function status_search(_id int = null, _search text = null) 
returns table (status_id int, value varchar(255))
as $$
begin

    return query 
    select x.status_id
            , x.value
    from todo.status x
    where  (_id is null or x.status_id = _id)
        and (_search is null or x.value ILIKE '%' || _search || '%');

end;
$$ language plpgsql stable;

-- add status
create or replace function status_add(_value text) 
returns table (status_id int)
as $$

    insert into todo.status
    (
        value
    )
    select _value
    returning status_id;

$$ language sql volatile;

-- edit status
create or replace function status_edit(_id int, _value text) 
returns table (status_id int)
as $$

    update todo.status
    set value = _value
    where status_id = _id
    returning status_id;

$$ language sql volatile;

-- delete status
create or replace function status_delete(_id int) 
returns table (status_id int)
as $$

    delete from todo.status
    where status_id = _id
    returning _id;

$$ language sql volatile;