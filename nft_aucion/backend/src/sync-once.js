/**
 * 仅执行一次历史同步（不写 WebSocket 监听），便于首次灌库或修复数据。
 */
import path from "path";
import { fileURLToPath } from "url";
import dotenv from "dotenv";
import { JsonRpcProvider } from "ethers";
import { initSchema, getLastSyncedBlock, setLastSyncedBlock } from "./db.js";
import { getNftContract } from "./indexer.js";
import { persistLog } from "./events.js";

const __dirname = path.dirname(fileURLToPath(import.meta.url));
dotenv.config({ path: path.join(__dirname, "../../.env") });
dotenv.config({ path: path.join(__dirname, "../.env") });

const SYNC_KEY = "mynft_last_block";
const CHUNK = 2000;

async function backfillRange(contract, fromBlock, toBlock) {
  const iface = contract.interface;
  const address = await contract.getAddress();
  const topics = [
    [
      iface.getEvent("Transfer").topicHash,
      iface.getEvent("Approval").topicHash,
      iface.getEvent("ApprovalForAll").topicHash,
      iface.getEvent("OwnershipTransferred").topicHash,
    ],
  ];
  const provider = contract.runner.provider;
  let cursor = fromBlock;
  while (cursor <= toBlock) {
    const end = Math.min(cursor + CHUNK - 1, toBlock);
    const logs = await provider.getLogs({
      address,
      topics,
      fromBlock: cursor,
      toBlock: end,
    });
    for (const log of logs) await persistLog(log);
    await setLastSyncedBlock(end, SYNC_KEY);
    console.log(`synced ${cursor}..${end}`);
    cursor = end + 1;
  }
}

function firstFromBlock() {
  const raw = process.env.NFT_DEPLOY_FROM_BLOCK;
  if (raw !== undefined && raw !== "") {
    const n = Number(raw);
    if (Number.isFinite(n) && n >= 0) return n;
  }
  return 0;
}

async function main() {
  await initSchema();
  const provider = new JsonRpcProvider(
    process.env.ETH_RPC_URL || process.env.SEPOLIA_RPC_URL
  );
  const contract = getNftContract(provider);
  const latest = await provider.getBlockNumber();
  let start = (await getLastSyncedBlock(SYNC_KEY)) ?? null;
  if (start === null) start = firstFromBlock();
  else start = start + 1;
  await backfillRange(contract, start, latest);
  console.log("done up to", latest);
}

main().catch((e) => {
  console.error(e);
  process.exit(1);
});
