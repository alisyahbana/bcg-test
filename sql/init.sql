CREATE SCHEMA `bcg_db`;

CREATE TABLE `bcg`.`catalog` (
   `id` int unsigned NOT NULL AUTO_INCREMENT,
   `sku` varchar(255) NOT NULL,
   `name` varchar(255) NOT NULL,
   `price` float unsigned NOT NULL,
   `qty` int unsigned NOT NULL,
   `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
   `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
   PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1