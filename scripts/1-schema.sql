CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table if not exists public.course
(
    id          uuid default uuid_generate_v4() not null,
    name        varchar                         not null,
    description text                            not null,
    constraint course_pk
    primary key (id)
    );

create unique index if not exists course_id_uindex
    on course (id);