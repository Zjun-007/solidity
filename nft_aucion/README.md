# NFT 拍卖系统 — Solidity + Go + Ethereum 事件索引

一套完整的 **ERC721 NFT 拍卖平台**，由 **Solidity 可升级合约** + **Go 后端 API** + **事件索引器** 组成。拍卖合约采用 **UUPS 代理模式**，支持升级；后端通过 **Gin** 提供 REST API，实时索引链上拍卖事件，并存储至 **MySQL**。

---

## 🎯 核心功能

| 功能 | 说明 |
|------|------|
| **创建拍卖** | NFT 所有者创建拍卖，设置起拍价、开始/结束时间 |
| **出价竞拍** | 参与者用 ETH 出价，自动跟踪最高出价者 |
| **结束结算** | 自动转账 NFT 给胜者，ETH 给卖家 |
| **事件索引** | 后台实时拉取链上事件（AuctionCreated、BidAuction 等），同步至数据库 |
| **REST API** | 提供拍卖列表、出价记录、统计、钱包 NFT 查询等接口 |
| **可升级设计** | 合约采用 UUPS 标准，无需迁移即可升级逻辑 |

---

## 📦 技术栈

## 📦 技术栈

| 组件 | 版本/框架 | 说明 |
|------|----------|------|
| **区块链** | Sepolia (ETH) | 测试网络 |
| **合约** | Solidity 0.8.24 | OpenZeppelin + UUPS |
| **合约工具** | Foundry | forge/cast 编译部署 |
| **后端** | Go 1.21+ | Gin Web、go-ethereum、GORM |
| **数据库** | MySQL 8.0+ | 存储拍卖、出价、索引状态 |
| **RPC** | Infura/Alchemy | ETH JSON-RPC 接口 |

---

## 📂 项目结构

```
├── src/
│   ├── NFTAuctionUUPS.sol         # 主拍卖合约（可升级）
│   ├── MyNFT.sol                  # 示例 ERC721 代币合约
│   ├── FeedAuction.sol            # 旧版单一拍卖合约（参考）
│   └── Counter.sol, Auction.sol   # 其他示例/测试
│
├── script/
│   ├── DeployNFTAuctionUUPS.s.sol # 部署脚本（UUPS 代理 + 实现）
│   ├── MyNFT.s.sol
│   └── ...
│
├── test/
│   ├── NFTAuctionUUPS.t.sol       # Foundry 测试
│   └── ...
│
├── internal/
│   ├── api/
│   │   ├── handler.go             # API 端点处理器
│   │   └── router.go              # Gin 路由定义
│   ├── blockchain/
│   │   ├── contract.go            # 合约实例化
│   │   ├── nftauctionuups_gen.go  # 合约 ABI 绑定（go-ethereum abigen 生成）
│   │   └── auction.go
│   ├── config/
│   │   └── config.go              # 环境变量加载
│   ├── db/
│   │   └── db.go                  # GORM MySQL 连接
│   ├── indexer/
│   │   └── indexer.go             # 事件拉取和同步器
│   ├── model/
│   │   └── uups.go                # 数据模型（Auction、Bid 等）
│   ├── repository/
│   │   └── uups_store.go          # 数据库操作（CRUD）
│   ├── service/
│   │   └── catalog.go             # 业务逻辑层
│   └── nft/
│       └── wallet.go              # NFT 钱包工具
│
├── migrations/
│   └── init.sql                   # MySQL 初始化脚本
│
├── abi/
│   ├── FeedAuction.json           # 旧合约 ABI
│   └── FeedAuction.abi
│
├── main.go                        # 应用入口
├── config.yaml                    # 配置文件（可选）
├── .env                           # 环境变量（本地开发）
├── go.mod / go.sum                # Go 依赖
├── foundry.toml                   # Foundry 配置
└── README.md                      # 本文件
```

---

## 🚀 快速开始

### 1. 环境配置

创建 `.env` 文件（复制 `.env.example` 或手动添加）：

