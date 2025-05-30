use sampledb;

# todoデータ 2 つ
insert into todos (id, title, done, created_at) values
(1, 'firstTodo', 'false', now());

insert into todos (id, title, done, created_at) values
(2, 'secondTodo', 'false', now());