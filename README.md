# ギリギリ飯

[![Build Frontend](https://github.com/momeemt/2000s/actions/workflows/frontend.yml/badge.svg?branch=main)](https://github.com/momeemt/2000s/actions/workflows/frontend.yml)

![](frontend/public/logo.png)

ギリギリ飯 is a service that provides a list of open restaurants in the vicinity of your current location.

https://girigirimeshi.netlify.com

## Dev
The frontend is implemented in Next.js and the backend in Go.

### Required
- make
- Docker

### Build Frontend

```sh
cd frontend
yarn install
yarn dev
```

### Build Backend

```sh
cd backend
make build
make up
```
