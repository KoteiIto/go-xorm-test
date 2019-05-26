create database if not exists `xorm_master`;
create database if not exists `xorm_user`;

drop table if exists `xorm_master`.`group`;
create table `xorm_master`.`group`
(
        id int not null,
        name varchar(255) not null,
        description text not null,
        created_at datetime not null,
        updated_at datetime not null,
        deleted_at datetime,
        primary key (id)
);

drop table if exists `xorm_user`.`group_member`;
create table `xorm_user`.`group_member`
(
        user_id bigint unsigned not null,
        group_id int not null,
        role enum('guest', 'admin') not null,
		version int not null,
        created_at datetime not null,
        updated_at datetime not null,
        deleted_at datetime,
        primary key (user_id, group_id),
        constraint uq_group_member_col_user_group UNIQUE (user_id, group_id)
);

drop table if exists `xorm_user`.`user`;
create table `xorm_user`.`user`
(
        id bigint unsigned not null AUTO_INCREMENT,
        email varchar(255) not null,
        password_digest char(100),
        version int not null,
        created_at datetime not null,
        updated_at datetime not null,
        deleted_at datetime,
        primary key (id),
        constraint uq_user_col_email UNIQUE (email)
);

truncate table `xorm_master`.`group`;
truncate table `xorm_user`.`group_member`;
truncate table `xorm_user`.`user`;