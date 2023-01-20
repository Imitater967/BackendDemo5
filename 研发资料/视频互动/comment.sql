use douyin;
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment` (
                           `video_id` INT NOT NULL,
                           `user_id` INT NOT NULL,
                           `id` INT NOT NULL AUTO_INCREMENT,
                           `content` VARCHAR(1023) NOT NULL,
                           `date` timestamp NOT NULL,
                           `deleted` tinyint NULL DEFAULT 0 COMMENT '1-已删除；0-未删除。默认为0',
                           PRIMARY KEY (`id`)
);