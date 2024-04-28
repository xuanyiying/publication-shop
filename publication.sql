-- create database book_store;
use book_store;

create table addresses
(
    id             bigint auto_increment
        primary key,
    name           varchar(255) not null,
    mobile         varchar(255) not null,
    address        varchar(255) not null,
    post_code      varchar(255) not null,
    created_at     datetime     not null,
    updated_at     datetime     not null,
    user_addresses bigint       null
)
    collate = utf8mb4_bin;

create table authors
(
    author_id int auto_increment comment '作者ID'
        primary key,
    name      varchar(100) not null comment '作者姓名',
    biography text         null comment '作者简介'
)
    comment '作者表';

create table books
(
    book_id          bigint auto_increment comment '书籍ID，主键'
        primary key,
    isbn             varchar(13)                        not null comment 'ISBN，唯一',
    title            varchar(255)                       not null comment '书名',
    author           varchar(100)                       not null comment '主编',
    translator       varchar(100)                       null comment '译者',
    publisher_id     bigint                             null comment '出版社ID',
    publisher        varchar(100)                       null comment '出版社',
    publication_year smallint                           null comment '出版年份',
    publication_date date                               not null comment '出版日期',
    edition          tinyint                            null comment '版次',
    category         varchar(50)                        null comment '类别',
    price            decimal(10, 2)                     null comment '价格',
    stock_quantity   int                                null comment '库存量',
    description      text                               null comment '书籍描述',
    added_on         datetime default CURRENT_TIMESTAMP null comment '添加时间',
    cover_image      varchar(255)                       null comment '图书封面图片URL',
    page_count       int                                not null comment '页数',
    language_id      bigint                             not null comment '语言',
    language         varchar(50)                        not null comment '语言',
    author_id        bigint                             not null comment '作者ID，外键约束确保其与实际作者表关联',
    category_id      bigint                             not null comment '图书类别ID，外键约束确保其与实际类别表关联',
    constraint isbn
        unique (isbn)
)
    comment '书籍表';

create table book_authors
(
    book_author_id bigint auto_increment
        primary key,
    author_id      bigint       not null comment '作者ID',
    isbn           varchar(22)  not null comment '国际标准书号',
    book_id        bigint       null comment '出版物ID',
    author         varchar(100) null comment '著者姓名',
    author_about   varchar(20)  null comment '作者简介'
)
    comment '图书作者' charset = utf8;

create table book_details
(
    detail_id   bigint auto_increment comment '详情ID'
        primary key,
    isbn        varchar(22)   not null comment '国际标准书号',
    detail_html varchar(4000) null comment '详情',
    detail_img  varchar(4000) null comment '详情',
    book_id     bigint        null comment '出版物ID'
)
    comment '图书详情' charset = utf8;

create table book_images
(
    img_id     bigint auto_increment
        primary key,
    img_url    varchar(256)  null comment '图片地址',
    isbn       bigint        not null comment '图书编码',
    img_encode varchar(4000) null,
    book_id    bigint        null comment '出版物ID',
    mainFlag   int           null
)
    comment '出版物图片';

create table book_stocks
(
    stock_id bigint auto_increment comment '库存Id'
        primary key,
    isbn     varchar(22) not null comment '国际标准书号',
    quantity bigint      null comment '数量',
    book_id  bigint      null comment '出版物ID'
)
    comment '出版物库存' charset = utf8;

create table book_x_catalogs
(
    book_catalog_id bigint auto_increment comment '出版物分类关系ID'
        primary key,
    catalog_id      bigint        not null comment '详情ID',
    isbn            varchar(22)   not null comment '国际标准书号',
    catalog_name    varchar(4000) null comment '详情',
    book_id         bigint        null comment '出版物ID'
)
    comment '出版物分类关系' charset = utf8;

create table book_x_classics
(
    book_classic_id bigint auto_increment
        primary key,
    classic_id      bigint        not null comment '详情ID',
    isbn            varchar(22)   not null comment '国际标准书号',
    classic_name    varchar(4000) null comment '详情',
    book_id         bigint        null comment '出版物ID'
)
    comment '出版物性质关系' charset = utf8;

create table book_x_languages
(
    book_language_id bigint auto_increment comment '出版物语言关系ID'
        primary key,
    language_id      bigint        not null comment '详情ID',
    isbn             varchar(22)   not null comment '国际标准书号',
    language_name    varchar(4000) null comment '详情',
    book_id          bigint        null comment '出版物ID'
)
    comment '出版物语言关系' charset = utf8;

create table cards
(
    id         bigint auto_increment
        primary key,
    name       varchar(255) not null,
    card_no    varchar(255) not null,
    ccv        varchar(255) not null,
    expires    varchar(255) not null,
    created_at datetime     not null,
    updated_at datetime     not null,
    user_cards bigint       null
)
    collate = utf8mb4_bin;

create table categories
(
    catalog_id        bigint auto_increment comment '分类ID'
        primary key,
    description      varchar(512)  not null comment '分类描述',
    catalog_name      varchar(100) null comment '分类名称',
    parent_catalog_id varchar(10)  not null comment '父编码'
)
    comment '分类' charset = utf8;

