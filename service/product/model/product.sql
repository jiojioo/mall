CREATE TABLE `product`
(
    `id`          bigint    NOT NULL AUTO_INCREMENT,
    `name`        varchar(255)        NOT NULL DEFAULT '' COMMENT '产品名称',
    `desc`        varchar(255)        NOT NULL DEFAULT '' COMMENT '产品描述',
    `stock`       int(10) unsigned    NOT NULL DEFAULT '0' COMMENT '产品库存',
    `amount`      int(10) unsigned    NOT NULL DEFAULT '0' COMMENT '产品金额',
    `status`      tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '产品状态',
    created_at    datetime     default CURRENT_TIMESTAMP not null,
    updated_at    datetime     default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP,
    deleted_at    datetime                 null,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;
create index idx_product_deleted_at
    on product (deleted_at);
