create table sys_menu
(
    id          bigint auto_increment comment '主键'
        primary key,
    menu_name   varchar(50)                            not null comment '菜单名称',
    menu_type   tinyint      default 1                 not null comment '菜单类型(1：目录   2：菜单   3：按钮)',
    status_id   tinyint      default 1                 not null comment '状态(1:正常，0:禁用)',
    sort        int          default 1                 not null comment '排序',
    parent_id   bigint                                 not null comment '父ID',
    menu_url    varchar(255) default ''                null comment '路由路径',
    api_url     varchar(255) default ''                null comment '接口URL',
    menu_icon   varchar(255)                           null comment '菜单图标',
    remark      varchar(255)                           null comment '备注',
    create_time datetime     default CURRENT_TIMESTAMP not null comment '创建时间',
    update_time datetime     default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '修改时间',
    constraint menu_name
        unique (menu_name)
)
    comment '菜单信息';

INSERT INTO sys_menu (id, menu_name, menu_type, status_id, sort, parent_id, menu_url, api_url, menu_icon, remark, create_time, update_time) VALUES (1, '首页', 1, 1, 0, 0, '/home', '', 'SmileOutlined', '首页', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
INSERT INTO sys_menu (id, menu_name, menu_type, status_id, sort, parent_id, menu_url, api_url, menu_icon, remark, create_time, update_time) VALUES (2, '权限管理', 1, 1, 1, 0, '', '', 'SettingOutlined', '权限管理', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
INSERT INTO sys_menu (id, menu_name, menu_type, status_id, sort, parent_id, menu_url, api_url, menu_icon, remark, create_time, update_time) VALUES (3, '用户管理', 2, 1, 1, 2, '/user', '/api/user/queryUserList', 'setting', '用户管理', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
INSERT INTO sys_menu (id, menu_name, menu_type, status_id, sort, parent_id, menu_url, api_url, menu_icon, remark, create_time, update_time) VALUES (4, '角色管理', 2, 1, 2, 2, '/role', '/api/queryRoleList', '', '角色管理', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
INSERT INTO sys_menu (id, menu_name, menu_type, status_id, sort, parent_id, menu_url, api_url, menu_icon, remark, create_time, update_time) VALUES (5, '菜单管理', 2, 1, 3, 2, '/menu', '/api/queryMenuList', 'setting', '菜单管理', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
INSERT INTO sys_menu (id, menu_name, menu_type, status_id, sort, parent_id, menu_url, api_url, menu_icon, remark, create_time, update_time) VALUES (6, '保存用户接口', 3, 1, 1, 3, '', '/api/user/addUser', '', '保存用户接口', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
INSERT INTO sys_menu (id, menu_name, menu_type, status_id, sort, parent_id, menu_url, api_url, menu_icon, remark, create_time, update_time) VALUES (7, '删除用户接口', 3, 1, 1, 3, '', '/api/user/deleteUserByIds', '', '删除用户接口', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
INSERT INTO sys_menu (id, menu_name, menu_type, status_id, sort, parent_id, menu_url, api_url, menu_icon, remark, create_time, update_time) VALUES (8, '更新用户接口', 3, 1, 1, 3, '', '/api/user/updateUser', '', '更新用户接口', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
INSERT INTO sys_menu (id, menu_name, menu_type, status_id, sort, parent_id, menu_url, api_url, menu_icon, remark, create_time, update_time) VALUES (9, '保存用户角色', 3, 1, 1, 3, '', '/api/user/updateUserRoleList', '', '保存用户角色接口', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
INSERT INTO sys_menu (id, menu_name, menu_type, status_id, sort, parent_id, menu_url, api_url, menu_icon, remark, create_time, update_time) VALUES (10, '用户关联的角色', 3, 1, 1, 3, '', '/api/user/queryUserRoleList', '', '获取用户关联的角色', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
INSERT INTO sys_menu (id, menu_name, menu_type, status_id, sort, parent_id, menu_url, api_url, menu_icon, remark, create_time, update_time) VALUES (11, '查询用户菜单接口', 3, 1, 1, 3, '', '/api/user/queryUserMenu', '', '查询用户菜单接口', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
INSERT INTO sys_menu (id, menu_name, menu_type, status_id, sort, parent_id, menu_url, api_url, menu_icon, remark, create_time, update_time) VALUES (12, '保存角色接口', 3, 1, 1, 4, '', '/api/role/addRole', '', '保存角色接口', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
INSERT INTO sys_menu (id, menu_name, menu_type, status_id, sort, parent_id, menu_url, api_url, menu_icon, remark, create_time, update_time) VALUES (13, '删除角色接口', 3, 1, 1, 4, '', '/api/role/deleteRoleByIds', '', '删除角色接口', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
INSERT INTO sys_menu (id, menu_name, menu_type, status_id, sort, parent_id, menu_url, api_url, menu_icon, remark, create_time, update_time) VALUES (14, '更新角色接口', 3, 1, 1, 4, '', '/api/role/updateRole', '', '更新角色接口', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
INSERT INTO sys_menu (id, menu_name, menu_type, status_id, sort, parent_id, menu_url, api_url, menu_icon, remark, create_time, update_time) VALUES (15, '设置权限弹窗', 3, 0, 1, 4, '', '/api/role/query_role_menu_view', '', '设置权限弹窗', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
INSERT INTO sys_menu (id, menu_name, menu_type, status_id, sort, parent_id, menu_url, api_url, menu_icon, remark, create_time, update_time) VALUES (16, '菜单角色关联', 3, 1, 1, 4, '', '/api/role/queryRoleMenuList', '', '菜单角色关联', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
INSERT INTO sys_menu (id, menu_name, menu_type, status_id, sort, parent_id, menu_url, api_url, menu_icon, remark, create_time, update_time) VALUES (17, '保存角色菜单关联', 3, 1, 1, 4, '', '/api/role/updateRoleMenuList', '', '角色菜单关联接口', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
INSERT INTO sys_menu (id, menu_name, menu_type, status_id, sort, parent_id, menu_url, api_url, menu_icon, remark, create_time, update_time) VALUES (18, '保存菜单接口', 3, 1, 1, 5, '', '/api/menu/addMenu', '', '保存菜单接口', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
INSERT INTO sys_menu (id, menu_name, menu_type, status_id, sort, parent_id, menu_url, api_url, menu_icon, remark, create_time, update_time) VALUES (19, '删除菜单接口', 3, 1, 1, 5, '', '/api/menu/deleteMenuById', '', '删除菜单接口', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
INSERT INTO sys_menu (id, menu_name, menu_type, status_id, sort, parent_id, menu_url, api_url, menu_icon, remark, create_time, update_time) VALUES (20, '更新菜单接口', 3, 1, 1, 5, '', '/api/menu/updateMenu', '', '更新菜单接口', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
INSERT INTO sys_menu (id, menu_name, menu_type, status_id, sort, parent_id, menu_url, api_url, menu_icon, remark, create_time, update_time) VALUES (21, '日志管理', 1, 1, 2, 0, '/log1', '', 'Setting', '', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
INSERT INTO sys_menu (id, menu_name, menu_type, status_id, sort, parent_id, menu_url, api_url, menu_icon, remark, create_time, update_time) VALUES (22, '登录日志', 2, 1, 1, 21, '/loginLog', '', 'Setting', '', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
INSERT INTO sys_menu (id, menu_name, menu_type, status_id, sort, parent_id, menu_url, api_url, menu_icon, remark, create_time, update_time) VALUES (23, '操作日志', 2, 1, 1, 21, '/sysLog', '', 'Setting', '', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
