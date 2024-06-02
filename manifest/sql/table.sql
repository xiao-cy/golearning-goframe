-- golearning.admin definition

CREATE TABLE `admin` (
  `adminId` bigint unsigned NOT NULL AUTO_INCREMENT,
  `mobile` varchar(32) NOT NULL,
  `password` varchar(32) NOT NULL,
  `department` varchar(16) NOT NULL,
  `adminName` varchar(16) NOT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`adminId`),
  KEY `idx_admin_mobile` (`mobile`),
  KEY `idx_admin_adminName` (`adminName`),
  KEY `idx_admin_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;