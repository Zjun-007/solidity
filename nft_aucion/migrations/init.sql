CREATE DATABASE IF NOT EXISTS nft_auction;
USE nft_auction;

CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    wallet_address VARCHAR(42) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS auctions (
    id INT AUTO_INCREMENT PRIMARY KEY,
    auction_id BIGINT UNSIGNED NOT NULL UNIQUE,
    nft_contract VARCHAR(42) NOT NULL,
    token_id BIGINT UNSIGNED NOT NULL,
    start_price VARCHAR(78) NOT NULL,
    highest_bid VARCHAR(78) DEFAULT '0',
    highest_bidder VARCHAR(42),
    end_time TIMESTAMP NOT NULL,
    ended BOOLEAN DEFAULT FALSE,
    creator VARCHAR(42) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_ended_endtime (ended, end_time)
);

CREATE TABLE IF NOT EXISTS bids (
    id INT AUTO_INCREMENT PRIMARY KEY,
    auction_id BIGINT UNSIGNED NOT NULL,
    bidder VARCHAR(42) NOT NULL,
    amount VARCHAR(78) NOT NULL,
    tx_hash VARCHAR(66),
    block_number BIGINT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_auction (auction_id),
    FOREIGN KEY (auction_id) REFERENCES auctions(auction_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS sync_logs (
    id INT AUTO_INCREMENT PRIMARY KEY,
    event_type VARCHAR(50) NOT NULL,
    block_number BIGINT NOT NULL,
    tx_hash VARCHAR(66) NOT NULL,
    processed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE KEY unique_event (event_type, tx_hash)
);