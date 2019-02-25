DROP table if exists profiles;
DROP table if exists users;

CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY ,
  username VARCHAR(32) NOT NULL UNIQUE ,
  password VARCHAR(255) NOT NULL,
  email VARCHAR(255)  NOT NULL UNIQUE ,
  created_at TIMESTAMP NOT NULL ,
  updated_at TIMESTAMP NOT NULL,
  balance DECIMAL(10,2) NOT NULL
);

CREATE TABLE IF NOT EXISTS profiles (
  id SERIAL PRIMARY KEY,
  user_id INTEGER NOT NULL,
  name VARCHAR(255),
  last_name VARCHAR(255),
  photo_path VARCHAR(255),
  about TEXT,
  FOREIGN KEY (user_id) references users(id)
);



insert into users(username, password, email, created_at, updated_at, balance)
values('test1', 'password', 'emai1l@gmail.com', NOW(), NOW(), 470);

insert into profiles(user_id, name, last_name, photo_path, about)
values(1, 'alex', 'test', 'photo', 'bio');


insert into users(username, password, email, created_at, updated_at, balance)
values('test2', 'password', 'emai2l@gmail.com', NOW(), NOW(), 125);

insert into profiles(user_id, name, last_name, photo_path, about)
values(1, 'peter', 'test', 'photo', 'bio');


insert into users(username, password, email, created_at, updated_at, balance)
values('test3', 'password', 'emai3l@gmail.com', NOW(), NOW(), 852);

insert into profiles(user_id, name, last_name, photo_path, about)
values(1, 'pavel', 'test', 'photo', 'bio');

begin;
commit;

begin;
rollback;


begin;
savepoint my_point;
rollback to savepoint my_point;
commit;


select * from users;
select * from profiles;

