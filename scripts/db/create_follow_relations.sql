CREATE TABLE IF NOT EXISTS follow_relations (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT,
    follower_id INT,
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (follower_id) REFERENCES users(user_id)
);