create table classics
(
    classic_id   bigint auto_increment comment '出版物性质ID'
        primary key,
    classic_name varchar(50) null comment '出版物性质名称'
)
    comment '出版物_性质' charset = utf8;

create table languages
(
    language_id   bigint auto_increment comment '出版物语言ID'
        primary key,
    language_name varchar(100) null comment '出版物语言名称',
    language_code varchar(100) null
)
    comment '维度_出版物_语言' charset = utf8;

create table orders
(
    order_id          bigint auto_increment comment '订单ID，主键'
        primary key,
    order_number      varchar(50)                                                                           null comment '订单编号',
    transaction_id    bigint                                                                                not null comment '交易ID',
    order_status      enum ('pending', 'processing', 'shipped', 'delivered', 'cancelled') default 'pending' not null comment '订单状态',
    delivered_address varchar(255)                                                                          null comment '配送地址',
    shipping_cost     decimal(10, 2)                                                                        null comment '运费',
    total_amount      decimal(10, 2)                                                                        null comment '总金额',
    placed_user_id    bigint                                                                                not null comment '下单用户ID',
    placed_at         datetime                                                                              null comment '下单时间',
    shipped_address   varchar(255)                                                                          null comment '发货地址',
    shipped_at        datetime                                                                              null comment '发货时间',
    payment_id        bigint                                                                                null comment '支付ID'
)
    comment '订单表';

create table payments
(
    payment_id     bigint auto_increment comment '支付ID'
        primary key,
    payment_number varchar(100)                                         not null comment '支付流水号',
    order_id       bigint                                               not null comment '订单ID，外键约束确保其与实际订单表关联',
    payment_method varchar(50)                                          not null comment '支付方式',
    transaction_id varchar(100)                                         null comment '交易ID',
    payment_at     datetime                                             not null comment '支付日期',
    amount_paid    decimal(10, 2)                                       not null comment '实付金额',
    payment_status enum ('unpaid', 'paid', 'refunded') default 'unpaid' not null comment '支付状态',
    constraint payment_number
        unique (payment_number)
)
    comment '支付表';

create table publishers
(
    publisher_id      bigint auto_increment comment '出版社ID，主键'
        primary key,
    publisher_name    varchar(100) not null comment '出版社名称',
    publisher_address varchar(255) null comment '出版社地址',
    publisher_phone   varchar(20)  null comment '出版社联系电话',
    publisher_email   varchar(255) null comment '出版社邮箱',
    publisher_website varchar(255) null comment '出版社网站'
)
    comment '出版社';

create table reviews
(
    review_id    bigint auto_increment comment '评论ID'
        primary key,
    book_id      bigint   not null comment '图书ID，外键约束确保其与实际图书表关联',
    user_id      bigint   not null comment '客户ID，外键约束确保其与实际客户表关联',
    rating       int      not null comment '评分（1-5星）',
    comment_text text     null comment '评论内容',
    created_at   datetime not null comment '创建时间'
)
    comment '评论表';

create table transactions
(
    transaction_id     bigint auto_increment comment '交易ID，主键'
        primary key,
    transaction_number varchar(50)                                                          not null comment '交易编号',
    transaction_type   enum ('buy', 'sell')                                                 not null comment '交易类型（买或卖）',
    user_id            bigint                                                               not null comment '用户ID',
    quantity           bigint                                                               not null comment '交易数量',
    transaction_status enum ('pending', 'completed', 'cancelled') default 'pending'         not null comment '交易状态',
    transaction_date   datetime                                   default CURRENT_TIMESTAMP null comment '交易日期',
    transaction_amount decimal(10, 2)                                                       null comment '交易金额',
    payment_id         bigint                                                               null comment '支付ID'
)
    comment '交易表';

create table transaction_items
(
    transaction_item_id bigint auto_increment comment '交易明细ID，主键'
        primary key,
    transaction_type    enum ('buy', 'sell') not null comment '交易类型（买或卖）',
    transaction_id      bigint               not null comment '交易ID',
    book_id             bigint               not null comment '书籍ID',
    quantity            int                  not null comment '交易数量',
    price               decimal(10, 2)       not null comment '交易价格',
    isbn                varchar(13)          not null comment 'ISBN',
    title               varchar(255)         not null comment '书名',
    author              varchar(100)         not null comment '主编',
    publisher_id        bigint               null comment '出版社ID',
    image_url           varchar(255)         null comment '图片URL'
)
    comment '交易项目表';

create table users
(
    user_id         bigint auto_increment comment '用户ID，主键'
        primary key,
    username        varchar(50)                        not null comment '用户名，唯一',
    email           varchar(255)                       not null comment '邮箱，唯一',
    password_hash   varchar(255)                       not null comment '密码哈希',
    first_name      varchar(50)                        null comment '名',
    last_name       varchar(50)                        null comment '姓',
    phone_number    varchar(20)                        null comment '联系电话',
    country         varchar(50)                        null comment '国家',
    city            varchar(50)                        null comment '城市',
    postal_code     varchar(20)                        null comment '邮政编码',
    university      varchar(255)                       null comment '大学',
    address         varchar(255)                       null comment '地址',
    registered_date datetime default CURRENT_TIMESTAMP null comment '注册日期',
    constraint email
        unique (email),
    constraint username
        unique (username)
)
    comment '用户表';

