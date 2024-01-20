<p align="center">
  <img alt="Shrine Logo" src="https://drive.google.com/uc?export=view&id=1p9ZayHmx9gldSom-VAnBPO9sHcGlSHIh" width="650px" />
  <h1 align="center">Shrine</h1>
</p>

## What is It?
Shrine is an authentication microservice that enables the rapid and secure creation of JWT tokens. 

It utilizes gRPC communication and stores data in Redis.

## Example Architecture

<p align="center">
  <img alt="Shrine Logo" src="https://drive.google.com/uc?export=view&id=1T3Os2tbp8wTVxQXaiEbB9y-7OSKOxOlu" width="800px" />
</p>

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

Access the **.env** file and include the **secret phrase (JWT_SECRET_KEY)** without spaces, this phrase will be used for JWT creation.

Inside the project folder, execute the following command:
```bash
  docker-compose up -d
```
Now Docker will download the Redis image and build our Shrine image. After finishing the service will start.

And that's it, the Shrine is up and running and ready to use.

For testing, you can use applications like **Postman**. Just import the **Token.proto** file.

## gRPC Documentation

<details>
<summary><h2>Token.proto</h2></summary>

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

</details>

## Why Does the Shrine Exist?

The Shrine was designed with the aim of creating a service that could be used by multiple applications, avoiding the need to develop a token management system for each of them. 

The choice to use the Go language stemmed from the desire to enhance my knowledge in it, coupled with its high execution speed. 

Additionally, the decision was made to implement gRPC and utilize the Redis database for the temporary persistence of these tokens.

## What Did I Learn?

With this project, I could realize how enjoyable it is to program using Golang, both in creating gRPC servers and integrating with Redis.

Beyond the technical aspects of development, I was able to enhance my architectural vision. Throughout the entire development process, I consistently thought about how the system could integrate with other applications and what role it should play within a complete architecture.

