CREATE TABLE videos (
    video_id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT,
    title VARCHAR(255),
    file_path VARCHAR(255),
    upload_time DATETIME,
    FOREIGN KEY (user_id) REFERENCES users(user_id)
);