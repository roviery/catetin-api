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

