-- Drop foreign key constraint on user_profiles table
ALTER TABLE public.user_profiles
    DROP CONSTRAINT IF EXISTS user_profiles_user_id_fk;

-- Drop the tables in reverse order to maintain referential integrity
DROP TABLE IF EXISTS public.user_profiles;
DROP TABLE IF EXISTS public.roles;
DROP TABLE IF EXISTS public.users;
