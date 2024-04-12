create table sys_operation_log
(
    id             bigint auto_increment comment '编号'
        primary key,
    user_name      varchar(50)   not null comment '用户名',
    operation      varchar(50)   not null comment '用户操作',
    method         varchar(200)  not null comment '请求方法',
    params         varchar(5000) null comment '请求参数',
    operation_time bigint        not null comment '执行时长(毫秒)',
    ip             varchar(64)   null comment 'IP地址',
    create_time    datetime      not null comment '创建时间'
)
    comment '操作日志';


create table sys_login_log
(
    id          bigint auto_increment comment '编号'
        primary key,
    user_name   varchar(50) not null comment '用户名',
    ip          varchar(64) null comment 'IP地址',
    create_by   varchar(50) not null comment '创建人',
    create_time datetime    null comment '创建时间'
)
    comment '登录日志';

