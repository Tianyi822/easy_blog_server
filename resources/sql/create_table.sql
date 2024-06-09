create table START_PAGE_USER
(
    uid         char(36)     not null comment '用户 ID，随机字符串序列',
    user_name   char(30)     not null comment '用户昵称',
    user_pwd    varchar(100) not null comment '用户密码',
    salt        char(32)     not null comment '加密盐',
    create_time timestamp    not null comment '创建时间(只有创建的时候会添加数据，添加后不允许修改)',
    update_time timestamp    not null comment '修改时间',
    primary key (uid, user_name, salt)
)
    comment '起始页用户表';


create table START_PAGE_BOOKMARK
(
    bm_id       char(36)             not null comment '书签 ID',
    uid         char(36)             not null comment '用户 ID',
    bm_url      varchar(100)         not null comment '书签地址',
    bm_name     varchar(100)         not null comment '书签名',
    bm_color    char(7)              not null comment '书签颜色',
    is_resident tinyint(1) default 0 not null comment '是否为常驻标签',
    create_time timestamp            not null comment '创建时间',
    update_time timestamp            not null comment '修改时间',
    primary key (bm_id, uid),
    constraint START_PAGE_BOOKMARK_START_PAGE_USER_uid_fk
        foreign key (uid) references START_PAGE_USER (uid)
            on update cascade on delete cascade
)
    comment '收藏书签表';


create table START_PAGE_BACKGROUND_IMG
(
    img_id      char(36)     not null comment '图片 ID',
    uid         char(36)     not null comment '用户 ID',
    img_url     varchar(100) not null comment '图片 URL',
    create_time timestamp    not null comment '创建时间',
    update_time timestamp    not null comment '修改时间',
    primary key (img_id, uid),
    constraint START_PAGE_BACKGROUND_IMG_START_PAGE_USER_uid_fk
        foreign key (uid) references START_PAGE_USER (uid)
            on update cascade on delete cascade
)
    comment '背景图';

