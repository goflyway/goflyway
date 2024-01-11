CREATE TABLE `{{.schema}}`.`{{.table}}` (
            `installed_rank` INT NOT NULL,
            `version` VARCHAR(50),
            `description` VARCHAR(200) NOT NULL,
            `type` VARCHAR(20) NOT NULL,
            `script` VARCHAR(1000) NOT NULL,
            `checksum` INT,
            `installed_by` VARCHAR(100) NOT NULL,
            `installed_on` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
            `execution_time` INT NOT NULL,
            `success` BOOL NOT NULL,
            CONSTRAINT `{{.table}}_pk`PRIMARY KEY (`installed_rank`)
) ENGINE=InnoDB;

CREATE INDEX `{{.table}}_s_idx` ON `{{.schema}}`.`{{.table}}` (`success`);