CREATE TABLE book(
    id int not null auto_increment primary key,
    name varchar(100) not null,
    pages_count int not null
);

CREATE TABLE author(
    id int not null auto_increment primary key,
    name varchar(50) not null,
    surname varchar(50) not null
);

CREATE TABLE book_author(
    book_id int not null,
    author_id int not null,
    FOREIGN KEY (book_id) REFERENCES book (id),
    FOREIGN KEY (author_id) REFERENCES author (id)
);

insert into book values (1,'Игры, в которые играют люди',563);
insert into book values (2,'За пределами игр и сценариев',500);

insert into author values (1,'Эрик','Берн');
insert into author values (2,'*Еще*','*Автор*');

insert into book_author values (1,1);
insert into book_author values (2,1);
insert into book_author values (2,2);

# select distinct a.name,a.surname,b.name 
# from book_author as ba 
# inner join author as a on a.id=ba.author_id 
# inner join book as b on b.id=ba.book_id;





