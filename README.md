# Go Sonarqube

Golang with sonarqube in local

## Description

Simple overview sonarqube usage with golang as example.

## Getting Started

### Prerequisite

- Docker
- sonar-scanner

### Installing
- run
```
git clone https://github.com/acephasanudin/go-sonarqube.git
```
- go to directory
```
cd go-sonarqube
```
- run docker compose
```
docker-compose up --build
```
- go to sonarqube page, login with
  - login: admin
  - password: admin
![image](https://github.com/acephasanudin/go-sonarqube/assets/70930571/21f757aa-ecd7-4dc3-b3aa-1dbb3386c0d7)
- change password
- create local project
![image](https://github.com/acephasanudin/go-sonarqube/assets/70930571/a1a2f949-cf12-450f-bdd6-698e30f6f1ee)
- fill display name, project key & main branch name
![image](https://github.com/acephasanudin/go-sonarqube/assets/70930571/5e52811b-f1cf-4998-bd78-d198ce0d2e16)
- use global setting
![image](https://github.com/acephasanudin/go-sonarqube/assets/70930571/378252d2-7cb4-4670-9db9-81a860138afc)
- after create project, choose locally
![image](https://github.com/acephasanudin/go-sonarqube/assets/70930571/1f753b37-4b5b-40f4-b5e0-a85f5fb33e66)
- generate token
![image](https://github.com/acephasanudin/go-sonarqube/assets/70930571/ea69d7c1-0aeb-4631-87bf-a7ed06c75063)
- run analysis on your project
```
sonar-scanner \
  -Dsonar.projectKey=Go-Sonarqube \
  -Dsonar.sources=. \
  -Dsonar.host.url=http://localhost:9000 \
  -Dsonar.token=<yourtoken>
```
![image](https://github.com/acephasanudin/go-sonarqube/assets/70930571/350484ed-d21d-4cdf-906c-435284d193de)
![image](https://github.com/acephasanudin/go-sonarqube/assets/70930571/df915e45-de0e-49bd-b6eb-a48cacb3cb79)
![image](https://github.com/acephasanudin/go-sonarqube/assets/70930571/f4e09739-5e79-4cd9-baf7-a31095d1cb2d)
- see the results in http://localhost:9000/dashboard?id=Go-Sonarqube
![image](https://github.com/acephasanudin/go-sonarqube/assets/70930571/4d72f139-8a10-482f-9ae0-b48b534abcdd)
![image](https://github.com/acephasanudin/go-sonarqube/assets/70930571/48c2a382-b94b-4ec7-9db2-44969ad30546)
