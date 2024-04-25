use book_store;
create table catalog
(
    catalog_id        bigint auto_increment comment '分类ID'
        primary key,
    catalog_code      varchar(10)  not null comment '分类编码',
    catalog_name      varchar(100) null comment '分类名称',
    parent_catalog_id varchar(10)  not null comment '父编码'
)
    comment '分类' charset = utf8;

create table classic
(
    classic_id   bigint auto_increment comment '出版物性质ID'
        primary key,
    classic_name varchar(50) null comment '出版物性质名称'
)
    comment '出版物_性质' charset = utf8;

create table language
(
    language_id   bigint auto_increment comment '出版物语言ID'
        primary key,
    language_name varchar(100) null comment '出版物语言名称',
    language_code varchar(100) null
)
    comment '维度_出版物_语言' charset = utf8;

create table publication_author
(
    author_id      bigint auto_increment comment '作者ID'
        primary key,
    isbn           varchar(22)  not null comment '国际标准书号',
    publication_id bigint       null comment '出版物ID',
    author         varchar(100) null comment '著者姓名',
    author_about   varchar(20)  null comment '作者简介'
)
    comment '图书作者' charset = utf8;

create table publication_detail
(
    detail_id      bigint auto_increment comment '详情ID'
        primary key,
    isbn           varchar(22)   not null comment '国际标准书号',
    detail_html    varchar(4000) null comment '详情',
    detail_img     varchar(4000) null comment '详情',
    publication_id bigint        null comment '出版物ID'
)
    comment '图书详情' charset = utf8;

create table publication_img
(
    img_id         bigint auto_increment
        primary key,
    img_url        varchar(256)  null comment '图片地址',
    isbn           bigint        not null comment '图书编码',
    img_encode     varchar(4000) null,
    publication_id bigint        null comment '出版物ID',
    mainFlag       int           null
)
    comment '出版物图片';

create table publication_info
(
    publication_id   bigint auto_increment comment '出版物ID'
        primary key,
    publication_name varchar(500)  null comment '出版物名称',
    org_id           varchar(20)   null comment '出版单位编码',
    published_times  varchar(50)   null comment '出版物版次',
    print_times      varchar(50)   null comment '出版物印次',
    price            decimal(4, 1) null comment '出版物单价(元)',
    introduction     varchar(2000) null comment '出版物内容简介',
    word_count       varchar(10)   null comment '出版物字数(千字)',
    isbn             varchar(20)   null comment '索书号',
    storage_by       varchar(20)   null comment '入库用户编码',
    storage_at       timestamp     null comment '出版物入库日期',
    modified_by      varchar(20)   null comment '更新用户编码',
    created_at       timestamp     null comment '出版物更新日期',
    modified_at      timestamp     null comment '出版物更新日期'
)
    comment '出版物_信息' charset = utf8;

create table publication_org
(
    org_id      bigint auto_increment comment '出版单位ID'
        primary key,
    org_code    varchar(20)  not null comment '出版单位编码',
    org_name    varchar(100) null comment '出版单位名称',
    org_address varchar(500) null comment '出版单位住所',
    modified_at timestamp    null comment '数据更新日期',
    created_at  timestamp    null comment '数据创建日期'
)
    comment '出版单位_基本信息' charset = utf8;

create table publication_x_catalog
(
    publication_catalog_id bigint auto_increment comment '出版物分类关系ID'
        primary key,
    catalog_id             bigint        not null comment '详情ID',
    isbn                   varchar(22)   not null comment '国际标准书号',
    catalog_name           varchar(4000) null comment '详情',
    publication_id         bigint        null comment '出版物ID'
)
    comment '出版物分类关系' charset = utf8;

create table publication_x_classic
(
    publication_classic_id bigint auto_increment
        primary key,
    classic_id             bigint        not null comment '详情ID',
    isbn                   varchar(22)   not null comment '国际标准书号',
    classic_name           varchar(4000) null comment '详情',
    publication_id         bigint        null comment '出版物ID'
)
    comment '出版物性质关系' charset = utf8;

create table publication_x_language
(
    publication_language_id bigint auto_increment comment '出版物语言关系ID'
        primary key,
    language_id             bigint        not null comment '详情ID',
    isbn                    varchar(22)   not null comment '国际标准书号',
    language_name           varchar(4000) null comment '详情',
    publication_id          bigint        null comment '出版物ID'
)
    comment '出版物语言关系' charset = utf8;

create table publication_stock
(
    stock_id bigint auto_increment
        primary key comment '库存Id',
    isbn                   varchar(22)   not null comment '国际标准书号',
    quantity          int8 null comment '数量',
    publication_id         bigint        null comment '出版物ID'
)
    comment '出版物库存' charset = utf8;
CREATE TABLE cart (
    -- 购物车ID，主键，自增整数
    cart_id INT AUTO_INCREMENT PRIMARY KEY,

    -- 用户ID，与用户表关联，标识拥有该购物车的用户
    user_id INT NOT NULL,

    -- 创建时间，记录购物车创建的时间
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    -- 更新时间，记录购物车最后一次更新的时间
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    -- 外键约束，确保引用的用户数据存在
    FOREIGN KEY (user_id) REFERENCES user(user_id)
) COMMENT '购物车表';
CREATE TABLE cart_item (
    -- 购物车项ID，主键，自增整数
    item_id INT AUTO_INCREMENT PRIMARY KEY,

    -- 购物车ID，与购物车表关联，标识该购物车项所属的购物车
    cart_id INT NOT NULL,

    -- 书籍ID，与书籍表关联，标识购物车项中的具体书籍
    book_id INT NOT NULL,

    -- 购买数量，用户选择购买该书籍的数量
    quantity INT NOT NULL DEFAULT 1,

    -- 选中状态，用于实现批量操作（如全选、反选），通常为布尔值，0表示未选中，1表示已选中
    is_selected BOOLEAN NOT NULL DEFAULT 0,

    -- 加入购物车的时间，用于记录用户添加商品到购物车的具体时间
    added_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    -- 外键约束，确保引用的数据存在
    FOREIGN KEY (cart_id) REFERENCES cart(cart_id),
    FOREIGN KEY (book_id) REFERENCES book(book_id)
) COMMENT '购物车项表';


