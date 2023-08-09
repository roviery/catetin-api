# Catetin API

## Tech Stack
- [Golang](https://go.dev/)
  - [Echo](https://echo.labstack.com/)
  - [JWT](https://echo.labstack.com/cookbook/jwt/)
- [MongoDB](https://www.mongodb.com/)
## How to install
1. Clone this repository (make sure you have go language installed in your device)
2. Make your own MongoDB in [MongoDB Atlas](https://www.mongodb.com/atlas/database) or [MongoDB Compass](https://www.mongodb.com/products/compass)
3. Create *config.json*
4. Put your MongoDB Connection URL in *config.json*
    ```
    {
      "mongo": {
        "uri": "mongodb+srv://user:password@example.wbbxw76.mongodb.net/test"
      }
    }
    ```
5. Run this command
    ```
    go mod tidy
    go run main.go
    ```

## Documentation

### **Users**
| Service | Method | URL Endpoint | Authorization | Request Body |
| :--- | :--- | :--- | :--- | :--- |
| Register | POST | /user/register | - | fullname, email, password |
| Login | POST | /user/login | Bearer Token | email, password |

### **Deadlines**
| Service | Method | URL Endpoint | Authorization | Request Body |
| :--- | :--- | :--- | :--- | :--- |
| Create Deadline | POST | /user/deadline | Bearer Token | deadline_date, task, priority |
| Get Deadline By ID | GET | /user/deadline/:id | Bearer Token | - |
| Get Deadlines By Owner ID | GET | /user/deadlines | Bearer Token | - |
| Update Deadline By ID | PUT | /user/deadline/:id | Bearer Token | deadline_date, task, priority |
| Delete Deadline By ID | DELETE | /user/deadline/:id | Bearer Token | - |

### **Quicknotes**
| Service | Method | URL Endpoint | Authorization | Request Body |
| :--- | :--- | :--- | :--- | :--- |
| Create Quicknote | POST | /user/quicknote | Bearer Token | note |
| Get Quicknote By ID | GET | /user/quicknote/:id | Bearer Token | - |
| Get Quicknotes By Owner ID | GET | /user/quicknotes | Bearer Token | - |
| Update Quicknote By ID | PUT | /user/quicknote/:id | Bearer Token | note |
| Delete Quicknote By ID | DELETE | /user/quicknote/:id | Bearer Token | - |

### **Finances**
| Service | Method | URL Endpoint | Authorization | Request Body |
| :--- | :--- | :--- | :--- | :--- |
| Create Finance | POST | /user/finance | Bearer Token | type, fund_allocation |
| Get Finance By ID | GET | /user/finance/:id | Bearer Token | - |
| Get Finances By Owner ID | GET | /user/finances | Bearer Token | - |
| Update Finance By ID | PUT | /user/finance/:id | Bearer Token | type, fund_allocation |
| Delete Finance By ID | DELETE | /user/finance/:id | Bearer Token | - |

### **Transactions**
| Service | Method | URL Endpoint | Authorization | Request Body |
| :--- | :--- | :--- | :--- | :--- |
| Create Transaction | POST | /user/transaction | Bearer Token | finance_id, title, expense |
| Get Transaction By ID | GET | /user/transaction/:id | Bearer Token | - |
| Get Transactions By Owner ID | GET | /user/transactions | Bearer Token | - |
| Update Transaction By ID | PUT | /user/transaction/:id | Bearer Token | finance_id, title, expense |
| Delete Transaction By ID | DELETE | /user/transaction/:id | Bearer Token | - |

### **Todos**
| Service | Method | URL Endpoint | Authorization | Request Body |
| :--- | :--- | :--- | :--- | :--- |
| Create Todo | POST | /user/todo | Bearer Token | todo, status |
| Get Todo By ID | GET | /user/todo/:id | Bearer Token | - |
| Get Todos By Owner ID | GET | /user/todos | Bearer Token | - |
| Update Todo By ID | PUT | /user/todo/:id | Bearer Token | todo, status |
| Delete Todo By ID | DELETE | /user/todo/:id | Bearer Token | - |
