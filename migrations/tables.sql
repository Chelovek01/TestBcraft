drop table if exists  recipe;
create table recipe (
    id integer,
    name text,
    description text,
    ingredients text,
    prepare_steps text

);

create table users (
    id int primary key,
    username text,
    password text

)