-- 29 使用join查询方式找出没有分类的电影id以及名称
select film.film_id, title
from film
         left join film_category
                   on film.film_id = film_category.film_id
where category_id is NULL;

-- 30
select title, description
from film
         left join film_category
                   on film.film_id = film_category.film_id
where category_id in (select category_id from category where name = 'Action');

-- 32
select concat(last_name, " ", first_name)
from employees

create table actor
(
    actor_id    smallint(5) primary key,
    first_name  varchar(45) not null,
    last_name   varchar(45) not null,
    last_update date        not null
);

insert into actor (actor_id, first_name, last_name, last_update)
values (3, "ED", "CHASE", "2006-02-15 12:34:33"),
       (2, "NICK", "WAHLBERG", "2006-02-15 12:34:33")

insert into actor (actor_id, first_name, last_name, last_update)
values
    (3, "ED", "CHASE", "2006-02-15 12:34:33")

