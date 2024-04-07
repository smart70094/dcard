# Dcard Ad API

## 設計
### 基礎架構
![](https://drive.google.com/u/2/uc?id=1Cw_rFg_RnuDN9NGBNc1CjMUgiXufEwJJ&export=download)
### 流程

![](https://drive.google.com/u/2/uc?id=1g4CkPX9lrH21ote-KpMkdYgNxtM3cifs&export=download)
- Eager Loading 將活躍廣告寫入至 hot data table 供後續資料庫查詢使用
- 因題目無特別要求，因此此處忽略當天投放的廣告，廣告投放後隔天才會生效，此限制能大幅提升緩存效果
- 因不在 API 的範圍此流程為示意使用
  
![](https://drive.google.com/u/2/uc?id=1H41jSA_8G6qw9179Tt-1JQlz7Beq_tmi&export=download)
- 將請求的內容值依順序串接得到一組唯一值供後續查詢使用
- 緩存設計為 Local Memory 與 Redis 緩存，透過唯一值存取
- 考量緩存穿透問題，導致查詢請求都跑到資料庫，透過分散式鎖與有限度的自璇鎖來初始化Redis資料與等待初始化過程
- 有限度的自璇鎖，每次等待時間會指數型增加分別為100、200、400，且會亂數加上0-100 ms

## 評分
### 正確性
#### Create Ad API
![](https://drive.google.com/u/2/uc?id=1-pHxbc3RTxw7gkE_gMajbxzqsNYGpK4Y&export=download)

#### Get Ad API
![](https://drive.google.com/u/2/uc?id=1XiKziJ1ZjAXyIEDq0IRWTTK3TTLk8EF4&export=download)

