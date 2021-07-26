create table labels
(
    address text not null,
    label   text not null
);

create unique index labels_address_idx on labels (address);
