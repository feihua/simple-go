create table sys_dict
(
    id          bigint auto_increment comment '编号'
        primary key,
    value       varchar(100)                       not null comment '数据值',
    label       varchar(100)                       not null comment '标签名',
    type        varchar(100)                       not null comment '类型',
    description varchar(100)                       not null comment '描述',
    sort        decimal                            not null comment '排序（升序）',
    remarks     varchar(255)                       null comment '备注',
    create_time datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    update_time datetime default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP comment '修改时间'
)
    comment '字典表';

INSERT INTO sys_dict (id, value, label, type, description, sort, remarks, create_time, update_time) VALUES (1, 'male', '男', 'sex', '性别', 0, 'test', current_timestamp, null);
INSERT INTO sys_dict (id, value, label, type, description, sort, remarks, create_time, update_time) VALUES (2, 'female', '女', 'sex', '性别', 1, 'test', current_timestamp, null);
