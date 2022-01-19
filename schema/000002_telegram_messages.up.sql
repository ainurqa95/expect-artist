CREATE TABLE telegram_message_types
(
    id serial not null unique,
    name        varchar(255) not null,
    code        varchar(255) not null unique 
);

CREATE TABLE telegram_messages
(
    id serial not null unique,
    user_id int references users (id) on delete cascade not null,
    chat_id int not null,
    message text,
    created_at   timestamp with time zone default CURRENT_TIMESTAMP not null,
    message_type_id int references telegram_message_types (id) on delete cascade not null
);