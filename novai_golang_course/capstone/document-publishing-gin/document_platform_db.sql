

CREATE TABLE public.documents (
    id integer NOT NULL,
    title character varying(100) NOT NULL,
    description text,
    author_id integer,
    price numeric(10,2) NOT NULL,
    file_path character varying(255) NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.documents OWNER TO postgres;

--
-- Name: documents_id_seq; Type: SEQUENCE; Schema: public; Owner: novanovriansyah
--

CREATE SEQUENCE public.documents_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.documents_id_seq OWNER TO postgres;

--
-- Name: documents_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: novanovriansyah
--

ALTER SEQUENCE public.documents_id_seq OWNED BY public.documents.id;


--
-- Name: purchases; Type: TABLE; Schema: public; Owner: novanovriansyah
--

CREATE TABLE public.purchases (
    id integer NOT NULL,
    user_id integer,
    document_id integer,
    purchased_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.purchases OWNER TO postgres;

--
-- Name: purchases_id_seq; Type: SEQUENCE; Schema: public; Owner: novanovriansyah
--

CREATE SEQUENCE public.purchases_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.purchases_id_seq OWNER TO postgres;

--
-- Name: purchases_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: novanovriansyah
--

ALTER SEQUENCE public.purchases_id_seq OWNED BY public.purchases.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: novanovriansyah
--

CREATE TABLE public.users (
    id integer NOT NULL,
    username character varying(50) NOT NULL,
    email character varying(100) NOT NULL,
    password_hash character varying(100) NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: novanovriansyah
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_id_seq OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: novanovriansyah
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: documents id; Type: DEFAULT; Schema: public; Owner: novanovriansyah
--

ALTER TABLE ONLY public.documents ALTER COLUMN id SET DEFAULT nextval('public.documents_id_seq'::regclass);


--
-- Name: purchases id; Type: DEFAULT; Schema: public; Owner: novanovriansyah
--

ALTER TABLE ONLY public.purchases ALTER COLUMN id SET DEFAULT nextval('public.purchases_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: novanovriansyah
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);






SELECT pg_catalog.setval('public.documents_id_seq', 1, false);


--
-- Name: purchases_id_seq; Type: SEQUENCE SET; Schema: public; Owner: novanovriansyah
--

SELECT pg_catalog.setval('public.purchases_id_seq', 1, false);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: novanovriansyah
--

SELECT pg_catalog.setval('public.users_id_seq', 1, false);


--
-- Name: documents documents_pkey; Type: CONSTRAINT; Schema: public; Owner: novanovriansyah
--

ALTER TABLE ONLY public.documents
    ADD CONSTRAINT documents_pkey PRIMARY KEY (id);


--
-- Name: purchases purchases_pkey; Type: CONSTRAINT; Schema: public; Owner: novanovriansyah
--

ALTER TABLE ONLY public.purchases
    ADD CONSTRAINT purchases_pkey PRIMARY KEY (id);


--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: novanovriansyah
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: novanovriansyah
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: users users_username_key; Type: CONSTRAINT; Schema: public; Owner: novanovriansyah
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_username_key UNIQUE (username);


--
-- Name: documents documents_author_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: novanovriansyah
--

ALTER TABLE ONLY public.documents
    ADD CONSTRAINT documents_author_id_fkey FOREIGN KEY (author_id) REFERENCES public.users(id);


--
-- Name: purchases purchases_document_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: novanovriansyah
--

ALTER TABLE ONLY public.purchases
    ADD CONSTRAINT purchases_document_id_fkey FOREIGN KEY (document_id) REFERENCES public.documents(id);


--
-- Name: purchases purchases_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: novanovriansyah
--

ALTER TABLE ONLY public.purchases
    ADD CONSTRAINT purchases_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);

