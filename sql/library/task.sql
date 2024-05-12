-- Библиотека:

-- Хранить историю движения книг. Иметь возможность по айди книги и дате понять в библиотеке она или в аренде.
-- Нельзя взять одну книгу одновременно.


-- Написать миграцию на создание необходимых таблиц.

create table books (
    id uuid not null default gen_random_uuid(),
    title text not null,
    created_at timestamp not null default now()
);


create table readers (
    id uuid not null default gen_random_uuid(),
    created_at timestamp not null default now()
);

create table reader_books (
    id uuid not null default gen_random_uuid(),
    reader_id uuid not null references readers(id),
    book_id uuid not null references books(id),
    created_at timestamp not null default now(),
    deleted_at timestamp
);


-- Написать запрос на получение просроченных книг и данных клиентов (на руках у клиента больше 1 месяца).

select b.*, r.* from books b
    join reader_books rb on b.id = rb.book_id
    join readers r on rb.reader_id = r.id
where
    rb.created_at < now() - '1 month' and
    rb.deleted_at is null;

select b.* from books b where id in (
    select r.id
        from reader_books rb join readers r on
            rb.reader_id = r.id
        where rb.created_at < now() - '1 month' and rb.deleted_at is null
);


-- Написать запрос на получение всех клиентов которые должны больше 3 книг, их данные и список этих книг.

-- Данные читателей и книг
select
    r.*, b.*
from readers r
    join reader_books rb on r.id = rb.reader_id
    join books b on rb.book_id = b.id

where r.id in (
    select r.id from readers r
        join reader_books rb on r.id = rb.reader_id
    where rb.deleted_at is null
        group by r.id
    having count(1) > 3
);

-- Только данные читателей
select
    r.*
from readers r where id in (
    select r.id from readers r join reader_books rb
        on r.id = rb.reader_id
    where rb.deleted_at is null
        group by r.id
    having count(1) > 3
);
