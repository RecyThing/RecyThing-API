[![codecov](https://codecov.io/gh/RecyThing/RecyThing-API/graph/badge.svg?token=ASZEGN68UP)](https://codecov.io/gh/RecyThing/RecyThing-API)

<p align="center">
  <img src="https://github.com/RecyThing/RecyThing-API/assets/66883583/ba81aec0-9b72-4633-a282-b20ac7c8a833" alt="Nama Alternatif">
</p>

---
## Table of Contents

- [Folder Structure](#folder-structure)
- [Tech Stack](#tech-stack)
- [API Documentation](#api-documentation)
- [Unit Testing](#unit-testing)
- [Running Project](#running-project)
- [How to deploy to Cloud Run](#how-to-deploy-to-cloud-run)
- [CI/CD Platform](#cicd-platform)
- [Team Back-End](#team-back-end)
- [References](#references)

## Folder Structure
```
├── app                            
│   ├── config                      # configuration connection any database
│   ├── database                    # any database connection
│   └── route                      
├── features
│   └── featureA
│       ├── dto                    
│       │   ├── request
│       │   │   ├── mapping.go
│       │   │   └── request.go
│       │   └── response
│       │       ├── mapping.go
│       │       └── response.go
│       ├── entity
│       │   ├── entity.go           
│       │   ├── interface.go        # interface contract for all function repository and service 
│       │   └── mapping.go          # mapping struct core to model or model to core
│       ├── handler
│       │   └── handler.go          # manages HTTP requests and responses.map HTTP input to calls to functions in the service layer.
│       ├── model
│       │   └── hook.go        
│       │   └── model.go            # structure database 
│       ├── repository
│       │   └── respository.go      # query for manipulating data
│       └── service
│           └── service.go          # contains the core business logic of the application and validation
├── utils                           # providing commonly used helper functions across different parts of the project
```
## Tech Stack
![header](https://github.com/RecyThing/RecyThing-API/assets/66883583/e358c9dd-ed14-4b24-8ec0-3d8ebbb40d65)
![header-sub-header-body-text-header-2](https://github.com/RecyThing/RecyThing-API/assets/66883583/57f609c9-3fbf-498e-840e-4744e2726fa1)
![header-sub-header-body-text-header-3](https://github.com/RecyThing/RecyThing-API/assets/66883583/f3c11700-2425-4b8d-b4d7-9010801f5832)
![Static Badge](https://img.shields.io/badge/RDS%20MYSQL-rds?style=for-the-badge&logo=amazonrds&logoColor=white&color=%234ba185)
![Static Badge](https://img.shields.io/badge/codecov-s?style=for-the-badge&logo=codecov&logoColor=white&color=%23F01F7A)
![Static Badge](https://img.shields.io/badge/mysql-s?style=for-the-badge&logo=mysql&logoColor=white&color=%234479A1)
![Static Badge](https://img.shields.io/badge/docker-s?style=for-the-badge&logo=docker&logoColor=white&color=%232496ED)
![Static Badge](https://img.shields.io/badge/cloudflare-s?style=for-the-badge&logo=cloudflare&logoColor=white&color=%23F38020)
![Static Badge](https://img.shields.io/badge/swagger-s?style=for-the-badge&logo=swagger&logoColor=%2385EA2D&color=black)
![Static Badge](https://img.shields.io/badge/openapi-s?style=for-the-badge&logo=openapiinitiative&logoColor=%2385EA2D&color=black)
![Static Badge](https://img.shields.io/badge/openai-s?style=for-the-badge&logo=openai&logoColor=white&color=%23412991)

## API Documentation

We use Swagger UI as the API documentation, running on Docker with JSON configuration located on an external server.

Pull a pre-built docker image of the swagger-ui directly from Docker Hub:
  ```
  docker pull swaggerapi/swagger-ui
  ```
Provide a URL to a swagger.json on an external host:
  ```
  docker run -p 80:8080 -e SWAGGER_JSON_URL=https://recything.apicode.my.id/swagger/apidoc.json swaggerapi/swagger-ui
  ```
Running on cloud run : 
  - Pull swaggerapi/swagger-ui,
    ```
    docker pull swaggerapi/swagger-ui
    docker tag swaggerapi/swagger-ui gcr.io/$PROJECT_ID/swagger-ui  #change $PROJECT_ID to your project
    ```
  - Push to container registry
    ```
    docker push gcr.io/$PROJECT_ID/swagger-ui #change $PROJECT_ID to your project
    ```
  - Deploy cloud run
    ```
    gcloud run deploy swagger-ui \
    --image=gcr.io/$PROJECT_ID/swagger-ui:latest \                                         #change $PROJECT_ID to your project
    --platform=managed \
    --allow-unauthenticated \
    --set-env-vars=SWAGGER_JSON_URL=https://recything.apicode.my.id/swagger/apidoc.json    #change to your external json
    ```
## Unit Testing

Using testify for testing code and generate mock with Mockery.

- Install Mockery
   ```
  go install github.com/vektra/mockery/v2@v2.38.0
  ```
- Generate Mock Auto Detect All Interface
  ```
  mockery -all   #The generated mocks will be available in a single file named 'mocks'.
  ```
- Test Command
  
  Running test and cover on directory test : 
  ```
  go test -v -cover   
  ```

  Running Test and cover with html view :
  ```
  go test -coverprofile=cover.out && go tool cover -html=cover.out
  ```

  Running Test All Sub Directories : 
  ```
  go list ./... | grep service | xargs -n1 go test -cover  
  ```
  #change grep with your folder name (e.g `grep handler`), It is recommended to use the same folder name.
  view like this :
 
  ![Screenshot from 2023-12-17 20-10-40](https://github.com/RecyThing/RecyThing-API/assets/66883583/ab647a0b-4c97-4c95-9400-b36897495296)
  
- Configuration Codecov

  Connect codecov first [here](https://app.codecov.io/), Look YAML configuration [here](https://github.com/RecyThing/RecyThing-API/blob/development/.github/workflows/codecov.yml)


## Running Project

Start project with Go :
```
go run main.go
```
## How to deploy to Cloud Run
  - First, build your Docker image and push it to the registry, just like you did with the Swagger UI configuration.
  - Deploy cloud run
    ```
    gcloud run deploy <Service-Name> \                            # e.g  gcloud run deploy recything
    --image=gcr.io/$PROJECT_ID/<Your-image>:latest \              #change $PROJECT_ID to your project
    --platform=managed \
    --allow-unauthenticated \
    --set-env-vars=MYSQL_PORT=3306                                #change to your env
    --set-env-vars=SERVER_PORT=  
    --set-env-vars=STORAGE=   
    ```

- Custom Domain

  Mapping domain with any DNS management , configuration mapping [here](https://console.cloud.google.com/run/domains). Click `Add Mapping`, And you can choose from 3 different ways of mapping domains, either using a custom domain with a load balancer, a custom domain with Cloud Run, or a custom domain with Firebase Hosting.

  ![Screenshot from 2023-12-18 11-06-44](https://github.com/RecyThing/RecyThing-API/assets/66883583/e8a32da4-0850-4ab1-87bf-d20e87dafda7)

  You will receive DNS records and add that configuration to the DNS Management in the platform where your domain is located


## CI/CD Platform
Cloud Build is a Google Cloud Platform CI/CD platform that is easily implemented with GitHub. Configure trigger [here](https://console.cloud.google.com/cloud-build/). We use cloudbuid.yaml for trigger CI/CD :

![Screenshot from 2023-12-18 10-50-19](https://github.com/RecyThing/RecyThing-API/assets/66883583/4635ccbe-7e4e-44b8-977b-b4e06bf5bba4)



## Team Back-End

| Name                           | University	                                         | 
| :----------------------------- | :--------------------------------------------------- | 
|	Hanief Fathul Bahri Ahmad      | Universitas Nusa Mandiri                              |	
|	Tiara Juli Arsita	             | Universitas Tadulako                                  |	
|	 Juan Azhar Adviseta Setiawan  | Universitas Pembangunan Nasional Veteran Yogyakarta   |	
|	Al Hilaluddin                  | Universitas Muslim Indonesia                          |	
|	Stanley	                       | Universitas Mikroskil                                 |	

## References

  - [Echo](https://echo.labstack.com/) : High performance, extensible, minimalist Go web framework
  - [Gorm](https://gorm.io/) : GORM facilitates the interaction with relational databases, simplifying the process of accessing and managing data
  - [Swagger-Ui](https://github.com/swagger-api/swagger-ui) : Api Documentation
  - [Go-openai](https://github.com/sashabaranov/go-openai) : Configuration Open-Ai
  - [Gorilla/schema](https://github.com/gorilla/schema) : Bind request struct
  - [Cloud Run](https://cloud.google.com/run?hl=id) : Deployment Platform
  - [Govalidator](https://github.com/asaskevich/govalidator) : Package of validators and sanitizers for strings, numerics, slices and structs
  - [Codecov](https://about.codecov.io/) : Coverage Reporting
  - [Testify](https://github.com/stretchr/testify) : A toolkit with common assertions and mocks that plays nicely with the standard library
  - [Mockery](https://github.com/vektra/mockery) : A mock code autogenerator for Go Interface






