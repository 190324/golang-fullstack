CREATE TABLE `products` ( 
    `id` INT NOT NULL AUTO_INCREMENT , 
    `no` VARCHAR(36) NOT NULL , 
    `name` VARCHAR(125) NOT NULL , 
    `brief` VARCHAR(255) NULL DEFAULT NULL , 
    `desp` TEXT NULL DEFAULT NULL , 
    `qty` INT NOT NULL DEFAULT '0' , 
    `price` FLOAT NOT NULL , 
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, 
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, 
    `deleted_at` TIMESTAMP NULL DEFAULT NULL , 
    PRIMARY KEY (`id`), UNIQUE (`no`)
) ENGINE = InnoDB;