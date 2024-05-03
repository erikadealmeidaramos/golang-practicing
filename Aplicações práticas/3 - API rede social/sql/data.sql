INSERT INTO users (name, nick, email, password) VALUES 
('Erika', 'erikinha' 'erika@gmail.com', '$2a$10$ABpoHWyjk.Jr0vGtGE3MMusGf98OAsbq95g2Ni0b86JYD/LnUv4pu'),
('João', 'joaozinho', 'joao@gmail.com', '$2a$10$ABpoHWyjk.Jr0vGtGE3MMusGf98OAsbq95g2Ni0b86JYD/LnUv4pu'),
('Maria', 'mariazinha', 'maria@gmail.com', '$2a$10$ABpoHWyjk.Jr0vGtGE3MMusGf98OAsbq95g2Ni0b86JYD/LnUv4pu');

INSERT INTO followers (user_id, follower_id) VALUES 
(1, 2),
(1, 3),
(2, 1),
(3, 1);