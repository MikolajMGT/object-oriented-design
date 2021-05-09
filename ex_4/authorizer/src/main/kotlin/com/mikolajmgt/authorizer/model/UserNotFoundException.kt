package com.mikolajmgt.authorizer.model

import org.springframework.http.HttpStatus
import org.springframework.web.bind.annotation.ResponseStatus

@ResponseStatus(code = HttpStatus.NOT_FOUND, reason = "user not found")
class UserNotFoundException : RuntimeException()
