ALTER TABLE books
    DROP CONSTRAINT books_name_author_unique;

ALTER TABLE books
    ALTER COLUMN created_at DROP DEFAULT;