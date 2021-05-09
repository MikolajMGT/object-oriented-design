package com.mikolajmgt.authorizer.api

import com.mikolajmgt.authorizer.model.User
import com.mikolajmgt.authorizer.model.UserNotFoundException
import com.mikolajmgt.authorizer.service.Authorizer
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.PostMapping
import org.springframework.web.bind.annotation.RequestBody
import org.springframework.web.bind.annotation.RequestMapping
import org.springframework.web.bind.annotation.RequestParam
import org.springframework.web.bind.annotation.RestController

@RestController
@RequestMapping("/authorizer")
class AuthorizerController @Autowired constructor(val authorizer: Authorizer) {

    @PostMapping("/register")
    fun register(@RequestBody user: User): User {
        return authorizer.register(user)
    }

    @GetMapping("/login")
    fun login(@RequestParam("nickname") nickname: String, @RequestParam("password") password: String): User {
        return authorizer.login(nickname, password) ?: throw UserNotFoundException()
    }

    @GetMapping("/logout")
    fun logout(@RequestParam("nickname") nickname: String): User {
        return authorizer.logout(nickname) ?: throw UserNotFoundException()
    }
}
