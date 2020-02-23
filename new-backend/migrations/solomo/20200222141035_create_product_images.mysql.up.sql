CREATE TABLE `product_images` ( 
    `id` INT NOT NULL AUTO_INCREMENT , 
    `product_id` INT NOT NULL , 
    `image` VARCHAR(255) NOT NULL , 
    `seq` INT NOT NULL DEFAULT '0' , 
    `is_main` TINYINT(1) NOT NULL DEFAULT '0' , 
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, 
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, 
    `deleted_at` TIMESTAMP NULL DEFAULT NULL , 
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;