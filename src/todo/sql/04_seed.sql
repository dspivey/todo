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
    'Test User',
    'tuser@doozer.com',
    'a94a8fe5ccb19ba61c4c0873d391e987982fbbd3',
    now()
);

-- insert test tasks
insert into todo.tasks
(
    value,
    priority_id,
    status_id,
    created_at
)
values
(
    'test task 1',
    (select p.priority_id from todo.priorities p where p.value = 'High'),
    (select s.status_id from todo.status s where s.value = 'Incomplete'),
    now()
),
(
    'test task 2',
    (select p.priority_id from todo.priorities p where p.value = 'High'),
    (select s.status_id from todo.status s where s.value = 'Incomplete'),
    now()
),
(
    'test task 3',
    (select p.priority_id from todo.priorities p where p.value = 'High'),
    (select s.status_id from todo.status s where s.value = 'Incomplete'),
    now()
);

-- link tasks to user
insert into todo.user_tasks
(
    task_id,
    user_id
)
select t.task_id, u.user_id
from todo.tasks t
inner join todo.users u on u.Email = 'tuser@doozer.com';