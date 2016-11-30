---------------------------------------------------------------------------------------------
-- TAGS
---------------------------------------------------------------------------------------------
set session search_path = todo;

-- get tags
create or replace function tags_search(_id int, _search text) 
returns table (tag_id int, value varchar(255))
as $$
begin

    return query 
    select x.tag_id
            , x.value
    from todo.tags x
    where   (_id is null or x.tag_id = _id)
        and (_search is null or x.value ILIKE '%' || _search || '%');

end;
$$ language plpgsql stable;

-- add tag
create or replace function tag_add(_value text) 
returns table (tag_id int)
as $$

    insert into todo.tags
    (
        value
    )
    select _value
    returning tag_id;

$$ language sql volatile;

-- edit tag
create or replace function tag_edit(_id int, _value text) 
returns table (tag_id int)
as $$

    update todo.tags
    set value = _value
    where tag_id = _id
    returning tag_id;

$$ language sql volatile;

-- delete tag
create or replace function tag_delete(_id int) 
returns table (tag_id int)
as $$

    delete from todo.tags
    where tag_id = _id
    returning _id;

$$ language sql volatile;