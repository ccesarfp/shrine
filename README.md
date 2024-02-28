<p align="center">
  <img alt="Shrine Logo" src="https://drive.google.com/uc?export=view&id=1p9ZayHmx9gldSom-VAnBPO9sHcGlSHIh" width="650px" />
  <h1 align="center">Shrine</h1>
</p>

## What is It?
Shrine is an authentication microservice that enables the rapid and secure creation of JWT tokens. 

It utilizes gRPC communication and stores data in Redis.

## Example Architecture

<p align="center">
  <img alt="Architecture Diagram" src="https://drive.google.com/uc?export=view&id=1a7duFkkFd4hPUVfX0raTU3NFmNjwq7i3" width="700px" />
</p>

## How it Works?

<details open>
<summary>First case</summary>
A user accesses for the first time (or is not authenticated), then their access data is sent to the main application, which redirects to the Shrine. The Shrine generates an opaque token and a JWT with the user's IP address data. Finally, the opaque token is returned to the user to be used as an access token.
</details>

<details>
<summary>Second case</summary>
After accessing for the first time, this user registers (or logs in) and their authentication data is sent to the main server. Once the main system confirms who the user is, it creates a JWT with all the data it will need for internal use and then directs this token to the Shrine, which updates the Opaque Token to store this new JWT.
</details>

<details>
<summary>Third case</summary>
This user has just accessed their profile and made a change to their name, so their updated data is sent to the main application (and their Opaque Token accompanies it). Upon arriving at the main application, this Opaque Token is redirected to the Shrine, which finds its previously stored JWT and sends it back to the main application, allowing it to continue saving the new data.
</details>

<details>
<summary>Fourth case</summary>
After a few days, the user accessed the application again to make a new change to their profile, but now, due to the time without access, their token was revoked. To deal with this, the Shrine notifies the main application, after trying to find its old Opaque Token, that it will take the user back to the authentication screen.
</details>

<details>
<summary>Fifth case</summary>
While processing data, the main application was unsure whether that user should still be accessing the system or not. Therefore, it forwards this Opaque Token to the Shrine, which checks its validity and notifies the main application about its current status.
</details>

## How to Install?

Firstly, ensure that you have Docker installed and running.

By default, Redis will be started on port **6379** and will be created with the password **123**.

Rename the file **.env.example** to **.env**.

Access the **.env** file and include the **secret phrase (JWT_SECRET_KEY)** without spaces, this phrase will be used for JWT creation.

Inside the project folder, execute the following command:
```bash
  docker-compose up
```
Now Docker will download the Redis image and build our Shrine image. After finishing the service will start.

And that's it, the Shrine is up and running and ready to use.

For testing, you can use applications like **Postman**. Just import the **Token.proto** file.

## gRPC Documentation

<details>
<summary>Token.proto</summary>

  #### Token  Service
  | Method | Request | Response | Description                                 |
  | --- | --- | --- |---|
  | CreateToken  | UserRequest | UserResponse | Create token using user data and return JWT |
  | UpdateToken | UserUpdateRequest | UserResponse | Receive opaque token and update linked jwt  |
  | GetJwt  | TokenRequest | TokenResponse | Receive opaque token and return user jwt    |
  | CheckTokenValidity  | TokenRequest | TokenStatus | Receive token and return if is valid        |
  
  <details>
  <summary>UserRequest</summary>
    
  Request message for CreateToken
  | Field | Type | Description |
  | --- | --- | --- |
  | hoursToExpire  | int32 | Token duration |
  
  </details>
  
  
  <details>
  <summary>UserUpdateRequest</summary>
    
  Request message for UpdateToken
  | Field | Type | Description |
  | --- | --- | --- |
  | token | string | User opaque token |
  | jwt | string | User jwt |
  | hoursToExpire | int32 | Token duration |
  
  </details>
  
  
  <details>
  <summary>TokenRequest</summary>
  
  Request message for GetJwt and CheckTokenValidity
  | Field | Type | Description |
  | --- | --- | --- |
  | token | string | Opaque token |
  
  </details>

  <details>
  <summary>UserResponse</summary>
  
  Response message for CreateToken and UpdateToken
  | Field | Type | Description |
  | --- | --- | --- |
  | token | string | User opaque token |
  
  </details>

  <details>
  <summary>TokenResponse</summary>

  Response message for GetJwt
  | Field | Type | Description |
  | --- | --- | --- |
  | jwt | string | User jwt |

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

