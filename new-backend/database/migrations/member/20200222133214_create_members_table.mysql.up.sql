CREATE TABLE `members` ( 
    `id` INT NOT NULL AUTO_INCREMENT , 
    `no` VARCHAR(30) NOT NULL , 
    `account` VARCHAR(255) NOT NULL , 
    `password` VARCHAR(255) NULL , 
    `image` VARCHAR(255) NULL , 
    `name` VARCHAR(100) NULL , 
    `nickname` VARCHAR(100) NULL , 
    `sex` TINYINT(1) NOT NULL DEFAULT '-1' , 
    `city` VARCHAR(100) NULL , 
    `area` VARCHAR(100) NULL , 
    `address` VARCHAR(255) NULL , 
    `mobile` VARCHAR(100) NULL , 
    `remark` TEXT NULL , 
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, 
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP NULL , 
    PRIMARY KEY (`id`), UNIQUE (`no`), UNIQUE (`account`)
) ENGINE = InnoDB;