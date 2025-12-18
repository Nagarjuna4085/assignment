# Stocky Assignment

A stock reward system built using Golang, Gin, PostgreSQL.

## Features
- Stock reward APIs
- Double-entry ledger
- Hourly mock price updates
- Idempotency handling
- Portfolio valuation



## API specifications (request/response payloads)
1. POST /reward
Request:
{
  "event_id": "evt-1",
  "user_id": "user_1",
  "stock_symbol": "RELIANCE",
  "quantity": 1.25,
  "timestamp": "2025-12-18T10:30:00Z"
}
Response:
{
  "status": "success"
}
Duplicate event handling: If the same event_id is submitted again:

{
  "message": "duplicate event ignored"
}

2.GET /today-stocks/{userId}
Request :
/today-stocks/user_1

Response:

{
  "user_id": "user_1",
  "date": "2025-12-18",
  "rewards": [
    {
      "ID": "uuid",
      "EventID": "evt-1",
      "UserID": "user_1",
      "StockSymbol": "RELIANCE",
      "Quantity": 1.25,
      "RewardedAt": "2025-12-18T10:30:00Z"
    }
  ]
}



3.GET /historical-inr/{userId}
Request : 
/historical-inr/user_1

Response
{
  "user_id": "user_1",
  "history": {
    "2025-12-17": 1839.28,
    "2025-12-16": 1125.50
  }
}

4.GET /stats/{userId}

Request: /stats/user_1

Response :
{
  "user_id": "user_1",
  "today_rewards": {
    "RELIANCE": 1.25
  },
  "current_portfolio_inr": 1839.28
}


5️ GET /portfolio/{userId} (Optional / Bonus)

Request :
{
  "user_id": "user_1",
  "portfolio": {
    "RELIANCE": 1.25,
    "TCS": 2.0
  },
  "total_inr": 2950.50
}

Response
{
  "user_id": "user_1",
  "portfolio": {
    "RELIANCE": 1.25,
    "TCS": 2.0
  },
  "total_inr": 2950.50
}


Database Schema

| Column       | Type          | Notes                       |
| ------------ | ------------- | --------------------------- |
| id           | UUID (PK)     | Primary key                 |
| event_id     | VARCHAR       | Unique, prevents duplicates |
| user_id      | VARCHAR       |                             |
| stock_symbol | VARCHAR       |                             |
| quantity     | NUMERIC(18,6) | Fractional shares allowed   |
| rewarded_at  | TIMESTAMP     | When reward happened        |
| created_at   | TIMESTAMP     | Auto-created                |


ledger_entries


| Column          | Type          | Notes                               |
| --------------- | ------------- | ----------------------------------- |
| id              | UUID (PK)     |                                     |
| reward_event_id | UUID (FK)     | FK to reward_events                 |
| stock_symbol    | VARCHAR       |                                     |
| stock_qty       | NUMERIC(18,6) | Positive for credit, negative debit |
| cash_inr        | NUMERIC(18,4) | Company outflow                     |
| fees_inr        | NUMERIC(18,4) | Brokerage, STT, GST, etc.           |
| entry_type      | VARCHAR       | "debit"/"credit"                    |
| created_at      | TIMESTAMP     |                                     |




Edge Cases & Scaling

Handled:

Duplicate rewards → checked via event_id

Rounding errors → INR rounded to 2 decimals

Price API downtime → uses last stored price

Not implemented yet (optional / future):

Refunds / adjustments → negative ledger entries

Stock splits / mergers → adjustment logic needed



Scaling considerations:

Use indexes on user_id, rewarded_at, stock_symbol

Background job to fetch hourly stock prices

Cache latest stock prices for fast computation

Ledger table ensures auditability and double-entry compliance


## Run
go run main.go
