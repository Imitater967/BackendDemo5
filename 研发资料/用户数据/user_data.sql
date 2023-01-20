DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment` (
                           `VideoId` INT NOT NULL,
                           `CommentId` INT NOT NULL AUTO_INCREMENT,
                           `Content` VARCHAR(1023) NOT NULL,
                           `CreateDate` DATE NOT NULL,
                           PRIMARY KEY (`CommentId`),
                           UNIQUE KEY `VideoId` (`VideoId`)
);