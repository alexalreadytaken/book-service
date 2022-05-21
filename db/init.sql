CREATE TABLE book(
    id int not null auto_increment primary key,
    name text not null,
    pages_count int not null
);

CREATE TABLE author(
    id int not null auto_increment primary key,
    name text not null,
    surname text not null
);

CREATE TABLE book_author(
    book_id int not null,
    author_id int not null,
    FOREIGN KEY (book_id) REFERENCES book (id),
    FOREIGN KEY (author_id) REFERENCES author (id)
);

insert into book values (1,'Games People Play',563);
insert into book values (2,'Beyond Games and Scripts',500);

insert into author values (1,'Eric','Berne');
insert into author values (2,'*more*','*author*');

insert into book_author values (1,1);
insert into book_author values (2,1);
insert into book_author values (2,2);

# select distinct a.name,a.surname,b.name 
# from book_author as ba 
# inner join author as a on a.id=ba.author_id 
# inner join book as b on b.id=ba.book_id;