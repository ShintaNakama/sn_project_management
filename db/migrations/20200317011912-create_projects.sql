
-- +migrate Up
create table if not exists projects
(
  id          int unsigned not null primary key auto_increment,
  name        varchar(128) not null,
  description text,
  created_at  datetime not null,
  updated_at  datetime not null,
  completed       tinyint(4) not null default 0
) character set utf8mb4 collate utf8mb4_bin;

-- +migrate Down
drop table if exists projects;
