create table sys_dept
(
    id          bigint auto_increment comment '编号'
        primary key,
    dept_name   varchar(50)                        not null comment '部门名称',
    parent_id   bigint                             not null comment '上级部门ID，一级部门为0',
    sort        int                                not null comment '排序',
    remark     varchar(255)                       null comment '备注',
    create_time datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    update_time datetime                           null on update CURRENT_TIMESTAMP comment '修改时间'
) comment '部门管理';

