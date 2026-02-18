-- +goose Up
INSERT INTO tours (created_at, updated_at, name, base_price)
VALUES
  (NOW(), NOW(), 'Grand Tour of Isla Nublar (spared no expense)', 10000.00),
  (NOW(), NOW(), 'Boat Tour of Amity Island (chum included)', 7000.00),
  (NOW(), NOW(), 'A Historic Tour Trough Hill Valley (please come in time)', 8500.00);


-- +goose Down
DELETE * FROM tours
WHERE name IN (
  'Grand Tour of Isla Nublar (spared no expense)',
  'Boat Tour of Amity Island (chum included)',
  'A Historic Tour Trough Hill Valley (please come in time)'
);
