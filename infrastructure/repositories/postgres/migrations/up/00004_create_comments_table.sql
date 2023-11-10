CREATE TABLE IF NOT EXISTS Comments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    comment_text TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    parent_id UUID REFERENCES Comments(id) DEFAULT NULL,
    creator_id UUID NOT NULL REFERENCES Users(id),
    thread_id UUID NOT NULL REFERENCES Threads(id)
);