```bash
# Ethereum RPC
ETH_RPC_URL=https://sepolia.infura.io/v3/YOUR_INFURA_KEY

# 数据库
DB_HOST=127.0.0.1
DB_PORT=3306
DB_USER=root
DB_PASSWORD=yourpassword
DB_NAME=nft_auction

# 部署的合约地址
NFT_AUCTION_PROXY_ADDRESS=0x3AC91EBd3383DD3EFe5762F96F0f30Ba2F90102D

# ERC721 NFT 合约地址（用于测试）
NFT_CONTRACT_ADDRESS=0xe248342E959a0ABdF99b08684a2EdFe7d1a1b010

# API 服务器端口
SERVER_PORT=8080

# 索引器块轮询间隔大小
INDEX_CHUNK_BLOCKS=1000

# 私钥（用于创建拍卖/出价等测试脚本）
PRIVATE_KEY=0x...
TOKEN_ID=1
```

### 2. 数据库初始化

```bash
# 创建数据库并导入初始化脚本
mysql -h 127.0.0.1 -u root -p -e "CREATE DATABASE nft_auction"
mysql -h 127.0.0.1 -u root -p nft_auction < migrations/init.sql
```

### 3. 启动 API 服务

```bash
# 编译并运行
go run main.go
```

服务会：
- 连接到 Ethereum RPC
- 初始化数据库连接
- 启动事件索引器（后台轮询链上事件）
- 在 `http://localhost:8080` 启动 API 服务器

**输出示例：**
```
2026/05/10 12:00:00 No .env file found, using system env
2026/05/10 12:00:00 API listening on :8080 (NFT auction proxy 0x3AC91EBd3383DD3EFe5762F96F0f30Ba2F90102D)
```

---

## 📡 API 文档

所有端点前缀为 `/api/v1`（除 `/health` 外）。

### 健康检查

```bash
GET /health
```

**响应：**
```json
{
  "status": "ok"
}
```

---

### 索引器状态

```bash
GET /api/v1/indexer/status
```

**说明：** 返回链上最新块号、已同步块号、延迟和最后错误信息。

**响应示例：**
```json
{
  "chain_head_block": 10824987,
  "lag": 10278987,
  "last_synced_block": 546000,
  "last_error": ""
}
```

| 字段 | 说明 |
|------|------|
| `chain_head_block` | 链上最新块高 |
| `last_synced_block` | 已同步到的块高 |
| `lag` | 落后块数 |
| `last_error` | 最后一次错误信息（无错时为空字符串） |

---

### 拍卖统计

```bash
GET /api/v1/stats
```

**说明：** 返回全局拍卖和出价数量。

**响应：**
```json
{
  "auction_total": 0,
  "bid_total": 0
}
```

---

### 拍卖列表

```bash
GET /api/v1/auctions
  ?status=<active|ended|cancelled>
  &seller=<address>
  &nft_contract=<address>
  &sort=<id|time>
  &page=<num>
  &page_size=<num>
```

**说明：** 分页列出所有拍卖（支持过滤）。

**参数：**
| 参数 | 类型 | 说明 | 默认值 |
|------|------|------|--------|
| `status` | string | 拍卖状态：`active`、`ended`、`cancelled` | 无（全部） |
| `seller` | string | 卖家地址 | 无 |
| `nft_contract` | string | NFT 合约地址 | 无 |
| `sort` | string | 排序字段：`id`、`time` | `id` |
| `page` | int | 页码（从 0 开始） | 0 |
| `page_size` | int | 每页条数 | 20 |

**响应示例：**
```json
{
  "items": [],
  "total": 0,
  "page": 0,
  "page_size": 20
}
```

---

### 拍卖出价记录

```bash
GET /api/v1/auctions/:id/bids
  ?page=<num>
  &page_size=<num>
```

**说明：** 列出指定拍卖 ID 的所有出价记录（从早到晚）。

**参数：**
| 参数 | 类型 | 说明 | 默认值 |
|------|------|------|--------|
| `page` | int | 页码（从 1 开始） | 1 |
| `page_size` | int | 每页条数 | 50 |

**响应示例：**
```json
{
  "auction_id": 1,
  "items": [
    {
      "id": 1,
      "auction_id": 1,
      "bidder": "0x...",
      "amount": "1500000000000000000",
      "timestamp": 1234567890,
      "tx_hash": "0x..."
    }
  ],
  "total": 1,
  "page": 1,
  "page_size": 50
}
```

---

### 钱包 NFT 查询

```bash
GET /api/v1/wallets/:address/nfts
  ?contracts=0x...,0x...,0x...
```

**说明：** 查询钱包中指定 ERC721 合约的 NFT 余额。

