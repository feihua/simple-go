create table sys_login_log
(
    id         bigint auto_increment comment '编号'
        primary key,
    user_name  varchar(50)                        not null comment '用户名',
    status     varchar(50)                        not null comment '登录状态',
    ip         varchar(64)                        not null comment 'IP地址',
    login_time datetime default CURRENT_TIMESTAMP not null comment '登陆时间'
) comment '系统登录日志';
