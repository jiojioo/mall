create table user
(
    id            bigint  NOT NULL  auto_increment,
    username      varchar(30)               unique             not null comment '用户名',
    password      varchar(255)                           not null comment '密码',
    realname      varchar(60)  not null comment '姓名',
    gender tinyint(3) not null default '0' comment '用户性别',
    phone         varchar(11)  not null comment '手机号',
    email         varchar(120) not null comment '邮箱',
    avatar_url VARCHAR(512) DEFAULT '' COMMENT '用户头像 URL',
    is_admin      tinyint unsigned default 0                 null comment '1管理员',
    pwd_err_times tinyint unsigned default 0                 null comment '密码错误次数',
    login_at      datetime null comment '上次登录时间',
    created_at    datetime     default CURRENT_TIMESTAMP not null,
    updated_at    datetime     default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP,
    deleted_at    datetime                 null,
    PRIMARY KEY (id)
) comment '用户表' charset = utf8mb4
                     row_format = DYNAMIC;
create index idx_username
    on user (username);

create index idx_user_deleted_at
    on user (deleted_at);