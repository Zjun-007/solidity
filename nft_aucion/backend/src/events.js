import { Interface } from "ethers";
import { MYNFT_ABI } from "./abi.js";
import { getPool } from "./db.js";

const iface = new Interface(MYNFT_ABI);

/**
 * @param {import("ethers").Log} log
 * @returns {object|null} row for nft_chain_events
 */
export function parseLogRow(log) {
  let parsed;
  try {
    parsed = iface.parseLog({ topics: log.topics, data: log.data });
  } catch {
    return null;
  }

  const rawObj = {};
  for (let i = 0; i < parsed.fragment.inputs.length; i++) {
    const inp = parsed.fragment.inputs[i];
    const key = inp.name || `arg${i}`;
    let v = parsed.args[inp.name] ?? parsed.args[i];
    if (typeof v === "bigint") v = v.toString();
    else if (typeof v === "string" && v.startsWith("0x")) v = v.toLowerCase();
    rawObj[key] = v;
  }

  const base = {
    block_number: log.blockNumber,
    tx_hash: log.transactionHash,
    log_index: log.index,
    event_name: parsed.name,
    token_id: null,
    from_addr: null,
    to_addr: null,
    owner_addr: null,
    approved_addr: null,
    operator_addr: null,
    approved_bool: null,
    previous_owner: null,
    new_owner: null,
    raw_args: JSON.stringify(rawObj),
  };

  if (parsed.name === "Transfer") {
    const [from, to, tokenId] = parsed.args;
    base.from_addr = String(from).toLowerCase();
    base.to_addr = String(to).toLowerCase();
    base.token_id = tokenId.toString();
  } else if (parsed.name === "Approval") {
    const [owner, approved, tokenId] = parsed.args;
    base.owner_addr = String(owner).toLowerCase();
    base.approved_addr = String(approved).toLowerCase();
    base.token_id = tokenId.toString();
  } else if (parsed.name === "ApprovalForAll") {
    const [owner, operator, approved] = parsed.args;
    base.owner_addr = String(owner).toLowerCase();
    base.operator_addr = String(operator).toLowerCase();
    base.approved_bool = approved ? 1 : 0;
  } else if (parsed.name === "OwnershipTransferred") {
    const [previousOwner, newOwner] = parsed.args;
    base.previous_owner = String(previousOwner).toLowerCase();
    base.new_owner = String(newOwner).toLowerCase();
  }

  return base;
}

export async function persistLogRow(row) {
  const p = getPool();
  await p.execute(
    `INSERT INTO nft_chain_events (
      block_number, tx_hash, log_index, event_name, token_id,
      from_addr, to_addr, owner_addr, approved_addr, operator_addr,
      approved_bool, previous_owner, new_owner, raw_args
    ) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?)
    ON DUPLICATE KEY UPDATE event_name = VALUES(event_name)`,
    [
      row.block_number,
      row.tx_hash,
      row.log_index,
      row.event_name,
      row.token_id,
      row.from_addr,
      row.to_addr,
      row.owner_addr,
      row.approved_addr,
      row.operator_addr,
      row.approved_bool,
      row.previous_owner,
      row.new_owner,
      row.raw_args,
    ]
  );
}

export async function persistLog(log) {
  const row = parseLogRow(log);
  if (!row) return;
  await persistLogRow(row);
}
