# NFT Auction Platform
基于以太坊的 NFT 拍卖系统，支持卖家创建拍卖、用户竞拍、自动退款等功能。项目包含 Solidity 智能合约（`FeedAuction` + `MyNFT`）、Go 后端 API 服务以及 Foundry 部署脚本。

## 项目结构
.
├── src/ # Solidity 合约
│ ├── FeedAuction.sol # 拍卖主合约
│ └── MyNFT.sol # 示例 ERC721 代币合约
├── script/ # Foundry 部署脚本
│ ├── FeedAuction.s.sol # 拍卖合约部署脚本
│ └── MyNFT.s.sol # NFT 合约部署脚本
├── out/ # 编译生成的 ABI 和 Bytecode
├── main.go   #主函数
├── internal/
│   ├── api/
│   │   ├── handler.go
│   │   └── router.go
│   ├── blockchain/
│   │   ├── client.go          #Client获取blockchain信息
│   │   ├── contract.go      # 由 abigen 生成
│   │   └── event_listener.go  #事件监听
│   ├── model/
│   │   └── models.go   #数据结构
│   ├── repository/
│   │   └── auction_repo.go   #数据访问
│   └── service/
│       └── auction_service.go   #业务逻辑
├── config/
│   └── config.yaml   #数据库&blockchain配置
├── migrations/
│   └── init.sql    #mysql表结构
├── go.mod
└── go.sum
├── foundry.toml # Foundry 配置
└── README.md

## 技术栈

- **智能合约**：Solidity ^0.8.24，OpenZeppelin 库
- **开发框架**：Foundry（编译、测试、部署）
- **后端**：Go 1.21+，Gin Web Framework，go-ethereum
- **交互工具**：Cast（Foundry 组件），Postman / cURL
- **测试网络**：Sepolia（Infura 节点）

## 智能合约说明

### FeedAuction（拍卖合约）

- 卖家创建拍卖时指定 NFT 合约地址、`tokenId`、起拍价、持续时间。
- 卖家需先调用 `depositNFT()` 将 NFT 转入合约，拍卖才会激活。
- 买家通过 `bid()` 出价，出价必须高于当前最高价且不低于起拍价。
- 被超越的买家可通过 `withdraw()` 取回被锁定的保证金。
- 拍卖结束后，卖家调用 `endAuction()`：若有最高出价者，NFT 转给赢家，卖家收到最高价 ETH；否则 NFT 退还给卖家。
- 卖家可在结束前调用 `cancelAuction()` 取消拍卖（需 NFT 尚未存入或已取回）。

### MyNFT（示例 ERC721）

- 简单的 NFT 铸造合约，仅合约拥有者可以调用 `mint()` 生成新 NFT。
- 提供标准 ERC721 方法：`ownerOf`、`balanceOf`、`approve`、`transferFrom` 等。

## 环境准备

### 1. 安装 Foundry

