-- Seeder script for public.donation_types table
INSERT INTO public.donation_types (uuid, name, status, created_at, updated_at, deleted_at)
VALUES ('12a7cfb4-19d3-4fe5-bd84-856c5f7ed6c1', 'Zakat', 1, now(), NULL, NULL),
       ('5efea85a-8f2e-4c7b-9f95-7f7e37f781b3', 'Infaq', 1, now(), NULL, NULL),
       ('d28dbd38-4823-4b88-a46d-91f5378a4d9e', 'Shadaqah', 1, now(), NULL, NULL),
       ('76a96a9e-6c75-4e44-aa02-67fc625485cd', 'Wakaf', 1, now(), NULL, NULL);
