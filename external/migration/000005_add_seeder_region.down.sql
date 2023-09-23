-- Delete data from region_provinces table
DELETE
FROM public.region_provinces;

-- Delete data from region_regencies table
DELETE
FROM public.region_regencies;

-- Reset sequences to their original values
-- Replace 'region_provinces_id_seq' and 'region_regencies_id_seq' with the correct sequence names if necessary
-- This assumes that the sequences were automatically created with bigserial columns
-- If you manually created the sequences or have different names, adjust accordingly
SELECT setval('public.region_provinces_id_seq', (SELECT MAX(id) FROM public.region_provinces));
SELECT setval('public.region_regencies_id_seq', (SELECT MAX(id) FROM public.region_regencies));
