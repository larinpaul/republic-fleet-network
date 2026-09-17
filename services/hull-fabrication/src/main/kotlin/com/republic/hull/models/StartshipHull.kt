package com.republic.hull.models

import jakarta.persistence.*
import java.time.LocalDateTime
import java.util.UUID

/**
 * The physical database entity representing a starship hull.
 * This maps directly to the 'starship_hulls' table in PostgreSQL.
 */
@Entity
@Table(name = "starship_hulls")
data class StarshipHull(

    @Id
    @GeneratedValue(strategy = GenerationType.UUID)
    val id: UUID = UUID.randomUUID(),

    @Column(nullable = false)
    val className: String, // e.g., "X-Wing", "A-Wing"

    @Column(nullable = false)
    val hullAlloy: String, // e.g., "Durasteel", "Beskar"

    @Column(nullable = false)
    val deflectorShields: Boolean = true,

    @Enumerated(EnumType.STRING)
    @Column(nullable = false)
    var status: HullStatus = HullStatus.FABRICATING,

    @Column(nullable = false, updatable = false)
    val fabricationTimestamp: LocalDateTime = LocalDateTime.now()
)

/**
 * The lifecycle states of a hull in the shipyard.
 */
enum class HullStatus {
    FABRICATING,
    HULL_COMPLETE_AWAITING_HYPERDRIVE,
    HYPERDRIVE_INSTALLED,
    DEPLOYED
}