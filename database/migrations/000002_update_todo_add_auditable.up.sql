ALTER TABLE todos
ADD COLUMN IF NOT EXISTS created_at TIMESTAMP DEFAULT NOW();
ALTER TABLE todos
ADD COLUMN IF NOT EXISTS updated_at TIMESTAMP DEFAULT NOW();