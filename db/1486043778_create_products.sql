CREATE TABLE products (
    id integer NOT NULL PRIMARY KEY, 
    scheme_id integer, 
    name character varying, 
    details json, 
    created_at timestamp without time zone, 
    updated_at timestamp without time zone
    );