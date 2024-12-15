CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS todo_items (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
    description TEXT NOT NULL,
    due_date TIMESTAMP
    WITH
        TIME ZONE NOT NULL,
        file_id TEXT,
        created_at TIMESTAMP
    WITH
        TIME ZONE DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP
    WITH
        TIME ZONE DEFAULT CURRENT_TIMESTAMP
);