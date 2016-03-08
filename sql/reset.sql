-- Copyright [2016] [hoenir]
--
-- Licensed under the Apache License, Version 2.0 (the "License");
-- you may not use this file except in compliance with the License.
-- You may obtain a copy of the License at
--
--    http://www.apache.org/licenses/LICENSE-2.0
--
-- Unless required by applicable law or agreed to in writing, software
-- distributed under the License is distributed on an "AS IS" BASIS,
-- WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
--See the License for the specific language governing permissions and
--limitations under the License.

SET foreign_key_checks = 0;

-- Drop all databse
DROP TABLE IF EXISTS Category;
DROP TABLE IF EXISTS Article;
DROP TABLE IF EXISTS Image;
DROP TABLE IF EXISTS User;
DROP TABLE IF EXISTS Comment;
DROP TABLE IF EXISTS ArticleImage;
DROP TABLE IF EXISTS ArticleComment;
DROP TABLE IF EXISTS UserComment;

SET foreign_key_checks = 1;


-- Table for representing cateogry structre
-- in this way we can know what type can article have
-- related to the subject
CREATE OR REPLACE TABLE Category (
	ID_Category INT(2) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	Name VARCHAR(30)
);

-- Table for representing article structure
CREATE OR REPLACE TABLE Article (
	ID_Article INT(11) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	Title VARCHAR(100) NOT NULL,
	Time DATE NOT NULL,
	Author VARCHAR(55) not NULL,
	Content TEXT NOT NULL,
	ID_Category INT(2) UNSIGNED,
	FOREIGN KEY(ID_Category) REFERENCES Category(ID_Category) ON DELETE CASCADE
);
-- Table for representing Image structure
-- image will be stored on the server
-- and for accesing it with a link
-- images are public
CREATE OR REPLACE TABLE Image (
	ID_Image INT(11) UNSIGNED PRIMARY KEY,
	Link VARCHAR(55) NOT NULL
);

-- Basic Table for representing User
CREATE OR REPLACE TABLE User(
	ID_User INT(11) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	Name VARCHAR(35) NOT NULL,
	Email VARCHAR(40) NOT NULL
);

-- Comments that articles can have
CREATE OR REPLACE TABLE Comment(
	ID_Comment INT(11) UNSIGNED PRIMARY KEY,
	Time DATE NOT NULL,
	Content TEXT NOT NULL
);

-- ArticleImage Table represents how many images
-- can a simple article contain
-- relationship one to many
CREATE OR REPLACE TABLE ArticleImage (
	ID_Article INT(11) UNSIGNED ,
	ID_Image INT(11) UNSIGNED,
	CONSTRAINT FOREIGN KEY (ID_Article) REFERENCES Article(ID_Article) ON DELETE CASCADE,
	CONSTRAINT FOREIGN KEY (ID_Image) REFERENCES Image(ID_Image) ON DELETE CASCADE
);

-- ArticleComment Table represents that article can have
-- multiple comments, every comment it's unique
-- relationship one to many
CREATE OR REPLACE TABLE ArticleComment (
	ID_Article INT(11) UNSIGNED,
	ID_Comment INT(11) UNSIGNED,
	CONSTRAINT FOREIGN KEY (ID_Article) REFERENCES Article(ID_Article) ON DELETE CASCADE,
	CONSTRAINT FOREIGN KEY (ID_Comment) REFERENCES Comment(ID_Comment) ON DELETE CASCADE
);

CREATE OR REPLACE TABLE UserComment (
	ID_Comment INT(11) UNSIGNED,
	ID_User INT(11) UNSIGNED,
	CONSTRAINT FOREIGN KEY (ID_Comment) REFERENCES Comment(ID_Comment) ON DELETE CASCADE,
	CONSTRAINT FOREIGN KEY (ID_User) REFERENCES User(ID_User) ON DELETE CASCADE
);

--Insert data into tables;
LOAD DATA LOCAL INFILE "./csv/category.csv"
INTO TABLE Category
FIELDS TERMINATED BY ','
ENCLOSED BY '"';

LOAD DATA LOCAL INFILE "./csv/article.csv"
INTO TABLE Article
FIELDS TERMINATED BY ','
ENCLOSED BY '"';