**参数：**
| 参数 | 类型 | 说明 | 必需 |
|------|------|------|------|
| `:address` | string | 钱包地址 | ✓ |
| `contracts` | string | 逗号分隔的 ERC721 合约地址 | ✓ |

**响应示例：**
```json
{
  "wallet": "0xbd5ef8efa4f5682d507740b38c0c4ace52a3e309",
  "contracts": [
    {
      "contract": "0xe248342e959a0abdf99b08684a2edfe7d1a1b010",
      "balance": "0",
      "enumerable": false,
      "note": "not ERC721Enumerable; cannot list tokenIds via RPC only"
    }
  ]
}
```

| 字段 | 说明 |
|------|------|
| `balance` | 该合约中的 NFT 数量 |
| `enumerable` | 是否实现 ERC721Enumerable 接口 |
| `note` | 额外说明（如非 Enumerable 则无法列出 tokenIds） |

---

## 🔗 智能合约

### NFTAuctionUUPS

主拍卖合约，采用 **UUPS 代理模式**（可升级）。

**关键接口：**

```solidity
// 创建拍卖
function createAuction(
    address nftContract,
    uint256 tokenId,
    uint256 startPrice,
    uint256 startTime,
    uint256 endTime
) external returns (uint256 auctionId);

// 出价
function bidAuction(uint256 auctionId) external payable;

// 结束拍卖（胜者取 NFT，卖家取 ETH）
function endAuction(uint256 auctionId) external nonReentrant;

// 取消拍卖（仅卖家在无出价时可用）
function cancelAuction(uint256 auctionId) external;

// 查询拍卖数据
function auctionData(uint256 auctionId) 
    external view returns (
        address seller,
        address nftContract,
        uint256 tokenId,
        uint256 startPrice,
        uint256 startTime,
        uint256 endTime,
        address highestBidder,
        uint256 highestBid,
        bool settled,
        bool cancelled
    );
```

**事件：**
```solidity
event AuctionCreated(uint256 indexed auctionId, address indexed seller, address nftContract, uint256 tokenId, uint256 startPrice, uint256 startTime, uint256 endTime);
event BidAuction(uint256 indexed auctionId, address indexed bidder, uint256 amount);
event AuctionEnded(uint256 indexed auctionId, address indexed winner, uint256 amount, address seller);
event AuctionCancelled(uint256 indexed auctionId, address indexed seller);
```

---

## 🔧 合约部署

### 使用 Foundry 部署

```bash
# 编译合约
forge build

# 部署到 Sepolia（使用 UUPS 代理 + 实现模式）
forge script script/DeployNFTAuctionUUPS.s.sol \
  --rpc-url https://sepolia.infura.io/v3/YOUR_KEY \
  --private-key YOUR_PRIVATE_KEY \
  --broadcast \
  --verify

# 查看部署结果
# 输出中会包含代理地址（PROXY）和实现地址（IMPLEMENTATION）
```

### 升级合约

如需升级逻辑，部署新的实现合约，然后调用代理的升级方法（仅 owner 可调）：

```bash
# 部署新实现
forge create src/NFTAuctionUUPSV2.sol:NFTAuctionUUPSV2 \
  --rpc-url https://sepolia.infura.io/v3/YOUR_KEY \
  --private-key YOUR_PRIVATE_KEY

# 然后通过代理升级（使用 cast 或前端调用）
cast send <PROXY_ADDRESS> "upgradeTo(address)" NEW_IMPLEMENTATION_ADDRESS \
  --rpc-url https://sepolia.infura.io/v3/YOUR_KEY \
  --private-key YOUR_PRIVATE_KEY
```

---

## 💾 数据库架构

### 主要表

