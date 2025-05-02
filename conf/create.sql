create table t_order
(
    id bigint auto_increment primary key,
    order_id varchar(255) not null,
    sku_id bigint not null,
    num int not null,
    price int not null,
    uid bigint not null,
    ctime timestamp default CURRENT_TIMESTAMP not null,
    utime timestamp default CURRENT_TIMESTAMP not null,
    constraint order_pk2 unique (order_id)
);


create table t_sku
(
    id bigint auto_increment primary key,
    name varchar(10) not null,
    price int null comment '分为单位',
    num int null,
    ctime timestamp default CURRENT_TIMESTAMP not null,
    utime timestamp default CURRENT_TIMESTAMP not null
);

create table t_user
(
    id bigint auto_increment primary key,
    name varchar(20) not null,
    ctime timestamp default CURRENT_TIMESTAMP not null,
    utime timestamp default CURRENT_TIMESTAMP not null
);