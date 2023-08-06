CREATE TABLE users (
  id UUID PRIMARY KEY,
  email VARCHAR(255) NOT NULL,
  name VARCHAR(255),
  socialName VARCHAR(255) NULL,
  birthDay DATE,
  UNIQUE(email)
);
