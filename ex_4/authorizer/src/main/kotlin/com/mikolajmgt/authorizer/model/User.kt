package com.mikolajmgt.authorizer.model

import javax.persistence.*

@Entity
class User {
    @Id
    @GeneratedValue(strategy=GenerationType.AUTO)
    val id: Long = 0

    @Column(unique = true)
    val email: String = ""

    @Column(unique = true)
    val nickname: String = ""

    @Column
    val password: String = ""
}