-- Seed the user that GetAuthToken authenticates as. Journals are owner-scoped
-- (journal.user_id REFERENCES usr(id)), so the owner must exist first.
-- Password is a bcrypt hash of 'testpassword'.
INSERT INTO usr (id, email, password, first_name, last_name, role_id)
VALUES (1, 'testuser@example.com', '$2a$10$yIvuMu8oh3LdVefo.zMEau7qqqsvbT5O.fwzDEAtes9L2Cp3Gf/s6', 'Test', 'User', (SELECT id FROM role WHERE role_name = 'user'));

-- Admin user (password is a bcrypt hash of 'testpassword').
INSERT INTO usr (id, email, password, first_name, last_name, role_id)
VALUES (2, 'adminuser@example.com', '$2a$10$yIvuMu8oh3LdVefo.zMEau7qqqsvbT5O.fwzDEAtes9L2Cp3Gf/s6', 'Admin', 'User', (SELECT id FROM role WHERE role_name = 'admin'));

-- usr.id is SERIAL; advance the sequence past the explicit ids so later sign-ups don't collide.
SELECT setval('usr_id_seq', (SELECT MAX(id) FROM usr));

INSERT INTO journal(title, description, completed, user_id) VALUES ('Psychopomp', 'Japanese Breakfast''s first album', 'true', 1);
INSERT INTO journal(title, description, completed, user_id) VALUES ('Soft Sounds from Another Planet', 'Absolute banger followup', 'false', 1);
INSERT INTO journal(title, description, completed, user_id) VALUES ('Jubilee', 'Here Michelle Zauner asks: what if joy was as complex as grief', 'unknown', 1);
