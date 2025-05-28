use sampledb;

# todoデータ 2 つ
insert into todos (title, done, created_at) values
('firstTodo', 'false', now());

insert into todos (title, done, created_at) values
('secondTodo', 'false', now());