**auctions**
```sql
CREATE TABLE auctions (
  id BIGINT PRIMARY KEY,
  seller VARCHAR(42),
  nft_contract VARCHAR(42),
  token_id BIGINT,
  start_price DECIMAL(50,0),
  start_time BIGINT,
  end_time BIGINT,
  highest_bidder VARCHAR(42),
  highest_bid DECIMAL(50,0),
  settled BOOLEAN,
  cancelled BOOLEAN,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

**bids**
```sql
CREATE TABLE bids (
  id BIGINT PRIMARY KEY AUTO_INCREMENT,
  auction_id BIGINT,
  bidder VARCHAR(42),
  amount DECIMAL(50,0),
  timestamp BIGINT,
  tx_hash VARCHAR(66),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (auction_id) REFERENCES auctions(id)
);
```

**indexer_state**
```sql
CREATE TABLE indexer_state (
  id INT PRIMARY KEY DEFAULT 1,
  last_synced_block BIGINT,
  last_error TEXT,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
```

---

## 🛠️ 常见操作

### 创建拍卖

```bash
# 使用提供的脚本（需要 .env 中的私钥）
go run create_auction.go
```

脚本会：
1. 检查 NFT 所有权
2. 批准拍卖合约使用 NFT
3. 调用 `createAuction`
4. 等待交易确认

### 查询链上拍卖数据

```bash
go run query_auction.go
```

### 手动测试 API

```bash
# 健康检查
curl http://localhost:8080/health

# 查询索引器状态
curl http://localhost:8080/api/v1/indexer/status

# 获取拍卖列表
curl http://localhost:8080/api/v1/auctions

# 获取统计信息
curl http://localhost:8080/api/v1/stats

# 查询钱包 NFT
curl "http://localhost:8080/api/v1/wallets/0xBd5eF8EfA4f5682d507740b38C0C4AcE52A3E309/nfts?contracts=0xe248342E959a0ABdF99b08684a2EdFe7d1a1b010"
```

---

## 🔍 事件索引工作流

1. **初始化**：应用启动时，从数据库读取 `last_synced_block`
2. **轮询**：后台定时从链上拉取 `[last_synced_block, chain_head]` 内的事件
3. **解析**：识别 `AuctionCreated`、`BidAuction`、`AuctionEnded`、`AuctionCancelled` 事件
4. **同步**：将事件数据写入 `auctions` 和 `bids` 表
5. **更新**：刷新 `indexer_state` 中的 `last_synced_block`

**关键配置：**
- `INDEX_CHUNK_BLOCKS`：每次轮询的块范围大小（默认 1000），可在 `.env` 配置

**特点：**
- 可容错：支持从任意块重新启动
- 增量同步：只查询新事件，避免重复
- 实时性：可配置轮询间隔以提高/降低实时性

---

## 📝 日志与调试

### 查看日志

应用直接打印到 stdout，可用 `go run main.go 2>&1 | tee app.log` 保存到文件。

**重要日志消息：**
- `indexer: starting block fetch` — 开始事件轮询
- `indexer: processing <N> events` — 处理事件数量
- `indexer: synced to block <N>` — 同步到某块
- `API listening on :<port>` — API 服务启动

### 调试模式

在 Go 代码中启用更详细的日志：
```go
// 在 internal/indexer/indexer.go 中添加
log.Printf("DEBUG: processing event %+v", event)
```

---

## ⚠️ 常见问题

### Q: 为什么 API 返回 0 个拍卖？
**A:** 索引器需要时间同步链上事件。检查 `/api/v1/indexer/status` 的 `last_synced_block` 是否接近 `chain_head_block`。如果落后很多，等待索引器追上。

### Q: 如何升级合约？
**A:** 部署新的实现合约，然后通过代理调用 `upgradeTo(newImplementation)` 方法（需要 owner 权限）。

### Q: 如何重置数据库？
**A:** 
```bash
mysql -u root -p -e "DROP DATABASE nft_auction; CREATE DATABASE nft_auction;"
mysql -u root -p nft_auction < migrations/init.sql
# 重启应用
```

### Q: NFT 合约是否需要特殊接口？
**A:** 需要支持标准 ERC721（`approve`、`transferFrom` 等）。如果要在 API 中查询余额，建议实现 `balanceOf` 和 `ownerOf`。

---

## 📚 参考资源

- [OpenZeppelin Contracts](https://github.com/OpenZeppelin/openzeppelin-contracts)
- [Foundry Book](https://book.getfoundry.sh/)
- [go-ethereum 文档](https://geth.ethereum.org/docs)
- [Gin Web Framework](https://gin-gonic.com/)
- [GORM 文档](https://gorm.io/)

---

## 📄 许可证

MIT License — 详见 [LICENSE](LICENSE) 文件。

---

## 👨‍💻 贡献

欢迎提交 Issue 和 Pull Request！

---

**最后更新：** 2026-05-10
