CREATE TABLE `t_project` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `project_name` varchar(64)  NOT NULL COMMENT '项目名称',
  `project_desc` varchar(128)  DEFAULT NULL COMMENT '项目描述',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `modified_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='项目表';


CREATE TABLE `t_project_module` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `module_name` varchar(64)  NOT NULL COMMENT '模块名称',
  `module_desc` varchar(128)  DEFAULT NULL COMMENT '模块描述',
  `project_id` bigint(20)  NOT NULL COMMENT '项目id',
  `project_name` varchar(64)  DEFAULT NULL COMMENT '项目名称',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `modified_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='项目模块表';