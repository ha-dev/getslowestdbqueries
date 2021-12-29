--
-- PostgreSQL database dump
--

-- Dumped from database version 11.8
-- Dumped by pg_dump version 11.8

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: querys; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.querys (
    id integer NOT NULL,
    query text,
    time_to_run character varying(255),
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone
);


ALTER TABLE public.querys OWNER TO postgres;

--
-- Name: querys_autoincreament; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.querys_autoincreament
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    MAXVALUE 9000000000
    CACHE 1;


ALTER TABLE public.querys_autoincreament OWNER TO postgres;

--
-- Name: querys_autoincreament; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.querys_autoincreament OWNED BY public.querys.id;


--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    MAXVALUE 99999999999999
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.querys.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id bigint DEFAULT nextval('public.users_id_seq'::regclass) NOT NULL,
    username character varying,
    password character varying,
    name character varying
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: querys id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.querys ALTER COLUMN id SET DEFAULT nextval('public.querys_autoincreament'::regclass);


--
-- Data for Name: querys; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.querys (id, query, time_to_run, created_at, updated_at, deleted_at) FROM stdin;
1	select *	1	\N	\N	\N
2	select * from users	1	2021-12-29 01:09:02.427231	2021-12-29 01:09:02.427231	\N
3	select * from users	1	2021-12-29 01:09:04.640742	2021-12-29 01:09:04.640742	\N
4	insert value 	1	\N	\N	\N
5	select name from users	1	\N	\N	\N
6	INSERT INTO users (username, password)\n\t    VALUES('test','1');	2	2021-12-29 14:23:05.439432	2021-12-29 14:23:05.439432	\N
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, username, password, name) FROM stdin;
1	094	$2a$12$oht9qHeDfkFbK1jdGf43sO5oixl9I8DUMyjWO5PUZJf7Tmes9LBsq	Lady Jennie Russel
3	test	1	\N
\.


--
-- Name: querys_autoincreament; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.querys_autoincreament', 6, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 3, true);


--
-- Name: querys querys_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.querys
    ADD CONSTRAINT querys_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

