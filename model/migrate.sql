drop table if exists articles;
drop table if exists users;
drop table if exists category;
drop table if exists tag;
drop table if exists tag_and_article;


create table users (
  id         int(10) unsigned not null auto_increment primary key,
  nickname   varchar(32) not null DEFAULT "" comment "昵称",
  realname   varchar(255) not null DEFAULT "" comment "真实姓名",
  avatar     varchar(255) not null DEFAULT "" comment "头像",
  email      varchar(255) not null unique comment "邮箱-用于登录",
  password   varchar(255) not null comment "密码",
  phone      char(11) not null unique comment "手机号",
  created_at int(10) unsigned not null DEFAULT 0,
  updated_at int(10) unsigned not null DEFAULT 0,
  deleted_at int(10) unsigned not null DEFAULT 0
)ENGINE=innodb DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT="管理员表";

create table articles (
  id         int(10) unsigned not null auto_increment primary key,
  user_id    int(10) unsigned not null DEFAULT 0 comment "用户id",
  cid        int(10) unsigned not null DEFAULT 0 comment "分类ID",
  `order`    int(10) unsigned not null DEFAULT 0 comment "排序",
  title      varchar(255) not null DEFAULT "" comment "标题",
  description varchar(1000) not null DEFAULT "" comment "描述",
  body       text not null comment "内容",
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

insert into users (phone,email,password)values("17710818223","17710818223@163.com","e10adc3949ba59abbe56e057f20f883e");