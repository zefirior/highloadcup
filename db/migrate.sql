create database if not exists highload;

grant all privileges on highload.* to python_app@'localhost' identified by 'password';

drop table if exists highload.account;
create table highload.account(
  id int not null ,
  fname varchar(50),
  sname varchar(50),
  email varchar(100) not null ,
  phone varchar(16),
  sex tinyint not null ,
  status tinyint not null ,
  birth int not null ,
  joined int not null ,
  country varchar(50),
  city varchar(50)
) ENGINE=MyISAM;

drop table if exists highload.interest;
create table highload.interest (
  interest varchar(100),
  account_id int not null
) ENGINE=MyISAM;

drop table if exists highload.premium;
create table highload.premium (
  account_id int not null,
  start int not null,
  finish int not null
) ENGINE=MyISAM;

drop table if exists highload.likes;
create table highload.likes (
  account_id int not null,
  ts int not null,
  like_account_id int not null
) ENGINE=MyISAM;

truncate highload.account;
truncate highload.interest;
truncate highload.premium;
truncate highload.likes;
