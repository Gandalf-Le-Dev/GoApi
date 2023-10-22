-- schema.sql

-- 1. Users Table
CREATE TABLE users (
  user_id SERIAL PRIMARY KEY,
  username VARCHAR(50) NOT NULL UNIQUE,
  email VARCHAR(100) NOT NULL UNIQUE,
  password_hash VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Insert default users
INSERT INTO users (username, email, password_hash) VALUES 
('john_doe', 'john@example.com', 'hashed_password1'),
('jane_doe', 'jane@example.com', 'hashed_password2');

-- 2. Categories Table
CREATE TABLE categories (
  category_id SERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL UNIQUE
);

-- Insert default categories
INSERT INTO categories (name) VALUES 
('Technology'),
('Lifestyle'),
('Travel');

-- 3. Posts Table
CREATE TABLE posts (
  post_id SERIAL PRIMARY KEY,
  user_id INT NOT NULL,
  category_id INT,
  title VARCHAR(255) NOT NULL,
  content TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES users(user_id),
  FOREIGN KEY (category_id) REFERENCES categories(category_id)
);

-- Insert default posts
INSERT INTO posts (user_id, category_id, title, content) VALUES 
(1, 1, 'Welcome to My Tech Blog', 'This is the first post on my tech blog. Stay tuned for more updates!'),
(2, 3, 'My Travel Adventures', 'Join me as I explore the world and share my adventures!');

-- 4. Comments Table
CREATE TABLE comments (
  comment_id SERIAL PRIMARY KEY,
  post_id INT NOT NULL,
  user_id INT NOT NULL,
  content TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (post_id) REFERENCES posts(post_id) ON DELETE CASCADE,
  FOREIGN KEY (user_id) REFERENCES users(user_id)
);

-- Insert default comments
INSERT INTO comments (post_id, user_id, content) VALUES 
(1, 2, 'Great post! Looking forward to more tech updates.'),
(2, 1, 'Your travel adventures sound amazing!');
