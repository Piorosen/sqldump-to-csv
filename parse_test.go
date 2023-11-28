package sqldumptoobject_test

import (
	"testing"

	sqlldump "github.com/Piorosen/sqldump-to-object"
)

func TestLibr(t *testing.T) {
	// sql, err := os.Open("test.sql")
	// sqldumptocsv
	// Parse
	// data.(*sqlparser.Insert)

	ta, err := sqlldump.ConvertTable("CREATE TABLE `user` ( " +
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
		t.Error(err, "not found test.sql", err)
	}

	ta, err = sqlldump.ConvertInsert(ta, "INSERT INTO `member` VALUES (10,'test','test','test','/src/assets/noProfile.png'),(11,'test1','test1','test1','/src/assets/noProfile.png'),(12,'test2','test2','test2','/src/assets/noProfile.png'),(13,'test3','test3','test3','/src/assets/noProfile.png'),(14,'test4','test4','test4','/src/assets/noProfile.png'),(15,'guddnr0421@naver.com','이형욱','752MBNTQFZ','/src/assets/noProfile.png');")

	if err != nil {
		t.Error(err, "not found test.sql", err)
	}
}

func TestWithgTableId(t *testing.T) {
	create := "CREATE TABLE IF NOT EXISTS `User` ( " +
		"	`id` int(11) NOT NULL AUTO_INCREMENT," +
		"	`std_id` varchar(10) NOT NULL DEFAULT ''," +
		"	`name` varchar(50) NOT NULL," +
		"	`nickname` varchar(50) NOT NULL DEFAULT ''," +
		"	`password` varchar(513) NOT NULL DEFAULT ''," +
		"	`department` varchar(30) NOT NULL DEFAULT ''," +
		"	`email` varchar(321) NOT NULL DEFAULT ''," +
		"	`sign_up` datetime NOT NULL DEFAULT current_timestamp()," +
		"	`authority` int(11) NOT NULL DEFAULT 0," +
		"	`hide` int(11) NOT NULL DEFAULT 0," +
		"	PRIMARY KEY (`id`)," +
		"	UNIQUE KEY `std_id` (`std_id`)" +
		"  ) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;"

	insert := "INSERT INTO `User` (`id`, `std_id`, `name`, `nickname`, `password`, `department`, `email`, `sign_up`, `authority`, `hide`) VALUES " +
		" (0, 'sys_admin', '관리자', 'sys_admin', '$2b$12$r1Xff2/rcRSSowW0hWsXi.QNxAny3s9H0qdTK5625LJKelz8CjFKy', '관라자', 'admin@admin.com', '2023-11-20 00:18:02', 0, 1), " +
		" (1, '20193156', '정민규', 'mQueue', '$2b$12$pIj2js8HWT6v/N0Ifv3/zeKX0GXpm9aXQIEjS5p.f4xKD79S9ITT2', '응용소프트웨어공학', 'tomorrow9913@gmail.com', '2023-11-17 16:29:14', 0, 0), " +
		" (2, '20203218', '림미선', 'temp', '$2b$12$r1Xff2/rcRSSowW0hWsXi.QNxAny3s9H0qdTK5625LJKelz8CjFKy', '응용소프트웨어공학', 'user@example.com', '2023-11-18 06:51:56', 0, 0), " +
		" (4, 'testtest', 'testtest', 'testtest', '$2b$12$E4sQQl5K1vnjJYuq2.Cnuuk.5b4VFyaD2PHdDmYusa9GcNr.WFrVe', 'testtest', 'testtest@testtest.com', '2023-11-19 17:26:12', 1, 0), " +
		" (5, '11111111', '11111111', '11111111', '$2b$12$Er8k6f2cvk6G.bm2c8Q4DuAXDgtl5dsiQs9LR04s90pLvOTcytPHG', '11111111', '11111111@gmail.com', '2023-11-19 17:35:58', 1, 0), " +
		" (6, '22222222', '22222222', '22222222', '$2b$12$BAXWP409nz/7nohwavRYZeLrtWxmvbOUOpCJJD1ctlZn9g0gcfAsG', '22222222', '22222222@gmail.com', '2023-11-19 17:36:56', 1, 0), " +
		" (7, '33333333', '33333333', '33333333', '$2b$12$snMMnk2EX4N7F9h6jdoaz.432HQSztYvwxZPpMqMQhb.PupaDLAma', '33333333', '33333333@33333333.com', '2023-11-19 17:39:17', 1, 0); "

	table, err := sqlldump.ConvertTable(create)
	if err != nil {
		t.Error(err, "fail to convert table", err)
	}

	table, err = sqlldump.ConvertInsert(table, insert)
	if err != nil {
		t.Error(err, "fail to insert data", err)
	}
}
