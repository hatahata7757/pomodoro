---- drop ----
DROP TABLE IF EXISTS `users`;
DROP TABLE IF EXISTS `tasks`:

---- create ----
create table IF NOT EXISTS `users`
(
  `id`               INT(20) AUTO_INCREMENT,
  `name`             VARCHAR(20) NOT NULL,
  `created_at`       Datetime DEFAULT NULL,
  `updated_at`       Datetime DEFAULT NULL,
    PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

create table IF NOT EXISTS `tasks`
(
  `id`               INT(20) AUTO_INCREMENT,
  `name`             VARCHAR(20) NOT NULL,
  `done`             BOOLEAN DEFAULT false,
  `user_id`          INT(20),
  `created_at`       Datetime DEFAULT NULL,
  `updated_at`       Datetime DEFAULT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY(`user_id`),
    REFERENCES `users`(`users_id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
