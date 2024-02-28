CREATE TABLE IF NOT EXISTS progress (
    user_id INT PRIMARY KEY,
    last_lesson VARCHAR(255), 
    last_position VARCHAR(255),   

    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
