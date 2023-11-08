CREATE TABLE `shares`
(
    `id`          bigint unsigned NOT NULL AUTO_INCREMENT,
    `number`      tinyint unsigned NOT NULL,
    `protocol`    varchar(100) NOT NULL,
    `uid`         char(36)     NOT NULL,
    `address`     varchar(100) NOT NULL,
    `port`        smallint unsigned NOT NULL,
    `security`    varchar(100) NOT NULL,
    `encryption`  varchar(100) NOT NULL,
    `public_key`  char(43)     NOT NULL,
    `header_type` varchar(100) NOT NULL,
    `fingerprint` varchar(100) NOT NULL,
    `network`     varchar(100) NOT NULL,
    `flow`        varchar(100) NOT NULL,
    `sni`         varchar(100) NOT NULL,
    `short_ids`   char(16)     NOT NULL,
    `remarks`     varchar(100) NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `share_UN` (`number`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
