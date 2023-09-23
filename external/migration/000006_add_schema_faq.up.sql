create table if not exists public.faq_categories
(
    id         smallserial primary key     not null,
    uuid       uuid                        not null,
    name       text unique                 not null,
    status     smallint    default 1       not null,
    created_at timestamptz default (now()) not null,
    updated_at timestamptz,
    deleted_at timestamptz
);

alter table public.region_provinces
    owner to "go-lazisnu-user";

create table if not exists public.faqs
(
    id              smallserial primary key     not null,
    uuid            uuid                        not null,
    faq_category_id int                         not null,
    question        text                        not null,
    answer          text                        not null,
    status          smallint    default 1       not null,
    created_at      timestamptz default (now()) not null,
    updated_at      timestamptz,
    deleted_at      timestamptz,
    constraint faq_questions_faq_category_id_faq_categories_id_fk
        foreign key (faq_category_id) references public.faq_categories (id)
);

alter table public.faqs
    owner to "go-lazisnu-user";
