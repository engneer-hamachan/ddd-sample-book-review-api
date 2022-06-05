DROP TABLE IF EXISTS users;
CREATE TABLE IF NOT EXISTS users (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  user_id text,
  name text,
  mail text,
  password text,
  created_at timestamp,
  updated_at timestamp
);

DROP TABLE IF EXISTS reviews;
CREATE TABLE IF NOT EXISTS reviews (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  review_id text,
  user_id text,
  book_title text,
  review_title text,
  publisher text,
  review text,
  readed_at timestamp,
  stars int,
  public_flg boolean,
  created_at timestamp,
  updated_at timestamp
);

DROP TABLE IF EXISTS comments;
CREATE TABLE IF NOT EXISTS comments (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  comment_id text,
  review_id text,
  user_id text,
  comment text,
  created_at timestamp,
  updated_at timestamp
);

DROP TABLE IF EXISTS review_likes;
CREATE TABLE IF NOT EXISTS review_likes (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  review_like_id text,
  review_id text,
  user_id text,
  created_at timestamp,
  updated_at timestamp
);

DROP TABLE IF EXISTS comment_likes;
CREATE TABLE IF NOT EXISTS comment_likes (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  comment_like_id text,
  comment_id text,
  user_id text,
  created_at timestamp,
  updated_at timestamp
);
