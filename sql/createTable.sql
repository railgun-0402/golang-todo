use sampledb;

# Todoデータを格納するためのテーブル
create table if not exists todos (
    id integer unsigned auto_increment primary key,
    title varchar(100) not null,
    done boolean,
    created_at datetime
);