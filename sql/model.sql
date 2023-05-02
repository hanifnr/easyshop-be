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
ALTER TABLE public.cust ADD COLUMN is_delete BOOLEAN DEFAULT FALSE;
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
ALTER TABLE public.shop ADD COLUMN is_active BOOLEAN;
ALTER TABLE public.shop RENAME COLUMN scrap_item_name TO scrape_item_name;
ALTER TABLE public.shop RENAME COLUMN scrap_item_price TO scrape_item_price;
ALTER TABLE public.shop ADD COLUMN idx BIGINT;

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
ALTER TABLE public.order ADD COLUMN is_delete BOOLEAN DEFAULT FALSE;
ALTER TABLE public.order ADD COLUMN passport VARCHAR;
ALTER TABLE public.order ADD COLUMN addr_id BIGINT  ;
ALTER TABLE public.order ADD COLUMN arrival_date TIMESTAMP;
ALTER TABLE public.order RENAME COLUMN status TO status_code;
ALTER TABLE public.order ALTER COLUMN status_code TYPE VARCHAR(2);
ALTER TABLE public.order ADD COLUMN exchange_rate NUMERIC(19,4) DEFAULT 1;
ALTER TABLE public.order ADD COLUMN shipping_cost NUMERIC(19,4);
ALTER TABLE public.order ADD COLUMN grand_total NUMERIC(19,4);
ALTER TABLE public.order ADD COLUMN voucher_id BIGINT;
ALTER TABLE public.order ADD COLUMN disc VARCHAR;
ALTER TABLE public.order ADD COLUMN disc_amount NUMERIC(19,4);
ALTER TABLE public.order ADD COLUMN tax_amount NUMERIC(19,4);
ALTER TABLE public.order ALTER COLUMN tax_amount TYPE VARCHAR;
ALTER TABLE public.order
ADD CONSTRAINT order_rel_voucher_fk
FOREIGN KEY (voucher_id) REFERENCES public.voucher (id) ON DELETE RESTRICT ON UPDATE CASCADE
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
ALTER TABLE public.orderd ADD COLUMN image VARCHAR;
ALTER TABLE public.orderd ADD COLUMN url VARCHAR;
ALTER TABLE public.orderd ADD COLUMN imported BOOLEAN DEFAULT FALSE;
ALTER TABLE public.orderd RENAME COLUMN productid TO product_id;
ALTER TABLE public.orderd ADD COLUMN arrived BOOLEAN DEFAULT FALSE;
ALTER TABLE public.orderd ADD COLUMN note TEXT;

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
ALTER TABLE public.purc ADD COLUMN is_delete BOOLEAN DEFAULT FALSE;
ALTER TABLE public.purc DROP COLUMN status;
ALTER TABLE public.purc ADD COLUMN is_active BOOLEAN DEFAULT TRUE;
ALTER TABLE public.purc ADD COLUMN refno VARCHAR;
ALTER TABLE purc ADD COLUMN imported BOOLEAN DEFAULT FALSE;

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
ALTER TABLE public.purcd ADD COLUMN imported BOOLEAN DEFAULT FALSE;

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
ALTER TABLE public.wh ADD COLUMN is_delete BOOLEAN DEFAULT FALSE;
ALTER TABLE public.wh DROP COLUMN status;
ALTER TABLE public.wh ADD COLUMN is_active BOOLEAN DEFAULT TRUE;
ALTER TABLE public.wh ADD COLUMN status_code VARCHAR;

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
ALTER TABLE public.passport DROP COLUMN name;
ALTER TABLE public.passport DROP COLUMN country_code;
ALTER TABLE public.passport DROP COLUMN issue_date;
ALTER TABLE public.passport DROP COLUMN exp_date;
ALTER TABLE public.passport ADD COLUMN status_residence VARCHAR;

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
ALTER TABLE public.order ADD COLUMN taxed BOOLEAN DEFAULT FALSE;

CREATE SEQUENCE public.status_id_seq;
CREATE TABLE public.status(
  id BIGINT NOT NULL DEFAULT nextval('public.status_id_seq'),
  code VARCHAR NOT NULL,
  name VARCHAR NOT NULL,
  CONSTRAINT status_pk PRIMARY KEY (id)
);
CREATE UNIQUE INDEX status_code_index ON public.status (code);
ALTER TABLE public.status ADD COLUMN idx BIGINT;
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

