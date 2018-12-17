create database if not exists highload;

grant all privileges on highload.* to python_app@'localhost' identified by 'password';

drop table if exists highload.accounts;
create table highload.accounts(
  id int not null ,
  fname varchar(50),
  sname varchar(50),
  email varchar(100) not null ,
  phone varchar(16),
  sex tynyint not null ,
  status tynyint not null ,
  birth timestamp not null ,
  joined timestamp not null ,
  country varchar(50),
  city varchar(50)
);

