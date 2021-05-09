package com.mikolajmgt.authorizer.model

import javax.persistence.Column
import javax.persistence.Entity
import javax.persistence.GeneratedValue
import javax.persistence.GenerationType
import javax.persistence.Id

@Entity
class User {
    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    val id: Long = 0

    @Column(unique = true)
    val email: String = ""

    @Column(unique = true)
    val nickname: String = ""

    @Column
    val password: String = ""
}
