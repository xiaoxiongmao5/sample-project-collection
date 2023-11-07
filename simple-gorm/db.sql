-- Active: 1691069230063@@127.0.0.1@3306@test_gorm
Create Table `user` (
	id int(11) NOT NULL AUTO_INCREMENT COMMENT '主键自增id',
    uid INT(11) NOT NULL COMMENT '用户id',
    keywords TEXT NOT NULL COMMENT '索引词',
    degree CHAR(2) NOT NULL DEFAULT "" COMMENT '学历',
    gender CHAR(1) NOT NULL DEFAULT "女" COMMENT '性别',
    city CHAR(2) NOT NULL COMMENT '城市',
    PRIMARY KEY(id),
    UNIQUE KEY `idx_uid`(uid)
) Engine=InnoDB AUTO_INCREMENT=1 DEFAULT charset=utf8 COMMENT='用户信息表'

INSERT user VALUES(NULL, 1, "社保|15|已交满", "", "男", "北京");
INSERT user VALUES(NULL, 2, "社保|15|已交满", "", "女", "北京");

ALTER TABLE `user` MODIFY degree CHAR(2) NOT NULL DEFAULT "" COMMENT '学历';
ALTER TABLE `user` MODIFY gender CHAR(1) NOT NULL DEFAULT "女" COMMENT '性别';

SELECT * from `user`;