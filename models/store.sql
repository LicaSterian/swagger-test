CREATE DATABASE bookstore DEFAULT CHARSET = utf8 COLLATE = utf8_unicode_ci;

USE bookstore;

CREATE TABLE books (
    id BINARY(16) NOT NULL,
    bookName VARCHAR(50),
    authorName VARCHAR(50),
    publishDate DATE,
    PRIMARY key (id)
);

CREATE TABLE reviewers (
    id BINARY(16) NOT NULL,
    name VARCHAR(50),
    PRIMARY key (id),
    UNIQUE key (name)
);

CREATE TABLE reviews (
    bookId BINARY(16) NOT NULL,
    reviewerId BINARY(16) NOT NULL,
    review VARCHAR(512),
    PRIMARY KEY (bookId, reviewerId)
);