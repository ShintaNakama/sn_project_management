create database if not exists sn_project_management character set utf8mb4 collate utf8mb4_bin;

use sn_project_management;

drop table if exists projects;
create table if not exists projects
(
  id          int unsigned not null primary key auto_increment,
  name        varchar(128) not null,
  description text,
  created_at  datetime not null,
  updated_at  datetime not null,
  completed       tinyint(4) not null default 0
) character set utf8mb4 collate utf8mb4_bin;

drop table if exists tasks;
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

drop table if exists completed_projects;
create table if not exists completed_projects
(
  id          int unsigned not null primary key auto_increment,
  project_id  int unsigned not null,
  created_at  datetime not null,
  CONSTRAINT fk_completed_projects_projects FOREIGN KEY (project_id) REFERENCES projects (id)
) character set utf8mb4 collate utf8mb4_bin;

drop table if exists completed_tasks;
create table if not exists completed_tasks
(
  id          int unsigned not null primary key auto_increment,
  project_id  int unsigned not null,
  task_id     int unsigned not null,
  created_at  datetime not null,
  CONSTRAINT fk_completed_tasks_projects FOREIGN KEY (project_id) REFERENCES projects (id),
  CONSTRAINT fk_completed_tasks_tasks FOREIGN KEY (task_id) REFERENCES tasks (id)
) character set utf8mb4 collate utf8mb4_bin;
