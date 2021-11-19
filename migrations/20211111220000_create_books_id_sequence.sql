-- +goose Up
CREATE SEQUENCE books_id_seq START WITH 1 OWNED BY books.id;
ALTER TABLE books
    ALTER COLUMN id SET DEFAULT nextval('books_id_seq');

-- +goose Down
DROP SEQUENCE books_id_seq;
ALTER TABLE books
    ALTER COLUMN id TYPE int ;
ALTER TABLE books
    ALTER COLUMN id SET not null ;