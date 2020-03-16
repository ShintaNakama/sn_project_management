
-- +migrate Up
create table if not exists tasks
(
  id          int unsigned not null primary key auto_increment,
  project_id  int unsigned not null,
  description text,
  created_at  datetime not null,
  updated_at  datetime not null,
  completed       tinyint(4) not null default 0,
  CONSTRAINT fk_tasks_projects FOREIGN KEY (project_id) REFERENCES projects (id)
) character set utf8mb4 collate utf8mb4_bin;

-- +migrate Down
drop table if exists tasks;
