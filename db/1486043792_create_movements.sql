CREATE TABLE movements (
    id integer NOT NULL PRIMARY KEY,
    details json,
    scheme_id integer,
    product_id integer,
    quantity numeric,
    store_from_id integer,
    store_to_id integer,
    user_from_id integer,
    user_to_id integer,
    ended boolean,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
    );