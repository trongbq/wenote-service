DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`
(
    `id`          INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `name`        VARCHAR(100)       NOT NULL DEFAULT '',
    `email`       VARCHAR(255)       NOT NULL,
    `password`    VARCHAR(255)       NOT NULL,
    `picture_url` VARCHAR(255)       NOT NULL DEFAULT '',
    `created_at`  TIMESTAMP          NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  TIMESTAMP          NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY `email` (`email`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

DROP TABLE IF EXISTS `oauth_token`;
CREATE TABLE `oauth_token`
(
    `id`            INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `user_id`       INT                NOT NULL,
    `access_token`  VARCHAR(255)       NOT NULL,
    `expires_at`    TIMESTAMP          NOT NULL,
    `refresh_token` VARCHAR(255)       NOT NULL,
    `created_at`    TIMESTAMP          NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`    TIMESTAMP          NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    KEY `user_id` (`user_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

DROP TABLE IF EXISTS `task`;
CREATE TABLE `task`
(
    `id`            VARBINARY(16) NOT NULL PRIMARY KEY,
    `user_id`       INT           NOT NULL,
    `task_group_id` VARBINARY(16),
    `task_goal_id`  VARBINARY(16),
    `content`       VARCHAR(100)  NOT NULL DEFAULT '',
    `note`          VARCHAR(255),
    `start`         TIMESTAMP,
    `reminder`      TIMESTAMP,
    `deadline`      TIMESTAMP,
    `order`         INT(10)       NOT NULL,
    `completed`     BOOLEAN                DEFAULT FALSE,
    `completed_at`  TIMESTAMP,
    `deleted`       BOOLEAN       NOT NULL DEFAULT FALSE,
    `deleted_at`    TIMESTAMP,
    `created_at`    TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`    TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

DROP TABLE IF EXISTS `checklist`;
CREATE TABLE `checklist`
(
    `id`           VARBINARY(16) NOT NULL PRIMARY KEY,
    `task_id`      VARBINARY(16) NOT NULL,
    `content`      VARCHAR(100)  NOT NULL DEFAULT '',
    `order`        INT(10)       NOT NULL,
    `completed`    BOOLEAN                DEFAULT FALSE,
    `completed_at` TIMESTAMP,
    `created_at`   TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`   TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

DROP TABLE IF EXISTS `tag`;
CREATE TABLE `tag`
(
    `id`         VARBINARY(16) NOT NULL PRIMARY KEY,
    `user_id`    INT           NOT NULL,
    `name`       VARCHAR(100)  NOT NULL DEFAULT '',
    `created_at` TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

DROP TABLE IF EXISTS `task_tag`;
CREATE TABLE `task_tag`
(
    `task_id`    VARBINARY(16) NOT NULL,
    `tag_id`     VARBINARY(16) NOT NULL,
    `created_at` TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`task_id`, `tag_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

DROP TABLE IF EXISTS `task_category`;
CREATE TABLE `task_category`
(
    `id`         VARBINARY(16) NOT NULL PRIMARY KEY,
    `user_id`    INT           NOT NULL,
    `name`       VARCHAR(100)  NOT NULL DEFAULT '',
    `order`      INT(10)       NOT NULL,
    `deleted`    BOOLEAN       NOT NULL DEFAULT FALSE,
    `deleted_at` TIMESTAMP,
    `created_at` TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

DROP TABLE IF EXISTS `task_goal`;
CREATE TABLE `task_goal`
(
    `id`           VARBINARY(16) NOT NULL PRIMARY KEY,
    `user_id`      INT           NOT NULL,
    `cat_id`       VARBINARY(16),
    `name`         VARCHAR(100)  NOT NULL DEFAULT '',
    `note`         VARCHAR(255),
    `start`        TIMESTAMP,
    `reminder`     TIMESTAMP,
    `deadline`     TIMESTAMP,
    `order`        INT(10)       NOT NULL,
    `completed`    BOOLEAN                DEFAULT FALSE,
    `completed_at` TIMESTAMP,
    `deleted`      BOOLEAN       NOT NULL DEFAULT FALSE,
    `deleted_at`   TIMESTAMP,
    `created_at`   TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`   TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

# Group of tasks
DROP TABLE IF EXISTS `task_group`;
CREATE TABLE `task_group`
(
    `id`           VARBINARY(16) NOT NULL PRIMARY KEY,
    `task_goal_id` VARBINARY(16),
    `user_id`      INT           NOT NULL,
    `name`         VARCHAR(100)  NOT NULL DEFAULT '',
    `order`        INT(10)       NOT NULL COMMENT 'Share order with task, treat this entity as a normal task',
    `created_at`   TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`   TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT fk_task_goal FOREIGN KEY (`task_goal_id`) REFERENCES task_group (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

ALTER TABLE `task`
    ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);
ALTER TABLE `checklist`
    ADD FOREIGN KEY (`task_id`) REFERENCES `task` (`id`);
ALTER TABLE `tag`
    ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);
ALTER TABLE `task_tag`
    ADD FOREIGN KEY (`task_id`) REFERENCES `task` (`id`);
ALTER TABLE `task_tag`
    ADD FOREIGN KEY (`tag_id`) REFERENCES `tag` (`id`);
ALTER TABLE `task_category`
    ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);
ALTER TABLE `task_goal`
    ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);
ALTER TABLE `task_group`
    ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);
ALTER TABLE `oauth_token`
    ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);
