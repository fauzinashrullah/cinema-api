-- USERS
CREATE TABLE users (
    id UUID PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- THEATERS
CREATE TABLE theaters (
    id UUID PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    city VARCHAR(100) NOT NULL
);

-- FILMS
CREATE TABLE films (
    id UUID PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    duration INT NOT NULL,
    description TEXT
);

-- SCHEDULES
CREATE TABLE schedules (
    id UUID PRIMARY KEY,
    film_id UUID NOT NULL,
    theater_id UUID NOT NULL,
    show_time TIMESTAMP NOT NULL,
    FOREIGN KEY (film_id) REFERENCES films(id),
    FOREIGN KEY (theater_id) REFERENCES theaters(id)
);

-- TRANSACTIONS
CREATE TABLE transactions (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    schedule_id UUID NOT NULL,
    total_amount DECIMAL(10, 2) NOT NULL,
    status VARCHAR(50) CHECK (status IN ('pending', 'paid', 'failed', 'refunded')),
    payment_time TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (schedule_id) REFERENCES schedules(id)
);

-- SEATS
CREATE TABLE seats (
    id UUID PRIMARY KEY,
    schedule_id UUID NOT NULL,
    seat_number VARCHAR(10) NOT NULL,
    status VARCHAR(50) CHECK (status IN ('available', 'locked', 'sold', 'refunded')),
    locked_until TIMESTAMP,
    transaction_id UUID,
    FOREIGN KEY (schedule_id) REFERENCES schedules(id),
    FOREIGN KEY (transaction_id) REFERENCES transactions(id),
    UNIQUE (schedule_id, seat_number)
);