-- Seeder script for public.faq_categories table
INSERT INTO public.faq_categories (uuid, name, status, created_at, updated_at, deleted_at)
VALUES ('512c6952-ad9c-48fa-ac1f-d224095fc7cc', 'Category 1', 1, now(), NULL, NULL),
       ('f5f813a2-dde2-4a2b-8722-7f19ffdd9747', 'Category 2', 1, now(), NULL, NULL),
       ('7e3cd4d7-03f6-4e98-b803-3b5721bcb1db', 'Category 3', 1, now(), NULL, NULL);

-- Seeder script for public.faqs table
INSERT INTO public.faqs (uuid, faq_category_id, question, answer, status, created_at, updated_at, deleted_at)
VALUES ('64d7e312-5702-4664-b79f-7aeb28a71e4b', 1, 'Question 1', 'Answer to Question 1', 1, now(), NULL, NULL),
       ('a3d8acd4-e968-4d1f-b60b-e5bc4211e743', 2, 'Question 2', 'Answer to Question 2', 1, now(), NULL, NULL),
       ('0c0a66c6-eee4-4a32-a416-4e5f83da0a67', 1, 'Question 3', 'Answer to Question 3', 1, now(), NULL, NULL);