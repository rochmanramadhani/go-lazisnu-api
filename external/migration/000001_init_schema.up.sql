create table if not exists public.users
(
    id                     bigserial primary key       not null,
    uuid                   uuid                        not null,
    company_id             bigint                      not null,
    role_id                smallint                    not null,
    name                   text                        not null,
    email                  text unique                 not null,
    password_hash          text                        not null,
    last_ip_address        text                        not null,
    last_ip_address_access timestamptz                 not null,
    is_contact_person      boolean     default false   not null,
    status                 smallint    default 1       not null,
    created_at             timestamptz default (now()) not null,
    updated_at             timestamptz,
    deleted_at             timestamptz
);

alter table public.users
    owner to "go-lazisnu-user";

create table if not exists public.user_profiles
(
    id         bigserial primary key       not null,
    user_id    bigint                      not null,
    address    text                        not null,
    phone      text                        not null,
    file_path  text                        not null,
    created_at timestamptz default (now()) not null,
    updated_at timestamptz,
    deleted_at timestamptz,
    constraint user_profiles_user_id_fk
        foreign key (user_id) references public.users (id)
);

alter table public.user_profiles
    owner to "go-lazisnu-user";

create table if not exists public.roles
(
    id          smallserial primary key     not null,
    uuid        uuid                        not null,
    name        text unique                 not null,
    description text                        not null,
    status      smallint    default 1       not null,
    created_at  timestamptz default (now()) not null,
    updated_at  timestamptz,
    deleted_at  timestamptz
);

alter table public.roles
    owner to "go-lazisnu-user";
