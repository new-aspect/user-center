--- 创建数据库
CREATE DATABASE user_center;

--- 创建表
-- 参考
       -- https://www.tutorialrepublic.com/sql-tutorial/sql-create-table-statement.php
       -- http://w3schools.com/mysql/mysql_create_table.asp
CREATE TABLE user (
  ID int,
  Username VARCHAR(255),
  Password VARCHAR(255),
  CreateAt DATETIME
);