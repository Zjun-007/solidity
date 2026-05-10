import path from "path";
import { fileURLToPath } from "url";
import dotenv from "dotenv";
import express from "express";
import cors from "cors";
import { JsonRpcProvider } from "ethers";
import { getPool, initSchema } from "./db.js";
import { getNftContract, runIndexer } from "./indexer.js";

const __dirname = path.dirname(fileURLToPath(import.meta.url));
dotenv.config({ path: path.join(__dirname, "../../.env") });
dotenv.config({ path: path.join(__dirname, "../.env") });

const app = express();
app.use(cors());
app.use(express.json());

const NFT_ADDR = (process.env.NFT_CONTRACT_ADDRESS || "").trim().toLowerCase();

/** MySQL 部分版本不支持预处理语句里的 LIMIT/OFFSET 占位符，改为校验后的字面量。 */
function clampPagination(limitRaw, offsetRaw, maxLimit = 200) {
  let limit = Math.floor(Number(limitRaw));
  if (!Number.isFinite(limit) || limit < 1) limit = 50;
  if (limit > maxLimit) limit = maxLimit;
  let offset = Math.floor(Number(offsetRaw));
  if (!Number.isFinite(offset) || offset < 0) offset = 0;
  return { limit, offset };
}

app.get("/health", (_req, res) => {
  res.json({ ok: true, contract: NFT_ADDR || null });
});

/** 合约元数据：name / symbol / owner */
app.get("/api/nft/contract", async (_req, res) => {
  try {
    const provider = new JsonRpcProvider(
      process.env.ETH_RPC_URL || process.env.SEPOLIA_RPC_URL
    );
    const c = getNftContract(provider);
    const [name, symbol, owner] = await Promise.all([
      c.name(),
      c.symbol(),
      c.owner(),
    ]);
    res.json({
      address: NFT_ADDR,
      name,
      symbol,
      owner: owner.toLowerCase(),
    });
  } catch (e) {
    res.status(500).json({ error: String(e.message || e) });
  }
});

/** 某地址持有数量 */
app.get("/api/nft/balance/:owner", async (req, res) => {
  try {
    const provider = new JsonRpcProvider(
      process.env.ETH_RPC_URL || process.env.SEPOLIA_RPC_URL
    );
    const c = getNftContract(provider);
    const owner = req.params.owner;
    const n = await c.balanceOf(owner);
    res.json({ owner: owner.toLowerCase(), balance: n.toString() });
  } catch (e) {
    res.status(500).json({ error: String(e.message || e) });
  }
});

/** tokenId 当前 owner；tokenURI 可能因未设置元数据而 revert */
app.get("/api/nft/token/:tokenId", async (req, res) => {
  try {
    const provider = new JsonRpcProvider(
      process.env.ETH_RPC_URL || process.env.SEPOLIA_RPC_URL
    );
    const c = getNftContract(provider);
    const tokenId = req.params.tokenId;
    const owner = await c.ownerOf(tokenId);
    let tokenURI = null;
    try {
      tokenURI = await c.tokenURI(tokenId);
    } catch {
      tokenURI = null;
    }
    let approved = null;
    try {
      const a = await c.getApproved(tokenId);
      approved = a === "0x0000000000000000000000000000000000000000" ? null : a.toLowerCase();
    } catch {
      approved = null;
    }
    res.json({
      tokenId,
      owner: owner.toLowerCase(),
      tokenURI,
      approvedSingle: approved,
    });
  } catch (e) {
    res.status(500).json({ error: String(e.message || e) });
  }
});

/** 单地址是否授权某 operator 全部 NFT */
app.get("/api/nft/is-approved-for-all", async (req, res) => {
  try {
    const owner = (req.query.owner || "").toString();
    const operator = (req.query.operator || "").toString();
    if (!owner || !operator) {
      return res.status(400).json({ error: "query owner and operator required" });
    }
    const provider = new JsonRpcProvider(
      process.env.ETH_RPC_URL || process.env.SEPOLIA_RPC_URL
    );
    const c = getNftContract(provider);
    const ok = await c.isApprovedForAll(owner, operator);
    res.json({
      owner: owner.toLowerCase(),
      operator: operator.toLowerCase(),
      isApprovedForAll: ok,
    });
  } catch (e) {
    res.status(500).json({ error: String(e.message || e) });
  }
});

/** 分页列出已索引链上事件（来自 MySQL） */
app.get("/api/nft/events", async (req, res) => {
  try {
    await initSchema();
    const { limit, offset } = clampPagination(req.query.limit, req.query.offset);
    const eventName = (req.query.type || "").toString().trim();
    const tokenId = (req.query.tokenId || "").toString().trim();

    const p = getPool();
    let sql =
      "SELECT id, block_number, tx_hash, log_index, event_name, token_id, from_addr, to_addr, owner_addr, approved_addr, operator_addr, approved_bool, previous_owner, new_owner, raw_args, created_at FROM nft_chain_events WHERE 1=1";
    const params = [];
    if (eventName) {
      sql += " AND event_name = ?";
      params.push(eventName);
    }
    if (tokenId) {
      sql += " AND token_id = ?";
      params.push(tokenId);
    }
    sql += ` ORDER BY block_number DESC, log_index DESC LIMIT ${limit} OFFSET ${offset}`;

    const [rows] = await p.execute(sql, params);
    let countSql =
      "SELECT COUNT(*) AS c FROM nft_chain_events WHERE 1=1";
    const countParams = [];
    if (eventName) {
      countSql += " AND event_name = ?";
      countParams.push(eventName);
    }
    if (tokenId) {
      countSql += " AND token_id = ?";
      countParams.push(tokenId);
    }
    const [countRows] = await p.execute(countSql, countParams);
    res.json({
      total: Number(countRows[0].c),
      limit,
      offset,
      items: rows,
    });
  } catch (e) {
    res.status(500).json({ error: String(e.message || e) });
  }
});

/** 某地址作为 from/to/owner 参与的事件 */
app.get("/api/nft/events/address/:addr", async (req, res) => {
  try {
    await initSchema();
    const addr = req.params.addr.toLowerCase();
    const { limit, offset } = clampPagination(req.query.limit, req.query.offset);
    const p = getPool();
    const [rows] = await p.execute(
      `SELECT id, block_number, tx_hash, log_index, event_name, token_id, from_addr, to_addr, owner_addr, approved_addr, operator_addr, approved_bool, previous_owner, new_owner, raw_args, created_at
       FROM nft_chain_events
       WHERE from_addr = ? OR to_addr = ? OR owner_addr = ? OR approved_addr = ? OR operator_addr = ? OR previous_owner = ? OR new_owner = ?
       ORDER BY block_number DESC, log_index DESC
       LIMIT ${limit} OFFSET ${offset}`,
      [addr, addr, addr, addr, addr, addr, addr]
    );
    res.json({ address: addr, limit, offset, items: rows });
  } catch (e) {
    res.status(500).json({ error: String(e.message || e) });
  }
});

const port = Number(process.env.SERVER_PORT || process.env.PORT || 8080);

async function main() {
  await initSchema();
  runIndexer({
    onError: (err) => console.error("[indexer]", err),
  }).catch((err) => console.error("[indexer fatal]", err));

  app.listen(port, () => {
    console.log(`MyNFT backend http://localhost:${port}`);
    console.log(`Contract ${NFT_ADDR}`);
  });
}

main();
