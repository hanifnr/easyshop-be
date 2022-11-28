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
ALTER TABLE public.cust ALTER COLUMN isactive SET DEFAULT TRUE;
ALTER TABLE public.cust ADD COLUMN isdelete BOOLEAN DEFAULT FALSE;
ALTER TABLE public.cust RENAME COLUMN isdelete TO is_delete;
ALTER TABLE public.cust DROP COLUMN passport;
ALTER TABLE public.cust DROP COLUMN status;
ALTER TABLE public.cust RENAME COLUMN isactive TO is_active;

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
ALTER TABLE public.shop ADD COLUMN scrape_image VARCHAR;
ALTER TABLE public.shop ADD COLUMN isactive BOOLEAN;
ALTER TABLE public.shop RENAME COLUMN scrap_item_name TO scrape_item_name;
ALTER TABLE public.shop RENAME COLUMN scrap_item_price TO scrape_item_price;
ALTER TABLE public.shop RENAME COLUMN isactive TO is_active;

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
ALTER TABLE public.order ADD COLUMN isdelete BOOLEAN DEFAULT FALSE;
ALTER TABLE public.order RENAME COLUMN isdelete TO is_delete;
ALTER TABLE public.order ADD COLUMN passport VARCHAR;
ALTER TABLE public.order ADD COLUMN addr_id BIGINT  ;
ALTER TABLE public.order ADD COLUMN arrival_date TIMESTAMP;
ALTER TABLE public.order RENAME COLUMN status TO status_code;
ALTER TABLE public.order ALTER COLUMN status_code TYPE VARCHAR(2);

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
ALTER TABLE public.orderd ADD COLUMN image VARCHAR;
ALTER TABLE public.orderd ADD COLUMN url VARCHAR;
ALTER TABLE public.orderd ADD COLUMN imported BOOLEAN DEFAULT FALSE;
ALTER TABLE public.orderd RENAME COLUMN productid TO product_id;

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
ALTER TABLE public.purc ADD COLUMN isdelete BOOLEAN DEFAULT FALSE;
ALTER TABLE public.purc RENAME COLUMN isdelete TO is_delete;
ALTER TABLE public.purc DROP COLUMN status;
ALTER TABLE public.purc ADD COLUMN is_active BOOLEAN DEFAULT TRUE;

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
ALTER TABLE public.purcd ADD COLUMN productid VARCHAR; 
ALTER TABLE public.purcd ADD COLUMN name VARCHAR;
ALTER TABLE public.purcd RENAME COLUMN productid TO product_id;

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
ALTER TABLE public.wh ADD COLUMN isdelete BOOLEAN DEFAULT FALSE;
ALTER TABLE public.wh RENAME COLUMN isdelete TO is_delete;
ALTER TABLE public.wh DROP COLUMN status;
ALTER TABLE public.wh ADD COLUMN is_active BOOLEAN DEFAULT TRUE;

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
ALTER TABLE public.whd RENAME COLUMN qtywh TO qty;

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
ALTER TABLE public.passport RENAME COLUMN isdelete TO is_delete;

CREATE SEQUENCE public.addr_id_seq;
CREATE TABLE public.addr(
  id BIGINT NOT NULL DEFAULT nextval('public.addr_id_seq'),
  cust_id BIGINT,
  name VARCHAR,
  email VARCHAR,
  phone_number VARCHAR,
  zip_code VARCHAR,
  country_code VARCHAR,
  province VARCHAR,
  city VARCHAR,
  full_address VARCHAR,
  created_at TIMESTAMP, 
  updated_at TIMESTAMP,
  isdelete BOOLEAN DEFAULT FALSE,
  CONSTRAINT addr_pk PRIMARY KEY (id)
);
ALTER TABLE public.addr
ADD CONSTRAINT addr_rel_cust_fk
FOREIGN KEY (cust_id) REFERENCES public.cust (id) ON DELETE RESTRICT ON UPDATE CASCADE
NOT DEFERRABLE;
ALTER TABLE public.addr RENAME COLUMN isdelete TO is_delete;
ALTER TABLE public.order
ADD CONSTRAINT order_rel_addr_fk
FOREIGN KEY (addr_id) REFERENCES public.addr (id) ON DELETE RESTRICT ON UPDATE CASCADE
NOT DEFERRABLE;
ALTER TABLE public.addr DROP COLUMN email;

CREATE SEQUENCE public.status_id_seq;
CREATE TABLE public.status(
  id BIGINT NOT NULL DEFAULT nextval('public.status_id_seq'),
  code VARCHAR NOT NULL,
  name VARCHAR NOT NULL,
  CONSTRAINT status_pk PRIMARY KEY (id)
);
CREATE UNIQUE INDEX status_code_index ON public.status (code);
ALTER TABLE public.order
ADD CONSTRAINT order_rel_status_fk
FOREIGN KEY (status_code) REFERENCES public.status (code) ON DELETE RESTRICT ON UPDATE CASCADE
NOT DEFERRABLE;

CREATE SEQUENCE public.order_log_id_seq;
CREATE TABLE public.order_log(
  id BIGINT NOT NULL DEFAULT nextval('public.order_log_id_seq'),
  CONSTRAINT order_log_pk PRIMARY KEY (id),
  order_id BIGINT NOT NULL,
  status_code VARCHAR(2) NOT NULL,
  note TEXT,
  date TIMESTAMP
);
ALTER TABLE public.order_log
ADD CONSTRAINT order_log_rel_order_fk
FOREIGN KEY (order_id) REFERENCES public.order (id) ON DELETE RESTRICT ON UPDATE CASCADE
NOT DEFERRABLE;
ALTER TABLE public.order_log
ADD CONSTRAINT order_log_rel_status_fk
FOREIGN KEY (status_code) REFERENCES public.status (code) ON DELETE RESTRICT ON UPDATE CASCADE
NOT DEFERRABLE;