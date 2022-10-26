CREATE SEQUENCE public.cust_id_seq;
CREATE TABLE public.cust(
  id BIGINT NOT NULL DEFAULT nextval('public.cust_id_seq'), 
  name VARCHAR, 
  email VARCHAR, 
  country_code VARCHAR, 
  phone_number VARCHAR, 
  passport VARCHAR, 
  status VARCHAR(1), 
  isactive BOOLEAN DEFAULT TRUE, 
  created_at TIMESTAMP, 
  updated_at TIMESTAMP,
  CONSTRAINT cust_pk PRIMARY KEY (id)
);

CREATE SEQUENCE public.shop_category_id_seq;
CREATE TABLE public.shop_category(
  id BIGINT NOT NULL DEFAULT nextval('public.shop_category_id_seq'), 
  name VARCHAR, 
  created_at TIMESTAMP, 
  updated_at TIMESTAMP, 
  CONSTRAINT shop_category_pk PRIMARY KEY (id)
);

CREATE SEQUENCE public.shop_id_seq;
CREATE TABLE public.shop(
  id BIGINT NOT NULL DEFAULT nextval('public.shop_id_seq'), 
  shop_category_id BIGINT, 
  name VARCHAR, 
  image_path VARCHAR, 
  baseurl VARCHAR, 
  keyurl VARCHAR, 
  scrap_item_name VARCHAR, 
  scrap_item_price VARCHAR, 
  created_at TIMESTAMP, 
  updated_at TIMESTAMP, 
  CONSTRAINT shop_pk PRIMARY KEY (id)
);
ALTER TABLE 
  public.shop 
ADD 
  CONSTRAINT shop_rel_shop_category_fk 
  FOREIGN KEY (shop_category_id) 
  REFERENCES public.shop_category (id) 
  ON DELETE RESTRICT 
  ON UPDATE CASCADE 
  NOT DEFERRABLE;

CREATE SEQUENCE public.order_id_seq;
CREATE TABLE public.order(
  id BIGINT NOT NULL DEFAULT nextval('public.order_id_seq'), 
  trxno VARCHAR, 
  date date, 
  cust_id BIGINT, 
  proof_link VARCHAR, 
  pick_date TIMESTAMP, 
  tracking_number VARCHAR, 
  status VARCHAR(1), 
  total NUMERIC(19, 4) DEFAULT 0, 
  created_at TIMESTAMP, 
  updated_at TIMESTAMP, 
  CONSTRAINT order_pk PRIMARY KEY (id)
);
ALTER TABLE 
  public.order 
ADD 
  CONSTRAINT order_rel_cust_fk 
  FOREIGN KEY (cust_id) 
  REFERENCES public.cust (id) 
  ON DELETE RESTRICT 
  ON UPDATE CASCADE 
  NOT DEFERRABLE;

CREATE TABLE public.orderd(
  order_id BIGINT, 
  dno INT, 
  shop_id BIGINT, 
  productid VARCHAR, 
  name VARCHAR, 
  qty NUMERIC(15, 4) DEFAULT 0, 
  qtypurc NUMERIC(15, 4) DEFAULT 0, 
  qtywh NUMERIC(15, 4) DEFAULT 0, 
  price NUMERIC(19, 4) DEFAULT 0, 
  subtotal NUMERIC(19, 4) DEFAULT 0, 
  CONSTRAINT orderd_pk PRIMARY KEY (order_id, dno)
);
ALTER TABLE 
  public.orderd 
ADD 
  CONSTRAINT orderd_rel_order_fk 
  FOREIGN KEY (order_id) 
  REFERENCES public.order (id) 
  ON DELETE RESTRICT 
  ON UPDATE CASCADE 
  NOT DEFERRABLE;
ALTER TABLE 
  public.orderd 
ADD 
  CONSTRAINT orderd_rel_shop_fk 
  FOREIGN KEY (shop_id) 
  REFERENCES public.shop (id) 
  ON DELETE RESTRICT 
  ON UPDATE CASCADE 
  NOT DEFERRABLE;

CREATE SEQUENCE public.purc_seq_id;
CREATE TABLE public.purc(
  id BIGINT NOT NULL DEFAULT nextval('public.purc_seq_id'), 
  trxno VARCHAR, 
  date date, 
  shop_id BIGINT, 
  status VARCHAR, 
  total NUMERIC(19, 4) DEFAULT 0, 
  created_at TIMESTAMP, 
  updated_at TIMESTAMP, 
  CONSTRAINT purc_pk PRIMARY KEY (id)
);
ALTER TABLE 
  public.purc 