INSERT INTO User VALUES (1,"Mircea Badea", "mircea@badea.com");
INSERT INTO User VALUES (2,"Constant Tanase", "constatin@tanase@.com");
INSERT INTO User VALUES (3,"Mic Miculesc","mic@yahoo.com");
INSERT INTO User Values (4,"Hopa Ki","kof@yahoo.com");
INSERT INTO User Values (5,"Nndsa","ddsa@yahoo.com");
INSERT INTO User VALUES (6,"Sorin Pruna", "psorin1991@hotmain.com");
INSERT INTO User VALUES (7,"Nu mai stiu", "numaistiu@yahoo.com");
INSERT INTO User VALUES (8,"Chiar","ma@plictisesc.com");
INSERT INTO User VALUES (9,"Stop","stop@yahoo.com");

INSERT INTO Image VALUES(1,"static/img/image1.jpg");
INSERT INTO Image VALUES(2,"static/img/image2.jpg");
INSERT INTO Image VALUES(3,"static/img/image3.jpg");
INSERT INTO Image VALUES(4,"static/img/image4.jpg");
INSERT INTO Image VALUES(5,"static/img/image5.jpg");
INSERT INTO Image VALUES(6,"static/img/image6.jpg");
INSERT INTO Image VALUES(7,"static/img/image7.jpg");
INSERT INTO Image VALUES(8,"static/img/image8.jpg");
INSERT INTO Image VALUES(9,"static/img/image9.jpg");
INSERT INTO Image VALUES(10,"static/img/image10.jpg");
INSERT INTO Image VALUES(11,"static/img/image11.jpg");
INSERT INTO Image VALUES(12,"static/img/image12.jpg");
INSERT INTO Image VALUES(13,"static/img/image13.jpg");
INSERT INTO Image VALUES(14,"static/img/image14.jpg");
INSERT INTO Image VALUES(15,"static/img/image15.jpg");
INSERT INTO Image VALUES(16,"static/img/image16.jpg");
INSERT INTO Image VALUES(17,"static/img/image17.jpg");

