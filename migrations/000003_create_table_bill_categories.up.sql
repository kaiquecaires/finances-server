CREATE TABLE bill_categories (
  id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
  description VARCHAR(255),
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW(),
  deleted_at TIMESTAMP NULL
);

CREATE TRIGGER update_bill_categories_updated_at
BEFORE UPDATE ON bill_categories
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();