ADD 
  CONSTRAINT purc_rel_shop_fk 
  FOREIGN KEY (shop_id) 
  REFERENCES public.shop (id) 
  ON DELETE RESTRICT 
  ON UPDATE CASCADE 
  NOT DEFERRABLE;

CREATE TABLE public.purcd(
  purc_id BIGINT, 
  dno INT, 
  order_id BIGINT, 
  order_dno INT, 
  qty NUMERIC(15, 4) DEFAULT 0, 
  qtywh NUMERIC(15, 4) DEFAULT 0, 
  price NUMERIC(19, 4) DEFAULT 0, 
  subtotal NUMERIC(19, 4) DEFAULT 0, 
  CONSTRAINT purcd_pk PRIMARY KEY (purc_id, dno)
);
ALTER TABLE 
  public.purcd 
ADD 
  CONSTRAINT purcd_rel_purc_fk 
  FOREIGN KEY (purc_id) 
  REFERENCES public.purc (id) 
  ON DELETE RESTRICT 
  ON UPDATE CASCADE 
  NOT DEFERRABLE;
ALTER TABLE 
  public.purcd 
ADD 
  CONSTRAINT purcd_rel_orderd_fk 
  FOREIGN KEY (order_id, order_dno) 
  REFERENCES public.orderd (order_id, dno) 
  ON DELETE RESTRICT 
  ON UPDATE CASCADE 
  NOT DEFERRABLE;

CREATE SEQUENCE public.wh_id_seq;
CREATE TABLE public.wh(
  id BIGINT NOT NULL DEFAULT nextval('public.wh_id_seq'), 
  trxno VARCHAR, 
  date date, 
  shop_id BIGINT, 
  status VARCHAR, 
  created_at TIMESTAMP, 
  updated_at TIMESTAMP, 
  CONSTRAINT wh_pk PRIMARY KEY (id)
);
ALTER TABLE 
  public.wh 
ADD 
  CONSTRAINT wh_rel_shop_fk 
  FOREIGN KEY (shop_id) 
  REFERENCES public.shop (id) 
  ON DELETE RESTRICT 
  ON UPDATE CASCADE 
  NOT DEFERRABLE;

CREATE TABLE public.whd(
  wh_id BIGINT, 
  dno INT, 
  purc_id BIGINT, 
  purc_dno INT, 
  qtywh NUMERIC(15, 4) DEFAULT 0, 
  CONSTRAINT whd_pk PRIMARY KEY (wh_id, dno)
);
ALTER TABLE 
  public.whd 
ADD 
  CONSTRAINT whd_rel_wh_fk 
  FOREIGN KEY (wh_id) 
  REFERENCES public.wh (id) 
  ON DELETE RESTRICT 
  ON UPDATE CASCADE 
  NOT DEFERRABLE;
ALTER TABLE 
  public.whd 
ADD 
  CONSTRAINT whd_rel_purcd_fk 
  FOREIGN KEY (purc_id, purc_dno) 
  REFERENCES public.purcd (purc_id, dno) 
  ON DELETE RESTRICT 
  ON UPDATE CASCADE 
  NOT DEFERRABLE;

CREATE TABLE public.ncount(
  code VARCHAR,
  name VARCHAR,
  number INTEGER,
  length INTEGER,
CONSTRAINT ncount_pk PRIMARY KEY (code)
);

CREATE SEQUENCE public.passport_id_seq;
CREATE TABLE public.passport(
id BIGINT NOT NULL DEFAULT nextval('public.passport_id_seq'),
  cust_id BIGINT, 
  name VARCHAR,
  country_code VARCHAR,
  number VARCHAR,
  nationality VARCHAR,
  birth_date DATE,
  issue_date DATE,
  exp_date DATE,
  created_at TIMESTAMP, 
  updated_at TIMESTAMP,
  isdelete BOOLEAN DEFAULT FALSE,
CONSTRAINT passport_pk PRIMARY KEY (id)
);
ALTER TABLE public.passport
ADD CONSTRAINT passport_rel_cust_fk
FOREIGN KEY (cust_id) REFERENCES public.cust (id) ON DELETE RESTRICT ON UPDATE CASCADE
NOT DEFERRABLE;
CREATE UNIQUE INDEX passport_cust_id_index ON public.passport (cust_id);