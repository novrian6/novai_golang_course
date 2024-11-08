-- Insert into users table
INSERT INTO public.users (username, email, password_hash, created_at) VALUES
    ('user1', 'user1@example.com', 'hashedpassword1', CURRENT_TIMESTAMP),
    ('user2', 'user2@example.com', 'hashedpassword2', CURRENT_TIMESTAMP),
    ('user3', 'user3@example.com', 'hashedpassword3', CURRENT_TIMESTAMP);

-- Insert into documents table with correct author_id references
INSERT INTO public.documents (title, description, author_id, price, file_path) VALUES
    ('Document Title 1', 'This is the description for document 1.', 1, 29.99, '/path/to/document1.pdf'),  -- Author is User 1
    ('Document Title 2', 'This is the description for document 2.', 2, 49.99, '/path/to/document2.pdf'),  -- Author is User 2
    ('Document Title 3', 'This is the description for document 3.', 1, 19.99, '/path/to/document3.pdf');  -- Author is User 1 again

-- Insert into purchases table
INSERT INTO public.purchases (user_id, document_id, purchased_at) VALUES
    (1, 1, CURRENT_TIMESTAMP),  -- User 1 purchases Document 1
    (2, 2, CURRENT_TIMESTAMP),  -- User 2 purchases Document 2
    (3, 1, CURRENT_TIMESTAMP);  -- User 3 purchases Document 1 again
