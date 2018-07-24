--changeset username:id
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';
CREATE TABLE `tablename` (
    `id` VARCHAR(36) NOT NULL,

    `name` VARCHAR(100) NOT NULL,

    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP DEFAULT 0,

    UNIQUE KEY (`name`),
    PRIMARY KEY (`id`)
);
--rollback DROP TABLE `tablename`;