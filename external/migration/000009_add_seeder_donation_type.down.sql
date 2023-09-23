-- Delete the donation_types that were inserted
DELETE
FROM public.donation_types;

-- Reset sequences to their original values
-- Replace 'donation_types_id_seq' with the correct sequence names if necessary
-- This assumes that the sequences were automatically created with bigserial columns
-- If you manually created the sequences or have different names, adjust accordingly
SELECT setval('public.donation_types_id_seq', (SELECT MAX(id) FROM public.donation_types));
