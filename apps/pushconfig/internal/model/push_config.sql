create table push_config
(
    id          bigint unsigned auto_increment comment '自增主键',
    user_id     varchar(32)                                 not null comment '用户id',
    config      json                                not null comment '推送配置，JSON格式存储',
    create_time timestamp default CURRENT_TIMESTAMP not null comment '创建时间',
    update_time timestamp default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
    primary key (id),
    constraint config_pusher_user_id_uindex
        unique (user_id)
);

