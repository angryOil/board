SET statement_timeout = 0;

--bun:split

CREATE TABLE "public"."board"
(
    id              SERIAL PRIMARY KEY,
    cafe_id         bigint        not null,
    board_type      bigint        not null,
    writer          bigint        not null,
    title           varchar(100)  not null,
    content         varchar(2000) not null,
    created_at      timestamptz,
    last_updated_at timestamptz
);
