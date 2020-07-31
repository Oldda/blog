drop table if exists articles;
drop table if exists users;


create table users (
  id         int(10) unsigned not null auto_increment primary key,
  nickname   varchar(32) not null unique,
  realname   varchar(255) not null unique,
  avatar     varchar(255) not null DEFAULT "",
  email      varchar(255) not null unique,
  password   varchar(255) not null,
  phone      char(11) not null unique,
  created_at int(10) unsigned not null DEFAULT 0,
  updated_at int(10) unsigned not null DEFAULT 0,
  deleted_at int(10) unsigned not null DEFAULT 0
)ENGINE=innodb DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

create table articles (
  id         int(10) unsigned not null auto_increment primary key,
  title      varchar(255) not null DEFAULT "",
  description varchar(1000) not null DEFAULT "",
  body       text,
  user_id    int(10) unsigned not null DEFAULT 0,
  created_at int(10) unsigned not null DEFAULT 0,
  updated_at int(10) unsigned not null DEFAULT 0,
  deleted_at int(10) unsigned not null DEFAULT 0 
)ENGINE=innodb DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;