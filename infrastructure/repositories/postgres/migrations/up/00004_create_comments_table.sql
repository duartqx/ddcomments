CREATE TABLE IF NOT EXISTS Comments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    comment_text TEXT,
    parent_id UUID NOT NULL REFERENCES Comments(id),
    creator_id UUID NOT NULL REFERENCES Users(id),
    thread_id UUID NOT NULL REFERENCES Threads(id)
);
