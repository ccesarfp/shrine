
# Shrine

The Shrine was created as a study proposal on Golang, gRPC, and Redis Database.

## How it Works?

<details open>
<summary>First case</summary>
A system sends a user's data to the Shrine. When this data arrives, it undergoes validation. If the data is filled out, a JWT is created and stored in Redis. Subsequently, the JWT is sent back to the system that made the initial request.
</details>

<details>
<summary>Second case</summary>
The same system needs to retrieve data from within the JWT. Therefore, it sends a request to the Shrine, which retrieves the data and returns it to the requesting system if the token is still valid. If it's not valid, the Shrine returns an error indicating that the token has expired.
</details>

<details>
<summary>Third case</summary>
If the system in question needs to check if the token still exists in Redis, it sends to the Shrine the user's ID and the name of the requesting system. Upon reaching the Shrine, these data are concatenated and searched in Redis. If not found, a "Content Not Found" error is returned. If found, the Shrine returns the token and related data.
</details>

<details>
<summary>Fourth case</summary>
The system only wants to verify the validity of the token. To do this, it sends the token to the Shrine, which checks its validity and returns "true" if it's valid or a validity error otherwise.
</details>

## How to Install?

Firstly, ensure that you have Docker installed and running.

By default, Redis will be started on port **6379** and will be created with the password **123**.

Rename the file **.env.example** to **.env**.

Access the .env file and include the **password (REDIS_PASSWORD)** and a **secret phrase (JWT_SECRET_KEY)** without spaces; this phrase will be used for JWT creation.

Inside the project folder, execute the following command:
```bash
  docker-compose up -d
```
Now, Docker will fetch the Redis image and start the service.

Once the process is complete, it's time to start the Shrine by running the following command:

```bash
  go run web/shrine/main.go
```
For testing, you can use applications like **Postman**. Just import the **Token.proto** file.

## gRPC Documentation

### Token.proto

#### Token  Service
| Method | Request | Response | Description |
| --- | --- | --- | --- |
| CreateToken  | UserRequest | TokenResponse | Create token using user data and return JWT |
| GetClaimsByToken | TokenRequest | UserResponse | Receive token and return all user data |
| GetClaimsByKey  | TokenRequestWithId | UserResponseWithToken | Receive token ID and return all user data |
| CheckTokenValidity  | TokenRequest | TokenStatus | Receive token and return if is valid |

<details>
<summary>UserRequest</summary>
  
Request message for CreateToken
| Field | Type | Description |
| --- | --- | --- |
| id | int64 | User id |
| name | string | User name |
| appOrigin  | string | Application that sent the request |
| accessLevel  | int32 | User access level |
| hoursToExpire  | int32 | Token duration |

</details>


<details>
<summary>TokenRequest</summary>
  
Request message for GetClaimsByToken and CheckTokenValidity
| Field | Type | Description |
| --- | --- | --- |
| token | string | User token |

</details>


<details>
<summary>TokenRequestWithId</summary>

Request message for GetClaimsByKey
| Field | Type | Description |
| --- | --- | --- |
| id | string | Token id |

</details>


<details>
<summary>TokenResponse</summary>

Response message for CreateToken
| Field | Type | Description |
| --- | --- | --- |
| token | string | User token |

</details>


<details>
<summary>UserResponse</summary>

Response message for GetClaimsByToken
| Field | Type | Description |
| --- | --- | --- |
| id | int64 | User id |
| name | string | User name |
| appOrigin   | int32 | Application that sent the request |
| accessLevel | int32 | User access level |

</details>

<details>
<summary>UserResponseWithToken</summary>
  
Response message for GetClaimsByKey
| Field | Type | Description |
| --- | --- | --- |
| id | int64 | User id |
| name | string | User name |
| accessLevel  | int32 | User access level |
| token | string | User token |

</details>


<details>
<summary>TokenStatus</summary>

Response message for CheckTokenValidity
| Field | Type | Description |
| --- | --- | --- |
| status | bool | Token status |

</details>
