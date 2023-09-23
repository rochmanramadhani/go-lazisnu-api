-- Delete the roles that were inserted
DELETE
FROM public.roles;

-- Reset sequences to their original values
-- Replace 'roles_id_seq' with the correct sequence names if necessary
-- This assumes that the sequences were automatically created with bigserial columns
-- If you manually created the sequences or have different names, adjust accordingly
SELECT setval('public.roles_id_seq', (SELECT MAX(id) FROM public.roles));
