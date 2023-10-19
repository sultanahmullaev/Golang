CREATE TABLE IF NOT EXISTS sports (
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    title text NOT NULL,
    description text NOT NULL,
    type text NOT NULL,
    brand text NOT NULL,
    sex text NOT NULL,
    sports_equipment_number integer NOT NULL,
    version integer NOT NULL DEFAULT 1
);

