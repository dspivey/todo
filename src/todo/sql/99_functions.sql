set session search_path = todo;

\echo dropping existing functions ----------------------------------------------------
-- delete any existing functions so we have a clean slate
do $$
declare
_stmt text;
begin

loop

  select 'drop function if exists ' || ns.nspname || '.' || proname || '(' || oidvectortypes(proargtypes) || ') cascade' into _stmt
  from pg_proc
  inner join pg_namespace ns on (pg_proc.pronamespace = ns.oid)
  where ns.nspname  in ('todo')
  order by proname
  limit 1;

  exit when _stmt is null;

  raise notice '%', _stmt;
  execute _stmt;
end loop;

end;
$$;

\ir 04_priorities.sql
\ir 05_sessions.sql
\ir 06_status.sql
\ir 07_tags.sql

\echo FINISHED
\q