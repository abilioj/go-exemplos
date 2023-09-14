create database api_todo;

create user user_todo;
ALTER user user_todo with encrypted password '123123';
grant all privileges on database api_todo to user_todo;
use api_todo;
grant all privileges on all tables in schema public to user_todo;

GRANT ALL PRIVILEGES ON api_todo.* TO 'user_todo'@'localhost';

CREATE TABLE IF NOT EXISTS todos (
  id serial PRIMARY KEY,
  title VARCHAR NOT NULL,
  description VARCHAR NOT NULL ,
  done bool NOT NULL
) ;