# gymshark-challenge

## Frontend
Developed in react. To start it, clone the repo and install dependencies using ```npm install```.
To run the application, run ```yarn start``` or ```npm start```.

The frontend is hosted on AWS using AWS Amplify which provides ci/cd.

Website URL: https://main.dcmc6mnsg1fp7.amplifyapp.com

## Backend
Developed in Go with AWS Lambda + AWS API Gateway.
The Backend can also be tested directly by cloning the repo and running ```go test```.

API URL: https://4qyor869oj.execute-api.eu-west-2.amazonaws.com/demo/get-needed-packs

To to use the API, send a POST request to the API URL with the below body:

```
{
    "packs": <array of available packs> e.g. [ 250, 500, 1000, 2000, 5000 ],
    "quantity": <number of quantity> e.g. 250
}

```
