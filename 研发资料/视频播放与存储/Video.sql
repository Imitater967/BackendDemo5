use douyin;
DROP TABLE IF EXISTS `video`;
CREATE TABLE `video` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `uploader` INT NOT NULL,
    `title` VARCHAR(16) NOT NULL,
    `date` timestamp NOT NULL,
    `deleted` tinyint NULL DEFAULT 0 COMMENT '1-已删除；0-未删除。默认为0',
    PRIMARY KEY (`id`)
);