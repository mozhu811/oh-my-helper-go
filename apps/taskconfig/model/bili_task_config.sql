create table bili_task_config
(
    id                   bigint unsigned not null auto_increment,
    dedeuserid           varchar(100)    not null,
    sessdata             varchar(100)    not null,
    bili_jct             varchar(100)    not null,
    donate_coins         int             not null,
    reserve_coins        int             not null,
    auto_charge          bool            not null,
    donate_gift          bool            not null,
    donate_gift_target   varchar(20)     not null,
    auto_charge_target   varchar(20)     not null,
    device_platform      varchar(20)     not null,
    donate_coin_strategy int             not null,
    user_agent           varchar(100)    not null,
    skip_task            bool            not null,
    follow_developer     bool            not null,
    create_time          timestamp       NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '数据创建时间[禁止在代码中赋值]',
    update_time          timestamp       NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '数据更新时间[禁止在代码中赋值]',
    primary key (`id`)
);

