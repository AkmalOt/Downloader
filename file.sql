select *from pg_available_extensions;
CREATE EXTENSION IF NOT EXISTS “uuid—ossp”;

SELECT uuid_generate_v1();

create table cloud_user (
                            id uuid primary key not null default uuid_generate_v1(),
                            name text not null,
                            login text not null unique,
                            password text not null,
                            active bool not null default true,
                            created_at timestamptz not null default current_timestamp,
                            updated_at timestamptz,
                            deleted_at timestamptz
);

create table cloud_tokens (
                              id uuid primary key not null default uuid_generate_v1(),
                              token text not null ,
                              user_id uuid not null references cloud_user,
                              expire timestamptz not null default current_timestamp + interval'1 hour',
                              created_at timestamptz not null default current_timestamp
);

-- select *from tokens where token = ?;

create table cloud_folders (
                               id uuid primary key not null default uuid_generate_v1(),
                               name text not null,
                               user_id uuid not null references cloud_user,
                               folder_id uuid references cloud_folders,
                               created_at timestamptz not null default current_timestamp,
                               deleted_at timestamptz
);

create table cloud_files (
                             id uuid primary key not null default uuid_generate_v1(),
                             name text not null,
                             url text not null,
                             user_id uuid not null references cloud_user,
                             folder_id uuid not null references cloud_folders,
                             created_at timestamptz not null default current_timestamp,
                             deleted_at timestamptz
);

select id from cloud_user where login= ?;
insert into cloud_folders (name, user_id, folder_id)
VALUES ('test4', '1bde500a-6f0c-11ed-a244-7c8ae16c8c64', '1be0126e-6f0c-11ed-a245-7c8ae16c8c64');

select * from cloud_folders where user_id= '1bde500a-6f0c-11ed-a244-7c8ae16c8c64';

insert into cloud_folders(name, user_id, folder_id)
values (?, ?, ?);

select * from cloud_folders where user_id= '1bde500a-6f0c-11ed-a244-7c8ae16c8c64'
                              and folder_id= '';

select coalesce(folder_id, '1be0126e-6f0c-11ed-a245-7c8ae16c8c64'), user_id from cloud_folders where folder_id= '';