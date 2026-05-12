INSERT INTO wallets (id, balance, created_at)
VALUES 
    ('550e8400-e29b-41d4-a716-446655440000', 10000, NOW()),
    ('771e8400-e29b-41d4-a716-446655441111', 0, NOW()),
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 500, NOW())
ON CONFLICT (id) DO NOTHING;