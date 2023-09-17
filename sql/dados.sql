USE devbook;

INSERT INTO usuarios(nome, nick, email, senha)
values
("Usuário 1", "usuario_1", "usuario1@gmail.com", "$2a$10$bQN6OQEWug3rQiXDSAJig.PD60WInM9U67OR08eYp1tSEWAI3VzNC"),
("Usuário 2", "usuario_2", "usuario2@gmail.com", "$2a$10$bQN6OQEWug3rQiXDSAJig.PD60WInM9U67OR08eYp1tSEWAI3VzNC"),
("Usuário 3", "usuario_3", "usuario3@gmail.com", "$2a$10$bQN6OQEWug3rQiXDSAJig.PD60WInM9U67OR08eYp1tSEWAI3VzNC"),
("Usuário 4", "usuario_4", "usuario4@gmail.com", "$2a$10$bQN6OQEWug3rQiXDSAJig.PD60WInM9U67OR08eYp1tSEWAI3VzNC");

INSERT INTO seguidores(usuario_id, seguidor_id)
values
(1, 2),
(3, 1),
(1, 3);

# $2a$10$bQN6OQEWug3rQiXDSAJig.PD60WInM9U67OR08eYp1tSEWAI3VzNC
# manumanu