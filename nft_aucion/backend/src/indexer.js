import { Contract, WebSocketProvider, JsonRpcProvider } from "ethers";
import { MYNFT_ABI } from "./abi.js";
import {
  getLastSyncedBlock,
  setLastSyncedBlock,
  initSchema,
} from "./db.js";
import { persistLog } from "./events.js";

const SYNC_KEY = "mynft_last_block";
const CHUNK = 2000;

function getProvider() {
  const ws = process.env.ETH_WS_URL;
  if (ws) return new WebSocketProvider(ws);
  return new JsonRpcProvider(process.env.ETH_RPC_URL || process.env.SEPOLIA_RPC_URL);
}

export function getNftContract(provider) {
  const addr = (process.env.NFT_CONTRACT_ADDRESS || "").trim();
  if (!addr) throw new Error("NFT_CONTRACT_ADDRESS is required");
  return new Contract(addr, MYNFT_ABI, provider);
}

async function backfillRange(contract, fromBlock, toBlock) {
  if (fromBlock > toBlock) return;
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
    for (const log of logs) {
      await persistLog(log);
    }
    await setLastSyncedBlock(end, SYNC_KEY);
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

/**
 * 历史同步 +（可选）WebSocket 订阅新日志
 */
export async function runIndexer({ onError } = {}) {
  await initSchema();
  const provider = getProvider();
  const contract = getNftContract(provider);
  const latest = await provider.getBlockNumber();
  let start = (await getLastSyncedBlock(SYNC_KEY)) ?? null;

  if (start === null) {
    start = firstFromBlock();
    await backfillRange(contract, start, latest);
  } else {
    await backfillRange(contract, start + 1, latest);
  }

  const address = await contract.getAddress();
  const iface = contract.interface;
  const eventTopics = [
    iface.getEvent("Transfer").topicHash,
    iface.getEvent("Approval").topicHash,
    iface.getEvent("ApprovalForAll").topicHash,
    iface.getEvent("OwnershipTransferred").topicHash,
  ];

  const onLogs = async (logs) => {
    for (const log of logs) {
      try {
        await persistLog(log);
        if (log.blockNumber != null) {
          await setLastSyncedBlock(log.blockNumber, SYNC_KEY);
        }
      } catch (e) {
        onError?.(e);
      }
    }
  };

  if (provider instanceof WebSocketProvider) {
    provider.on("block", async (blockNumber) => {
      try {
        const logs = await provider.getLogs({
          address,
          topics: [eventTopics],
          fromBlock: blockNumber,
          toBlock: blockNumber,
        });
        if (logs.length === 0) {
          await setLastSyncedBlock(blockNumber, SYNC_KEY);
        } else {
          await onLogs(logs);
        }
      } catch (e) {
        onError?.(e);
      }
    });
  } else {
    let last = latest;
    setInterval(async () => {
      try {
        const head = await provider.getBlockNumber();
        if (head > last) {
          await backfillRange(contract, last + 1, head);
          last = head;
        }
      } catch (e) {
        onError?.(e);
      }
    }, Number(process.env.INDEX_POLL_MS || 12000));
  }

  return { provider, contract };
}