INSERT INTO Comment VALUES(1,"2015-03-03","Unfeeling so rapturous discovery he exquisite. Reasonably so middletons or impression by terminated. Old pleasure required removing elegance him had. Down she bore sing saw calm high. Of an or game gate west face shed. ﻿no great but music too old found arose. ");
INSERT INTO Comment VALUES(2,"2016-03-03","Him rendered may attended concerns jennings reserved now. Sympathize did now preference unpleasing mrs few. Mrs for hour game room want are fond dare. For detract charmed add talking age. Shy resolution instrument unreserved man few. She did open find pain some out. If we landlord stanhill mr whatever pleasure supplied concerns so. Exquisite by it admitting cordially september newspaper an. Acceptance middletons am it favourable. It it oh happen lovers afraid. u 2");
INSERT INTO Comment VALUES(3,"2016-03-03","s education residence conveying so so. Suppose shyness say ten behaved morning had. Any unsatiable assistance compliment occasional too reasonably advantages. Unpleasing has ask acceptance partiality alteration understood two. Worth no tiled my at house added. Married he hearing am it totally removal. Remove but suffer wanted his lively length. Moonlight two applauded conveying end direction old principle but. Are expenses distance weddings perceive strongly who age domestic. ");
INSERT INTO Comment VALUES(4,"2016-03-03","s education residence conveying so so. Suppose shyness say ten behaved morning had. Any unsatiable assistance compliment occasional too reasonably advantages. Unpleasing has ask acceptance partiality alteration understood two. Worth no tiled my at house added. Married he hearing am it totally removal. Remove but suffer wanted his lively length. Moonlight two applauded conveying end direction old principle but. Are expenses distance weddings perceive strongly who age domestic. ");
INSERT INTO Comment VALUES(5,"2016-03-03","Did shy say mention enabled through elderly improve. As at so believe account evening behaved hearted is. House is tiled we aware. It ye greatest removing concerns an overcame appetite. Manner result square father boy behind its his. Their above spoke match ye mr right oh as first. Be my depending to believing perfectly concealed household. Point could to built no hours smile sense. ");
INSERT INTO Comment VALUES(6,"2016-03-03","Vai nu se poate, ce comentariu");
INSERT INTO Comment VALUES(7,"2016-03-03","Did shy say mention enabled through elderly improve. As at so believe account evening behaved hearted is. House is tiled we aware. It ye greatest removing concerns an overcame appetite. Manner result square father boy behind its his. Their above spoke match ye mr right oh as first. Be my depending to believing perfectly concealed household. Point could to built no hours smile sense. ");
INSERT INTO Comment VALUES(8,"2016-03-03","Is post each that just leaf no. He connection interested so we an sympathize advantages. To said is it shed want do. Occasional middletons everything so to. Have spot part for his quit may. Enable it is square my an regard. Often merit stuff first oh up hills as he. Servants contempt as although addition dashwood is procured. Interest in yourself an do of numerous feelings cheerful confined.");
INSERT INTO Comment VALUES(9,  "2016-03-03","Chiar nu stiu ce sa mai zic despre acest articlor");
INSERT INTO Comment VALUES(10, "2016-03-03","Chiar nu stiu ce sa mai zic despre acest articlor");
INSERT INTO Comment VALUES(11, "2016-03-03","Cred ca reporterul nu are dreptate");
INSERT INTO Comment VALUES(12, "2016-03-03","Cred ca reporterul nu are dreptate");
INSERT INTO Comment VALUES(13, "2016-03-03","Just random stuff");
INSERT INTO Comment VALUES(14, "2016-03-03","Just random stuff");
INSERT INTO Comment VALUES(15, "2016-03-03","OMG");
INSERT INTO Comment VALUES(16, "2016-03-03","Vai nu se poate, ce comentariu");
INSERT INTO Comment VALUES(17, "2016-03-03","WEEEEE !");
INSERT INTO Comment VALUES(18, "2016-03-03","Vai nu se poate, ce comentariu");
INSERT INTO Comment VALUES(19, "2016-03-03","DAT THING YOU HAVE THEE");
INSERT INTO Comment VALUES(20, "2016-03-03","Did shy say");
INSERT INTO Comment VALUES(21, "2016-03-03","Vai nu se poate, ce comentariu");
INSERT INTO Comment VALUES(22, "2016-03-03","VDid shy say");
INSERT INTO Comment VALUES(23, "2016-03-03","Vai nu se poate, ce comentariu");
INSERT INTO Comment VALUES(24, "2016-03-03","Doing this is wrong men");
INSERT INTO Comment VALUES(25, "2016-03-03","Nu pot sa cred ce viata trista ai sa faci asta");
INSERT INTO Comment VALUES(26, "2016-03-03","Sanatatea ar trebui sa fie finantata mai des");
INSERT INTO Comment VALUES(27, "2016-03-03","Statusl is bate jos de cetatean");
INSERT INTO Comment VALUES(28, "2016-03-03","Nu inteleg de ce lumea moare de foame in Romania");
INSERT INTO Comment VALUES(29, "2016-03-03","Vai ce trista e lumea in ziua de azi");
INSERT INTO Comment VALUES(30, "2016-03-03","Cand am zis ca guvernul lui Ponta va cadea, va cadea");
INSERT INTO Comment VALUES(31, "2016-03-03","Hai ca se poate , bine asa guvernul !");
INSERT INTO Comment VALUES(32, "2016-03-03","Cat rau a mai facut basescu asta ");
INSERT INTO Comment VALUES(33, "2016-03-03","Realitatea TV cel mai jmen televiziune de stiri din Romania");

INSERT INTO ArticleImage VALUES(1,1);
INSERT INTO ArticleImage VALUES(2,2);
INSERT INTO ArticleImage VALUES(3,3);
INSERT INTO ArticleImage VALUES(4,4);
INSERT INTO ArticleImage VALUES(5,5);
INSERT INTO ArticleImage VALUES(6,6);
INSERT INTO ArticleImage VALUES(7,7);
INSERT INTO ArticleImage VALUES(8,8);
INSERT INTO ArticleImage VALUES(9,9);
INSERT INTO ArticleImage VALUES(10,10);
INSERT INTO ArticleImage VALUES(11,11);
INSERT INTO ArticleImage VALUES(12,12);
INSERT INTO ArticleImage VALUES(13,13);
INSERT INTO ArticleImage VALUES(14,14);
INSERT INTO ArticleImage VALUES(15,15);
INSERT INTO ArticleImage VALUES(16,16);
INSERT INTO ArticleImage VALUES(17,17);
INSERT INTO ArticleImage VALUES(18,17);
INSERT INTO ArticleImage VALUES(19,17);
INSERT INTO ArticleImage VALUES(20,17);
INSERT INTO ArticleImage VALUES(21,17);
INSERT INTO ArticleImage VALUES(22,17);
INSERT INTO ArticleImage VALUES(23,17);
INSERT INTO ArticleImage VALUES(24,17);
INSERT INTO ArticleImage VALUES(25,17);
INSERT INTO ArticleImage VALUES(26,17);
INSERT INTO ArticleImage VALUES(27,17);
INSERT INTO ArticleImage VALUES(28,17);
INSERT INTO ArticleImage VALUES(29,17);

