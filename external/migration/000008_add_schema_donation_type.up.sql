create table if not exists public.donation_types
(
    id         smallserial primary key     not null,
    uuid       uuid                        not null,
    name       text unique                 not null,
    status     smallint    default 1       not null,
    created_at timestamptz default (now()) not null,
    updated_at timestamptz,
    deleted_at timestamptz
);

alter table public.faqs
    owner to "go-lazisnu-user";
