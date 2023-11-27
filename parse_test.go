package sqldumptoobject_test

import (
	"fmt"
	"testing"

	"github.com/blastrain/vitess-sqlparser/sqlparser"
)

func TestLibr(t *testing.T) {
	// sql, err := os.Open("test.sql")
	// sqldumptocsv
	// Parse
	data, err := sqlparser.Parse("INSERT INTO `member` VALUES (10,'test','test','test','/src/assets/noProfile.png'),(11,'test1','test1','test1','/src/assets/noProfile.png'),(12,'test2','test2','test2','/src/assets/noProfile.png'),(13,'test3','test3','test3','/src/assets/noProfile.png'),(14,'test4','test4','test4','/src/assets/noProfile.png'),(15,'guddnr0421@naver.com','이형욱','752MBNTQFZ','/src/assets/noProfile.png');")
	// data.(*sqlparser.Insert)
	if err != nil {
		t.Error(err, "not found test.sql")
	}
	fmt.Printf("%s\n", data.(*sqlparser.Insert).Table.Name)

	ss, err := sqlparser.Parse("CREATE TABLE `user` ( " +
		"		`userId` varchar(30) NOT NULL COMMENT '유저아이디', " +
		"		`userName` text NOT NULL COMMENT '유저이름', " +
		"		`userEmail` text NOT NULL COMMENT '유저이메일', " +
		"		`userPass` text NOT NULL COMMENT '유저비번', " +
		"		`userSex` text DEFAULT NULL COMMENT '유저성별', " +
		"		`userWeight` int(11) DEFAULT NULL COMMENT '유저몸무게', " +
		"		`userHeight` int(11) DEFAULT NULL COMMENT '유저키', " +
		"		`userBirthday` date DEFAULT NULL COMMENT '유저생년월일', " +
		"		`userProfile` text DEFAULT NULL COMMENT '유저프로필이미지경로' " +
		"	  ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='유저정보';")

	if err != nil {
		t.Error(err, "not found test.sql")
	}

	fmt.Printf("%s\n", ss.(*sqlparser.Insert).Table.Name)

	// // data.Format()
	// // sqlparser.ColName
	// fmt.Printf("stmt : %+v", data)
}

// CREATE TABLE `user` (
// 	`userId` varchar(30) NOT NULL COMMENT '유저아이디',
// 	`userName` text NOT NULL COMMENT '유저이름',
// 	`userEmail` text NOT NULL COMMENT '유저이메일',
// 	`userPass` text NOT NULL COMMENT '유저비번',
// 	`userSex` text DEFAULT NULL COMMENT '유저성별',
// 	`userWeight` int(11) DEFAULT NULL COMMENT '유저몸무게',
// 	`userHeight` int(11) DEFAULT NULL COMMENT '유저키',
// 	`userBirthday` date DEFAULT NULL COMMENT '유저생년월일',
// 	`userProfile` text DEFAULT NULL COMMENT '유저프로필이미지경로'
//   ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='유저정보';

// CREATE TABLE IF NOT EXISTS `Container_Management` (
// 	`id` int(11) NOT NULL AUTO_INCREMENT,
// 	`submit_id` int(11) DEFAULT NULL,
// 	`status` int(11) DEFAULT NULL,
// 	`container_id` varchar(32) NOT NULL,
// 	`create_time` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
// 	`delete_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
// 	PRIMARY KEY (`id`) USING BTREE,
// 	KEY `FK_Container_Management_Submit` (`submit_id`),
// 	KEY `FK_Container_Management_Container_Status` (`status`),
// 	CONSTRAINT `FK_Container_Management_Container_Status` FOREIGN KEY (`status`) REFERENCES `Container_Status` (`id`) ON UPDATE NO ACTION,
// 	CONSTRAINT `FK_Container_Management_Submit` FOREIGN KEY (`submit_id`) REFERENCES `Submit` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
//   ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='와! 이것은 컨테이너를 관리합니다.';

// INSERT INTO `member` VALUES (10,'test','test','test','/src/assets/noProfile.png'),(11,'test1','test1','test1','/src/assets/noProfile.png'),(12,'test2','test2','test2','/src/assets/noProfile.png'),(13,'test3','test3','test3','/src/assets/noProfile.png'),(14,'test4','test4','test4','/src/assets/noProfile.png'),(15,'guddnr0421@naver.com','이형욱','752MBNTQFZ','/src/assets/noProfile.png');

// INSERT INTO `usermeasuredata` (`measureId`, `userId`, `deviceType`, `deviceAddress`, `measureTime`, `elapsedTime`, `totalDistance`, `instananeousPace`) VALUES
// ('measure_1650697217438', 'tjdgns52', 'XEBROW', 'D4:58:C8:93:C3:12', '2022-04-23 16:00:17', 141, 209, 166),
// ('measure_1650697453247', 'tjdgns53', 'XEBROW', 'D4:58:C8:93:C3:12', '2022-04-23 16:04:13', 261, 1386, 275),
// ('measure_1650703475104', 'tjdgns52', 'XEBROW', 'D4:58:C8:93:C3:12', '2022-04-23 17:44:35', 18, 43, 244),
// ('measure_1651227662851', 'tjdgns52', 'XEBROW', 'E9:0C:19:FC:59:12', '2022-04-29 19:21:02', 169, 659, 123);
