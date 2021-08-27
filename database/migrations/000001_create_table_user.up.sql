CREATE TABLE IF NOT EXISTS users (
    id serial primary key,
    username text not null,
    password text not null,
    email text not null,    
    activated bool not null default false, 
    created_at timestamp,
    updated_at timestamp
);


