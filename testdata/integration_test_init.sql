-- Seed the user that GetAuthToken authenticates as. Journals are owner-scoped
-- (journal.user_id REFERENCES usr(id)), so the owner must exist first.
INSERT INTO usr (id, email, password, first_name, last_name, role)
VALUES (1, 'testuser@example.com', 'testpassword', 'Test', 'User', 'user');

-- usr.id is SERIAL; advance the sequence past the explicit id so later sign-ups don't collide.
SELECT setval('usr_id_seq', (SELECT MAX(id) FROM usr));

INSERT INTO journal(title, description, completed, user_id) VALUES ('Psychopomp', 'Japanese Breakfast''s first album', 'true', 1);
INSERT INTO journal(title, description, completed, user_id) VALUES ('Soft Sounds from Another Planet', 'Absolute banger followup', 'false', 1);
INSERT INTO journal(title, description, completed, user_id) VALUES ('Jubilee', 'Here Michelle Zauner asks: what if joy was as complex as grief', 'unknown', 1);
