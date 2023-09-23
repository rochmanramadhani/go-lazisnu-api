create table if not exists public.region_provinces
(
    id   smallserial primary key not null,
    name text                    not null
);

alter table public.region_provinces
    owner to "go-lazisnu-user";

create table if not exists public.region_regencies
(
    id          smallserial primary key not null,
    province_id smallint                not null,
    name        text                    not null,
    constraint region_regencies_province_id_fk
        foreign key (province_id) references public.region_provinces (id)
);

alter table public.region_regencies
    owner to "go-lazisnu-user";

create table if not exists public.region_districts
(
    id         smallserial primary key not null,
    regency_id smallint                not null,
    name       text                    not null,
    constraint region_districts_regency_id_fk
        foreign key (regency_id) references public.region_regencies (id)
);

alter table public.region_districts
    owner to "go-lazisnu-user";

create table if not exists public.region_villages
(
    id          smallserial primary key not null,
    district_id smallint                not null,
    name        text                    not null,
    constraint region_villages_district_id_fk
        foreign key (district_id) references public.region_districts (id)
);

alter table public.region_villages
    owner to "go-lazisnu-user";
