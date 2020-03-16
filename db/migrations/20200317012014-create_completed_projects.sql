
-- +migrate Up
create table if not exists completed_projects
(
  id          int unsigned not null primary key auto_increment,
  project_id  int unsigned not null,
  created_at  datetime not null,
  CONSTRAINT fk_completed_projects_projects FOREIGN KEY (project_id) REFERENCES projects (id)
) character set utf8mb4 collate utf8mb4_bin;

-- +migrate Down
drop table if exists completed_projects;
