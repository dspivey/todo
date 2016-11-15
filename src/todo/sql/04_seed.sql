-- insert initial priority values
insert into todo.priorities
(
    value
) 
values
(
    'High'
)
,
(
    'Medium'
)
,
(
    'Low'
);

-- insert initial status values
insert into todo.status
(
    value
) 
values
(
    'Incomplete'
)
,
(
    'Complete'
)
,
(
    'Archived'
);

-- insert initial tag values
insert into todo.tags
(
    value
) 
values
(
    'soco'
)
,
(
    'doozer'
)
,
(
    'honda'
);

-- insert test user
insert into todo.users
(
    name,
    email,
    password,
    created_at
)
values
(
    'Test User'
    ,'tuser@doozer.com'
    -- password is 'test'
    , 'a94a8fe5ccb19ba61c4c0873d391e987982fbbd3' 
    , now()
);

-- insert test tasks
insert into todo.tasks
(
    value,
    user_id,
    priority_id,
    status_id,
    created_at
)
values
(
    'test task 1'
    ,(select u.user_id from todo.users u where u.email = 'tuser@doozer.com')
    ,(select p.priority_id from todo.priorities p where p.value = 'High')
    ,(select s.status_id from todo.status s where s.value = 'Incomplete')
    ,now()
),
(
    'test task 2'
    ,(select u.user_id from todo.users u where u.email = 'tuser@doozer.com')
    ,(select p.priority_id from todo.priorities p where p.value = 'High')
    ,(select s.status_id from todo.status s where s.value = 'Incomplete')
    ,now()
),
(
    'test task 3'
    ,(select u.user_id from todo.users u where u.email = 'tuser@doozer.com')
    ,(select p.priority_id from todo.priorities p where p.value = 'High')
    ,(select s.status_id from todo.status s where s.value = 'Incomplete')
    ,now()
);