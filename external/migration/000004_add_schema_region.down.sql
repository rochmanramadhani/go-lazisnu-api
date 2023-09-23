-- Drop foreign key constraints on region_villages, region_districts, and region_regencies tables
ALTER TABLE public.region_villages
    DROP CONSTRAINT IF EXISTS region_villages_district_id_fk;
ALTER TABLE public.region_districts
    DROP CONSTRAINT IF EXISTS region_districts_regency_id_fk;
ALTER TABLE public.region_regencies
    DROP CONSTRAINT IF EXISTS region_regencies_province_id_fk;

-- Drop the tables in reverse order to maintain referential integrity
DROP TABLE IF EXISTS public.region_villages;
DROP TABLE IF EXISTS public.region_districts;
DROP TABLE IF EXISTS public.region_regencies;
DROP TABLE IF EXISTS public.region_provinces;
