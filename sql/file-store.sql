SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for file
-- ----------------------------
DROP TABLE IF EXISTS `file`;
CREATE TABLE `file`  (
    `id` varchar(255) NOT NULL COMMENT 'ID',
    `user_id` varchar(255) NOT NULL COMMENT '用户名ID',
    `folder_id` varchar(255) NULL DEFAULT NULL COMMENT '文件夹ID',
    `file_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '文件名',
    `file_hash` varchar(255) NULL DEFAULT NULL COMMENT '文件哈希值',
    `file_path` varchar(255) NOT NULL DEFAULT '/' COMMENT '文件存储路径',
    `postfix` varchar(255) NULL DEFAULT NULL COMMENT '文件后缀',
    `size` int(11) NULL DEFAULT NULL COMMENT '文件大小',
    `type` int(11) NULL DEFAULT NULL COMMENT '文件类型',
    `download_num` int(11) NULL DEFAULT 0 COMMENT '下载次数',
    `created_at` datetime DEFAULT NULL COMMENT '更新时间',
    `updated_at` datetime DEFAULT NULL COMMENT '上传时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for folder
-- ----------------------------
DROP TABLE IF EXISTS `folder`;
CREATE TABLE `folder`  (
    `id` varchar(255) NOT NULL COMMENT 'ID',
    `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '文件夹名称',
    `user_id` varchar(255) NOT NULL COMMENT '用户id',
    `parent_id` varchar(255) NOT NULL COMMENT '父文件夹id',
    `created_at` datetime DEFAULT NULL,
    `updated_at` datetime DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for file_store
-- ----------------------------
DROP TABLE IF EXISTS `file_store`;
CREATE TABLE `file_store`  (
    `id` varchar(255) NOT NULL COMMENT 'ID',
    `user_id` varchar(255) NULL DEFAULT NULL COMMENT '用户id',
    `current_size` int(11) NULL DEFAULT 0 COMMENT '当前容量（单位KB）',
    `max_size` int(11) NULL DEFAULT 1048576 COMMENT '最大容量（单位KB）',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;



-- ----------------------------
-- Table structure for share
-- ----------------------------
DROP TABLE IF EXISTS `share`;
CREATE TABLE `share`  (
    `id` varchar(255) NOT NULL COMMENT 'ID',
    `code` varchar(10) NULL DEFAULT NULL COMMENT 'Code',
    `user_id` varchar(255) NULL DEFAULT NULL COMMENT '用户id',
    `file_id` varchar(255) NOT NULL COMMENT '文件id',
    `created_at` datetime DEFAULT NULL COMMENT '更新时间',
    `updated_at` datetime DEFAULT NULL COMMENT '上传时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
    `id` varchar(255) NOT NULL COMMENT 'ID',
    `user_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '用户名',
    `pass_word` varchar(50) NOT NULL COMMENT '密码',
    `image_path` varchar(255) NULL DEFAULT '' COMMENT '头像地址',
    `register_time` datetime(0) NULL DEFAULT NULL COMMENT '注册时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
