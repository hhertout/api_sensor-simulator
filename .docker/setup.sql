-- create the databases
CREATE DATABASE IF NOT EXISTS goapp;

-- create the users for each database
CREATE USER 'goapp'@'%' IDENTIFIED BY 'goapp';
GRANT CREATE, ALTER, INDEX, LOCK TABLES, REFERENCES, UPDATE, DELETE, DROP, SELECT, INSERT ON `projectone`.* TO 'goapp'@'%';

FLUSH PRIVILEGES;