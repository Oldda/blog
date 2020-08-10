drop table if exists articles;
drop table if exists users;
drop table if exists category;
drop table if exists tag;
drop table if exists tag_and_article;


create table users (
  id         int(10) unsigned not null auto_increment primary key,
  nickname   varchar(32) not null,
  realname   varchar(255) not null,
  avatar     varchar(255) not null DEFAULT "",
  email      varchar(255) not null unique,
  password   varchar(255) not null,
  phone      char(11) not null unique,
  created_at int(10) unsigned not null DEFAULT 0,
  updated_at int(10) unsigned not null DEFAULT 0,
  deleted_at int(10) unsigned not null DEFAULT 0
)ENGINE=innodb DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT="管理员表";

create table articles (
  id         int(10) unsigned not null auto_increment primary key,
  cid        int(10) unsigned not null DEFAULT 0 comment "分类ID",
  title      varchar(255) not null DEFAULT "",
  description varchar(1000) not null DEFAULT "",
  body       text,
  user_id    int(10) unsigned not null DEFAULT 0,
  created_at int(10) unsigned not null DEFAULT 0,
  updated_at int(10) unsigned not null DEFAULT 0,
  deleted_at int(10) unsigned not null DEFAULT 0 
)ENGINE=innodb DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT="文章表";

create table category (
  id         int(10) unsigned not null auto_increment primary key,
  pid        int(10) unsigned not null DEFAULT 0,
  `order`      int(10) unsigned not null DEFAULT 0,
  title      varchar(255) not null DEFAULT "",
  article_count int(10) unsigned not null DEFAULT 0,
  created_at int(10) unsigned not null DEFAULT 0,
  updated_at int(10) unsigned not null DEFAULT 0,
  deleted_at int(10) unsigned not null DEFAULT 0 
)ENGINE=innodb DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT="分类表";

create table tag (
  id         int(10) unsigned not null auto_increment primary key,
  `order`      int(10) unsigned not null DEFAULT 0,
  title      varchar(255) not null DEFAULT "",
  article_count int(10) unsigned not null DEFAULT 0,
  created_at int(10) unsigned not null DEFAULT 0,
  updated_at int(10) unsigned not null DEFAULT 0,
  deleted_at int(10) unsigned not null DEFAULT 0 
)ENGINE=innodb DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT="标签表";

create table tag_and_article (
  id         int(10) unsigned not null auto_increment primary key,
  article_id int(10) unsigned not null DEFAULT 0,
  tag_id     int(10) unsigned not null DEFAULT 0,
  created_at int(10) unsigned not null DEFAULT 0,
  updated_at int(10) unsigned not null DEFAULT 0,
  deleted_at int(10) unsigned not null DEFAULT 0 
)ENGINE=innodb DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT="标签文章关系表";