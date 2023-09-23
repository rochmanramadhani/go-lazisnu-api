-- Delete data from user_profiles table
DELETE
FROM public.user_profiles;

-- Delete data from users table
DELETE
FROM public.users;

-- Reset sequences to their original values
-- Replace 'users_id_seq' and 'user_profiles_id_seq' with the correct sequence names if necessary
-- This assumes that the sequences were automatically created with bigserial columns
-- If you manually created the sequences or have different names, adjust accordingly
SELECT setval('public.users_id_seq', (SELECT MAX(id) FROM public.users));
SELECT setval('public.user_profiles_id_seq', (SELECT MAX(id) FROM public.user_profiles));