CREATE SEQUENCE public.email_verif_id_seq;
CREATE TABLE public.email_verif(
  id BIGINT NOT NULL DEFAULT nextval('public.email_verif_id_seq'),
	email         VARCHAR NOT NULL,
	verif_code    VARCHAR,
	verified      BOOLEAN,
	verified_at   TIMESTAMP,
	generated_at  TIMESTAMP,
  CONSTRAINT email_verif_pk PRIMARY KEY (id)
);
ALTER TABLE public.email_verif ADD COLUMN auth_code VARCHAR;
ALTER TABLE public.email_verif ADD COLUMN wait_time int;
ALTER TABLE public.email_verif ADD COLUMN type VARCHAR;

CREATE SEQUENCE public.product_id_seq;
CREATE TABLE public.product(
  id          BIGINT NOT NULL DEFAULT nextval('public.product_id_seq'),
  code        VARCHAR NOT NULL,
  name        VARCHAR NOT NULL,
	image       VARCHAR,
	price       VARCHAR,
	price_tax   VARCHAR,
	size        VARCHAR,
  is_delete BOOLEAN DEFAULT FALSE,
  created_at TIMESTAMP, 
  updated_at TIMESTAMP,
  CONSTRAINT product_pk PRIMARY KEY (id)
);

CREATE SEQUENCE public.firebase_token_id_seq;
CREATE TABLE public.firebase_token(
  id  BIGINT NOT NULL DEFAULT nextval('public.firebase_token_id_seq'),
  uid VARCHAR NOT NULL,
  token VARCHAR NOT NULL,
  type VARCHAR NOT NULL,
  is_delete BOOLEAN DEFAULT FALSE,
  created_at TIMESTAMP, 
  updated_at TIMESTAMP,
  CONSTRAINT token_pk PRIMARY KEY (id)
);
ALTER TABLE public.firebase_token RENAME COLUMN is_deleted TO is_delete;


CREATE SEQUENCE public.partnership_type_id_seq;
CREATE TABLE public.partnership_type(
  id BIGINT NOT NULL DEFAULT nextval('public.partnership_type_id_seq'),
  code VARCHAR(10) NOT NULL,
  name VARCHAR(25) NOT NULL,
  CONSTRAINT partnership_type_pk PRIMARY KEY (id)
);

CREATE SEQUENCE public.partnership_id_seq;
CREATE TABLE public.partnership(
  id BIGINT NOT NULL DEFAULT nextval('public.partnership_id_seq'),
  name VARCHAR NOT NULL,
  partnership_type_id BIGINT,
  social_media VARCHAR,
  phone_number VARCHAR,
  is_delete BOOLEAN DEFAULT FALSE,
  created_at TIMESTAMP, 
  updated_at TIMESTAMP,
  CONSTRAINT partnership_pk PRIMARY KEY (id)
);
ALTER TABLE public.partnership ADD COLUMN email VARCHAR;
ALTER TABLE public.partnership ADD COLUMN approved BOOLEAN;
ALTER TABLE public.partnership ADD COLUMN note TEXT;
ALTER TABLE public.partnership
ADD CONSTRAINT partnership_rel_partnership_type_fk
FOREIGN KEY (partnership_type_id) REFERENCES public.partnership_type (id) ON DELETE RESTRICT ON UPDATE CASCADE
NOT DEFERRABLE;

CREATE SEQUENCE public.voucher_id_seq;
CREATE TABLE public.voucher(
  id BIGINT NOT NULL DEFAULT nextval('public.voucher_id_seq'),
  code VARCHAR NOT NULL,
  amount VARCHAR NOT NULL,
  qty NUMERIC(16,4),
  startdate TIMESTAMP,
  enddate TIMESTAMP,
  partnership_id BIGINT,
  note TEXT,
  is_delete BOOLEAN DEFAULT FALSE,
  created_at TIMESTAMP, 
  updated_at TIMESTAMP,
  CONSTRAINT voucher_pk PRIMARY KEY (id)
);
ALTER TABLE public.voucher
ADD CONSTRAINT voucher_rel_partnership_fk
FOREIGN KEY (partnership_id) REFERENCES public.partnership (id) ON DELETE RESTRICT ON UPDATE CASCADE
NOT DEFERRABLE;
ALTER TABLE public.voucher ADD COLUMN qty_used NUMERIC(16,4);

