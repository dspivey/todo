set session search_path = todo;

---------------------------------------------------------------------------------------------
-- PRIORITIES
---------------------------------------------------------------------------------------------

-- get priorities
create or replace function priorities_search(_id int = null, _search text = null) 
returns table (priority_id int, value varchar(255))
as $$
begin

    return query 
    select x.priority_id
            , x.value
    from todo.priorities x
    where   (_id is null or x.priority_id = _id) 
        and (_search is null or x.value LIKE '%' || _search || '%');

end;
$$ language plpgsql stable;

-- add priority
create or replace function priority_add(_value text) 
returns table (priority_id int)
as $$

    insert into todo.priorities
    (
        value
    )
    select _value
    returning priority_id;

$$ language sql volatile;

-- edit priority
create or replace function priority_edit(_id int, _value text) 
returns table (priority_id int)
as $$

    update todo.priorities
    set value = _value
    where priority_id = _id
    returning priority_id;

$$ language sql volatile;

-- delete priority
create or replace function priority_delete(_id int) 
returns table (priority_id int)
as $$

    delete from todo.priorities 
    where priority_id = _id
    returning _id;

$$ language sql volatile;