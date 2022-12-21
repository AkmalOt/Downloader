select password, id from cloud_user where login = ?;
select name, password, id from cloud_user where login = ?;


select *from pg_available_extensions;
CREATE EXTENSION IF NOT EXISTS 'uuid—ossp';

SELECT uuid_generate_v4();

create table users (
                            id uuid primary key not null default uuid_generate_v4(),
                            name text not null,
                            login text not null unique,
                            password text not null,
                            active bool not null default true,
                            created_at timestamptz not null default current_timestamp,
                            updated_at timestamptz,
                            deleted_at timestamptz
);

create table tokens (
                              id uuid primary key not null
                                  default uuid_generate_v4(),
                              token text not null ,
                              user_id uuid not null references users,
                              expire timestamptz not null default current_timestamp + interval'1 hour',
                              created_at timestamptz not null default current_timestamp
);

-- select *from tokens where token = ?;

create table folders (
                               id uuid primary key not null default uuid_generate_v4(),
                               name text not null,
                               user_id uuid not null references users,
                               folder_id uuid references folders,
                               created_at timestamptz not null default current_timestamp,
                               deleted_at timestamptz
);

create table files (
                             id uuid primary key not null default uuid_generate_v4(),
                             name text not null,
                             user_id uuid not null references users,
                             folder_id uuid not null references folders,
                             active bool not null default true,
                             created_at timestamptz not null default current_timestamp,
                             deleted_at timestamptz
);

-- select id from cloud_user where login= ?;
-- insert into cloud_folders (name, user_id, folder_id)
-- VALUES ('test4', '1bde500a-6f0c-11ed-a244-7c8ae16c8c64', '1be0126e-6f0c-11ed-a245-7c8ae16c8c64');

-- select * from cloud_folders where user_id= '1bde500a-6f0c-11ed-a244-7c8ae16c8c64';

-- insert into cloud_folders(name, user_id, folder_id)
-- values (?, ?, ?);
--
-- select * from cloud_folders where user_id= '1bde500a-6f0c-11ed-a244-7c8ae16c8c64'
--                               and folder_id= '';
--
-- select id, coalesce((select coalesce(folder_id, null))::text, ' ') from cloud_folders
-- where user_id= '1bde500a-6f0c-11ed-a244-7c8ae16c8c64';
--
--
-- select id from cloud_folders where folder_id= '';
-- select name, user_id from cloud_folders where user_id= '1bde500a-6f0c-11ed-a244-7c8ae16c8c64' and folder_id= NULL;
--
-- select user_id from cloud_files where name='PList.pdf' and folder_id='f098d552-7166-11ed-9ee2-7c8ae16c8c64';
--
-- select * from cloud_files where folder_id='f098d552-7166-11ed-9ee2-7c8ae16c8c64';
--
-- select * from cloud_files where folder_id= ?
--
-- select *from cloud_files where folder_id='f098d552-7166-11ed-9ee2-7c8ae16c8c64' order by folder_id limit 3 offset 1
--
-- update cloud_files set name= ? where id = '77f3b308-723b-11ed-abcd-7c8ae16c8c64'
--
-- DELETE FROM cloud_files where id='cddc941a-7966-11ed-bbb3-7c8ae16c8c64';
--
-- select *from cloud_files where folder_id='1be0126e-6f0c-11ed-a245-7c8ae16c8c64';


CREATE TABLE access (
                        id uuid primary key not null default uuid_generate_v4(),
                        user_id uuid not null references users,
                        file_id uuid not null references files,
                        access_to uuid not null references users,
                        active bool not null default true,
                        created_at timestamptz not null default current_timestamp,
                        updated_at timestamptz,
                        expire timestamptz not null default current_timestamp
);

