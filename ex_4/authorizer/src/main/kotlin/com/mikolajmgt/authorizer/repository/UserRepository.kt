package com.mikolajmgt.authorizer.repository

import com.mikolajmgt.authorizer.model.User
import org.springframework.data.repository.CrudRepository;
import org.springframework.lang.Nullable

interface UserRepository: CrudRepository<User, Int> {
    @Nullable
    fun findByNicknameAndPassword(nickname: String, password: String): User

    @Nullable
    fun findByNickname(nickname: String): User
}