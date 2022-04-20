CREATE TABLE IF NOT EXISTS `runoob_tbl`(
   `runoob_id` INT UNSIGNED AUTO_INCREMENT,
   `runoob_title` VARCHAR(100) NOT NULL,
   `runoob_author` VARCHAR(40) NOT NULL,
   `submission_date` DATE,
   PRIMARY KEY ( `runoob_id` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `mf_fd_cache` (
  `id` bigint(18) NOT NULL AUTO_INCREMENT,
  `dep` varchar(3) NOT NULL DEFAULT '',
  `arr` varchar(3) NOT NULL DEFAULT '',
  `flightNo` varchar(10) NOT NULL DEFAULT '',
  `flightDate` date NOT NULL DEFAULT '1000-10-10',
  `flightTime` varchar(20) NOT NULL DEFAULT '',
  `isCodeShare` tinyint(1) NOT NULL DEFAULT '0',
  `tax` int(11) NOT NULL DEFAULT '0',
  `yq` int(11) NOT NULL DEFAULT '0',
  `cabin` char(2) NOT NULL DEFAULT '',
  `ibe_price` int(11) NOT NULL DEFAULT '0',
  `ctrip_price` int(11) NOT NULL DEFAULT '0',
  `official_price` int(11) NOT NULL DEFAULT '0',
  `uptime` datetime NOT NULL DEFAULT '1000-10-10 10:10:10',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uid` (`dep`,`arr`,`flightNo`,`flightDate`,`cabin`),
  KEY `uptime` (`uptime`),
  KEY `flight` (`dep`,`arr`),
  KEY `flightDate` (`flightDate`)
) ENGINE=InnoDB  DEFAULT CHARSET=gbk;