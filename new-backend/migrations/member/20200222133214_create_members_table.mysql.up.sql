CREATE TABLE `solomo`. ( 
    `id` INT NOT NULL AUTO_INCREMENT , 
    `no` VARCHAR(30) NOT NULL , 
    `name` VARCHAR(100) NULL , 
    `nickname` VARCHAR(100) NULL , 
    `sex` TINYINT(1) NOT NULL DEFAULT '-1' , 
    `city` VARCHAR(100) NULL , 
    `area` VARCHAR(100) NULL , 
    `address` VARCHAR(255) NULL , 
    `mobile` VARCHAR(100) NULL , 
    `remark` TEXT NULL , 
    `created_at` TIMESTAMP NOT NULL , 
    `updated_at` TIMESTAMP NOT NULL , 
    `deleted_at` TIMESTAMP NULL , 
    PRIMARY KEY (`id`), UNIQUE (`no`)
) ENGINE = InnoDB;