CREATE INDEX IF NOT EXISTS sports_title_idx ON sports USING GIN (to_tsvector('simple', title));
