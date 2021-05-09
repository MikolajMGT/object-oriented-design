package com.mikolajmgt.authorizer.service

import com.mikolajmgt.authorizer.model.User
import com.mikolajmgt.authorizer.repository.UserRepository
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.stereotype.Service

@Service
class Authorizer @Autowired constructor(val userRepository: UserRepository) {

    fun register(user: User): User {
        return userRepository.save(user)
    }

    fun login(nickname: String, password: String): User? {
        return userRepository.findByNicknameAndPassword(nickname, password)
    }

    fun logout(nickname: String): User? {
        return userRepository.findByNickname(nickname)
    }
}
