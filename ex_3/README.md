**Endpoints**

- /user/get
- /user/create
- /user/update
- /user/delete

**Examples**

- /user/get

  request:
  
  /user/get?id=1
  
  response:

  {
  "id": 1,
  "email": "test@test.coma",
  "username": "test",
  "password": "qwerty123"
  }

- /user/create

  request: 

  {
  "email": "test@test.com",
  "username": "test",
  "password": "qwerty123"
  }

  response:

  {
  "id": 1,
  "email": "test@test.com",
  "username": "test",
  "password": "qwerty123"
  }

- /user/update

  request:

  {
  "id": 1,
  "email": "test@test.com",
  "username": "another",
  "password": "one"
  }

  response:

  {
  "id": 1,
  "email": "test@test.com",
  "username": "another",
  "password": "one"
  }

- /user/delete

  request:

  {
  "id": 1
  }

  response:

  {
  "id": 1,
  "email": "test@test.coma",
  "username": "test",
  "password": "qwerty123"
  }