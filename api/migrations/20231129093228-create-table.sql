
-- +migrate Up
-- ==========================
-- master
-- ==========================
CREATE TABLE IF NOT EXISTS `pets` (
  `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` VARCHAR(255) NULL COMMENT '名前',
  `birth_day` DATETIME NULL COMMENT '誕生日',
  `sex` BIGINT NULL COMMENT '性別',
  `now_weight` BIGINT NULL COMMENT '現在体重',
  `target_weight` BIGINT NULL COMMENT '目標体重',
  `created_at` DATETIME NOT NULL COMMENT '作成日時',
  `updated_at` DATETIME NOT NULL COMMENT '更新日時',
  PRIMARY KEY (`id`))
ENGINE = InnoDB
COMMENT = 'ペットマスタ';

CREATE TABLE IF NOT EXISTS `schedules` (
  `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `pet_id` BIGINT NOT NULL COMMENT 'ペットID',
  `title` VARCHAR(255) NULL COMMENT 'タイトル',
  `date` DATETIME NOT NULL COMMENT '日時',
  `location` VARCHAR(45) NULL COMMENT '場所',
  `created_at` DATETIME NOT NULL COMMENT '作成日時',
  `updated_at` DATETIME NOT NULL COMMENT '更新日時',
  PRIMARY KEY (`id`),
  CONSTRAINT `fk_schedule1`
    FOREIGN KEY (`pet_id`)
    REFERENCES `pets` (`id`)
    -- ON DELETE CASCADE
    -- ON UPDATE NO ACTION
  )
ENGINE = InnoDB
COMMENT = '予定';

CREATE TABLE IF NOT EXISTS `conditions` (
  `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `pet_id` BIGINT NOT NULL COMMENT 'ペットID',
  `food` BIGINT NULL COMMENT '食事の調子',
  `sweat` BIGINT NULL COMMENT '汗の調子',
  `toilet` BIGINT NULL COMMENT '排泄の調子',
  `created_at` DATETIME NOT NULL COMMENT '作成日時',
  `updated_at` DATETIME NOT NULL COMMENT '更新日時',
  PRIMARY KEY (`id`, `pet_id`),
  CONSTRAINT `fk_schedule1`
    FOREIGN KEY (`pet_id`)
    REFERENCES `pets` (`id`)
    -- ON DELETE CASCADE
    -- ON UPDATE NO ACTION
  )
ENGINE = InnoDB
COMMENT = '体調履歴';


-- +migrate Down
DROP TABLE `schedules`;
DROP TABLE `pets`;

