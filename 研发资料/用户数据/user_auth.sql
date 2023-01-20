use gostudy;
DROP TABLE IF EXISTS `user_auth`;
CREATE TABLE `user_auth` (
                             `username` VARCHAR(50) NOT NULL,
                             `password` VARCHAR(50) NOT NULL,
                             `UserId` INT NOT NULL AUTO_INCREMENT,
                             `Token` VARCHAR(50) DEFAULT NULL,
                             `Expire` INT DEFAULT NULL,
                             PRIMARY KEY (`UserId`),
                             UNIQUE KEY `username` (`username`)
);