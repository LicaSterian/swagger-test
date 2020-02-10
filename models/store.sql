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

CREATE TABLE books_reviewers (
    bookId BINARY(16) NOT NULL,
    reviewerId BINARY(16) NOT NULL,
    PRIMARY KEY (bookId, reviewerId)
);

