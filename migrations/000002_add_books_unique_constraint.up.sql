ALTER TABLE books
    ALTER COLUMN created_at SET DEFAULT now();

ALTER TABLE books
    ADD CONSTRAINT books_name_author_unique
        UNIQUE (name, author);