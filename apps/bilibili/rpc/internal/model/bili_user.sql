create table `bilibili_user`
(
    id              bigint unsigned auto_increment,
    dedeuserid      tinytext       null,
    username        tinytext       null,
    coins           decimal(10, 2) null,
    current_exp     int            null,
    next_exp        int            null,
    is_login        bool           null,
    upgrade_days    int            null,
    level           int            null,
    medals          text           null,
    vip_status      tinyint        null,
    vip_type        tinyint        null,
    vip_label_theme varchar(20)    null,
    sign            varchar(100)   null,
    last_run_time   datetime       null,
    create_time     datetime       null,
primary key (`id`)
);

