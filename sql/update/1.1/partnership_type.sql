INSERT INTO partnership_type(code,name) VALUES 
('TRAVEL', 'Travel Agent'),
('HOTEL', 'Hotel');
UPDATE partnership_type SET code = 'TRVLA' WHERE code = 'TRAVEL';
INSERT INTO partnership_type(code,name) VALUES 
('TRVLG', 'Travel Guide');