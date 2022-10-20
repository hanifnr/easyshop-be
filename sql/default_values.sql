TRUNCATE shop_category CASCADE;
SELECT setval('public.shop_category_id_seq', 1, false);
SELECT setval('public.shop_id_seq', 1, false);


INSERT INTO shop_category(name, created_at, updated_at) 
VALUES 
('Drugstore', NOW(), NOW()),
('Department store', NOW(), NOW()),
('Electronic', NOW(), NOW()),
('Apparel', NOW(), NOW()),
('Hobby Shop', NOW(), NOW());

INSERT INTO shop(name, image_path, baseurl, keyurl, shop_category_id, created_at, updated_at)
VALUES
(
    'Matsumoto Kiyoshi',
    'logo_matsukiyo.png',
    'https://www.matsukiyo.co.jp/store/online',
    'https://www.matsukiyo.co.jp/store/online/p/',
    1,
    NOW(),
    NOW()
  ),
  (
    'Sun Drug',
    'logo_sundrug.png',
    'https://ec.sundrug.co.jp/catalog/searchresult/',
    'https://ec.sundrug.co.jp/catalog/category/0/',
    1,
    NOW(),
    NOW()
  ),
  (
    'Welcia',
    'logo_welcia.png',
    'https://www.e-welcia.com/',
    'https://www.e-welcia.com/products/detail.php?product_id=',
    1,
    NOW(),
    NOW()
  ),
  (
    'Tokyu Hands',
    'logo_tokyuhands.png',
    'https://hands.net/cate/',
    'https://hands.net/goods/',
    2,
    NOW(),
    NOW()
  ),
  (
    'Loft',
    'logo_loft.png',
    'https://www.loft.co.jp/store/',
    'https://www.loft.co.jp/store/g/',
    2,
    NOW(),
    NOW()
  ),
  (
    'Odakyu',
    'logo_odakyu.png',
    'https://shop.odakyu-dept.co.jp/ec/top',
    'https://shop.odakyu-dept.co.jp/product/detail/s/',
    2,
    NOW(),
    NOW()
  ),
  (
    'Takashimaya',
    'logo_takashimaya.png',
    'https://www.takashimaya.co.jp/shopping/',
    'https://www.takashimaya.co.jp/shopping/product.html?p_cd=',
    2,
    NOW(),
    NOW()
  ),
  (
    'Mitsukoshi Isetan',
    'logo_mitsukoshi.png',
    'https://www.mistore.jp/shopping',
    'https://www.mistore.jp/shopping/product/',
    2,
    NOW(),
    NOW()
  ),
  (
    'Bic Camera',
    'logo_biccamera.png',
    'https://www.biccamera.com/bc/main/',
    'https://www.biccamera.com/bc/item/',
    3,
    NOW(),
    NOW()
  ),
  (
    'Yamada Denki',
    'logo_yamada.png',
    'https://www.yamada-denkiweb.com/',
    'https://www.yamada-denkiweb.com/',
    3,
    NOW(),
    NOW()
  ),
  (
    'Yodobashi Camera',
    'logo_yodobashi.png',
    'https://www.yodobashi.com/',
    'https://www.yodobashi.com/product/',
    3,
    NOW(),
    NOW()
  ),
  (
    'ABC Mart',
    'logo_abcmart.png',
    'https://www.abc-mart.net/shop/',
    'https://www.abc-mart.net/shop/g/',
    4,
    NOW(),
    NOW()
  ),
  (
    'Muji',
    'logo_muji.png',
    'https://www.muji.com/jp/ja/store',
    'https://www.muji.com/jp/ja/store/cmdty/detail/',
    4,
    NOW(),
    NOW()
  ),
  (
    'Uniqlo',
    'logo_uniqlo.png',
    'https://www.uniqlo.com/jp/ja/',
    'https://www.uniqlo.com/jp/ja/products/',
    4,
    NOW(),
    NOW()
  ),
  (
    'Onitsuka Tiger',
    'logo_onitsuka.png',
    'https://www.onitsukatigermagazine.com/store/',
    'https://www.onitsukatigermagazine.com/store/p/',
    4,
    NOW(),
    NOW()
  ),
  (
    'Yoshida Kaban',
    'logo_yoshidakaban.png',
    'https://www.yoshidakaban.com/',
    'https://www.yoshidakaban.com/product/',
    4,
    NOW(),
    NOW()
  ),
  (
    'Sanrio',
    'logo_sanrio.png',
    'https://shop.sanrio.co.jp/',
    'https://shop.sanrio.co.jp/item/detail/',
    5,
    NOW(),
    NOW()
  ),
  (
    'Bandai',
    'logo_bandai.png',
    'https://p-bandai.jp/',
    'https://p-bandai.jp/item/',
    5,
    NOW(),
    NOW()
  );