```bash
curl -L https://foundry.paradigm.xyz | bash
foundryup
2. 安装 Go 依赖
bash
go mod tidy
3. 设置环境变量
创建 .env 文件：

bash
PRIVATE_KEY=你的钱包私钥（64位十六进制，无0x前缀）
SEPOLIA_RPC_URL=https://sepolia.infura.io/v3/你的Infura项目ID
ETHERSCAN_API_KEY=你的Etherscan API密钥（用于合约验证）
NFT_CONTRACT_ADDRESS=0x...   # MyNFT 合约地址（部署后填写）
TOKEN_ID=1
START_PRICE=1000000000000000   # 0.001 ETH
DURATION=604800                 # 7天
加载变量：

bash
source .env
部署合约
部署 MyNFT
bash
forge create src/MyNFT.sol:MyNFT \
  --rpc-url $SEPOLIA_RPC_URL \
  --private-key $PRIVATE_KEY
记录输出的合约地址，设为环境变量 NFT_CONTRACT_ADDRESS。

铸造 NFT 给卖家
bash
cast send $NFT_CONTRACT_ADDRESS "mint(address)" $YOUR_ADDRESS \
  --rpc-url $SEPOLIA_RPC_URL \
  --private-key $PRIVATE_KEY
记录获得的 tokenId（通常从 1 开始）。

部署 FeedAuction
bash
forge create src/FeedAuction.sol:FeedAuction \
  --constructor-args $NFT_CONTRACT_ADDRESS $TOKEN_ID $START_PRICE $DURATION \
  --rpc-url $SEPOLIA_RPC_URL \
  --private-key $PRIVATE_KEY
记录拍卖合约地址，设为 AUCTION_CONTRACT。

授权并存入 NFT
bash
# 授权拍卖合约操作该 NFT
cast send $NFT_CONTRACT_ADDRESS "approve(address,uint256)" $AUCTION_CONTRACT $TOKEN_ID \
  --rpc-url $SEPOLIA_RPC_URL \
  --private-key $PRIVATE_KEY

# 存入 NFT
cast send $AUCTION_CONTRACT "depositNFT()" \
  --rpc-url $SEPOLIA_RPC_URL \
  --private-key $PRIVATE_KEY
验证合约（可选）
bash
forge verify-contract $AUCTION_CONTRACT src/FeedAuction.sol:FeedAuction \
  --chain-id 11155111 \
  --constructor-args $(cast abi-encode "constructor(address,uint256,uint256,uint256)" $NFT_CONTRACT_ADDRESS $TOKEN_ID $START_PRICE $DURATION) \
  --etherscan-api-key $ETHERSCAN_API_KEY
启动后端 API 服务
bash
go run main.go
服务默认监听 http://localhost:8080。

API 文档
健康检查
GET /health → {"status":"ok"}

拍卖信息查询
GET /api/v1/auction?address=<合约地址>

响应示例：

json
{
  "seller": "0x...",
  "highestBidder": "0x...",
  "highestBid": "1000000000000000",
  "startPrice": "1000000000000000",
  "endTime": 1744567890,
  "isActive": true,
  "canceled": false,
  "ended": false,
  "nftDeposited": true
}
存入 NFT（卖家）
POST /api/v1/auction/deposit

json
{
  "private_key": "卖家私钥（64位无0x）",
  "contract_address": "0x拍卖合约地址"
}
出价
POST /api/v1/auction/bid

json
{
  "private_key": "竞拍者私钥",
  "contract_address": "0x拍卖合约地址",
  "value_eth": "0.001"
}
结束拍卖（卖家）
POST /api/v1/auction/end

json
{
  "private_key": "卖家私钥",
  "contract_address": "0x拍卖合约地址"
}
取回被超越的保证金
POST /api/v1/auction/withdraw

json
{
  "private_key": "竞拍者私钥",
  "contract_address": "0x拍卖合约地址"
}
取消拍卖（卖家，未结束前）
POST /api/v1/auction/cancel

json
{
  "private_key": "卖家私钥",
  "contract_address": "0x拍卖合约地址"
}
使用 Postman 测试
安装 Postman（Windows / macOS / Linux）。

创建集合 FeedAuction API。

配置环境变量 base_url = http://localhost:8080。

按照 API 文档逐一创建请求并测试。

也可以使用 curl 命令快速测试，例如：

bash
curl -X POST http://localhost:8080/api/v1/auction/bid \
  -H "Content-Type: application/json" \
  -d '{"private_key":"...","contract_address":"0x...","value_eth":"0.001"}'

## 常见问题
1. depositNFT 报错 Auction expired
原因：拍卖持续时间已过，未能在结束前存入 NFT。

解决：重新部署拍卖合约，设置更长的 DURATION（如 7 天），并尽快调用 depositNFT。

2. execution reverted: NFT not yet deposited
原因：未调用 depositNFT 或 NFT 未被转入。

解决：先执行 approve 再调用 depositNFT，确认 isActive() 返回 true。

3. 私钥解码失败（Failed to decode private key）
检查私钥是否为 64 位十六进制字符串，不含 0x 前缀。

使用 cast wallet address --private-key $PRIVATE_KEY 验证。

4. Etherscan 验证超时
国内网络可能需要代理，设置 https_proxy 环境变量。

或者使用 forge verify-contract --show-standard-json-input 生成 JSON 文件后手动上传。

开发与扩展
添加更多 NFT 相关 API（mint、ownerOf、balanceOf、approve）：参考 internal/api/handler.go 中的实现模板。

支持多拍卖合约动态管理：可通过数据库记录拍卖信息，允许用户创建多个拍卖。

增加事件监听服务：实时同步链上 Bid、AuctionEnded 等事件并推送至前端。

许可证
MIT
