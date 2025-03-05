-- Active: 1740897367239@@127.0.0.1@3306@gorm
CREATE DATABASE IF NOT EXISTS cart;

/*create table cart if not exists cart
(
    id         int auto_increment,
    user_id    int      not null,
    product_id int      not null,
    qty        int      not null,
    created_at datetime not null default current_timestamp,
    updated_at datetime not null default current_timestamp on update current_timestamp,
    constraint cart_pk primary key (id)
);*/