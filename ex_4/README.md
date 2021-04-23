Available endpoints:

1. register new user: /authorizer/register

example input:

{
"email": "test@test.com",
"nickname": "test",
"password": "test123"
}

example output:

{
"id": 1,
"email": "test@test.com",
"nickname": "test",
"password": "test123"
}

2) login user: /authorizer/login?nickname=test&password=test123

example output:

{
"id": 1,
"email": "test@test.com",
"nickname": "test",
"password": "test123"
}

3) logout user: /authorizer/logout?nickname=test

{
"id": 1,
"email": "test@test.com",
"nickname": "test",
"password": "test123"
}