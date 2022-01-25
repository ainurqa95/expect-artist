CREATE table cities
(
    id serial not null unique,
    name varchar(255) unique
);

CREATE TABLE users
(
    id            serial       not null unique,
    email         varchar(255) unique,
    password_hash varchar(255),
    telegram_id   int unique,
    created_at    timestamp with time zone default CURRENT_TIMESTAMP not null,
    city_id       int references cities (id) on delete cascade
);


CREATE TABLE artists
(
    id          serial       not null unique,
    name        varchar(255) not null,
    description text,
    sorting int default 0
);

CREATE TABLE genres
(
    id          serial       not null unique,
    name        varchar(255) not null,
    code        varchar(255) not null unique
);

CREATE TABLE artist_genre
(
    artist_id int references artists (id) on delete cascade not null,
    genre_id  int references genres (id) on delete cascade not null
);
create unique index artist_genre_idx on artist_genre using btree (artist_id, genre_id);    

CREATE TABLE events
(
    id          serial       not null unique,
    title       varchar(255) not null,
    description varchar(255) not null,
    happen_at   timestamp with time zone default CURRENT_TIMESTAMP not null,
    buy_link    varchar(500) default null,
    artist_id   int references artists (id) on delete cascade      not null
);

CREATE TABLE subscriptions
(
    artist_id int references artists (id) on delete cascade not null,
    user_id  int references users (id) on delete cascade not null
);
create unique index subscriptions_idx on subscriptions using btree (artist_id, user_id);    

CREATE TABLE notification_types
(
    id          serial       not null unique,
    name       varchar(255) not null
);

CREATE TABLE notifications
(
    id      serial                                           not null unique,
    type_id int references notification_types (id) on delete cascade not null,
    user_id int references users (id) on delete cascade not null,
    message text
);