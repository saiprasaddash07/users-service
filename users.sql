CREATE TABLE users(
    userId int auto_increment primary key,
    firstName varchar(50) NOT NULL,
    lastName varchar(50) NOT NULL,
    password varchar(100) NOT NULL,
    mobileNo varchar(15) NOT NULL,
    email varchar(50) NOT NULL,
    gender int default 1 NOT NULL,
    isDeleted varchar(10) default 'false' NOT NULL,
    createdAt datetime default CURRENT_TIMESTAMP NOT NULL,
    updatedAt datetime default CURRENT_TIMESTAMP NOT NULL ON UPDATE CURRENT_TIMESTAMP
);
CREATE INDEX `userId` ON users(userId);
CREATE INDEX `email` ON users(email);
CREATE INDEX `gender` ON users(gender);
CREATE INDEX `mobileNo` ON users(mobileNo);
CREATE INDEX `name` ON users(firstName, lastName);