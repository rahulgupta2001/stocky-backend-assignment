# Stocky Backend Assignment

## Overview
Stocky is a backend service for a hypothetical platform where users earn **shares of Indian stocks** (e.g., Reliance, TCS, Infosys) as incentives for onboarding, referrals, or trading milestones.  

This backend handles:
- Rewarding users with stocks.
- Tracking internal ledger for fees.
- Providing user portfolio and daily stats.
- Calculating current INR value of holdings.

---

## Tech Stack
- **Language:** Golang  
- **Framework:** Gin (`github.com/gin-gonic/gin`)  
- **Database:** PostgreSQL  
- **Logging:** Logrus (`github.com/sirupsen/logrus`)  
- **UUID generation:** `github.com/google/uuid`  

---

## Setup Instructions

### 1. Clone the repository
```bash
git clone https://github.com/rahulgupta2001/stocky-backend-assignment.git
cd stocky-backend-assignment