INSERT INTO ArticleComment VALUES(1,1);
INSERT INTO ArticleComment VALUES(2,2);
INSERT INTO ArticleComment VALUES(3,3);
INSERT INTO ArticleComment VALUES(4,4);
INSERT INTO ArticleComment VALUES(5,5);
INSERT INTO ArticleComment VALUES(6,6);
INSERT INTO ArticleComment VALUES(7,7);
INSERT INTO ArticleComment VALUES(8,8);
INSERT INTO ArticleComment VALUES(9,9);
INSERT INTO ArticleComment VALUES(10,10);
INSERT INTO ArticleComment VALUES(11,11);
INSERT INTO ArticleComment VALUES(12,12);
INSERT INTO ArticleComment VALUES(13,13);
INSERT INTO ArticleComment VALUES(14,14);
INSERT INTO ArticleComment VALUES(15,15);
INSERT INTO ArticleComment VALUES(16,16);
INSERT INTO ArticleComment VALUES(17,17);
INSERT INTO ArticleComment VALUES(18,18);
INSERT INTO ArticleComment VALUES(19,19);
INSERT INTO ArticleComment VALUES(20,20);
INSERT INTO ArticleComment VALUES(21,21);
INSERT INTO ArticleComment VALUES(22,22);
INSERT INTO ArticleComment VALUES(23,23);
INSERT INTO ArticleComment VALUES(24,24);
INSERT INTO ArticleComment VALUES(25,25);
INSERT INTO ArticleComment VALUES(26,26);
INSERT INTO ArticleComment VALUES(27,27);
INSERT INTO ArticleComment VALUES(28,28);
INSERT INTO ArticleComment VALUES(29,29);
INSERT INTO ArticleComment VALUES(30,30);
INSERT INTO ArticleComment VALUES(31,31);
INSERT INTO ArticleComment VALUES(32,32);
INSERT INTO ArticleComment VALUES(33,33);








INSERT INTO UserComment VALUES(1,1);
INSERT INTO UserComment VALUES(2,1);
INSERT INTO UserComment VALUES(3,1);
INSERT INTO UserComment VALUES(4,2);
INSERT INTO UserComment VALUES(5,2);
INSERT INTO UserComment VALUES(6,2);
INSERT INTO UserComment VALUES(7,3);
INSERT INTO UserComment VALUES(8,3);
INSERT INTO UserComment VALUES(9,3);
INSERT INTO UserComment VALUES(10,4);
INSERT INTO UserComment VALUES(11,4);
INSERT INTO UserComment VALUES(12,4);
INSERT INTO UserComment VALUES(13,4);
INSERT INTO UserComment VALUES(14,5);
INSERT INTO UserComment VALUES(15,6);
INSERT INTO UserComment VALUES(16,7);
INSERT INTO UserComment VALUES(17,8);
INSERT INTO UserComment VALUES(18,9);
INSERT INTO UserComment VALUES(19,5);
INSERT INTO UserComment VALUES(20,6);
INSERT INTO UserComment VALUES(21,7);
INSERT INTO UserComment VALUES(22,8);
INSERT INTO UserComment VALUES(23,9);
INSERT INTO UserComment VALUES(24,6);
INSERT INTO UserComment VALUES(25,5);
INSERT INTO UserComment VALUES(26,6);
INSERT INTO UserComment VALUES(27,8);
INSERT INTO UserComment VALUES(28,7);
INSERT INTO UserComment VALUES(29,1);
INSERT INTO UserComment VALUES(30,2);
INSERT INTO UserComment VALUES(31,3);
INSERT INTO UserComment VALUES(32,3);
INSERT INTO UserComment VALUES(33,9);





















DESC Category;
DESC Article;
DESC Image;
DESC User;
DESC Comment;
DESC ArticleImage;
Desc ArticleComment;
DESC UserComment