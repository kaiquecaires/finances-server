CREATE TABLE user_passwords (
  id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
  user_id UUID REFERENCES users(id) ON DELETE CASCADE,
  password VARCHAR(255) NOT NULL,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW(),
  deleted_at TIMESTAMP NULL
);

CREATE TRIGGER update_user_passwords_updated_at
BEFORE UPDATE ON user_passwords
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();
