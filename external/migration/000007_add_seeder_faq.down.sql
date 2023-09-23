-- Delete data from faq_categories table
DELETE
FROM public.faq_categories;

-- Delete data from faqs table
DELETE
FROM public.faqs;

-- Reset sequences to their original values
-- Replace 'faq_categories_id_seq' and 'faqs_id_seq' with the correct sequence names if necessary
-- This assumes that the sequences were automatically created with smallserial columns
-- If you manually created the sequences or have different names, adjust accordingly
SELECT setval('public.faq_categories_id_seq', (SELECT MAX(id) FROM public.faq_categories));
SELECT setval('public.faqs_id_seq', (SELECT MAX(id) FROM public.faqs));
