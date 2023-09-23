-- Drop foreign key constraints on faq_questions
ALTER TABLE public.faq
    DROP CONSTRAINT IF EXISTS faq_questions_faq_category_id_faq_categories_id_fk;

-- Drop the tables in reverse order to maintain referential integrity
DROP TABLE IF EXISTS public.faq_categories;
DROP TABLE IF EXISTS public.faqs;
