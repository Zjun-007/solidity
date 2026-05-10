import mysql from "mysql2/promise";

let pool;

export function getPool() {
  if (!pool) {
    pool = mysql.createPool({
      host: process.env.DB_HOST || "localhost",
      port: Number(process.env.DB_PORT || 3306),
      user: process.env.DB_USER,
      password: process.env.DB_PASSWORD,
      database: process.env.DB_NAME,
      waitForConnections: true,
      connectionLimit: 10,
    });
  }
  return pool;
}

export async function initSchema() {
  const p = getPool();
  await p.execute(`
    CREATE TABLE IF NOT EXISTS chain_sync_state (
      \`key\` VARCHAR(64) PRIMARY KEY,
      value BIGINT UNSIGNED NOT NULL
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
  `);
  await p.execute(`
    CREATE TABLE IF NOT EXISTS nft_chain_events (
      id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
      block_number BIGINT UNSIGNED NOT NULL,
      tx_hash VARCHAR(66) NOT NULL,
      log_index INT UNSIGNED NOT NULL,
      event_name VARCHAR(64) NOT NULL,
      token_id VARCHAR(78) NULL,
      from_addr VARCHAR(42) NULL,
      to_addr VARCHAR(42) NULL,
      owner_addr VARCHAR(42) NULL,
      approved_addr VARCHAR(42) NULL,
      operator_addr VARCHAR(42) NULL,
      approved_bool TINYINT(1) NULL,
      previous_owner VARCHAR(42) NULL,
      new_owner VARCHAR(42) NULL,
      raw_args JSON NULL,
      created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
      UNIQUE KEY uk_tx_log (tx_hash, log_index),
      KEY idx_block (block_number),
      KEY idx_event (event_name),
      KEY idx_token (token_id)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
  `);
}

export async function getLastSyncedBlock(defaultKey = "mynft_last_block") {
  const p = getPool();
  const [rows] = await p.execute(
    "SELECT value FROM chain_sync_state WHERE `key` = ?",
    [defaultKey]
  );
  if (!rows.length) return null;
  return Number(rows[0].value);
}

export async function setLastSyncedBlock(block, defaultKey = "mynft_last_block") {
  const p = getPool();
  await p.execute(
    "INSERT INTO chain_sync_state (`key`, value) VALUES (?, ?) ON DUPLICATE KEY UPDATE value = VALUES(value)",
    [defaultKey, block]
  );
}
