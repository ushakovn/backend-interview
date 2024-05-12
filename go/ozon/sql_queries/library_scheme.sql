-- Библиотека
-- Сущности: автор, книга, читатель

create table authors (
    id serial primary key,
    full_name varchar(128)
);

create table authors_ratings(
    id serial primary key,
    rating int,
    author_id int references authors
);

create table books (
    id serial primary key,
    full_name varchar(256),
    authors_set_id int references authors_set,
    unique(full_name, authors_set_id)
);

create table booked_books(
    id serial primary key,
    book_id int references books
);

create table authors_set (
    id serial primary key,
    author_id int references authors,
    book_id int references books
);

create table readers (
    id serial primary key,
    book_id int references books
);

-- Выбрать названия всех книг которые на руках

select
    full_name
from books as b
    inner join booked_books as bb on b.id = bb.book_id;

-- Выбрать названия всех книг в библиотеке у которых больше трех авторов

select
    b.full_name
from books as b
    inner join authors_set as s on b.authors_set_id = s.id
    inner join authors as a on s.author_id = a.id
group by b.full_name
    having(count(author_id) > 3);

-- Выбрать имена топ 3 читаемых авторов на данных момент

select
    a.full_name
from authors as a
    inner join authors_ratings as ar on a.id = ar.author_id
order by ar.rating desc
    limit 3;


-- Создание таблицы работников

create table employee (
    id serial primary key,
    sex int check (sex in ('m', 'f')),
    salary int not null,
    age int not null check (age between 18 and 80),
    created_at timestamp
);

-- Построить оптимальный индекс для запроса

-- Полный составной индекс по трем столбцам

create index salary_male_employee_index on employee(sex, salary, age);

-- Частичный индекс для запроса

create index salary_male_employee_part_index on employee(created_at)
    where sex = 'm' and salary > 300000 and age = 20;





