DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `username`   varchar(128) NOT NULL DEFAULT '' COMMENT 'Username',
    `password`   varchar(128) NOT NULL DEFAULT '' COMMENT 'Password',
--     `follow_count` bigint(20) NOT NULL DEFAULT '0' COMMENT 'Follow count',
--     `follower_count` bigint(20) NOT NULL DEFAULT '0' COMMENT 'Follower count',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'User account create time',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'User account update time',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'User account delete time',
    PRIMARY KEY (`id`),
    KEY          `idx_username` (`username`) COMMENT 'Username index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='User account table';

DROP TABLE IF EXISTS `follows`;
CREATE TABLE `follows` (
                           `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'PK',
                           `user_id` bigint(20) NOT NULL COMMENT 'UserId',
                           `follower_id` bigint(20) NOT NULL COMMENT 'Follower id',
                           `cancel` tinyint(4) NOT NULL DEFAULT '0' COMMENT 'Follow is 0,delete is otel-collector-config.yaml',
                           PRIMARY KEY (`id`),
                           UNIQUE KEY `userIdToFollowerIdIdx` (`user_id`,`follower_id`) USING BTREE,
                           KEY `FollowerIdIdx` (`follower_id`) USING BTREE COMMENT 'Follow id index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='Follow table';

DROP TABLE IF EXISTS `message`;
CREATE TABLE `message` (
                           `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
                           `user_id` bigint(20) NOT NULL COMMENT 'UserId',
                           `to_user_id` bigint(20) NOT NULL COMMENT 'ToUserId',
                           `content` varchar(255) NOT NULL COMMENT 'Message Content',
                           `created_time` bigint unsigned NOT NULL COMMENT 'User account create time',
--                            `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'User account update time',
--                            `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'User account delete time',
                           PRIMARY KEY (`id`),
                           KEY `userIdToUserIdIdx` (`user_id`,`to_user_id`) USING BTREE COMMENT 'userIdToUserIdIdx id index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='Message table';

DROP TABLE IF EXISTS `Video`;
CREATE TABLE `Video` (
                         `id` int NOT NULL AUTO_INCREMENT,
                         `user_id` int DEFAULT NULL,
                         `play_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
                         `cover_url` varchar(255) DEFAULT NULL,
                         `favourite_count` int(10) unsigned zerofill DEFAULT NULL,
                         `comment_count` int DEFAULT NULL,
                         `is_favourite` int(10) unsigned zerofill DEFAULT NULL,
                         `title` varchar(255) DEFAULT NULL,
                         `pub_time` int DEFAULT NULL,
                         PRIMARY KEY (`id`)
)
    ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

SET FOREIGN_KEY_CHECKS = 1;

DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment` (
                           `comment_id` int NOT NULL AUTO_INCREMENT,
                           `user_id` int DEFAULT NULL,
                           `video_id` int DEFAULT NULL,
                           `content` varchar(255) DEFAULT NULL,
                           `create_date` varchar(255) DEFAULT NULL,
                           `pub_time` int DEFAULT NULL,
                           PRIMARY KEY (`comment_id`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

SET FOREIGN_KEY_CHECKS = 1;

DROP TABLE IF EXISTS `favorite`;
CREATE TABLE `favorite` (
                            `user_id` int DEFAULT NULL,
                            `video_id` int DEFAULT NULL,
                            `favorite_time` int DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

SET FOREIGN_KEY_CHECKS = 1;