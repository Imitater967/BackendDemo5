use douyin;
DROP TABLE IF EXISTS `video_favorite`;
CREATE TABLE `video_favorite` (
    `video` INT NOT NULL,
    `user` INT NOT NULL
);