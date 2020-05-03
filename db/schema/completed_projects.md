# completed_projects

## Description

<details>
<summary><strong>Table Definition</strong></summary>

```sql
CREATE TABLE `completed_projects` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `project_id` int(10) unsigned NOT NULL,
  `created_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_completed_projects_projects` (`project_id`),
  CONSTRAINT `fk_completed_projects_projects` FOREIGN KEY (`project_id`) REFERENCES `projects` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin
```

</details>

## Columns

| Name | Type | Default | Nullable | Children | Parents | Comment |
| ---- | ---- | ------- | -------- | -------- | ------- | ------- |
| id | int(10) unsigned |  | false |  |  |  |
| project_id | int(10) unsigned |  | false |  | [projects](projects.md) |  |
| created_at | datetime |  | false |  |  |  |

## Constraints

| Name | Type | Definition |
| ---- | ---- | ---------- |
| fk_completed_projects_projects | FOREIGN KEY | FOREIGN KEY (project_id) REFERENCES projects (id) |
| PRIMARY | PRIMARY KEY | PRIMARY KEY (id) |

## Indexes

| Name | Definition |
| ---- | ---------- |
| fk_completed_projects_projects | KEY fk_completed_projects_projects (project_id) USING BTREE |
| PRIMARY | PRIMARY KEY (id) USING BTREE |

## Relations

![er](completed_projects.png)

---

> Generated by [tbls](https://github.com/k1LoW/tbls)