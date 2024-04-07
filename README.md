# Dcard Ad API

## 開發環境
OS：Windows 10
Database：PostgreSQL、Redis (由 Docker 管理)
IDE：Goland
Test Tool：Postman、K6

## 設計
### 基礎架構
![](https://drive.google.com/u/2/uc?id=1BrIZw3UwW3FRL8eU_FvTH2DHUpOlvQAK&export=download)
### ER Model
![](https://drive.google.com/u/2/uc?id=1w4b1ztXmgL8KP7mh-rYxCmAmJhJ05OSD&export=download)
- 熱門緩存資料雖違反正規化，但在廣告資料資料眾多時，因主要侷限範圍都在熱門緩存資料表，而筆數又是固定，查詢速度將不會随之增加

#### Ad Table
存放廣告資料
| Field    | Description            |
|----------|------------------------|
| id       | 主鍵，自增碼            |
| title    | 廣告標題 |
| startAt  | 廣告投放起始日 |
| endAt    | 廣告投放截止日 |

#### AdCondition Table
存放廣告投放條件
| Field     | Description                            |
|-----------|----------------------------------------|
| id        | 主鍵，自增碼                          |
| age_start | 年齡起始值，大於等於此值的用戶將符合條件 |
| age_end   | 年齡結束值，小於等於此值的用戶將符合條件 |
| gender    | 用戶性別，可能的值有男性、女性   |
| country   | 用戶所在國家，以兩個字母的代碼表示       |
| platform  | 用戶設備平台，可能的值有Android、IOS等  |
| ad_id     | 廣告ID，對應到廣告表中的ID               |

#### HotAd Table
存放熱門廣告資料
| Field     | Description                            |
|-----------|----------------------------------------|
| id        | 主鍵，自增碼                          |
| title     | 廣告標題                               |
| startAt   | 廣告投放起始日                         |
| endAt     | 廣告投放截止日                         |
| ad_id     | 廣告ID，對應到廣告表中的ID            |

#### HotAd Table
存放熱門廣告投放條件資料
| Field           | Description                            |
|-----------------|----------------------------------------|
| id              | 主鍵，自增碼                          |
| age_start       | 年齡起始值，大於等於此值的用戶將符合條件 |
| age_end         | 年齡結束值，小於等於此值的用戶將符合條件 |
| gender          | 用戶性別，可能的值有男性、女性   |
| country         | 用戶所在國家，以兩個字母的代碼表示       |
| platform        | 用戶設備平台，可能的值有Android、IOS等  |
| hot_ad_id       | 熱門廣告ID，對應到熱門廣告表中的ID       |
| ad_condition_id | 廣告條件ID，對應到廣告條件表中的ID       |


### 流程
#### Create Ad API 流程
![](https://drive.google.com/u/2/uc?id=14mH75y8GhZ7yUPkL5iPT2SSE_50yLcWP&export=download)
- 檢查資料合法性
- 將資料寫入 Ad 與  AdCondition 表
- 若有錯誤則透過事務 Rollback
#### Get Ad API 流程
<p align="center">
    <img src="https://drive.google.com/u/2/uc?id=1N0rsANCv-qSmgTzkoI0LiwugCATkca3R&export=download" />
</p>
- Eager Loading 將活躍廣告寫入至 hot data table 供後續資料庫查詢使用
- 因題目無特別要求，因此此處忽略當天投放的廣告，廣告投放後隔天才會生效，此限制能大幅提升緩存效果
- 因不在 API 的範圍此流程為示意使用

![](https://drive.google.com/u/2/uc?id=1Gjsjsqsc5qafWmcReNUMRDQgzQiXJqC6&export=download)
- 將請求的內容值依順序串接得到一組唯一值供後續查詢使用
- 緩存設計為 Local Memory 與 Redis 緩存，透過上述的唯一值存取
- 考量緩存穿透問題，導致查詢請求都跑到資料庫，透過分散式鎖與有限度的自璇鎖來初始化Redis資料與等待初始化處理
- 有限度的自璇鎖，每次等待時間會指數型增加分別為100、200、400ms，且會亂數加上0-100 ms

## 評分
### 正確性
#### Create Ad API
![](https://drive.google.com/u/2/uc?id=1-9VNPVNlSoJyfplCTGNAfDwy-9EyjnKg&export=download)

#### Get Ad API
![](https://drive.google.com/u/2/uc?id=1-iXcWo5bkADwOPhX1waeDy2hCVNpuzwz&export=download)

### 效能
![](https://drive.google.com/u/2/uc?id=1QJ0NhvC9rVzvBMOTSc6eFeweoddzr9WQ&export=download)
- 以資料已熱加載為前提，個人電腦處理的處理量達6595/s
- 使用生產環境的計算能力依規格是有機會單台突破10000/s
- Local Memory 的 Cache若都被擊中，則 Server 之間的資源是獨立的，互不相影響，因此使用 Server Cluster，僅管生產環境的計算能力與個人電腦相同，也可很快破10000/s
- 若 Server Cluster 的 Local Memory 沒被擊中，也有 redis 緩存，快速同步結果到Local Memory

### 可讀性
![](https://drive.google.com/u/2/uc?id=1zK3B142B8vftbr0p1A85Rd6o6CTMqvKN&export=download)

多數實作風格方式是參考 Uncle Bob 相關書籍
- package 以 by feature 設計，以模組為主建立資料夾，能快速清楚了解專案模組架構與藉由package來限制方法被存取的使用範圍
- Handler 主要職責為分發工作
- Converter 與外部交互的資料格式轉換會由這一層處理
- Service 主要處理邏輯
- Dao 處理資料庫相關工作
- Util 額外共用程式碼

### 測試
- 以商業邏輯為主的測試，因此主要注重在使用者正常使用的情況下，是否有正確的回傳 Http Status 200與400
- Http Status 500 理論上是使用者操作不出來的，多數是相關服務出現錯誤，因此不在測試的重點範圍內

#### Create Ad API
- 測試參數全部為合法值，正常回傳
- 測試 Gender 值不在合法範圍內，是否有回傳錯誤
- 測試 Country 值不在合法範圍內，是否有回傳錯誤
- 測試 Platform 值不在合法範圍內，是否有回傳錯誤
- 測試 EndAt 日期在 StartAt之前，是否有回傳錯誤

#### Get Ad API
- 測試參數全部為合法值，正常回傳