CREATE SEQUENCE public.voucher_log_id_seq;
CREATE TABLE public.voucher_log(
  id BIGINT NOT NULL DEFAULT nextval('public.voucher_log_id_seq'),
  voucher_id BIGINT NOT NULL,
  cust_id BIGINT NOT NULL,
  redeem_at TIMESTAMP,
  CONSTRAINT voucher_log_pk PRIMARY KEY (id)
);
ALTER TABLE public.voucher_log
ADD CONSTRAINT voucher_log_rel_voucher_fk
FOREIGN KEY (voucher_id) REFERENCES public.voucher (id) ON DELETE RESTRICT ON UPDATE CASCADE
NOT DEFERRABLE;
ALTER TABLE public.voucher_log
ADD CONSTRAINT voucher_log_rel_cust_fk
FOREIGN KEY (cust_id) REFERENCES public.cust (id) ON DELETE RESTRICT ON UPDATE CASCADE
NOT DEFERRABLE;

CREATE SEQUENCE public.req_order_id_seq;
CREATE TABLE public.req_order(
  id BIGINT NOT NULL DEFAULT nextval('public.req_order_id_seq'),
  email VARCHAR NOT NULL,
  created_at TIMESTAMP, 
  CONSTRAINT req_order_pk PRIMARY KEY (id)
);

CREATE SEQUENCE public.req_orderd_id_seq;
CREATE TABLE public.req_orderd(
  req_order_id BIGINT NOT NULL,
  dno INT NOT NULL,
  url VARCHAR NOT NULL,
  approved BOOLEAN,
  approval_note TEXT,
  note TEXT,
  CONSTRAINT req_orderd_pk PRIMARY KEY (req_order_id, dno)
);
ALTER TABLE public.req_orderd
ADD CONSTRAINT req_orderd_rel_req_order_fk
FOREIGN KEY (req_order_id) REFERENCES public.req_order (id) ON DELETE RESTRICT ON UPDATE CASCADE
NOT DEFERRABLE;

-- CREATE SEQUENCE public.grp_id_seq;
-- CREATE TABLE public.grp(
--   id BIGINT NOT NULL DEFAULT nextval('public.grp_id_seq'),
--   code VARCHAR NOT NULL,
--   name VARCHAR NOT NULL,
--   CONSTRAINT grp_pk PRIMARY KEY (id)
-- );

-- CREATE SEQUENCE public.prv_id_seq;
-- CREATE TABLE public.prv(
--   id BIGINT NOT NULL DEFAULT nextval('public.prv_id_seq'),
--   code VARCHAR NOT NULL,
--   name VARCHAR NOT NULL,
--   CONSTRAINT prv_pk PRIMARY KEY (id)
-- );

-- CREATE SEQUENCE public.grpprv_id_seq;
-- CREATE TABLE public.grpprv(
--   id BIGINT NOT NULL DEFAULT nextval('public.grpprv_id_seq'),
--   grp_id BIGINT NOT NULL,
--   prv_id BIGINT NOT NULL,
--   enable BOOLEAN
--   CONSTRAINT grpprv_pk PRIMARY KEY (id)
-- );

-- CREATE SEQUENCE public.grpprv_id_seq;
-- CREATE TABLE public.grpprv(
-- id BIGINT NOT NULL DEFAULT nextval('public.grpprv_id_seq'),
-- CONSTRAINT grpprv_pk PRIMARY KEY (id)
-- );

-- CREATE SEQUENCE public.user_id_seq;
-- CREATE TABLE public.user(
--   id BIGINT NOT NULL DEFAULT nextval('public.user_id_seq'),
--   name VARCHAR NOT NULL,
--   password VARCHAR NOT NULL,
--   email VARCHAR NOT NULL,
--   created_at TIMESTAMP, 
--   updated_at TIMESTAMP,
--   isdelete BOOLEAN DEFAULT FALSE,
--   CONSTRAINT user_pk PRIMARY KEY (id)
-- );