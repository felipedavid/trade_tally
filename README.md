# Trade Tally
A website to help you to track your portfolio in the stock market.

### Functionalities
- Keep historic information of the stock market
- Show realtime stocks information
- Allow users to register bought stocks and keep track of the performance of their portfolio

### Database model
```mermaid
erDiagram
    stock||--o{stock_price : has
    
    stock {
       id integer PK 
       created_at text
       updated_at text
       symbol text
       company text
    }
    
    stock_price {
        id integer PK
        created_at text
        updated_at text
        stock_id bigint FK
        date text
        open integer
        high integer
        low integer
        close integer
        version bigint
    }
    
    user {
        id integer PK
        created_at text
        updated_at text
        first_name text
        last_name text 
        phone_number text
        email text
        hashed_password text
    }
    
```