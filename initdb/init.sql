CREATE TABLE roles (
	created_at int8 NULL,
	created_by varchar(255) NULL,
	modified_at int8 NULL,
	modified_by varchar(255) NULL,
	deleted_at int8 NULL,
	deleted_by varchar(255) NULL,
	role_id varchar(255) NOT NULL,
	role_name varchar(255) NOT NULL,
	CONSTRAINT roles_pk PRIMARY KEY (role_id),
	CONSTRAINT unique_role_name UNIQUE (role_name)
);

CREATE TABLE navigation (
	created_at int8 NULL,
	created_by varchar(255) NULL,
	modified_at int8 NULL,
	modified_by varchar(255) NULL,
	deleted_at int8 NULL,
	deleted_by varchar(255) NULL,
	navigation_id varchar(255) NOT NULL,
	navigation_name varchar(255) NOT NULL,
	parent_navigation_id varchar(255) NULL,
	sort_order int4 NOT NULL,
	url_path varchar(255) NULL,
	is_active bool DEFAULT true NULL,
	CONSTRAINT navigation_pk PRIMARY KEY (navigation_id)
);

CREATE TABLE roles_navigation (
	role_navigation_id serial4 NOT NULL,
	role_id varchar(255) NOT NULL,
	navigation_id varchar(255) NOT NULL,
	allow_read bool NOT NULL,
	allow_create bool NOT NULL,
	allow_update bool NOT NULL,
	allow_delete bool NOT NULL,
	allow_approval bool NOT NULL,
	CONSTRAINT roles_navigation_pk PRIMARY KEY (role_navigation_id)
);


-- public.roles_navigation foreign keys

ALTER TABLE public.roles_navigation ADD CONSTRAINT roles_navigation_navigation_id_fk FOREIGN KEY (navigation_id) REFERENCES navigation(navigation_id);
ALTER TABLE public.roles_navigation ADD CONSTRAINT roles_navigation_role_id_fk FOREIGN KEY (role_id) REFERENCES roles(role_id);